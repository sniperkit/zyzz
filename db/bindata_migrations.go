// Code generated by go-bindata.
// sources:
// db/migrations/0-pragma.sql
// db/migrations/1-initial.sql
// db/migrations/2-userchanges.sql
// db/migrations/3-event.sql
// db/migrations/4-stream-activelisteners.sql
// DO NOT EDIT!

package db

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

var _dbMigrations0PragmaSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\xcd\xaa\x83\x30\x10\x46\xf7\x79\x8a\x6f\x7f\xc9\x13\x04\x17\x17\x74\xe7\xf2\x5e\xba\x94\x69\x1d\xd3\xf8\x33\x13\x92\x88\xf8\xf6\xc5\x82\x50\x84\xae\xbf\x73\x66\x8e\xb5\xf8\x59\x82\x4f\x54\x18\xff\x11\xa2\x25\x91\x64\x7a\x94\xa0\x62\xac\x05\x0b\xdd\x67\xc6\xa0\x89\x83\x17\x4c\xbc\x23\xaf\x31\x6a\x2a\x26\x26\xf2\x0b\x9d\x53\x37\xf1\x9e\x51\x41\xc5\x1d\x5e\xe6\x82\x51\x93\xd0\x8c\x45\x7b\x46\x51\x6c\x29\x14\x06\x3d\x99\x7a\xcc\xea\x4f\x1f\xa3\xae\x07\xd7\xbd\xb9\x0a\xb7\xdf\xd6\x19\xf3\xd9\x55\xeb\x26\x97\xb2\x2f\xbf\x87\xc1\x99\x73\xbb\x9c\xad\x9b\xb6\xf9\x6b\xdc\x2b\x00\x00\xff\xff\x07\xb2\xc4\x19\xf0\x00\x00\x00")

func dbMigrations0PragmaSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations0PragmaSql,
		"db/migrations/0-pragma.sql",
	)
}

func dbMigrations0PragmaSql() (*asset, error) {
	bytes, err := dbMigrations0PragmaSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/0-pragma.sql", size: 240, mode: os.FileMode(420), modTime: time.Unix(1487862705, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrations1InitialSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x92\xcf\x4e\xf4\x30\x0c\xc4\xef\x79\x0a\x1f\xbf\x4f\xb0\x12\xf7\x5e\x79\x05\xce\x91\xdb\x0c\x5b\x8b\xfc\x93\xe3\xc0\xee\xdb\xa3\xb6\x08\x75\x61\x7b\xe5\x16\xf9\x37\x96\xc7\x13\x9f\x4e\xf4\x90\xe4\xac\x6c\xa0\x97\xea\x26\xc5\xf2\x32\x1e\x23\xa8\x99\x82\x93\x6f\xd0\x77\xe8\x3f\x47\x44\x24\x81\x24\x1b\xce\x50\xaa\x2a\x89\xf5\x4a\x6f\xb8\x12\x77\x2b\x92\x27\x45\x42\xb6\xc7\x55\xd9\xe6\xa2\xe6\x25\x90\xe1\x62\x94\x8b\x51\xee\x31\x6e\x2c\x73\xc2\xbd\xfa\x5c\x9a\x1d\xb1\x65\xac\x66\x8e\x5e\xea\x3d\x8c\xcb\x01\x76\xff\x07\x77\xbb\x56\x6f\x7f\xb3\x4d\x40\x9b\x54\xaa\x49\xc9\x3f\xc8\x66\x27\x78\x36\x0a\x6c\x30\x49\xf8\x56\x50\xc0\x2b\xf7\x68\x34\x75\x55\x64\xf3\x0b\x6d\xc6\xa9\x6e\xbd\xb5\x8f\x51\xda\x8c\xd5\xfa\xef\xae\xa7\x2f\xbf\x7d\x5c\x86\x8f\x87\xb2\x35\x16\xb7\xff\xfd\xe7\xf2\x91\x5d\xd0\x52\x77\x31\x0d\xfb\xc2\xcd\x39\x0c\x9f\x01\x00\x00\xff\xff\x17\xa3\x19\xd1\x39\x02\x00\x00")

