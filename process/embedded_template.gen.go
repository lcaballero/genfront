// Code generated by go-bindata.
// sources:
// .files/dump_struct.fm
// .files/req_rest_methods.fm
// .files/struct_sql_tomap.fm
// DO NOT EDIT!

package process

import (
	"fmt"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name string
	size int64
	mode os.FileMode
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

var _dump_structFm = []byte(`{{ range .names }}
{{ . }}
{{ end }}`)

func dump_structFmBytes() ([]byte, error) {
	return _dump_structFm, nil
}

func dump_structFm() (*asset, error) {
	bytes, err := dump_structFmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dump_struct.fm", size: 36, mode: os.FileMode(420), modTime: time.Unix(1452553056, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _req_rest_methodsFm = []byte(`---
methods:
  - OPTIONS
  - GET
  - HEAD
  - POST
  - PUT
  - DELETE
  - TRACE
  - CONNECT
---
package {{ .GOPACKAGE }}
{{ .GEN_TAGLINE }}
// {{ getenv "GOLINE" }}

const (
{{ range .methods }}	{{ . }} = "{{ . }}"
{{ end }})

// Methods for the Rest state{{ range .methods }}
func (r *Rest) {{ . | title }}() *Rest {
	return r.Method({{ . }})
}{{ end }}

// Methods for Req state{{ range .methods }}
func (r *Req) {{ . | title }}() *Req {
	return r.Method({{ . }})
}{{ end }}
`)

func req_rest_methodsFmBytes() ([]byte, error) {
	return _req_rest_methodsFm, nil
}

func req_rest_methodsFm() (*asset, error) {
	bytes, err := req_rest_methodsFmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "req_rest_methods.fm", size: 477, mode: os.FileMode(420), modTime: time.Unix(1452550578, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _struct_sql_tomapFm = []byte(`package {{ .GOPACKAGE }}
{{ .GEN_TAGLINE }}

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

// Methods for the Rest state
func (r *{{ .structName }}) ToMap() map[string]interface{} {
  m := make(map[string]interface{}){{ range .names }}
  m["{{ . }}"] = r.{{ . }}{{ end }}
  return m
}

// Private function (to this fil).  Don't use this function
// directly, it's intended for use only in this file.
func __snakeToPascal(sk string) string {
	parts := strings.Split(sk, "_")
	for i,p := range parts {
		parts[i] = strings.Title(p)
	}
	return strings.Join(parts, "")
}

// For known named parameter types, this strips known prefixes.
func __removePrefix(prop string) string {
	hasPrefix := strings.HasPrefix(prop, "$") ||
		strings.HasPrefix(prop, ":") ||
		strings.HasPrefix(prop, "@") ||
		strings.HasPrefix(prop, "?")
	if hasPrefix {
		return prop[1:]
	}
	return prop
}

// Returns an array of values from the {{ .structName }} instance as designated
// in the props array.  The string of the props array should conform to the
// possible named value syntax which sqlite accepts.
// (See: https://www.sqlite.org/c3ref/bind_parameter_name.html and
// https://www.sqlite.org/c3ref/bind_blob.html)
func (e *{{ .structName }}) Parameters(props []string) (rs []interface{}, err error) {
	rs = make([]interface{}, len(props))
	for i,p := range props {
		p = __removePrefix(p)
		switch p { {{ range .names }}
		case "{{ . }}":
			rs[i] = e.{{ . }}{{ end }}
		default:
			err = errors.New(fmt.Sprintf("{{ .structName }} doesn't have a property named: %s", p))
		}
	}
	return rs, err
}

// Fills pointer array with pointers to receiver fields.
func (e *{{ .structName }}) FromColumns(cols []string, ptrs []interface{}) error {
	if len(cols) != len(ptrs) {
		return errors.New("Column length doesn't equal pointer array length")
	}
	snakeToPascal := __snakeToPascal
	for i,c := range cols {
		pascal := snakeToPascal(c)
		switch pascal { {{ range .names }}
		case "{{ . }}":
			ptrs[i] = &e.{{ . }}{{ end }}
		default:
			return errors.New(fmt.Sprintf(
				"Provided a column that doesn't exist in structure: %s",
				c))
		}
	}
	return nil
}

func To{{ .structName }}Rows(rows *sql.Rows) ([]*{{ .structName }}, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	ptrs := make([]interface{}, len(cols))
	w := make([]*{{ .structName }}, 0)

	for rows.Next() {
		e := &{{ .structName }}{}
		if err := e.FromColumns(cols, ptrs); err != nil {
			return nil, err
		}
		if err := rows.Scan(ptrs...); err != nil {
			return nil, err
		}
		w = append(w, e)
	}
	return w, nil
}`)

func struct_sql_tomapFmBytes() ([]byte, error) {
	return _struct_sql_tomapFm, nil
}

func struct_sql_tomapFm() (*asset, error) {
	bytes, err := struct_sql_tomapFmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "struct_sql_tomap.fm", size: 2595, mode: os.FileMode(420), modTime: time.Unix(1452550503, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"dump_struct.fm": dump_structFm,
	"req_rest_methods.fm": req_rest_methodsFm,
	"struct_sql_tomap.fm": struct_sql_tomapFm,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"dump_struct.fm": &bintree{dump_structFm, map[string]*bintree{
	}},
	"req_rest_methods.fm": &bintree{req_rest_methodsFm, map[string]*bintree{
	}},
	"struct_sql_tomap.fm": &bintree{struct_sql_tomapFm, map[string]*bintree{
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

