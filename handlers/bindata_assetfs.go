// Code generated by go-bindata.
// sources:
// ../openapi.yaml
// DO NOT EDIT!

package handlers

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _openapiYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\xdd\x6f\xe4\xb6\x11\x7f\xd7\x5f\x31\x48\x0b\xf8\x25\xeb\xf5\x7d\xb8\x0d\xf4\xe6\xf3\x35\xa9\x81\xb8\x67\xd4\x06\xfa\x50\x14\x0d\x57\x9a\xdd\x65\x8e\x22\x15\x92\xf2\xda\x28\xfa\xbf\x07\x24\xf5\x41\x4a\x5c\x69\xe5\xf3\xad\x73\x87\xdb\x27\x9b\x9a\x19\x0e\x7f\x9c\x2f\x0e\x29\x4a\xe4\xa4\xa4\x29\xbc\x39\x3d\x3b\x3d\x4b\x12\xca\xd7\x22\x4d\x00\x34\xd5\x0c\x53\xb8\xdb\x22\xdc\x72\xf2\x11\x17\xb7\x28\xef\x51\xc2\xc5\xcd\x55\x02\x90\xa3\xca\x24\x2d\x35\x15\x3c\x42\x43\x15\x10\x50\xee\xef\xb5\x90\x20\x38\xa3\x1c\x81\xc8\x8c\xe4\x08\x1b\x52\x20\x2c\x40\x19\x8e\xd3\x04\xe0\x1e\xa5\xb2\x72\xde\x9e\x9e\x9f\x9e\x25\x00\x8c\x66\xc8\x15\x1a\x2d\x00\x38\x29\x30\x85\xeb\xab\xbb\x04\x20\x13\x5c\x93\x4c\xfb\x1f\xae\xee\x09\x87\x9b\x4a\x6d\x3f\x52\x6e\x87\x2b\xc9\x52\xd8\x6a\x5d\xaa\x74\xb9\xd4\x3b\xaa\x35\xca\xd3\x4c\x14\x4b\x43\xf9\x1e\xef\x91\x89\x12\x65\x92\x38\xf5\x94\x91\xb5\x70\x4c\x4b\x52\x52\x2b\x22\x58\xdc\x65\x25\x25\x72\x0d\x94\x2b\x4d\x78\x86\x1d\xbd\x99\x24\x5d\x2e\x99\xc8\x08\xdb\x0a\xa5\xd3\x1f\xce\x7e\x38\x8b\x0b\xf9\xfb\xdd\xdd\xcd\x1e\x09\x6a\x86\x88\xdb\x4e\x46\x52\x12\xbd\xb5\xda\x2f\x0d\x9e\xca\x61\xb2\xc1\x1a\x1c\x00\x55\x15\x05\x91\x8f\x29\xfc\x84\x1a\x08\x30\xaa\x34\x88\xb5\x05\x5f\xd5\x24\x81\x7c\x4b\xc6\x98\x23\x00\xc1\xdb\x1d\xac\x89\x35\xd9\xa8\x46\xb6\x59\xc0\x4f\x9e\xa0\x92\x48\x52\xa0\xae\xe1\x6c\x28\xdc\x06\x31\x5a\x50\xdd\x8e\x02\x50\x9e\xc2\x6f\x15\xca\x47\x6f\x2c\x54\xc4\x2a\x60\xd9\x40\x0b\x90\x98\x21\xbd\x47\xa0\x46\x23\x89\xaa\x14\x5c\xa1\xc7\xab\xb2\x2d\x16\x24\xf5\x46\x00\xf4\x63\x89\x29\x50\xae\x71\xd3\xea\xef\x7e\x6b\x21\x0b\xa2\xed\xb7\x37\xaf\x07\xba\x2a\x21\x35\xe5\x9b\xf9\xda\xd6\x8c\x76\xe2\x83\x74\x53\x5a\x86\x13\x01\x20\xaf\x8a\x90\xd4\x28\xa6\x0a\x22\xf5\x60\x54\x12\x9e\x8b\xa2\x1e\x6e\x40\xf1\xc0\x7f\x7d\x76\x96\x4e\xe1\x6b\xec\x81\x37\x5e\x4a\x78\x0e\x24\xcf\xa9\x21\x21\x0c\x4c\x10\x30\x40\x51\xc1\x3d\x31\xc6\xfd\x90\xeb\x50\x49\x52\x96\x8c\x66\x96\x74\xf9\xab\x12\xbc\xbf\x84\x18\x06\xe6\xf7\x67\x89\xeb\x14\x4e\xfe\xb4\xcc\x44\x51\x0a\x8e\x5c\xab\xa5\xa3\x55\x4b\xab\xe2\x49\xcb\xf2\x36\x5c\x4d\x8c\xb3\xc5\x60\x79\xc5\xef\x09\xa3\xf9\x4d\x6b\x92\x4e\x4e\x69\xbc\xab\xef\x1a\xb7\x9a\x48\xe3\x1c\x1c\x77\xd6\xee\x63\x7e\x71\x29\x91\x68\xf4\x88\x7c\xdf\x80\x1d\xd5\x5b\xd8\xd0\x7b\xe4\x50\x90\xd2\x73\x04\x8b\xa8\x35\x62\x75\x80\x03\x49\xfc\xad\x42\xa5\xdf\x89\xfc\xb1\x23\x31\x83\x54\x62\x9e\x82\x96\x55\x67\x56\x91\x5d\xf0\xf7\xe0\x61\xb1\xdb\xed\x16\x66\xfb\x16\x95\x64\xc8\x33\x91\x63\x1e\xc2\x1f\xdf\x12\x67\x98\x62\xf5\x2b\x66\x7d\x83\x2b\xa5\x89\x99\x9a\xfa\x26\xd6\xfc\xec\x1a\x87\xc3\x3d\x14\x6f\x18\x79\x34\xa8\x38\xb7\x36\x39\xa1\x87\xf9\x50\x93\x98\xfb\xba\x5f\xdc\x89\xbb\x5f\x41\x39\x2d\xaa\x22\x85\x57\x83\x8f\x3b\x9a\xeb\xed\xa4\xb6\xd7\xa4\x74\x94\x9f\x43\x39\xf2\xe0\x94\x7b\x7d\x7e\x3e\xf8\xbc\x45\xba\xd9\x4e\xa3\x69\xf4\x73\xa4\xc7\x56\x10\x39\x59\x31\xfc\xef\x8e\x30\x16\xb1\x85\x41\x5d\x40\x15\xac\x84\x60\x48\x78\xe7\x1a\x40\x79\x6e\x8c\x15\x15\xec\xb6\xa8\xb7\x28\x4d\x98\x27\x79\x0e\x56\xaa\xf9\x47\x6f\xd1\x73\x37\x09\x5c\x98\x54\xb0\x77\xad\xf5\x14\x51\x75\xd6\xa4\x62\xba\xe7\x41\xee\xd7\xba\xd7\x80\x6f\x31\x48\x5a\xcd\x78\xdc\x28\x16\xe1\x66\x44\xe3\xf1\xab\xbd\xf1\xf8\xaa\x8b\xb5\x40\x56\xa2\xd2\x90\xd9\x90\x93\xf7\xdd\xe3\xa8\xe1\xf7\x79\xa3\xaf\xf9\x9d\xcf\x91\xe3\x0a\xc9\xbf\x49\x29\xa4\x2f\xe1\xcd\x5e\x14\x0d\x03\xcd\x10\x2a\x4e\xee\x09\x65\xc6\x4a\xbf\x87\x8c\x70\x63\x3a\x59\x3f\x84\xbf\x04\xa8\xed\x5a\x5c\xbd\xb6\xfc\x1f\xcd\xff\x3f\x5e\xb4\xd1\x81\x61\x90\xbd\x59\x6a\x8c\x1e\x56\x8f\x40\x73\xe4\x9a\xae\xcd\xa2\xc4\xd3\x2b\xba\xd8\xfa\x3a\x4a\x6b\x37\x57\xef\x4f\x9e\x56\x96\x0c\xdd\xc0\x44\x01\x17\x01\xfa\x39\xe9\x0b\xf7\x84\xb7\x67\x6f\x0f\x97\x63\x54\xf8\x87\xd0\x3f\x8a\x8a\xe7\xcf\xe5\x4c\x39\x32\xd4\x38\x30\xba\xf7\x76\x78\xc4\xca\x02\x82\x2f\xc3\xac\x3e\x58\xdb\x71\x75\x5a\xa0\xad\x39\x0c\x19\x0b\x73\x58\xbc\x5c\xb0\x75\x98\xe6\x5f\xa1\x95\xc1\x27\x84\x6c\xb7\x2b\x5d\x08\x90\x26\xbf\x02\x17\xbb\x97\x0d\xdd\x19\x29\x49\x46\xf5\xe3\xfe\xc0\x5d\x77\x40\x1a\xc2\x3d\x3e\x71\xeb\x9f\xac\x07\x81\x5c\x85\x32\x4e\xe1\xb2\xfe\x0b\xa8\x72\xb5\x51\x55\xac\xd0\x5a\xb0\x28\x91\x63\x6e\xa0\xe0\x98\x19\x09\x0a\x72\x7a\x4f\x73\xcc\xcd\x31\xc5\x55\xdb\xb5\xa1\x07\xa7\xf9\x99\x5e\xf4\x4f\xd4\x95\xe4\xad\x46\x46\x64\xaf\x3d\x70\xd4\x0d\x69\xf0\xb0\x7b\xd2\x34\xac\xe2\xfb\x31\x4c\x2b\x13\x7d\x8d\xf1\xad\x89\xe5\xd8\x4f\x07\x56\x4d\xcb\x3d\x2a\xbe\x06\xb4\x5e\xa9\xb2\x5c\x49\x41\xf2\x8c\x34\xa7\xe8\xe8\x79\xfa\x5d\x43\x03\x04\x0a\x54\x8a\x6c\xf0\x80\xa4\x10\x00\x12\x11\x61\x8f\x07\x8c\x41\x59\x9f\x21\x6d\x27\xc8\x46\x05\x1b\xd5\x4b\xa2\x14\xe6\xd1\x54\xf4\x2c\xe9\xe6\x8b\x3d\x99\xd7\xf0\x4d\x9f\x26\x3b\x98\xdb\x4d\xb6\x47\x74\x0f\x70\x85\x0c\xb3\x48\x9a\x0c\x55\x8c\x74\xb5\xc6\xcf\x59\xa1\x91\x3c\x2d\xb5\xdb\xa6\x30\x35\x51\xb0\xd3\x5e\xa2\xaa\x98\x86\x35\x23\x9b\x97\xf0\x9f\xd6\x88\xbf\x65\xf5\x20\xab\x7b\x1b\xf4\xc7\xc8\xe7\x5e\x7c\x73\xee\x35\xd2\x47\xff\xd9\x76\x4c\x1d\x95\xeb\x01\x16\xa4\x9c\x1b\xde\x9a\x2e\xfb\x1e\x39\xc7\xaf\x8f\x2f\x42\x65\x3a\x55\x8e\xba\x23\xce\x95\x9f\xb9\xe5\x0b\x2f\xef\x2e\xcb\x92\xf2\xcd\x7e\x93\xba\xa1\x7c\xb3\x28\x05\xdf\x34\x69\xa6\x8b\x9e\x87\x97\x25\x46\x48\xc3\x6f\x62\x78\xb6\xc5\xec\x23\x54\xe5\x27\x56\x26\x37\x4e\xad\xc1\x7d\xcb\xd1\x6c\xc2\x28\x70\x92\x74\x1f\xd2\x24\x19\x3a\x88\xb3\xfd\x46\x30\xe5\x29\x94\xa4\xed\xd0\xb9\x4b\x1d\x9a\x27\x71\x25\xe2\x8d\xd2\x58\x7b\x34\x9a\xef\x07\x97\x2a\xbd\x12\x24\x19\x40\x3e\xb0\xd3\x34\x26\xaa\xa6\xf2\x96\x9a\xec\x41\x7e\x0c\xf5\x18\xe2\x07\xc5\x44\x00\xcf\x86\xa3\x0a\x5e\x70\x0b\x9a\xe4\x84\x35\x47\x15\x34\xc4\x47\x56\xd3\xf7\xd6\xa8\x9e\x76\x4f\x4c\xda\x59\x1b\x92\x23\x6a\x97\x34\x9c\xce\x66\x9d\xaa\xed\x66\x47\xaa\xb9\x91\xb2\x46\x85\x37\xb7\xdd\x95\x5d\xdf\x34\xbd\x58\xe1\xdf\xf1\x0e\xbb\xd9\x0b\xc8\x44\xc5\x9b\xff\x63\x05\xe4\xc6\x57\xb7\x53\x99\x48\x49\xfc\x7b\x50\xaa\xb1\x50\x87\x83\x13\xf6\xd2\x06\xd7\x46\xfb\xee\x80\x07\x6a\x0c\x6f\x36\xe2\x37\x1a\x76\x95\x7b\x67\x68\xae\xf5\x1d\xaa\x96\x76\x80\xed\xe1\x73\xd6\xee\xdd\x1c\x44\xe7\x6f\x72\x26\x8a\x42\xf0\xc8\x31\x70\x72\xb3\x49\xa5\xb7\x42\x06\xbb\x6d\x5f\x4f\x78\x23\xf5\x03\x0b\x6f\x64\x55\x51\x96\x8f\x58\x80\x13\x3a\x34\x81\x5e\x81\x1f\x3c\xd4\x18\xa1\x6b\x5e\x78\x4c\xd1\x59\xb5\xf6\x52\xb5\xae\xf4\x54\x90\x87\xe8\x7a\xa7\x99\x18\xb6\x6d\xea\x98\xf2\xa2\xe1\xdd\xd0\xa2\x7f\x41\xb7\x00\x49\x34\x8e\x60\x4e\xf3\x51\x77\x88\x1e\x70\x61\xb6\x57\x4c\xfb\x5d\x7b\xec\x3b\x86\xfb\x35\x73\x39\x07\xa4\xbc\x7f\xbe\x9c\x37\xe7\xe0\x76\x77\xf2\x56\x77\x9e\xfc\xe1\xf5\xec\xf4\xb5\xec\xbc\x19\x8c\x8d\x8c\xef\x8e\x67\x45\x87\x8b\xb7\x1f\xde\x85\x1d\x9c\xf9\xfe\x33\x76\xbe\x8e\xb9\x8f\xaa\xb2\x0c\x95\x1a\xb1\xf9\x9a\x62\xef\x8a\x7f\x64\x64\x63\x9b\x9e\xb2\x42\xa0\x6b\x4f\x03\xaa\x1c\x33\xe6\xf6\xa5\xc5\xba\x26\x5c\x13\xa6\x86\x94\x6b\x42\x19\xe6\x03\xd4\x9a\x7b\x63\xfb\xe1\x32\x68\xed\x3e\x21\x82\xd7\x16\xad\xa2\x0d\xe0\x18\x3c\x3d\x92\x18\x3e\x59\x4f\xa9\x81\x16\x4f\xea\x0c\xaf\x1e\x7b\xb4\x84\x31\xb1\xeb\x11\xaf\x85\x6c\xd6\xe2\x3d\x5c\x0b\x01\x74\x12\x22\x56\x97\x8b\x6a\xc5\xd0\x01\x1b\x14\x92\x53\xa8\x5e\x76\x09\x31\x47\xa3\x9a\xad\x2c\xed\x16\xab\xad\x90\xda\xa7\x06\x8d\x0f\x7a\x0c\x5f\x91\xfb\x59\xd0\xa3\x8e\x42\x2d\xf2\xfd\x9e\x67\xd7\x10\x0a\x9c\xeb\xd9\x66\xfa\xfd\x4d\x13\xbb\x36\xb7\x56\xef\xc3\x68\x26\xac\x2f\x8e\x9e\x2f\x19\x46\x6e\xc2\x46\x93\xe2\xf1\x93\x98\xfd\xf0\xc1\x6f\xd4\xcc\x5f\x77\xf3\xc6\xb1\xd7\xf5\xb0\x26\x66\x9f\x70\xb5\xab\x1a\xc1\xa0\xe6\x1d\x01\x42\x84\x5a\x0e\xd4\xfa\xb9\xae\xea\xf7\xa8\xf3\x94\xe2\x5b\x70\xfc\xb0\xee\x9f\xb1\xe3\xdd\xa3\xa6\x26\xb7\xaf\x76\x4f\x66\xb1\x5c\x94\x25\x9b\xc9\x72\x29\x64\xa9\x66\xf2\x5c\x8b\x6a\x2e\xcb\xbf\x88\x46\x59\x20\x13\x7c\x2e\x1f\x63\x1d\x47\x41\xca\xa9\x76\x4f\xab\x23\x29\x4f\x6a\x5f\xa4\xd2\xc5\xcd\xd0\x2a\x83\xf2\x36\x7c\xdf\xb9\x80\xef\xb8\x90\x7a\xfb\x9d\x3f\x82\x44\xe9\x60\x40\x89\xaa\x47\xb2\x43\x43\xe2\x26\x15\xbd\x4c\xee\xdb\x48\xcf\x3e\x0e\xef\x75\x78\xcf\xe7\xce\x3c\x50\x22\x6f\xc2\x78\xc5\x6c\x7b\x37\x75\x59\x37\x69\xb9\xaf\xec\xdc\xd0\x48\x2c\xc8\x43\x6f\x04\x1f\x48\x51\x1a\xc6\x7f\xbf\xfe\x1e\xde\xfc\xa7\x5d\x8e\x3a\x78\x3d\xa3\x37\xea\x42\xd7\xdb\x62\xcd\x7b\x66\xa0\xb0\x3c\xa7\xf6\x4d\xbb\xe1\x30\x89\xf5\x17\xfb\x54\xfd\x97\x91\x78\x10\xbc\xfb\xed\x9d\x1b\x72\x31\x1a\x28\xac\x5a\x07\x1a\x9c\x53\xf1\xee\xb1\xf4\x3c\x23\x8c\xb8\xd3\xdc\x57\x5e\x6b\x33\x17\x61\x84\x9a\x40\x55\xd5\xb0\xda\x10\x30\x13\x56\xcb\xd3\x83\x95\x98\xb1\xaf\x13\xd6\x19\xa8\xd6\xa0\xba\x20\x39\x13\x55\xc7\xd4\x83\x35\xb3\x83\x5f\x27\xae\x4f\x31\x57\x9b\x4a\x66\x02\x6b\x79\x7a\xb8\x16\x66\xec\x19\x60\xb5\xff\x36\x09\xe3\x8f\x0b\xf5\x2c\x13\x6e\xf9\xfa\x89\x70\x92\xbb\x61\xa8\x37\xab\x4b\xe2\x33\x77\xac\x63\xec\x6d\xdb\xae\xfd\xf0\xcd\x25\x54\x8b\x32\x63\xb3\xf1\x65\x6c\x80\x2c\x63\xdf\x30\x6d\x30\x6d\x04\x85\xb8\x86\x25\x97\x57\x70\xfd\xe5\x6d\x0c\xec\x0b\x5e\xef\x44\x7b\x52\x6a\x79\xdb\xa2\xe9\xcd\x2b\x7f\xc2\x3b\x0f\xa9\x43\xaa\x4e\x9b\x74\x83\x92\xd2\xe5\x8b\x60\xc8\x86\xba\xb0\x12\x35\x25\x50\x58\x89\x12\xc6\x7a\x03\x8d\xa3\xd5\x05\xea\x8d\x68\x2e\x3f\xa7\xad\xcc\x5e\x37\x36\x0f\x7e\xd6\x52\x14\xd3\x9d\xef\x52\xb4\xab\x8c\x99\x50\xe9\xcd\xbe\x6f\x3f\x20\x7e\xc4\xbc\xee\xce\x00\x53\x7a\x5f\x1f\x7a\x64\x1c\x6d\x0f\xc7\xf4\xff\xdc\xcd\x4c\xcf\xa2\x5e\xbd\xee\x4a\xfd\xcf\xdf\xe3\xf4\x26\xfe\xeb\x79\xf2\x7b\x00\x00\x00\xff\xff\x72\x69\x65\xd8\x80\x3a\x00\x00")

func openapiYamlBytes() ([]byte, error) {
	return bindataRead(
		_openapiYaml,
		"openapi.yaml",
	)
}

func openapiYaml() (*asset, error) {
	bytes, err := openapiYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "openapi.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"openapi.yaml": openapiYaml,
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
	"openapi.yaml": &bintree{openapiYaml, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