func dbMigrations1InitialSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations1InitialSql,
		"db/migrations/1-initial.sql",
	)
}

func dbMigrations1InitialSql() (*asset, error) {
	bytes, err := dbMigrations1InitialSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/1-initial.sql", size: 569, mode: os.FileMode(420), modTime: time.Unix(1489040733, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrations2UserchangesSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcf\x41\x0a\xc2\x30\x10\x85\xe1\x7d\x4f\x31\x7b\xe9\x09\xba\xf5\x0a\xae\x65\x9a\x3c\x25\x34\x99\x94\xe9\x44\x3d\xbe\x08\x2a\x11\xc5\x64\xff\xf3\x1e\xdf\x38\xd2\x2e\x85\xb3\xb2\x81\x0e\xeb\xc0\xd1\xa0\x64\x3c\x47\x50\xd9\xa0\xc4\xde\x93\xcb\xb1\x24\x21\x24\x0e\x91\x0c\x37\x9b\xfe\x76\x12\xdc\x22\x9c\xd0\x91\xf2\x85\x8d\xb5\x68\xcf\xec\x69\x0e\xbe\x67\xd2\x39\x6c\xdb\xd1\xf2\x02\x79\xe6\x43\x8d\xdc\xe7\xab\x7c\x0f\x78\xcd\xeb\x87\xf3\xc7\x49\xdd\xbc\x8c\x8d\xec\xed\x6b\x74\x0f\x5b\x6b\xaa\x72\x4d\xf7\x00\x00\x00\xff\xff\x74\x40\x31\x50\xb6\x01\x00\x00")

func dbMigrations2UserchangesSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations2UserchangesSql,
		"db/migrations/2-userchanges.sql",
	)
}

func dbMigrations2UserchangesSql() (*asset, error) {
	bytes, err := dbMigrations2UserchangesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/2-userchanges.sql", size: 438, mode: os.FileMode(420), modTime: time.Unix(1489040733, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrations3EventSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x54\x31\x0e\xdb\x30\x0c\xdc\xfd\x0a\x8e\x31\xda\x00\xdd\xb3\xf6\x0b\x9d\x05\xd9\xba\x38\x42\x2d\xca\xa0\xa8\xa4\xf9\x7d\x21\x2b\x6e\xed\xc4\xe9\x90\xa1\x9b\x71\xc7\xb3\x78\x47\x4a\xc7\x23\x7d\x09\x7e\x10\xab\xa0\x1f\x53\xd3\x0b\xca\x97\xda\x6e\x04\xe1\x0a\xd6\x43\x43\x44\xe4\x1d\x79\x56\x0c\x10\x9a\xc4\x07\x2b\x77\xfa\x89\x3b\xd9\xac\xd1\x73\x2f\x08\x60\xfd\x3a\x57\xa6\x4b\x14\x35\xde\x91\xe2\x97\x12\x47\x25\xce\xe3\x58\x39\xb6\x01\x7b\xb8\x43\xea\xc5\x4f\xea\x23\x3f\x31\xb5\x1f\x67\xac\x92\xb3\x0a\xf5\x01\x7f\x2a\xc8\xe1\x6c\xf3\xa8\xd4\x67\x11\xb0\x9a\xc2\x26\xb5\x61\x7a\x74\xa2\x56\x74\x56\x2c\xd2\x8a\x83\xdd\x0e\x2a\x99\xd9\xf3\x60\x38\xde\x8a\xd5\x0a\x06\xab\xfd\xa5\x7a\x5f\x01\x59\x46\x4a\x2a\x9e\x87\xa6\x3d\x35\xdb\xcc\x92\x0a\x6c\xf8\xdf\xa1\xed\xd0\xb3\xfb\x0f\x93\x03\xbb\xcf\x33\xd7\x9c\x8a\xed\x57\xc9\xb7\x47\x49\xee\x4a\xe3\x1d\xc4\xf4\x31\xb3\xfe\xb3\x78\xca\xdd\xe8\xd3\xc5\x94\xc8\xf7\x4c\x2e\xff\x7a\x5b\x30\x8f\xc3\x24\xc8\x15\x62\xea\x40\xf6\x56\x2c\xbe\x21\x55\x2c\xa7\xa9\x0c\xe7\xcd\x01\xf3\x1d\x29\xda\xba\x10\x4f\xec\x39\x0a\xfc\xc0\x65\xec\x87\xbf\xe7\xb4\x24\x38\x43\xc0\x3d\x12\xe5\x04\x39\x78\xd7\xbe\x0a\x9e\x7b\xdf\xc8\x36\xe4\xbe\x7e\x69\x6d\xa3\xab\x77\x7a\xd9\xb7\xb6\x2c\xf0\xfa\x09\xf8\x1e\x6f\xdc\x38\x89\xd3\xfa\x09\x38\xad\x91\x7a\xf0\xe9\x77\x00\x00\x00\xff\xff\x0b\x69\x10\xc3\x38\x04\x00\x00")

