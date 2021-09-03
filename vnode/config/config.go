package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/MOACChain/MoacLib/log"
)

// Configuration to support the info in vnodeconfig.json
type Configuration struct {
	VssBaseAddr string `json:VssBaseAddr`
	VnodeIP     string `json:VnodeIP`
	VnodePort   string `json:VnodePort`
	ChainId     int    `json:ChainId`
}

// GetConfiguration: read config from .json file
// 1) No config file, using default value, don't create new file;
// 2) has config file, error in reading config, stop and display correct info;
// 3) has config file, no error in reading config, go ahead load and run;
func GetConfiguration(configFilePath string) (*Configuration, error) {
	//check the file path
	filePath, _ := filepath.Abs(configFilePath)
	log.Debugf("load vnode config file from %v", filePath)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		//If no config file exists, return nil
		log.Infof("%v not exists, use default settings", configFilePath)
		return nil, nil
	}

	if _, err := os.Stat(configFilePath); err != nil {

		log.Errorf("Open %v error: \n%v\n", configFilePath, err)

		return nil, err
	}

	return GetUserConfig(configFilePath)
}

//GetUserConfig -- read the user configuration in json file
func GetUserConfig(filepath string) (*Configuration, error) {
	// Open our jsonFile
	jsonFile, err := os.Open(filepath)
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// initialize the Confgiruation array
	var conf Configuration

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &conf)

	return &conf, nil
}
