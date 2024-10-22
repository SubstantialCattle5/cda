// Code generated by go-bindata.
// sources:
// functions/dccheck.sh
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

var _functionsDccheckSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x55\xc1\x6e\xe3\x36\x10\xbd\xeb\x2b\x5e\x68\x6d\x6d\xb7\x91\x6c\x67\xd1\x1e\x92\x38\x68\x9b\x64\x17\x3d\x14\x5b\xa4\xed\x29\x08\x0c\x4a\xa2\x2c\xc2\x14\x69\x90\x54\x36\x5d\xaf\xff\xbd\x20\x69\xd3\xb2\x77\x13\x5f\x2c\xcd\x70\x86\x6f\xe6\xbd\x19\x0d\xce\x26\x05\x97\x93\x82\x9a\x26\x49\x06\xe8\x0c\x5d\x32\xb4\xcc\xb8\xff\xc4\xbf\x8d\xc6\xd8\x24\x00\xc0\xca\x46\x81\xfc\xeb\x6c\x97\x48\xa7\xb8\x56\x6b\xa6\xa9\xe5\x4a\xde\xe0\x7a\xad\x94\xc8\x24\x6d\xd9\x0d\xae\x79\x4b\x97\x2c\xbc\x90\x5e\x24\x10\x03\x2e\x31\x64\x2f\x6b\xa5\xed\x10\xb5\xd2\xa8\x3b\x21\x10\x0c\xe7\x7b\x4f\x56\xf1\xba\x0e\xee\x9e\x61\x97\xef\x85\x5b\xcc\x92\xad\x43\xfc\xa1\x93\xa5\x4b\x09\xab\x50\x36\xac\x5c\x81\xd7\xd0\x45\x85\x52\xb5\x2d\x95\x95\x3b\x6c\xac\x49\xbc\x6f\xa1\x8b\x6a\xb1\x73\xc4\xba\x78\x8d\xb3\x78\x3a\x7b\xf6\xc1\x3f\xdc\x4c\x2a\xf6\x3c\x91\x9d\x10\x57\xb0\x0d\x93\xfe\xe4\xa1\x96\x7b\xad\x95\xbe\xc4\x50\x17\xd5\x30\xc6\x4a\x65\x51\xab\x4e\x56\x39\xfe\x12\x8c\x1a\x86\x96\xae\x18\x4c\xa7\x19\x6e\xd9\xba\x41\x67\xb9\xe0\x96\x33\x03\xaa\x19\xb8\x34\x96\x0a\xc1\x2a\xb8\x60\x5a\x96\xcc\x18\x5e\x08\x96\x93\xc3\x65\xa1\x50\xf7\x58\xf3\x50\xef\x2d\x15\x65\x27\xa8\x65\x0e\x16\xcc\x5a\x70\x57\xdc\xde\xb8\x08\x86\x58\x9b\x50\x25\x15\x70\xe4\x38\x3a\xe6\x24\x9d\x91\x9e\xdd\x13\xb5\x73\x5c\x90\xc4\x7b\x06\xf8\xc0\x6c\xd9\x84\xec\xfc\x0b\x43\xad\x55\xeb\x7b\xc2\x65\xad\xfc\x11\x67\x9e\xcf\xa6\xfb\xe6\x3d\x3e\x22\xfb\x02\x92\x3a\x33\xc1\xd3\xd3\x5b\x0d\xbb\x55\x9d\x08\x8d\xaa\x98\x65\xba\xe5\x72\x77\x8b\xaa\xfd\x8d\x1e\x51\x8e\xfb\x17\x6e\xb9\x5c\xbe\xda\x89\x1d\xd2\x43\x2f\x66\xd3\x77\xfb\x14\x2e\x5d\xbf\x76\xa6\x4b\x26\xed\x3c\x1d\x05\x20\xa6\xa4\x82\xcd\xa7\x57\xf0\x78\xf1\x23\xa6\xf9\x0c\x13\xcc\x08\xbe\xa2\x28\xc7\x49\x4f\xb0\xe9\x2e\x96\xf8\xce\x9f\xe1\x9f\x4f\x77\x9f\x2e\x83\xc8\x54\x67\xbd\x34\xd5\x33\xd3\xb5\x50\x9f\xc1\x5c\x7d\xc9\x00\xbf\x69\x4d\xff\x73\x5a\x6c\x94\x70\x2d\xab\xf8\x33\xaf\x3a\x2a\x50\x36\x9d\x5c\xa1\xa1\xa6\x61\x4e\x8c\x9d\x5c\x2d\xc2\xcb\x7c\x34\xfe\x46\xc7\x47\x1c\x97\xaa\x2d\xb8\x64\x55\xb8\xd9\x74\x6d\x8f\xee\xbd\x6f\xb1\xf7\x45\xe6\xa3\xc7\xdd\xb2\xe0\x72\xdd\xb9\x1e\xac\x35\x97\xb6\x06\x79\x67\x08\x48\xba\xe9\xe3\x78\xfc\xf5\x69\x4b\xc6\x87\xf2\x33\x09\x92\x7e\x27\x8b\x6b\x94\x69\xe8\xc5\xcf\xbf\x38\x28\xdb\x24\x09\xe3\xb9\xa0\xf2\x80\xe2\x54\x80\x71\xf0\x4f\x15\xd8\x53\xe6\xc5\x6b\xca\x7c\xdf\x77\x04\x7d\x3b\x01\xa6\xa3\x53\xd5\x3b\xc6\x76\xf9\x5c\x75\x31\x07\x39\x62\xf5\xa1\x93\x92\xcb\x25\xd2\x4d\x44\xb5\xf5\x13\x18\xc1\x7b\x66\xd3\xcd\x3e\xd7\x76\x92\x6e\x62\xae\x6d\x9e\xe7\xfd\xa5\xf6\xb7\xbb\xd9\x69\xd5\x53\xf5\x87\x3b\x86\x82\x1a\x56\x41\x49\xa4\x9b\x08\x77\xfb\x31\xcf\x63\xe8\x00\xf7\xbe\x67\x78\xf8\xfd\x2e\x94\x7a\x1e\x0a\x03\x97\x7e\x8d\x75\x72\x65\xce\x3d\x28\xd7\x77\x30\x5a\x36\xc1\x1a\xe6\x4f\xd2\xb5\x69\x94\x5d\xd4\x5c\xb0\x39\xc9\x27\x1f\xef\x72\x6e\x54\xc8\x1d\xf2\x64\xc5\xf1\xe5\x8e\xec\xa3\xb0\x2d\x41\x96\xd5\x5c\x58\xa6\xe7\xc3\x48\xe7\x10\x5f\x31\x8a\x33\xf7\xb9\xe1\x82\x41\x33\x5a\x21\xd3\x01\x89\x0b\x75\x6d\xb8\x42\xa5\xe2\x39\x2f\xb7\x9e\x92\x7e\x9a\x8f\x48\xea\x1e\x77\x72\x72\xbf\x4a\xc9\x30\x95\xe3\xb0\xc6\xfe\xa4\x5c\xc2\x94\x9a\xaf\x2d\x84\x5a\xf2\x32\x49\xdc\x2a\x01\x49\x07\x04\x99\x64\x78\x8f\xfe\x22\xf1\x5f\xa2\xc4\x4d\xfe\x89\x94\x8e\x05\x74\x22\x1b\xb7\x2e\xfd\x27\xc1\x11\x1a\x56\x98\x5f\xb9\x3e\xde\x79\x07\x78\xe8\xa4\x67\x2e\x88\x38\xe8\x20\x4e\x5f\x54\xc4\xf7\x45\x0e\x92\x46\x34\xe4\x55\xf1\xbd\x35\xad\xc9\xff\x01\x00\x00\xff\xff\x83\x5a\x1a\x1a\x7f\x07\x00\x00")

func functionsDccheckShBytes() ([]byte, error) {
	return bindataRead(
		_functionsDccheckSh,
		"functions/dccheck.sh",
	)
}

func functionsDccheckSh() (*asset, error) {
	bytes, err := functionsDccheckShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "functions/dccheck.sh", size: 1919, mode: os.FileMode(509), modTime: time.Unix(1720757990, 0)}
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
	"functions/dccheck.sh": functionsDccheckSh,
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
	"functions": &bintree{nil, map[string]*bintree{
		"dccheck.sh": &bintree{functionsDccheckSh, map[string]*bintree{}},
	}},
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

