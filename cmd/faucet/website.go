// Code generated by go-bindata.
// sources:
// faucet.html
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _faucetHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x59\x7b\x73\xdb\x36\x12\xff\xdb\xf9\x14\x5b\x5e\x5a\x4b\x63\x93\xb4\xe3\x4c\xda\x91\x49\x75\x32\x69\x2f\xed\xcd\xf5\x31\x6d\x3a\x77\x9d\xb6\x73\x03\x12\x2b\x12\x31\x08\xb0\xc0\x52\xb2\xea\xd1\x77\xbf\x01\x40\x52\x94\x6c\xa7\xb9\x4b\xf3\x87\x42\x00\xbb\xbf\x7d\x81\xfb\xa0\xb3\x8f\xbe\xf8\xee\xd5\x9b\x9f\xbf\xff\x12\x6a\x6a\xe4\xf2\x49\xe6\xfe\x03\xc9\x54\x95\x47\xa8\xa2\xe5\x93\x93\xac\x46\xc6\x97\x4f\x4e\x4e\xb2\x06\x89\x41\x59\x33\x63\x91\xf2\xa8\xa3\x55\xfc\x59\xb4\x3f\xa8\x89\xda\x18\x7f\xef\xc4\x3a\x8f\xfe\x1d\xff\xf4\x32\x7e\xa5\x9b\x96\x91\x28\x24\x46\x50\x6a\x45\xa8\x28\x8f\xbe\xfe\x32\x47\x5e\xe1\x84\x4f\xb1\x06\xf3\x68\x2d\x70\xd3\x6a\x43\x13\xd2\x8d\xe0\x54\xe7\x1c\xd7\xa2\xc4\xd8\x2f\xce\x41\x28\x41\x82\xc9\xd8\x96\x4c\x62\x7e\x19\x2d\x9f\x38\x1c\x12\x24\x71\x79\x77\x97\x7c\x8b\xb4\xd1\xe6\x66\xb7\x5b\xc0\x6b\x41\x5f\x75\x05\xfc\x9d\x75\x25\x52\x96\x06\x12\x4f\x2d\x85\xba\x81\xda\xe0\x2a\x8f\x9c\xce\x76\x91\xa6\x25\x57\x6f\x6d\x52\x4a\xdd\xf1\x95\x64\x06\x93\x52\x37\x29\x7b\xcb\x6e\x53\x29\x0a\x9b\xd2\x46\x10\xa1\x89\x0b\xad\xc9\x92\x61\x6d\x7a\x95\x5c\x25\x9f\xa6\xa5\xb5\xe9\xb8\x97\x34\x42\x25\xa5\xb5\x11\x18\x94\x79\x64\x69\x2b\xd1\xd6\x88\x14\x41\xba\xfc\xff\xe4\xae\xb4\xa2\x98\x6d\xd0\xea\x06\xd3\xe7\xc9\xa7\xc9\x85\x17\x39\xdd\x7e\xb7\x54\x27\xd6\x96\x46\xb4\x04\xd6\x94\xef\x2d\xf7\xed\xef\x1d\x9a\x6d\x7a\x95\x5c\x26\x97\xfd\xc2\xcb\x79\x6b\xa3\x65\x96\x06\xc0\xe5\x07\x61\xc7\x4a\xd3\x36\x7d\x96\x3c\x4f\x2e\xd3\x96\x95\x37\xac\x42\x3e\x48\x72\x47\xc9\xb0\xf9\x97\xc9\x7d\x2c\x86\x6f\x8f\x43\xf8\x57\x08\x6b\x74\x83\x8a\x92\xb7\x36\x7d\x96\x5c\x7e\x96\x5c\x0c\x1b\xf7\xf1\xbd\x00\x17\x34\x27\xea\x24\x59\xa3\x21\x51\x32\x19\x97\xa8\x08\x0d\xdc\xb9\xdd\x93\x46\xa8\xb8\x46\x51\xd5\xb4\x80\xcb\x8b\x8b\x8f\xaf\x1f\xda\x5d\xd7\x61\x9b\x0b\xdb\x4a\xb6\x5d\xc0\x4a\xe2\x6d\xd8\x62\x52\x54\x2a\x16\x84\x8d\x5d\x40\x40\xf6\x07\x3b\x2f\xb3\x35\xba\x32\x68\x6d\x2f\xac\xd5\x56\x90\xd0\x6a\xe1\x6e\x14\x23\xb1\xc6\x87\x68\x6d\xcb\xd4\x3d\x06\x56\x58\x2d\x3b\xc2\x23\x45\x0a\xa9\xcb\x9b\xb0\xe7\x5f\xe3\xa9\x11\xa5\x96\xda\x2c\x60\x53\x8b\x9e\x0d\xbc\x20\x68\x0d\xf6\xf0\xd0\x32\xce\x85\xaa\x16\xf0\xa2\xed\xed\x81\x86\x99\x4a\xa8\x05\x5c\xec\x59\xb2\x74\x70\x63\x96\x86\x8c\xf5\xe4\x24\x2b\x34\xdf\xfa\x18\x72\xb1\x86\x52\x32\x6b\xf3\xe8\xc8\xc5\x3e\x13\x1d\x10\xb8\x04\xc4\x84\x1a\x8e\x0e\xce\x8c\xde\x44\xe0\x05\xe5\x51\x50\x22\x2e\x34\x91\x6e\x16\x70\xe9\xd4\xeb\x59\x8e\xf0\x64\x2c\xab\xf8\xf2\xd9\x70\x78\x92\xd5\x97\x03\x08\xe1\x2d\xc5\x3e\x3e\x63\x64\xa2\x65\x26\x06\xde\x15\x83\x15\x8b\x0b\x46\x75\x04\xcc\x08\x16\xd7\x82\x73\x54\x79\x44\xa6\x43\x77\x8f\xc4\x12\xa6\x79\x6f\x48\x7b\x2f\x3b\xaa\x51\x39\x3b\x09\x79\x9f\x04\xe1\x18\xb6\x12\x54\x77\x45\xcc\x24\x3d\x0a\x9e\xa5\xf5\xe5\x60\x52\xca\xc5\xba\xf7\xc8\xe4\xf1\xc8\x39\x8f\xdb\xff\x19\xf4\x0f\x7a\xb5\xb2\x48\xf1\xc4\x1d\x13\x62\xa1\xda\x8e\xe2\xca\xe8\xae\x1d\xcf\x4f\x32\xbf\x0b\x82\xe7\x51\x25\x2c\x45\x40\xdb\xb6\xf7\x5d\x34\x9a\xa4\x4d\x13\xbb\xd0\x19\x2d\x23\x68\x25\x2b\xb1\xd6\x92\xa3\xc9\xa3\xde\x27\xaf\x85\x25\xf8\xe9\x87\x7f\x42\x1f\x60\xa1\x2a\xd8\xea\xce\xc0\x37\x9a\x95\xdf\x6a\x8e\xc0\x38\x77\x97\x3b\x49\x92\x89\x6c\x7f\xd3\xef\x6b\x17\x17\xa4\xf6\x54\x27\x59\xd1\x11\xe9\x91\xb0\x20\x05\x05\xa9\x98\xe3\x8a\x75\x92\x80\x1b\xdd\x72\xbd\x51\x31\xe9\xaa\x72\x05\x31\x58\x10\x98\x22\xe0\x8c\x58\x7f\x94\x47\x03\xed\x10\x14\x66\x5b\xdd\x76\x6d\x1f\x96\xb0\x89\xb7\x2d\x53\x1c\xb9\x0b\xa5\xb4\x18\x2d\x5f\x8b\x35\x42\x83\xf0\xcd\x77\x2f\x5f\x9d\x1c\x07\xba\x64\x06\x29\x9e\x62\x3e\x10\xe8\xa0\x4b\xb0\x08\xfa\x7f\x59\x27\x07\xa4\xd1\x82\x06\x55\x07\x07\xab\xd8\xb8\x24\x14\x2d\xef\xee\x0c\x53\x15\xc2\x53\xc1\x6f\xcf\xe1\x29\x6b\x74\xa7\x08\x16\x39\x24\x2f\xfd\xa3\xdd\xed\x0e\xd0\x01\x32\x29\x96\x19\x7b\xd7\xbb\x00\x5a\x95\x52\x94\x37\x79\x44\x02\x4d\x7e\x77\xe7\xc0\x77\xbb\x6b\xb8\xbb\x13\x2b\x78\x9a\xfc\x80\x25\x6b\xa9\xac\xd9\x6e\x57\x99\xe1\x39\xc1\x5b\x2c\x3b\xc2\xd9\xfc\xee\x0e\xa5\xc5\xdd\xce\x76\x45\x23\x68\x36\xb0\xbb\x7d\xc5\x77\x3b\xa7\x73\xaf\xe7\x6e\x07\xa9\x03\x55\x1c\x6f\xe1\x69\xf2\x3d\x1a\xa1\xb9\x85\x40\x9f\xa5\x6c\x99\xa5\x52\x2c\x7b\xbe\x43\x27\xa5\x9d\xdc\x5f\x97\xd4\xdd\x97\xf1\x66\xfb\x17\xc5\xab\x3a\xd5\xf4\x81\x7b\x5f\xc5\xa3\xf6\xfd\x75\xb0\x82\xf0\x06\xb7\x79\x74\x77\x37\xe5\xed\x4f\x4b\x26\x65\xc1\x9c\x5f\x82\x69\x23\xd3\x1f\xe8\xae\xe9\x5a\x58\xdf\x78\x2d\x07\x0d\xf6\x6a\xbf\xe7\x8b\x7c\x94\xe5\x48\xb7\x0b\xb8\x7a\x36\x49\x71\x0f\xbd\xe3\x2f\x8e\xde\xf1\xab\x07\x89\x5b\xa6\x50\x82\xff\x8d\x6d\xc3\xe4\xf0\xdc\xbf\x2c\x93\x77\xef\x98\x29\x76\x09\x7d\x54\x6d\x2c\x0c\x17\xd7\xa0\xd7\x68\x56\x52\x6f\x16\xc0\x3a\xd2\xd7\xd0\xb0\xdb\xb1\x38\x5e\x5d\x5c\x4c\xf5\x76\x0d\x23\x2b\x24\xfa\x7c\x62\xf0\xf7\x0e\x2d\xd9\x31\x8f\x84\x23\xff\xeb\xd2\x09\x47\x65\x91\x1f\x79\xc3\x49\x74\xae\xf5\x54\x93\xd0\x8f\xce\x7c\x50\xf7\x95\xd6\x63\xbd\x99\xaa\xd1\x43\x4f\x4a\x63\xb4\xcc\xc8\xec\xe9\x4e\x32\xe2\xff\x53\xbd\x30\xae\x1f\x7c\xac\x5c\x84\x84\xe6\x6c\x6f\x11\x4d\x68\x46\xdc\x95\x05\xbf\xcc\x52\xe2\x1f\x20\xd9\x5d\xc2\x82\x59\x7c\x1f\xf1\xbe\x2d\xd8\x8b\xf7\xcb\x0f\x95\x5f\x23\x33\x54\x20\x7b\xbc\xa2\x4d\x14\x58\x75\x8a\x4f\xec\x77\xa9\xf3\x43\xe5\x77\x4a\xac\xd1\x58\x41\xdb\xf7\x55\x00\xf9\x5e\x83\xb0\x3e\x54\x21\x4b\xc9\xbc\xfb\xaa\x4d\x17\x7f\xd1\xbb\xfd\x67\xed\xcb\xd5\xf2\x2b\xbd\x01\xae\xd1\x02\xd5\xc2\x82\x6b\x3e\x3e\xcf\xd2\xfa\x6a\x24\x69\x97\x6f\xdc\x81\xf3\x29\xac\x42\xfb\x21\x2c\x98\x4e\xf9\xb2\xab\x15\x50\x8d\x87\x9d\x8b\x0a\x4f\x09\xbc\xd1\xae\xfb\x5b\xa3\x22\x68\x98\x14\xa5\xd0\x9d\x05\x56\x92\x36\x16\x56\x46\x37\x80\xb7\x35\xeb\x2c\x39\x20\x97\x3c\xd8\x9a\x09\xe9\xdf\x24\x1f\x50\xd0\x06\x58\x59\x76\x4d\xe7\xba\x57\x55\x01\x2a\xdd\x55\x75\x50\x85\x34\x84\xaa\x24\xb5\xaa\x46\x75\x6c\xcb\x1a\x60\x44\xac\xbc\xb1\xe7\x30\xa4\x04\x60\x06\x81\x04\x72\xc7\xd5\xf7\x10\xac\x2c\x7d\x25\x4b\xe0\xa5\xda\x6a\x85\x50\xb3\xb5\xd7\xe3\x88\x00\x1a\xb6\x1d\x80\x7a\xb5\x36\x82\x6a\x11\xec\x6e\xd1\x34\x6e\x1a\xe1\x20\x45\x23\xc8\x26\x59\xda\x4e\x3d\xa7\x0f\x59\xcf\xc1\x8a\xa6\x95\x5b\x28\x0d\x32\x42\x60\x90\xb1\xa3\x41\xd2\xb5\x45\x49\xe8\xe7\xfc\x28\x12\x01\x31\x53\xb9\x31\xfd\x3f\xac\xd0\x1d\x2d\x0a\xc9\xd4\x8d\x6b\x13\xc6\x56\xc8\xd5\x34\xaf\xd4\xc3\x4d\x10\xb4\xcc\x3a\x0d\x85\x22\xed\x95\xee\xe7\x72\x0b\x33\xb7\x5a\x09\x89\x7e\x74\xf7\xb7\x40\x9d\x3a\x8b\xdd\x7c\x35\x3f\x87\x52\xb7\xdb\xc0\xed\xf9\x9c\x6a\xd6\xf7\x5d\x23\x14\x2b\xf4\x1a\x21\x34\x75\x85\xbe\x05\xa6\x38\xac\x84\x41\x60\x1b\xb6\xfd\x08\x7e\xd6\x1d\x94\x4c\x01\x19\x56\xde\x04\xd9\x9d\x31\xee\x3e\xb4\xa8\x5c\xc6\xdf\x87\xa8\x40\xa9\x37\x9e\x24\xa0\xad\x04\x4a\x1f\x2f\x8b\x08\xb5\xde\x40\xd3\x95\xde\x40\x17\x28\x74\x07\x1b\x26\x08\x3a\x45\x42\x06\xbb\xa9\x33\x0a\x4a\xdd\xe0\x41\x14\xee\x95\xec\x0c\x9b\xe5\x1b\x67\xf7\xbd\xbb\x3c\x16\x5b\x30\xf8\x2a\x90\x43\x6b\x34\x61\xe9\x86\x22\x60\x15\x13\xca\x3a\x3b\x7d\x9c\xb1\x79\x8f\x62\x3c\x3e\xf5\x0f\xfb\x29\xd4\x1f\xa7\x29\xbc\x96\xba\x60\x12\xd6\x2e\xc7\x14\xd2\xbd\x86\x1a\x5c\xbb\x7b\xe0\x2d\x4b\x8c\x3a\x0b\x7a\xe5\x77\x83\xe6\x8e\x7f\xcd\x8c\xbb\xed\xd8\xb4\x04\x79\x3f\x43\xb9\x3d\x8b\x66\xdd\x4f\x86\x6e\xe9\x1a\xae\x70\xde\x0b\xfd\x02\x57\x42\x85\xa0\xae\x3a\x15\xcc\xa3\x9a\x11\x84\x16\xc4\x02\xf3\xc1\x86\xce\x48\xe8\x23\x1d\x20\x47\x01\x9e\x0e\xf2\x91\x7d\x76\xcf\xcf\xfd\x43\xef\xa3\x79\x3f\x03\x06\x98\xc4\xa2\xe2\xb3\x7f\xfc\xf8\xdd\xb7\x89\x25\x23\x54\x25\x56\xdb\xd9\x5d\x67\xe4\x02\x9e\xce\xa2\xbf\xf9\xd1\x60\xfe\xcb\xc5\x6f\xc9\x9a\xc9\x0e\xcf\xbd\x01\x0b\xff\x7b\x4f\xcc\x39\xf4\x8f\x0b\x38\x94\xb8\x9b\xcf\xaf\x1f\xee\xd7\x26\xed\xa5\x41\x8b\x34\x73\x84\x63\x24\x77\xd7\x87\x4e\x62\xd0\x20\xd5\xda\xdf\x45\x83\xa5\x56\x0a\x4b\x82\xae\xd5\xaa\xf7\x09\x48\x6d\xed\xe0\x98\x3d\xc5\xc4\x37\x83\xf1\x62\x05\xb3\x21\x5c\x1f\xc3\x33\xc8\x73\xb8\x18\xce\x7a\xcf\x40\x0e\x0a\x37\xf0\x2f\x2c\x7e\xd4\xe5\x0d\xd2\x2c\xda\x58\x97\x16\x22\x38\x03\xa9\x4b\xe6\xf0\x92\x5a\x5b\x82\x33\x88\x52\xd6\x8a\x68\x1e\x26\xe9\x1d\xb8\xfe\xf8\xcf\xc1\xde\x0b\x2b\x7c\x6b\x08\x9a\x9e\x9d\x85\x6b\x33\x84\x4e\xab\x06\xad\x65\x15\x4e\x2d\xf4\x49\x7e\x34\xc5\x39\xa2\xb1\x15\xe4\xe0\x43\xdc\x32\x63\x31\x90\x24\xae\xad\xe8\xa5\x78\x77\x78\xb2\x3c\x07\xd5\x49\x39\xf2\x9f\x18\x74\x2f\x73\x4f\xb6\x7b\x72\x40\x9e\x84\x24\xfc\x51\x9e\x83\x2b\xb2\x2e\x46\x7c\xcf\xe9\xae\x4f\xe8\x06\xe6\x89\xab\xf3\x7b\x8e\xf9\x08\x77\x0f\x0d\xf9\x9f\xc1\x21\x3f\xc6\x43\xfe\x08\xa0\x6f\xbe\xde\x85\x17\x9a\xb5\x09\x9c\xdf\x78\x04\x4d\x75\x4d\x81\xe6\x5d\x70\xa1\xf9\xea\xe1\xbc\xab\xbf\x56\x34\xe1\x3d\x87\xcb\x17\xf3\x47\xd0\xd1\x18\xfd\x28\xb8\xd2\xb4\x9d\xdd\x49\xb6\x75\x55\x07\x4e\x49\xb7\xaf\x7c\xb3\x74\x7a\x0e\x4e\xd6\x02\x46\x84\x73\x3f\x04\x2f\xe0\xd4\xaf\x4e\x77\x8f\x48\xb3\x5d\x59\xba\x7a\xf4\x21\xf2\x7a\x8c\x51\x62\xbf\x7e\x54\xe6\x58\x5f\x0e\x84\xc2\x27\x9f\xc0\xbd\xd3\xc3\x2b\xe8\xee\x70\x5f\x28\x21\x87\x28\xea\xe1\x4f\x56\xda\xc0\xcc\x1d\x8a\xfc\xe2\x1a\x44\x36\x85\x49\x24\xaa\x8a\xea\x6b\x10\x67\x67\x7b\xa4\x93\x01\xe6\x2c\x87\xc8\x8d\x03\x19\xf1\xa5\xef\xcb\x42\xf3\xf6\x6b\xe4\xc6\xbf\xca\xe8\x4e\xf1\x85\x4b\xb9\xb3\xd3\x7d\x33\x30\xe9\x03\xce\x0e\x54\xfe\x45\xfc\x96\x74\x16\x8d\xaf\xdc\x67\x10\x25\xad\xaa\x3e\xf7\x43\xe3\x8b\xe7\xa7\xf3\x6b\xd8\x63\xfa\x51\x72\x01\xa5\x1b\xac\xae\x21\x0c\x27\xbe\x47\x84\x71\xac\xf2\xab\x42\x1b\x8e\x26\x36\x8c\x8b\xce\x2e\xe0\x79\x7b\x7b\xfd\xeb\x30\x76\xfa\x4e\xd6\xeb\xdd\x1a\x5c\x3e\xa4\xcb\xd0\x2e\x9d\x41\x94\xa5\x8e\x68\x60\x19\xad\x9c\x7e\x31\x84\x07\x7a\x70\x18\xbf\xe7\xf5\xfb\x8d\xe0\x5c\xa2\x53\xc2\x0b\x0c\x1f\x5e\x79\x67\x7c\xe2\x9a\x85\xf5\xec\x58\x0f\x12\x0d\xce\x93\x4e\x89\xdb\xd9\x3c\xee\x69\x86\xf5\x39\x9c\x5a\x97\x9f\xb9\x3d\x9d\x27\x75\xd7\x30\x25\xfe\xc0\x99\x6b\xe8\xe7\x41\x6f\xa7\xb1\xeb\xd2\xc7\x68\xef\x26\x2f\xda\x38\x60\xce\x93\x9a\x1a\x39\x8b\x32\xf2\x5f\x25\x9d\x72\x63\x88\x3d\x4a\xd8\x3e\xbc\x91\xbb\xc3\x1c\x5a\x4a\x6d\xf1\xa8\x46\x80\x45\x7a\x23\x1a\xd4\x1d\xcd\xc6\x3a\x72\xee\x86\xde\x8b\xf9\x35\xec\xf6\x1f\x6f\xd3\x14\xbe\xb4\x6e\x8e\x10\xb6\x06\x06\x1b\x2c\xac\xcf\xef\xd0\xf3\xf8\x72\x1e\xca\xf6\xcb\xef\xbf\x9e\x94\xee\x11\x75\xe6\x95\x1b\x3f\x5e\x3f\x54\x27\x1f\xfc\x5a\xbe\xd9\x6c\x92\x4a\xeb\x4a\x86\xef\xe4\x63\x21\x75\xd5\x23\x79\xeb\x66\x55\xbb\x55\x25\x70\x5c\xa1\x59\x4e\xe0\xfb\xea\x9a\xa5\xe1\x3b\x6e\x96\x86\xbf\x51\xfd\x37\x00\x00\xff\xff\x99\x81\x6f\xb0\xb4\x1a\x00\x00")

func faucetHtmlBytes() ([]byte, error) {
	return bindataRead(
		_faucetHtml,
		"faucet.html",
	)
}

func faucetHtml() (*asset, error) {
	bytes, err := faucetHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "faucet.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"faucet.html": faucetHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"faucet.html": &bintree{faucetHtml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