func dbMigrations3EventSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations3EventSql,
		"db/migrations/3-event.sql",
	)
}

func dbMigrations3EventSql() (*asset, error) {
	bytes, err := dbMigrations3EventSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/3-event.sql", size: 1080, mode: os.FileMode(420), modTime: time.Unix(1489266434, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrations4StreamActivelistenersSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xcc\xb1\x0d\x02\x31\x0c\x05\xd0\x3e\x53\xfc\x1e\x9d\x44\x7f\x2d\x2b\x50\x23\x73\x31\x28\x92\x63\x47\xce\x0f\xac\x4f\x4b\x01\x0b\xbc\x6d\xc3\xa9\xb7\x67\x0a\x15\xd7\x51\xc4\xa8\x09\xca\xdd\x14\x93\xa9\xd2\x8b\xd4\x8a\x23\x6c\x75\x87\x1c\x6c\x2f\xbd\x59\x9b\x54\xd7\x9c\x68\x4e\x78\x10\xbe\xcc\x50\xf5\x21\xcb\x88\xf3\x5e\xca\x37\x7b\x89\xb7\xff\x82\x6b\xc6\xf8\x27\xef\x9f\x00\x00\x00\xff\xff\xd7\x36\xcd\xf0\x98\x00\x00\x00")

func dbMigrations4StreamActivelistenersSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations4StreamActivelistenersSql,
		"db/migrations/4-stream-activelisteners.sql",
	)
}

func dbMigrations4StreamActivelistenersSql() (*asset, error) {
	bytes, err := dbMigrations4StreamActivelistenersSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/4-stream-activelisteners.sql", size: 152, mode: os.FileMode(420), modTime: time.Unix(1489843670, 0)}
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
	"db/migrations/0-pragma.sql": dbMigrations0PragmaSql,
	"db/migrations/1-initial.sql": dbMigrations1InitialSql,
	"db/migrations/2-userchanges.sql": dbMigrations2UserchangesSql,
	"db/migrations/3-event.sql": dbMigrations3EventSql,
	"db/migrations/4-stream-activelisteners.sql": dbMigrations4StreamActivelistenersSql,
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
	"db": &bintree{nil, map[string]*bintree{
		"migrations": &bintree{nil, map[string]*bintree{
			"0-pragma.sql": &bintree{dbMigrations0PragmaSql, map[string]*bintree{}},
			"1-initial.sql": &bintree{dbMigrations1InitialSql, map[string]*bintree{}},
			"2-userchanges.sql": &bintree{dbMigrations2UserchangesSql, map[string]*bintree{}},
			"3-event.sql": &bintree{dbMigrations3EventSql, map[string]*bintree{}},
			"4-stream-activelisteners.sql": &bintree{dbMigrations4StreamActivelistenersSql, map[string]*bintree{}},
		}},
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

