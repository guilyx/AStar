// Code generated by "esc -o bindata/esc.go -pkg=bindata templates"; DO NOT EDIT.

package bindata

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return []os.FileInfo(fis[0:limit]), nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/templates/call.tmpl": {
		name:    "call.tmpl",
		local:   "templates/call.tmpl",
		size:    241,
		modtime: 1540446832,
		compressed: `
H4sIAAAAAAAC/0SOQWrDQAxFryKMFy0YHaDQA3hTSlvatRjLrsCeFo2SEITuHsY4mdWHP2/el/vEs2SG
LtG6dhHuF7FfwA9OLGfW2sgM+c8Ax/JpekoWYYbunKf6eicBI1qLb7RxxJO7Ul4Yehmg5xVeXgHfSWlj
Yy2HvZeIAR5/296PitUbzJB0KU2/K+riTuPX9Z9xLN+kQpOkCMTG7vF85C0AAP//ZQi8iPEAAAA=
`,
	},

	"/templates/function.tmpl": {
		name:    "function.tmpl",
		local:   "templates/function.tmpl",
		size:    2392,
		modtime: 1540483824,
		compressed: `
H4sIAAAAAAAC/7RWTW/jNhA9S79i1sgupMJh7g58aJC06KFx4QTNoSgKRh65RGnKJUcJDIL/vSBFfVpO
c9kcImlIznvz5g0Ta3dYCoWwKGtVkKjUwrnU2mu4KmG1BuZcmvolsJY9o6FHfkDnMoIfCA0JtWfPOdg0
8UfeBf0NbIsFijfUzqVJCIsS2C/miXRdUAh20Z8Eyp1pYgmdjghliIAJm33euFtztcfJgcTa8O1JBnqn
I8YlfwTVLn51mG1o8D559ax8mb9xzQ9IqANYoMb1fkRsQOv8RAAMoTN2A8QxvhfUeNH/+HMAo/gBPaxQ
+3h4RuaWO1e7XuuJXFHa5tEpIk2vWZvyXNAL4n0gWZIEvfyvmTMD3bZoakmmxXnhij6SrIPcItVamQet
q6jBO1f0oDW8VpWc6OyFvLmB5839ZgU/7nbgtYaCGzQstKGsNFgrSsgqDeypfm2akamKvKCP/B/c5blz
8NcSiHyTrA3JYynNdptC/GlZdpmcI7atVUbEfEeX4IdqOkYQOcN1L/tcty+M1VkPA83Qn9MRw2aunfsW
qUeF2e9c1uicbVNcmLbEWtZM/wqIWOMjNpjBZZ+gn71kZiDPPiLezAi1Zb5oQV31o9FareHb64nQsLu6
LFHbzwDGUWnau1HyNHRTfh7fKAwq5dAxIzwcJSeEhW4cvICrMvi2Xym4lE34EosZGyeijF2bEnMOUOum
q3Mgt50pM7/vyxqUkLl/ErF2OmKbiYWUZbYY5jqgMXyPsRT0O2ANX9+W0B7/+rZYjuCFOtZd8aj1cgCW
945ob4nRuIe1yZjoULC1zc1UVIqEqjEWNuewDy11DnnJU0H1nyvqB6fzGHsK12+W3w62NKoOL6zed9Jg
xLjjRhSDP0xdc6/KOX/5oRxxGOoshcJpoz/N5zvhf9FYSiyI3SMeH/6tucy6DMsxoXzIqOveZ3zYEo5k
f60liaMckY18eq/+j1Evkrz8D8PEp+ALGl7XLnVp2vr0vwAAAP//X+Qs81gJAAA=
`,
	},

	"/templates/header.tmpl": {
		name:    "header.tmpl",
		local:   "templates/header.tmpl",
		size:    142,
		modtime: 1540481163,
		compressed: `
H4sIAAAAAAAC/0TMMQ7CMAyF4d2nsDrBQC7BxIK4gkUebYXiViGb9e6OlAi6/bL1voiM1+rQaYFl1ImU
iGo+Q9N1KwXePmRE6g941gspuz3fNkMj0mMkKbKWfatNT4dw65cB3K2AHJO2/DhSzv/6BgAA///GzMM9
jgAAAA==
`,
	},

	"/templates/inline.tmpl": {
		name:    "inline.tmpl",
		local:   "templates/inline.tmpl",
		size:    49,
		modtime: 1540446006,
		compressed: `
H4sIAAAAAAAC/6quTklNy8xLVVDKzMvJzEtVqq1VqK4uSc0tyEksSVVQSk7MyVFS0AOLpual1NYCAgAA
//+q60H/MQAAAA==
`,
	},

	"/templates/inputs.tmpl": {
		name:    "inputs.tmpl",
		local:   "templates/inputs.tmpl",
		size:    152,
		modtime: 1540446821,
		compressed: `
H4sIAAAAAAAC/0yNMQoCQQxFrxKWLSUHEDyAneAJIptZptgomWz1yd1lRoupEh7/vw9sWqopLdU+Z7Ql
E1gLXW/E/a2F7B3Ez/MV2qJlRrDJoRcC1LZ/Zi388GpxH5IOXWzXwcXl0FD/dcX3xsCgfWLyzOcbAAD/
/468z9qYAAAA
`,
	},

	"/templates/message.tmpl": {
		name:    "message.tmpl",
		local:   "templates/message.tmpl",
		size:    201,
		modtime: 1540446006,
		compressed: `
H4sIAAAAAAAC/zyN4WqDQBCE//sUiyi0oPsAhT5A/xRpS/9f4mgW9GLuTkNY9t2DB/HXDDPDN6o9BvGg
ckaMbkRJrVmhKgP5ayL+XU8JMUWz+sakCt+bqd4lXYh/cIZsCHvCf48F/O+mFWZ8DPnbzTB7y0Tugvj0
5Zd1B6oG50dQJQ1VmOjjk7hzwc1ICLmXgSoxa16/9XZws7wXqi1l+wwAAP//kC65UskAAAA=
`,
	},

	"/templates/results.tmpl": {
		name:    "results.tmpl",
		local:   "templates/results.tmpl",
		size:    168,
		modtime: 1540446006,
		compressed: `
H4sIAAAAAAAC/1yNTQrCQAyFr/Iosyw9gOBS3HsDoRkJlAy8ma5C7i6pRcFVfr4vee6rVDXBROn7NvoU
AXc+7SUoOqPIhssVy+ODI9y1omjEDHexNTf3NrBkc85a82DstH4jG1MW8uQ4hMbv0385A3/uUd8BAAD/
/7BPz2GoAAAA
`,
	},

	"/templates": {
		name:  "templates",
		local: `templates`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"templates": {
		_escData["/templates/call.tmpl"],
		_escData["/templates/function.tmpl"],
		_escData["/templates/header.tmpl"],
		_escData["/templates/inline.tmpl"],
		_escData["/templates/inputs.tmpl"],
		_escData["/templates/message.tmpl"],
		_escData["/templates/results.tmpl"],
	},
}