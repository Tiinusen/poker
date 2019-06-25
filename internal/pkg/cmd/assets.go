// Package cmd Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// ../../../assets/certificates/poker.crt
// ../../../assets/certificates/poker.key
// ../../../assets/certificates/root.poker.crt
package cmd

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _pokerCrt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x95\xb9\x0e\xb3\x4a\x12\x46\x73\x9e\x62\x72\x34\x32\x06\xcc\x12\x4c\xd0\x4d\x43\xbb\xcd\x66\x30\x7b\x66\xc0\xb4\x17\x30\x66\x31\x0d\x7e\xfa\xd1\xfd\xef\x64\x77\x2a\x3c\x2a\xa9\xa4\xaa\xa3\xfa\xfe\xfd\x57\x41\x13\x13\xef\x5f\x86\x19\x46\xc4\x22\x06\x88\xcc\x3f\x94\x73\x09\x31\xf3\x9f\x61\x80\xd7\x97\x02\x46\x20\xa0\x24\x88\x4f\xd5\x9c\xb4\x06\x0a\x70\xda\x57\x82\xf1\xe6\x73\x7e\x5e\x9e\xc0\x83\xf4\x35\xdc\x5f\x0f\xac\x33\x01\x82\x60\xb2\x00\x02\x09\xe7\x86\x2e\x33\x83\x1c\x25\x41\x80\x4c\xf6\xd9\x4a\x51\x17\x9c\x37\x5c\xae\x62\xb2\xb9\x50\xce\x50\x64\x1e\x5c\x94\x6f\x5e\x64\xca\x1e\xca\x99\x9f\xf4\x19\x8a\x08\x73\x23\xf2\x87\x71\x1e\x22\xab\x6b\xf5\xcc\x44\xa6\xef\x02\x86\xc1\x3e\x36\xc1\xea\x5a\x15\xd6\xc7\x22\x23\x8c\x52\xf3\xe1\x02\x01\x1b\x97\x01\x5f\x48\x29\xa1\xc0\x84\x20\x88\x01\x90\x09\x44\x0c\x30\x8e\x52\xd3\x06\x3d\x81\x20\x40\x00\x59\xcb\xc1\x53\x6c\xeb\x10\xe9\x03\x12\x3c\x79\xbb\x6e\xe3\x6f\x7f\xbb\xc8\xdf\x8e\x06\xc7\x70\x7c\x9c\xc6\x8c\x88\xeb\xeb\xbb\x15\x77\x7d\x69\x7f\xbb\x9f\x06\xb8\x42\x04\xe5\x8a\xc2\x4a\x26\xf5\x8b\x30\xf9\x53\x7e\x26\x7a\x47\xe5\x32\x16\x16\x79\x19\xbb\x62\x13\x06\xba\x3b\x26\xe1\xf2\x7d\x59\x67\x1c\x30\x19\x62\x52\x56\xf7\x24\x6a\x79\x2c\xeb\x77\x95\x43\xeb\x96\xd7\x65\xee\x1e\xe2\x03\x85\x2f\xbd\x3e\x93\xcb\x24\xb8\xcc\xae\x76\x60\xa9\x54\xec\xc6\x8d\x2d\x7b\x77\x28\x29\x71\x05\xea\x22\xde\x94\xef\x69\x63\x49\xd9\x3e\x87\x60\x3b\x3f\x1f\x29\x77\x44\x7b\x13\xdb\xf9\x64\x8c\xce\x07\x3e\xd4\x8f\xaa\x0c\x3e\x8e\xfd\x87\x17\xb6\xc6\xe8\x23\x39\x30\xd1\xa3\xe9\x70\x81\xdb\x3a\xee\xd4\xa1\x5a\xdd\xa9\xe8\xb3\x3e\xd9\xea\x1c\xd5\xd5\xe2\x72\x27\xa5\xcc\x0e\xfb\xca\x0b\xf2\x71\xea\xe6\xbd\x54\xd4\x69\x72\x7b\x8b\x49\xc1\xac\x50\x29\x7c\xb1\xe3\x2f\x1e\xa3\xb8\xcb\x4a\xdf\xc7\x15\xca\x96\x48\xa2\xf7\xbb\x1e\xf4\xc8\x5c\x9a\xc7\xba\x71\x9f\xdc\xfe\x09\x0f\x5a\x75\xc7\x82\x2d\xa7\x58\xda\x6e\xe3\x08\xf9\x5f\x44\xab\x39\x95\x9c\xfb\x0e\x50\x17\x02\x80\x9f\xb4\x74\x19\x2d\x01\x43\x34\x47\x49\x28\x9c\x41\x70\xdc\xfd\xb5\x79\x0e\x50\x5f\x76\xa1\xf0\xd7\xf9\xea\x53\x10\xa4\x2e\x0c\xb0\x61\x4c\x18\x04\xb1\x05\x99\x0b\x21\xa5\x23\xa4\xa6\x05\x83\x0a\x81\x27\xa8\x21\xf5\x92\x63\x20\x9b\x16\x0d\xe2\x0a\x9d\x2e\x9c\xd6\x33\x1f\x3a\x09\x4f\xb2\xe8\xa2\x90\x25\xf3\xce\xef\x73\x48\x5d\x76\x64\x7f\x26\x3d\x21\xa4\xcc\xea\x41\x5c\x65\x8f\xa9\x88\x93\x41\xf8\x1c\xae\xba\xd5\xb9\xaf\x5a\x56\x9f\x45\x52\x72\x32\xed\xd9\xf9\xef\xe6\x10\x22\xca\xbc\x9e\xfc\x4f\x22\xdb\x40\xd9\xdf\x46\x3a\x19\xdc\x4a\xe9\x7e\x70\x22\xfc\x0f\xc6\x39\x91\x6d\x18\xe9\xba\xe4\xa2\x35\x5d\xb1\xfe\xab\x49\x65\x36\x0c\x00\x10\xfd\x1f\xf7\x81\x6f\x00\x6a\x02\x61\xde\xd0\x77\x48\x73\xec\x6a\x33\x77\xad\xca\x9c\x8f\x81\x53\xeb\xee\x41\x3b\x13\xfc\x1d\xf4\xf9\xbe\xc5\xf0\x79\x9b\x3a\x2b\x6c\x90\x28\xee\x01\x51\x5f\x1f\xe7\x52\xa4\x71\x63\x39\xa3\x66\x1f\xad\x2e\x0c\x29\x91\xa4\x72\x69\x0f\xdc\xfb\x26\x54\x5a\xb4\x5a\xe4\xe0\xd8\x37\x7b\xa2\xcb\x8e\x4f\x78\x33\xd5\x9e\xb2\xf2\x1c\x1a\x9e\x0a\xe3\x45\x29\x9a\xd9\x3f\x8a\xa5\xd9\xc5\x67\xf1\x74\x94\x44\xf9\xe0\x79\x3f\xde\xf7\x02\xd1\x6b\xb9\x33\x1e\xb4\xc7\x3a\xb6\xae\x22\xcd\xe0\x49\x22\x4d\x14\x0c\x4f\xda\x94\x64\x9a\x08\x30\x76\x1f\xa1\xb9\x2e\xaf\x8f\xbf\x69\xce\xb6\x2f\xb2\x6e\xf6\xda\x9b\x65\x3d\xde\x1a\x91\x52\xb5\x1f\x86\x2b\xd7\x09\x3b\xde\x51\x50\x72\xa3\xb2\x57\x1e\x9b\x3d\x31\xbf\xd7\xf4\x75\xb2\x88\x46\x2e\x66\x7c\x06\x78\xf7\x52\x7f\x3b\x83\x3f\x28\xed\x63\x35\x59\xfa\xde\x6a\x5d\x17\x42\xfb\xf8\x6a\x75\x92\xca\x47\x4e\x0f\x6d\x45\xaf\x6b\x4f\x9b\x16\x3d\x8a\x9a\x5b\xf6\x10\xe2\x29\x11\x7d\xe3\xe4\x65\xda\xba\xa6\x27\xb5\xb1\xf4\xc5\xcf\x95\x6f\xcd\xfb\x7d\xaa\xe4\xe7\x5f\x15\x9e\x74\x2a\x41\x22\xd1\x02\x0a\x9c\xd6\xe9\xfd\xad\x79\xae\x2e\x9c\x59\x2d\x2a\x36\x35\x7c\xf7\xfe\xc4\x71\x91\xab\xf4\xab\xd8\x79\xe0\xef\x95\x41\x12\xbc\xc4\x17\x7f\xee\xbd\x66\xfb\xfb\xfa\xbd\x39\x5f\x5f\x48\x97\xfb\xcd\x6b\x2b\xee\x7b\x7b\xe1\x1f\xb9\x9a\x6c\xc8\x9c\x06\x17\xdd\xcb\xfc\x58\x86\x64\x5b\xd3\x3b\x98\xdf\x45\x63\x67\x42\xff\x75\xe8\x3b\x60\xbe\xec\x80\xf4\xac\x88\xa2\x1e\x6d\x9d\xf4\x7d\xda\xa1\x73\xf6\xf0\xc4\x21\xef\x42\xe1\x12\xeb\x16\x8f\x92\x65\xb8\x0a\x9a\x31\x57\xa5\xf6\x2e\x9d\xfa\x20\xcf\xae\x7c\x40\x91\x5a\x97\x6d\xee\xca\x0f\x75\x06\xdd\x5c\x7f\x48\xd0\xf6\xbf\xd1\x08\x9c\x42\x18\xf8\x8b\xce\xad\x78\xfb\x96\xe9\x38\xce\x53\x24\x1e\x93\x86\xa9\x04\x96\xd2\xfb\x01\xe5\x60\xad\x51\x92\x9c\xf6\x26\x0e\x19\xd6\x72\x53\xa5\x61\x73\x63\xf1\x7e\x97\x0e\xbe\x0b\xaf\xcd\x20\xd3\x36\xb8\x78\x1c\x3e\x48\xfe\xf2\x81\xca\xdd\xea\xf9\x73\xff\x43\x13\xe9\x5a\xeb\x5e\x64\xba\x8a\x3e\xb7\xd7\xe7\x48\x9b\xc4\x7f\x3b\x7e\xdc\x56\x4a\x93\x5c\x4f\x09\xef\x15\x86\x51\x2f\x85\xbe\x4c\xfa\xa2\x44\x94\xab\x7e\xc2\xed\x7d\x73\xdd\xbd\x3a\xda\xd5\xd2\x69\xfb\x5f\xaf\xd2\xab\x1c\xf4\x13\xb1\xde\x27\xf0\x1f\xee\xcf\xf7\x37\x3d\xf4\xcf\x44\xf8\x6f\x00\x00\x00\xff\xff\x71\x46\xa9\x01\x2e\x06\x00\x00")

