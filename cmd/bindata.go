// Code generated by go-bindata.
// sources:
// ../devctl.sh
// DO NOT EDIT!

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

var _devctlSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x53\xcd\x6e\xd3\x4a\x14\xde\xcf\x53\x9c\xeb\x5a\x72\x73\x25\x27\x71\x7b\x17\xb7\x0a\x61\x41\xe2\xc2\x22\xa4\x52\xa0\x05\x54\xaa\x64\xe2\x39\x8e\x47\x1d\x8f\xa3\x99\x71\xd3\x2a\xf5\x3b\xb0\x40\x42\x42\x95\x60\xc7\x7b\xf5\x09\x78\x04\x34\x89\x4d\xe2\xb6\x6a\xc5\x82\x59\xcd\xb1\xcf\x37\xdf\x8f\xce\xd9\xf9\xa7\x35\xe5\xb2\x35\xa5\x3a\x21\x64\xcc\xf0\x22\x32\x62\xac\x72\x39\x8e\xb2\x34\xa5\x92\x75\x27\xe6\x6a\x8e\xe0\x53\x58\xff\x83\x6b\xa0\x8b\x73\xf0\x86\x87\x4b\x98\x2b\x2e\x0d\xb8\xc3\x43\x28\x3c\xb8\x06\x43\xb9\x00\x5f\x42\x30\x21\x84\xc7\x70\x6a\xef\x8e\xdb\x0f\x4f\x7a\x6f\x07\xe3\xd1\xf1\x70\xdc\x0f\x4f\x1c\x38\xeb\x80\x49\x50\x12\x00\x91\x45\x54\x40\xc5\x39\xa7\x26\x21\x50\x2b\xbb\x93\x92\xd3\xf7\x99\xba\x52\xb9\x84\x88\xc1\x8c\x9b\x24\x9f\x36\xa3\x2c\x6d\x4d\x23\x2a\xd8\x02\x85\x68\x3d\xae\x6d\xb2\xf5\xee\xb6\x35\x67\x96\x81\x7d\xd5\x5d\x6e\xb3\x16\xad\x94\x72\xd9\x9c\x65\x0e\x89\x39\x21\xeb\x3f\xbb\x0d\x58\x12\x80\x88\x6a\x04\xc7\x0d\x1c\xe0\xd6\x81\xf5\x40\x99\xcf\xf0\xa2\xb1\xaa\x00\xf0\x72\x9e\x29\x03\x75\xd3\x5d\xa3\x72\x2c\x1b\xd6\xa6\xeb\x9e\xed\xf9\x4b\xb6\xed\xd9\x01\x9d\xa0\x10\x51\x82\xd1\x39\x30\xae\xe9\x54\x60\xf7\x4d\x2f\x68\x1f\xb4\xcb\x0e\x9d\xe5\x2a\xb2\xc6\x96\xb5\x1c\xd6\x45\x53\x27\x4e\xd9\x57\xc5\x84\x51\x92\x8d\xb9\x8c\x33\x70\x06\x19\x65\xc8\xac\xfc\xd2\x42\xd5\xab\xd0\xe4\x4a\x96\x45\xa7\xb3\x49\x4b\x5f\x69\x83\x69\x15\x58\x2e\x35\xde\xcd\xeb\x49\xdd\xc1\x3d\xdd\xaf\x8e\x5e\x87\x45\xab\xb9\x96\xf0\x27\xc2\xd7\x6a\x9e\xd4\x8e\x9a\x46\x84\x00\xc4\xac\xeb\xb8\xbb\xe9\xb9\xc1\x74\x0e\x2d\x93\xce\x4b\x2e\x3f\x66\xfe\x7b\x7b\x1a\xf6\x0d\x95\x82\x1f\x5b\x59\x31\x2b\x1c\x0b\xc3\x0b\x2a\x6c\xfd\xc0\x14\x16\xe0\xfe\xeb\xc0\xff\xcf\xb7\xba\x17\x09\x17\x08\xd3\x9c\x0b\xc3\x25\x28\xa4\x0c\x7c\x05\x82\x4b\xec\x00\xcb\x56\xb2\xaa\x07\xed\xc7\xc2\x52\xb2\x4c\x22\x3c\xdb\xe6\xac\x8b\x28\x36\xdb\xbd\x0a\x74\x8c\x4a\x65\x6a\xb7\x61\xa7\xda\x2e\xeb\x29\xb8\x01\xf8\x12\xa1\x0d\x67\x9b\x15\xbd\x93\x5c\x6c\x37\xdc\x71\xf7\xc0\x5e\x90\x59\x5e\x14\x1a\xef\x37\xea\x3c\x8a\x50\xeb\x55\x6f\x79\x8f\xf3\x55\xbc\x31\xb7\x52\x86\xbd\xae\xf7\xf1\x32\x98\x9e\xb6\x53\x8f\xbc\x1c\x85\xe1\xb0\xac\xf7\xf7\x52\x8f\x8c\xc2\x7e\x55\x06\xa9\x47\x5e\x0c\x8e\xc3\xb2\x3e\xf8\x2f\xf5\xc8\x87\x70\x30\x38\x7a\x57\x75\xec\xa7\xde\xc6\xda\x36\x79\xb9\xb2\xf6\x13\xf8\xab\x21\x59\x11\x15\xb7\x37\x9f\xc1\x5d\x0e\x7b\x05\xb8\x41\x2d\x97\xdf\x16\x4b\x64\x0d\x3b\x0a\xfb\xc5\xed\xcd\x97\x47\x90\x76\xac\x1e\x42\x5a\xf9\xc5\xcf\x6f\x9f\x7e\x3c\x82\x5d\x50\x25\xb9\x9c\x3d\x04\x5f\xbb\x2d\x6e\xbf\x7e\xaf\xe1\x7f\x05\x00\x00\xff\xff\x96\x25\x31\x16\xb8\x05\x00\x00")

func devctlShBytes() ([]byte, error) {
	return bindataRead(
		_devctlSh,
		"devctl.sh",
	)
}

func devctlSh() (*asset, error) {
	bytes, err := devctlShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "devctl.sh", size: 1464, mode: os.FileMode(493), modTime: time.Unix(1515302141, 0)}
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
	"devctl.sh": devctlSh,
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
	"devctl.sh": &bintree{devctlSh, map[string]*bintree{}},
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

