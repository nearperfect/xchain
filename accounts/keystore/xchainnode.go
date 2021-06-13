package keystore

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/log"
	"github.com/MOACChain/xchain/accounts"
)

type KeyStoreInfo struct {
	Address    common.Address
	Passphrace string
	KeyStore   string
}

var iv = []byte{
	0xBA, 0x37, 0x2F, 0x02, 0xC3, 0x92, 0x1F, 0x7D,
	0x7A, 0x3D, 0x5F, 0x06, 0x41, 0x9B, 0x3F, 0x2D,
	0xBA, 0x37, 0x2F, 0x02, 0xC3, 0x92, 0x1F, 0x7D,
	0x7A, 0x3D, 0x5F, 0x06, 0x41, 0x9B, 0x3F, 0x2D,
}

var passphraceFileName = "passphrace"
var XBasePath = ""

func SetXBasePath(prefix string) {
	XBasePath, _ = filepath.Abs(filepath.Join(prefix, "keystorex"))
}

func EncryptX(text string, key []byte) (string, error) {
	var iv = key[:aes.BlockSize]
	encrypted := make([]byte, len(text))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypter.XORKeyStream(encrypted, []byte(text))
	return hex.EncodeToString(encrypted), nil
}

func DecryptX(encrypted string, key []byte) (string, error) {
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	src, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	var iv = key[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var block cipher.Block
	block, err = aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypter.XORKeyStream(decrypted, src)
	return string(decrypted), nil
}

/*
 * Save the passphrace with keystore file
 * for later usage.
 */
func SaveXPassphrace(passphrase string) error {
	//Encrypt the passphrase into a file
	encrypt, _ := EncryptX(passphrase, iv)

	//Save with keystore
	stat, err := os.Stat(XBasePath)
	if err != nil {
		ferr := os.MkdirAll(XBasePath, os.ModePerm)
		if ferr != nil {
			return ferr
		} else {
			log.Info("Create default keystore dir.")
		}
	}

	//Check again for the base path
	stat, err = os.Stat(XBasePath)

	if err == nil && stat.IsDir() {
		filepath := filepath.Join(XBasePath, passphraceFileName)
		err = ioutil.WriteFile(filepath, []byte(encrypt), os.ModePerm)
		if err != nil {
			log.Error("write file err")
			return err
		} else {
			log.Info("Passphrace saved successfully!")
			return nil
		}
	}

	return errors.New(
		"Basepath error: Either nots exist or is a file. Please remove the file and try again!",
	)
}

//Get the passphrace from file
func GetXPassphrace() string {
	filepath := filepath.Join(XBasePath, passphraceFileName)
	fd, err := os.OpenFile(filepath, os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Errorf("GetPassphrace() open file error: %v", err)
		return ""
	}
	defer fd.Close()

	// get the file string
	buf_len, _ := fd.Seek(0, os.SEEK_END)
	fd.Seek(0, os.SEEK_SET)
	bufio := make([]byte, buf_len)
	fd.Read(bufio)
	encrypt := string(bufio[:])
	decrypt, _ := DecryptX(encrypt, iv)

	return decrypt
}

func CreateXAccount() error {
	passphrase := GetXPassphrace()
	if len(passphrase) == 0 {
		return errors.New("Need to have a valid passphrase, program exit!")
	}
	// Create an encrypted keystore with standard crypto parameters
	ks := NewKeyStore(XBasePath, StandardScryptN, StandardScryptP)
	// Create account
	a, err := ks.NewAccount(passphrase)
	log.Debug("Created account", "Address:", a.Address, "URL:", a.URL)
	if err != nil {
		log.Errorf("Failed to create new account: %v", err)
		return err
	}

	return nil
}

// scan and locate the keystore file
func scanXKeyFile() (string, error) {
	files, err := ioutil.ReadDir(XBasePath)

	if err != nil {
		return "", err
	}

	for _, fi := range files {
		path := filepath.Join(XBasePath, fi.Name())
		// Skip editor backups and UNIX-style hidden files.
		if strings.HasSuffix(fi.Name(), "~") || strings.HasPrefix(fi.Name(), ".") {
			continue
		}
		// Skip misc special files, directories (yes, symlinks too).
		if fi.IsDir() || fi.Mode()&os.ModeType != 0 {
			continue
		}

		// Only to get one account
		find, err := regexp.MatchString("^UTC--.*", fi.Name())
		if find && err == nil {
			log.Debug("Found keystore on account scan", "path", path)
			return path, nil
		}
	}
	return "", errors.New("KeyFile not found.")
}

func GetOrCreateXKeyStore() (*KeyStoreInfo, error) {
	var (
		buf        = new(bufio.Reader)
		address    common.Address
		jsonString string
		keyJSON    struct {
			Address string `json:"address"`
		}
	)
	// Create base path
	os.MkdirAll(XBasePath, os.ModePerm)
	log.Infof("XBasepath: %v", XBasePath)

	path, err := scanXKeyFile()
	if err != nil {
		log.Debug("Failed to get keystore file try to create")
		err := CreateXAccount()
		if err != nil {
			return nil, err
		}
		path, err = scanXKeyFile()

		if err != nil {
			log.Debug("Failed to get keystore file after creating")
			return nil, err
		}
	}

	fd, err := os.Open(path)
	if err != nil {
		log.Infof("Keystore PATH:%v\n", path)
		log.Debug("Failed to open keystore file", "err", err)
		return nil, err
	}

	buf.Reset(fd)
	// Parse the address.
	keyJSON.Address = ""
	err = json.NewDecoder(buf).Decode(&keyJSON)
	address = common.HexToAddress(keyJSON.Address)

	// get the file string, return error
	buf_len, berr := fd.Seek(0, os.SEEK_END)
	if berr != nil {
		log.Infof("Failed to find the length of the key store file")
		return nil, err
	}
	fd.Seek(0, os.SEEK_SET)
	bufio := make([]byte, buf_len)
	fd.Read(bufio)
	jsonString = string(bufio[:])

	switch {
	case err != nil:
		log.Debug("Failed to decode keystore key", "err", err)
	case (address == common.Address{}):
		log.Debug("Failed to decode keystore key", "err", "missing or zero address")
	default:
	}
	fd.Close()

	return &KeyStoreInfo{
		Address:    address,
		Passphrace: GetXPassphrace(),
		KeyStore:   jsonString,
	}, nil
}

func GetXDecryptedKey(a accounts.Account, passphrase string) (accounts.Account, *Key, error) {
	ks := NewKeyStore(XBasePath, StandardScryptN, StandardScryptP)
	return ks.getDecryptedKey(a, passphrase)
}

////////////////////////////
// Vss key section
////////////////////////////

func vssKeyFileName(vssBaseAddr common.Address) string {
	return filepath.Join(XBasePath, fmt.Sprintf("%x.vsskey", vssBaseAddr))
}

func GetVSSKey(vssBaseAddr common.Address) []byte {
	keyFile := vssKeyFileName(vssBaseAddr)
	content, _ := ioutil.ReadFile(keyFile)
	return content
}

func PutVSSKey(vssBaseAddr common.Address, data []byte) error {
	keyFile := vssKeyFileName(vssBaseAddr)
	err := ioutil.WriteFile(keyFile, data, 0644)
	return err
}

// compare two bls signatures, 1: > , 0: = , -1: <
func CmpSigs(sig1, sig2 []byte) int {
	if len(sig1) != len(sig2) {
		return 1
	}

	for i, b1 := range sig1 {
		b2 := sig2[i]
		if uint8(b1) > uint8(b2) {
			return 1
		} else if uint8(b1) < uint8(b2) {
			return -1
		}
	}

	return 0
}