func pokerCrtBytes() ([]byte, error) {
	return bindataRead(
		_pokerCrt,
		"poker.crt",
	)
}

func pokerCrt() (*asset, error) {
	bytes, err := pokerCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "poker.crt", size: 1582, mode: os.FileMode(436), modTime: time.Unix(1561488413, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pokerKey = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xd5\xb5\x12\xa4\xea\x02\x04\xe0\x9c\xa7\xd8\x9c\xba\x85\x0c\x32\x13\xfe\xb8\xbb\x67\xc0\xe0\xee\x30\x4f\x7f\x6b\x37\x3e\x9d\x76\xd2\xd5\xc9\xf7\xbf\xbf\x61\x78\x51\x36\xfe\x38\x2e\xf8\x63\x39\x72\x00\x3c\xfe\x8f\xca\xc7\xff\x1a\x48\x97\x65\x7e\xba\x64\x06\x00\x95\x05\x36\x0f\x2e\x70\x67\xb0\x57\x34\x75\xe1\x23\x69\xf5\xe1\x0a\x7d\xd8\x96\xf7\x27\xea\x0e\x6d\x9e\x78\x06\x4d\x89\x26\xc5\x1e\x63\x4b\xb4\x6d\x88\xcb\x8c\x84\x72\xe4\x8d\x00\xf1\x5b\x89\x6f\x1b\x8d\x4c\x56\x4a\x58\xdd\x54\x43\x2a\x93\x65\x1b\xa7\x29\xdc\x71\x15\x70\xe2\xf9\xae\x4e\xd6\x8d\xf9\x19\x3d\x33\xde\xdd\xe9\x26\x70\x1a\xf1\x25\x07\x3e\x59\x42\xf5\x61\xc5\xc5\xf5\xce\x07\x29\xc4\x39\x33\x30\x63\x90\x58\x91\xf7\xf0\xab\xc1\xe9\xec\x68\x5d\xda\xcb\xbc\x5b\x41\x7a\x0e\x2e\xb6\xbf\x70\x2f\x01\x29\x0c\xf4\x63\x6d\xf2\x4d\x08\xc9\x98\xea\x20\xbd\x25\x88\xbe\xbe\x3e\x0e\xd3\x0c\x1a\x58\x1e\xca\x8e\x0f\xaa\x80\x97\xb6\xee\xb9\xa9\xf5\x13\x7b\xf9\x55\x30\xcf\x00\x62\x24\xeb\xc1\x99\xa3\x40\x39\xa8\x51\xf7\x34\x31\x15\x28\x21\x1f\x43\x08\xbc\x22\xed\xc7\x16\x03\x06\x7f\x23\xce\x17\x55\x5a\x59\xbf\x5f\x3c\x0a\xfa\x68\xfe\xf4\x61\xce\xf8\xc5\xe0\xed\xf3\xd9\xb5\xb9\xcc\xcc\x3d\xfe\x6b\xeb\x11\x60\x34\xfa\x21\xe4\xb8\xf4\x55\x70\x43\xda\x8b\xd8\xf2\x25\x64\xb7\x8f\x3c\x49\x4a\x8d\xe7\xda\x13\x18\xef\x71\xa1\xae\x72\x43\x09\x49\x0b\xf6\x87\x28\x2f\x99\x03\x36\x60\xc0\x24\x33\xc0\x66\x6f\x60\xe9\x93\x2c\x8d\xe5\xae\x4e\x39\x44\x90\xaa\x12\xb6\x70\x18\xed\x5e\x20\x17\xee\x0b\x4e\x80\xb7\x88\xfe\xc0\xcf\xbe\x71\xf7\x57\x06\x94\x7e\x5d\x4e\xf2\xa7\x7d\xed\x95\x40\xf7\x75\xd9\x07\xb8\x1f\x9f\x13\x71\xef\x81\x36\x7c\x06\x0a\x0c\x4f\x59\x5f\x7a\x4a\xb6\x94\x17\x21\x27\x96\x1e\x2a\x2a\xd2\xf6\x65\xc3\x2e\x8f\xe8\x72\xb1\xe7\x77\xb9\x7d\x14\x6d\x0a\x35\x7c\xcf\xdf\x2d\x96\xb4\xc2\xbb\x59\xfb\xcb\xad\xe9\xf4\x37\xf9\x10\xc7\x07\x6f\x67\x53\xc5\xde\xf4\xb9\x21\x6a\x43\x8b\x65\xca\xcd\x42\xea\xc8\x11\x5e\x06\x56\x31\x52\x12\x8b\x1d\x7f\x37\x32\x21\x1a\x40\xe8\xc4\xf4\x6d\xfe\x32\x9c\xeb\xba\x17\x7b\x9e\xe1\x6d\x43\x71\x2b\x96\xe6\xc7\x37\xe1\xb2\x4d\x1d\x9c\x43\x4e\x75\xdb\x6c\x5a\xdc\x6d\x37\x52\x88\x8d\xe1\x7f\xf0\x2b\xfb\x9a\x34\x3d\xb4\xbf\xe3\xa8\xf2\x4b\xef\x53\xe3\xf0\x25\xcb\x5c\xd8\xf9\x37\x05\x07\x64\xbc\x85\x8f\x94\xd1\x1e\x39\xbc\x58\x9c\xda\x68\x3f\x25\xd8\x8a\xf7\xbc\xbd\x67\x23\x41\x55\x9b\xd2\xfa\x2a\x35\x7f\x8b\xd9\xe5\xcd\x1d\x39\x58\xa8\x1e\x55\x62\x75\x23\x7b\xb7\xad\xe5\x5a\x07\x10\xac\x3a\x96\xd9\x11\x0c\x98\x44\x06\x98\xd9\xfb\x5e\x95\x7e\xa5\xee\x91\x67\x0a\x1f\x46\x62\xe7\x19\x18\xf5\x73\xc9\x0a\x57\xe4\x4c\x63\xb0\x06\x0c\x0e\xfa\xd2\x5f\xc4\xaf\x36\x7b\xeb\x14\x53\xe8\x64\xd3\x60\xcf\x5c\x18\x49\x37\x31\x93\xf3\x5c\xed\xed\x2b\x0f\xd6\xbd\x5f\x14\x4f\x5c\xea\x09\xaf\x95\xaa\x20\x1a\x63\x7c\x0b\x82\xfd\x0b\x0a\x8c\xb9\x22\x84\xcd\x72\x01\x5f\xb9\x3d\xea\x76\xe8\x60\x0d\x8e\x10\x4f\x35\xb4\xf7\x74\xb6\x2c\x74\x27\xa3\x4d\xe1\xa5\x20\x69\x53\x39\x59\x3d\xb2\x15\xee\xd7\x96\x10\x98\x5d\x97\x3a\xd3\xb6\xda\x0d\x2e\x4d\x3c\x77\xe4\xdf\x62\xc3\x8b\x21\xe5\x5b\x3f\x9c\x94\x91\x41\x6f\x78\x4f\x3a\x2f\xdf\xbc\x9e\xaf\xe0\x36\x70\x1f\xef\x03\x4b\x43\xdc\xa9\x8d\x00\xe1\x29\x9a\x75\x19\xcf\xf2\x70\xe7\xde\xd6\xda\xcc\x9c\xe6\x5a\x82\x94\x0b\x6b\xa8\x53\x5f\x91\x1e\x7c\xa2\x41\x2e\x32\x6b\x4b\x2a\x16\xf5\xe3\x54\xe8\x6f\x61\x57\x8a\x36\xe1\x75\x33\x71\x6f\x77\x4a\x47\x95\x7a\x42\x4f\xaa\x0b\xf4\x10\xbf\x5b\x26\x67\xc6\x8f\x1e\xaf\xc9\x4f\x20\x09\x46\x66\xa9\xd6\x8e\x77\xfd\x99\xbc\x6d\xce\x0a\x86\x77\xa6\x87\x24\x50\x04\x9b\x55\x85\x84\x8d\x56\x4e\x50\x5d\xff\x7b\x32\xd0\x31\x39\xf7\x88\x87\xb2\x45\xcd\xd9\x5f\x47\xd2\x51\x78\x08\x99\x7d\x6b\x74\x2b\xfa\xd8\x13\x91\x1d\xba\xda\x38\x27\x9f\xd9\x62\x85\x86\x5e\x6e\x3f\xf5\x3e\x71\x62\x9a\xd3\x29\x8b\x1d\x69\x06\x1f\x88\xbc\xf3\x91\x28\xd0\x25\x8d\xdd\x25\xef\xd2\xad\x2c\x40\x4b\x3a\xc4\xaa\x31\x86\xdd\xb7\x7f\x8e\xee\xdc\xf2\x25\xf8\xf4\x14\x8b\xd2\xbf\x05\x2f\xba\x4d\x52\x89\x74\x33\xdd\x4c\x52\x91\xa2\x79\x34\x7f\x6e\xac\x5f\x5b\x86\xbc\xd1\x64\x31\x35\x51\x3c\x0b\x25\xa4\xaf\x17\xea\x46\x0d\x0e\xf7\x8c\x96\x54\xb3\x00\xa1\x87\xec\xcd\x56\x31\x33\xad\x1e\x17\xcd\xee\x37\xd5\x68\xe4\x87\xb9\x31\xcc\xca\xa6\x3a\x22\xe0\x35\x32\xe1\x87\xb2\x7f\x8c\xed\xce\x08\x44\x22\x99\xcf\x71\xf5\x93\x67\x23\x9e\x80\x66\xe7\xcf\x31\xed\xf4\x9e\xa0\xd0\xc6\xa7\x69\x03\x3e\x5e\xa2\xb1\x78\x48\x36\xf8\xef\x56\xdf\x34\x42\xf1\x4f\xbf\x71\x9a\xd1\xc0\x72\xca\x4c\x90\x15\xa2\x3f\xc2\xa1\x65\xda\x07\x7e\x38\x6a\x2e\x90\xa6\xa2\x55\x6d\x94\xbc\x4d\x02\x1c\x52\xf5\x34\xd3\x13\xd8\x8d\x09\x80\xa9\x78\x3b\x9a\x9e\x45\xac\x46\x34\xc8\xf7\x2e\xd2\xcc\x43\x97\xee\x8a\x82\x9c\x24\xc2\x80\xad\x32\x15\xc7\x0d\x58\x6f\x7d\xcf\x2f\xc0\x9c\x6b\x4c\xd6\x0e\xab\x63\xf8\xe4\x91\x0a\x84\x34\x46\x99\x49\x89\x54\x59\xd7\x6b\xf8\x9c\x6f\x67\x6f\x73\x05\x7c\x65\xbf\xd1\x7c\x41\xde\xee\x54\x72\x1f\xde\x34\xac\x31\x0c\xfd\x51\x43\xc4\xa2\x60\x7c\x04\x85\x31\x68\xb5\x87\x7b\x85\x6b\xe4\x70\x49\x21\x19\xaa\x1b\x8e\x50\xd5\x3c\x49\xe4\x6f\xb2\xf0\x8b\x31\x76\x2e\x99\x14\xa1\x8f\xd2\xd3\xe6\x17\x7b\xd0\x21\x6d\xb4\xd0\xb0\x9e\x32\xc8\xd3\xab\x1a\x45\xdf\x15\xba\x74\xb7\x3a\xa4\xbe\xd7\x48\xf6\x12\xf0\x5d\x26\xe2\x8c\x8e\x2d\x0d\xa1\x7f\xa4\xf0\x06\xf7\xdf\xd4\xfc\x3f\x00\x00\xff\xff\x8c\x44\x41\x5d\x8b\x06\x00\x00")

func pokerKeyBytes() ([]byte, error) {
	return bindataRead(
		_pokerKey,
		"poker.key",
	)
}

func pokerKey() (*asset, error) {
	bytes, err := pokerKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "poker.key", size: 1675, mode: os.FileMode(436), modTime: time.Unix(1561488413, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rootPokerCrt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x95\xb9\xee\x83\x4e\x96\x85\x73\x9e\x62\x72\x34\xb2\x8d\x59\x83\x09\x0a\xaa\x30\x50\x2c\x66\xb1\x31\x64\x66\x07\xb3\xef\xf0\xf4\xa3\xdf\xbf\x93\x96\xba\x6f\x78\xa4\x7b\x74\x92\x4f\xdf\xff\xfe\x9d\x88\x1e\xaa\xf9\x3f\x12\x72\x3c\x55\x56\x25\xe0\xa1\x7f\x52\xc2\x50\x55\xc4\x56\x92\x04\x66\x9c\x83\x4d\x15\x41\xae\x8a\xc0\x03\xa6\x98\xff\x86\xe2\x57\x3e\x84\xed\x2a\x02\x7b\x92\x01\x04\x6f\xc3\x31\x36\x64\x07\xf0\x6d\xdb\x10\x6d\xfd\x11\x51\xc2\x95\xd0\x5b\x71\xfd\x52\xef\xc3\x10\xe9\x0f\xf4\x10\x63\xc0\xe0\x30\x3d\x44\x9b\x50\xdd\x8d\x77\xf7\x81\x9e\xba\x19\x9e\xfa\xef\xd9\x26\x7b\xc8\x33\x44\xf4\x00\xb7\x17\x22\xc0\x6e\xe0\xb8\x11\xd6\x44\x62\xb6\x88\x9a\xeb\xf8\x6f\x4d\xa9\x6e\xd0\x0e\x34\xdc\x85\x6a\xb1\xc6\x26\xb0\x91\x28\xda\x00\xe6\xb9\xfa\x04\x50\x92\x40\xde\x49\xf9\xdf\x52\xeb\x03\x09\x9f\x77\x24\x8b\xd1\xe5\x72\x6e\x4a\xb7\x7c\xa7\x41\x4a\x91\xb7\xc0\x56\x99\x8c\xba\xb3\x5f\xb1\x9c\x0a\xd2\xd0\x87\xfb\xd9\xe2\x49\xb5\xb1\x25\x44\x0e\x38\x45\xe1\xb6\x6a\xc9\x60\xbe\x94\xc7\x49\xe4\xe1\x38\x25\xd3\x08\xc0\xab\x8c\x66\xee\xde\xcb\xf5\xc8\x57\x8a\xb2\x0e\xf1\x4c\x2a\x5a\xc4\x3f\xef\x59\x2b\xbd\xfa\xc5\x3f\xbf\x5e\xfc\xfc\xec\x47\x2f\x37\xda\x46\x51\x3f\xe6\xa2\xd8\xdc\xb3\x20\x7a\xdb\x0f\x71\xa4\xb1\x46\xd8\x52\xd6\x58\x80\xae\xb8\xad\x89\xf7\x51\xa6\x9c\x87\x51\xf6\x3d\x75\x0b\x7e\x36\xb2\x63\xf5\x39\xa2\x13\x05\x52\xc1\x72\x7d\xca\x2b\x99\xe6\x98\xf7\xbd\xe2\x66\x10\xfa\xd5\x0d\x45\x11\x8f\xcd\xa7\x1a\xdc\x6e\xa9\xee\x05\x3e\x2e\xb7\xd9\xe5\x26\x38\xed\x57\xdb\xd7\xb8\x67\x03\x82\xfd\xc2\x51\x8a\x29\xde\xcf\x4b\x2b\x86\x4b\xeb\x69\x58\x05\x64\x94\x75\xdf\x9e\xb8\x2f\x40\xc8\xea\xd7\xab\xa2\x25\x7c\x49\x55\xc0\x07\xad\xae\xb5\x7c\xab\x24\xdb\x15\x71\x68\xdf\xad\x06\x9e\x72\x66\xc6\x5d\xee\xe3\xd5\x50\xee\x85\x9c\x6c\x4a\x0e\x6b\x09\x79\x5d\x15\x2d\xc4\xc3\xff\xae\x3e\x7c\xe3\x7a\x37\x4b\x8e\x61\xa7\xb7\x4b\x32\xd9\x56\x88\xb3\x27\x0a\x54\xaf\x69\x09\xe3\x0f\x8b\x5d\x36\xf0\xd5\x59\x7e\xc6\xad\xf5\x17\x4b\xc9\x40\x27\x97\x3c\xd9\xf7\xc6\x6b\x08\x2d\xd0\x8b\x63\xa6\xdf\xc9\x6b\x4d\xa0\x13\x72\x8d\x77\xb1\x1e\x49\x9c\x9e\x9f\x2a\x99\xbf\x5c\x45\xe1\xdf\x9c\x4c\x70\xe0\x66\xce\xed\x13\xd5\x6c\xb9\x5b\x1e\xb8\xd8\x0c\xa7\x94\x6c\xe7\xab\x44\x64\x7c\xe1\x25\x01\xfe\x80\x86\x14\x16\xf2\x42\x76\x6e\x3c\x62\xeb\xc8\x20\x15\x60\x61\x37\xaf\xa5\xc6\x23\x79\x91\x12\xaf\xe0\x8f\x2b\x62\xbe\x07\xdf\x67\x1c\x7f\xaf\xce\xf0\x43\xbd\x7c\xe9\x20\xe0\x54\xbd\xf5\xe0\x46\x0d\x71\xd0\xeb\xbd\x9e\x48\xc0\x58\x84\xf1\xb7\x47\xa6\xff\x8e\x07\xf0\x61\x6f\xb7\xb1\xdc\xf4\x7c\x49\x0b\xe7\x4d\x37\x3d\x9f\x46\xf2\x00\x3a\x09\x19\xe7\x11\x3a\xcb\x40\x9c\x09\x2f\x91\xa2\xe1\xbb\x77\x1d\xbb\x54\x8d\x9d\x0b\xff\xa9\x2d\x6e\xa4\x2b\x86\xea\x5e\xd3\xcc\xf3\xf9\x9a\x5e\x1d\xdc\xd2\x74\x74\x05\x2f\xd6\x80\xa6\x7e\x90\xf0\xa5\x4a\x74\x78\xb9\x67\x26\xe1\x5b\x82\xb4\x87\x85\xda\xd9\x85\x2b\xdc\xae\xce\x8a\xe2\x6f\x7a\x4e\xfa\x60\xb4\x74\x6b\x9d\x8b\xe8\x0b\x7c\xcf\x97\x3a\xc8\x0d\x11\x80\x47\xe5\x78\x22\x34\x00\xfd\xc7\x41\x02\x37\x24\x12\x97\xcd\x46\x7f\x14\x8a\x15\x70\xc5\xdc\x7c\x2b\x8e\x21\x82\x8c\x47\x12\x04\x0f\x60\x2b\x17\x90\x23\x60\x88\xd7\x7f\x3d\xe4\xb6\x2f\x8a\xce\x9e\xe2\xbd\x76\xfc\xd1\x6d\xeb\xf1\x4a\xf8\xc1\xe1\xdc\xab\xc5\xac\xdf\x6b\x29\x55\xff\x05\x63\x60\x49\x7f\x25\xa5\x9a\x6d\x33\x9d\xfd\xbe\xd0\x3a\xf8\xfd\xd6\x81\x5f\xa9\xd7\x3f\x37\x23\x4a\x2a\x14\xaf\xf7\xd5\x20\x49\x33\x07\xc2\xf3\xb6\x2b\x29\x54\x39\x3d\x89\xab\x44\xdd\x0d\xe7\x78\x72\x72\x2b\xcc\x88\xa9\x17\xc9\x7e\xe7\x5e\xec\x43\x6a\x30\xeb\x89\x36\x1d\x77\x56\xda\xb5\x26\x72\x4f\x96\xea\xe9\x80\x7e\x64\x79\x0a\xb0\xf8\xe7\xd6\x8c\x17\x53\x3c\x7f\xf9\x39\x07\xe7\x4b\xea\x9c\x1c\xe7\x7d\xd7\x72\xfc\x32\x01\x4d\x63\xf3\x65\xb5\x2a\x9c\x47\x1f\x36\x80\x6f\xaf\x29\x89\x39\x79\x54\x92\xae\x95\x8e\xda\x66\xbc\x38\x6e\x5c\x27\x54\x3b\xa3\x62\xf8\x2b\xbd\x6b\x86\x2c\x3e\x08\xf0\xc3\xba\x70\xd4\xe4\x0e\x71\x1b\xa7\x3e\x79\x4e\x6d\x1b\x2f\xd4\xa8\x14\x2a\x3e\x1c\x62\xef\xdc\x9f\x15\x65\x7d\xd6\x85\x61\x1d\x52\x4e\x13\x32\xe4\xab\x60\x3a\xa6\xec\x8d\x87\x5b\x56\x17\xee\xe4\x39\xe4\x3c\xc9\x46\xd9\xaf\xb5\x08\xf8\xd5\xe5\x4e\x3f\x9c\x82\xc1\x48\x69\xf6\xdb\x11\x34\x6f\xe4\x1c\x13\x35\x82\xb1\x46\xf6\x4e\x19\x25\x9b\x4a\xbc\xd3\x06\x68\x8f\xc9\xa2\x3f\x65\xaf\x4b\xd2\xe0\xb8\x5c\xc4\x4a\x68\x59\xcc\x1f\x73\xf1\x43\x7c\x99\x9f\xd2\xfd\x6e\x0c\xba\x4b\x13\x39\x30\xf4\x87\x12\xcf\xdb\x6d\x13\x58\x84\xc1\xfc\xa4\xe6\x39\x09\xf0\xee\xe0\x5b\xc5\xf9\x03\x8c\x99\x9b\xcb\xca\xed\xca\xdc\xf6\xc6\xe1\xac\x92\x3d\xf4\x96\x57\xe9\xb4\x4d\x6d\x49\x14\x16\xe2\x92\x07\x9b\x8f\xce\x71\xec\x92\xe1\xb2\x46\x6b\x6c\x57\x5b\x8a\x41\x72\xd1\x64\x49\x0e\x02\xea\x1a\x4a\x87\xdb\x60\x44\x8f\x42\x85\xab\x7d\x6e\x0a\xc3\x50\x39\xa5\xec\xa0\x96\x9b\x88\xa9\x08\x9b\x15\xca\xe2\xe9\x2b\xf6\x13\xef\xd6\x52\x7e\x49\xbb\xf3\xa5\x48\x62\x7f\x33\xa0\xc4\xa4\x37\xc5\xe6\xd4\x9b\xea\xd9\x75\xf4\x77\x23\xb9\x5f\x9a\x9e\xea\xaf\x66\x29\x8f\x2d\x4b\x61\x39\xfb\x89\x10\x57\x67\x6b\xc6\x20\x14\x30\x4a\xc8\x0e\x7c\x91\x5c\xf4\x54\x00\x9b\xc7\xf8\xfc\x4d\x8a\x56\xa6\xd6\xa7\x1a\xb5\x57\xd7\x69\x02\x52\xc5\xa9\x6f\x7d\x23\xe5\x64\x84\xc3\x53\x9b\x10\x9e\x4a\x82\x62\x07\xf4\xf6\xcb\x0f\x9f\xdf\x1e\xcd\x29\x71\xe3\xa7\xe0\x6a\xf8\x12\x50\xb4\xcd\x5f\xfa\xd7\x93\x86\x5f\x31\x5d\x12\x58\xcc\x4d\x0d\x23\xaf\x3c\x4c\x64\x38\xab\xf5\xaa\xbb\x5d\x6e\xf4\xaa\x25\xbe\x9f\xf0\xbc\xfa\x4c\x6d\xbd\xf4\x5f\xe7\x7a\x94\xda\xe6\xff\x47\xfc\x23\x2d\x64\xc2\xff\x14\xd9\xff\x07\x00\x00\xff\xff\x1d\x2f\x1a\xe1\xe5\x06\x00\x00")

func rootPokerCrtBytes() ([]byte, error) {
	return bindataRead(
		_rootPokerCrt,
		"root.poker.crt",
	)
}

func rootPokerCrt() (*asset, error) {
	bytes, err := rootPokerCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "root.poker.crt", size: 1765, mode: os.FileMode(436), modTime: time.Unix(1561488403, 0)}
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
	"poker.crt":      pokerCrt,
	"poker.key":      pokerKey,
	"root.poker.crt": rootPokerCrt,
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
	"poker.crt":      &bintree{pokerCrt, map[string]*bintree{}},
	"poker.key":      &bintree{pokerKey, map[string]*bintree{}},
	"root.poker.crt": &bintree{rootPokerCrt, map[string]*bintree{}},
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
