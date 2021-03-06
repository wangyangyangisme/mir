// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/chi_iface.tmpl (1.521kB)
// templates/echo_iface.tmpl (1.032kB)
// templates/gin_iface.tmpl (1.04kB)
// templates/httprouter_iface.tmpl (1.486kB)
// templates/iris_iface.tmpl (1.077kB)
// templates/macaron_iface.tmpl (1.641kB)
// templates/mux_iface.tmpl (1.557kB)

package generator

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _chi_ifaceTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x53\xbd\x6e\xe3\x3c\x10\xec\xf9\x14\x0b\x57\xd2\x87\x84\xea\x03\xa4\xf8\x90\xdf\x2b\xfc\x03\xc7\x87\xab\x19\x6a\x2d\x11\xb1\x48\x1d\xb9\x4a\x60\x10\x7c\xf7\x03\x29\x39\x96\x03\xfb\x02\x5d\x65\x73\xb9\x3b\xb3\x33\x1c\x15\x05\xdc\x99\x12\xa1\x42\x8d\x56\x10\x96\xf0\xba\x87\xca\x5c\x37\xca\x72\xb8\x5f\xc2\x62\xb9\x81\x87\xfb\x1f\x1b\xce\x58\x2b\xe4\x9b\xa8\x10\xbc\x07\xbe\x7a\xab\x16\xa2\x41\x08\x81\x31\xd5\xb4\xc6\x12\x64\x0c\x00\x60\xa6\x91\x8a\x9a\xa8\x9d\xb1\xfe\x5c\x29\xaa\xbb\x57\x2e\x4d\x53\x54\xe6\x5a\xd6\xaa\x90\xb5\x9a\xb1\x9c\x31\xef\xd5\x16\xb4\xa1\x87\xa6\xa5\xfd\x0b\x59\xe0\x77\xa6\x69\x50\x13\x84\x50\x14\xe0\xfd\xe1\x1c\x82\xf7\xa8\xcb\x10\x18\xed\xdb\xc8\xcf\x37\xfb\x16\x23\x7f\x08\xa0\x34\xa1\xdd\x0a\x89\xe0\xcf\x21\xd6\x42\x69\x08\x21\xae\x12\xa5\xa6\x63\x6b\xcd\xbb\x2a\x11\x1a\x55\x96\x3b\xfc\x10\x16\x1d\x6c\x8d\x05\x59\xab\xb4\x73\x64\x8e\x8d\x21\x64\x79\x2c\xf2\xf9\xb1\x91\x1d\x56\xf1\xde\x0a\x5d\x21\xf0\x47\x85\xbb\xd2\xc5\x1d\x2f\xcb\x19\xe8\xc7\x92\x0e\x4c\x73\xa4\xda\x94\xbd\x98\x2c\x1a\xc7\xd7\xe8\x5a\xa3\x1d\xfe\xb2\x8a\xd0\x5e\xc1\x7f\x43\xf5\x77\x87\x8e\x72\xef\x71\xe7\xb0\x87\xfc\xb7\xf1\x61\xfd\xfe\x37\x30\x56\x14\xb0\xc6\x4a\x39\x42\x7b\x62\xed\x0b\xda\x77\xa1\x09\xec\x70\xf9\xc5\x78\x37\x5c\x93\x49\xc6\x6d\x3b\x2d\xff\x8a\x93\x25\x83\xf9\xda\x74\x69\x2d\x77\x0a\x97\x9f\x7d\xbe\x27\x6b\xba\x76\xf0\xcf\xa6\x41\xb8\xb9\x4d\x30\x0b\xfc\xe8\x91\xb2\x3c\x39\x69\xf9\xdc\x74\x9a\xb2\x99\xf7\xfd\x54\x08\xb3\xab\x61\xe6\xc4\xb3\x23\x8c\x3d\x7a\x71\x29\x36\x6c\x78\xb8\xce\x21\xc8\x54\x8b\x39\xe9\x21\xd2\xdd\x38\x42\x37\xb7\xe0\xf8\x38\x3b\xec\xc8\xc7\x7f\x3a\xcc\x46\xcd\x9c\xf3\xfc\xf3\x0d\x06\x92\x4f\x9f\xd3\x88\x03\xa5\xb7\x26\xba\x3b\xd0\x5d\x0c\xdc\x33\x51\xfb\xbf\xde\x03\x8f\x7f\xfa\x3c\x9c\x58\x36\x64\xe4\xb1\xd3\x32\xd9\x73\x6c\x4b\x1e\xc5\xd2\x4a\x50\x9d\x0e\x49\xc0\x38\x53\x67\xbc\xe3\xcf\x28\xca\xec\x9b\xb1\xb1\xf6\x27\xa4\x29\xed\x2b\xe3\xa6\xf5\x77\x93\xda\xef\x71\x87\x84\x93\x08\x04\xc9\x7a\xca\xc0\xb2\x25\x65\xb4\x9b\x32\x72\x67\xb4\x46\x39\x49\xc8\xc6\x0a\xf9\xad\x8e\x2f\xdf\x7b\x60\x7f\x02\x00\x00\xff\xff\x49\x65\xd4\xe6\xf1\x05\x00\x00"

func chi_ifaceTmplBytes() ([]byte, error) {
	return bindataRead(
		_chi_ifaceTmpl,
		"chi_iface.tmpl",
	)
}

func chi_ifaceTmpl() (*asset, error) {
	bytes, err := chi_ifaceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "chi_iface.tmpl", size: 1521, mode: os.FileMode(0664), modTime: time.Unix(1584082042, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x6c, 0x6, 0x66, 0xa7, 0xe9, 0x9d, 0xe, 0x9e, 0xb4, 0xe, 0x6, 0x95, 0x91, 0x0, 0xe0, 0xca, 0x3f, 0x77, 0xb9, 0x7b, 0x68, 0xf2, 0x78, 0xa2, 0xd7, 0xd0, 0xe6, 0x16, 0x8e, 0x8, 0xf6, 0xc7}}
	return a, nil
}

var _echo_ifaceTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xcf\x8e\xd3\x30\x10\xc6\xef\x7e\x8a\x51\x4f\x09\x5a\x9c\x0b\x27\x24\x0e\x55\xb7\x0b\x1c\xf6\x8f\xd8\x72\x42\x1c\xdc\x78\x9a\x58\x6d\xec\xc8\x9e\x14\x22\xcb\xef\x8e\x6c\xa7\xdd\x74\xb5\x05\x4e\xb1\xc7\x93\x6f\xbe\xf9\xcd\x54\x15\xac\x8c\x44\x68\x50\xa3\x15\x84\x12\xb6\x23\x34\xe6\x7d\xa7\x2c\x87\xdb\x47\x78\x78\xdc\xc0\xfa\xf6\xeb\x86\x33\xd6\x8b\x7a\x2f\x1a\x04\xef\x81\x3f\xed\x9b\x07\xd1\x21\x84\xc0\x98\xea\x7a\x63\x09\x0a\x06\x00\xb0\x68\x14\xb5\xc3\x96\xd7\xa6\xab\x0e\x62\xeb\x48\xd4\xfb\x0a\xeb\xd6\x54\xc7\x0f\x0b\x56\x32\xe6\xbd\xda\x81\x36\xb4\xee\x7a\x1a\x9f\xc9\x02\x5f\x99\xae\x43\x4d\x10\x42\x55\x81\xf7\xa7\x7b\x08\xde\xa3\x96\x21\x30\x1a\xfb\x58\x95\x6f\xc6\x1e\x63\xd5\x10\x40\x69\x42\xbb\x13\x35\x82\x7f\x4b\xb1\x15\x4a\x67\xbd\x7c\xec\xad\x39\x2a\x89\xd0\x0a\x2d\x0f\x68\x1d\xd4\x29\xbc\x33\x16\xa2\xb7\xe4\x3c\x56\x8e\xd1\x10\x8a\x12\x7e\xfc\x8c\x71\x7e\xaf\xa4\x3c\xe0\x2f\x61\xf1\x6e\xd0\x35\x3b\x19\xf2\xde\x0a\xdd\x20\xf0\x3b\x85\x07\xe9\xa2\xd3\xeb\x4d\x45\xed\x57\x8d\x9d\xea\xdd\x23\xb5\x46\xe6\x96\x8a\x54\x70\x65\x34\xe1\x6f\x2a\x01\xad\x35\xd6\x7b\x3c\x38\xcc\x12\xff\x97\x3e\xd9\xcb\xdf\xc0\x58\x55\xc1\x37\x6c\x94\x23\xb4\x17\x00\x9f\xd1\x1e\x85\x26\xb0\xd3\xe3\x2b\xbc\x6e\x7a\x26\x93\xf9\xec\x06\x5d\xff\x55\xa8\x40\x78\x97\x1c\xad\xeb\xd6\xdc\x80\xbb\xd4\x2b\xdf\x9c\xd2\x67\x6b\x86\x7e\x02\xd4\xc0\xc7\x4f\x80\x39\x54\x2c\xbc\xcf\xa7\x10\x16\xe5\x9c\x41\xce\x7a\x69\xf3\xda\xdc\x27\xe4\x83\xc3\xd9\xa0\xad\x19\x08\x6d\x42\xdf\x9d\xc7\xea\xa2\xa2\xe3\xf3\xd9\xa7\x8c\x86\x7f\x77\x58\xcc\xf2\x38\xe7\xe5\x19\xec\xa4\x7f\x86\x97\xa4\x1d\x28\xbd\x33\x11\xd9\x54\xe9\xea\x96\x7c\x21\xea\x97\x7a\x04\x1e\x0f\x79\xa8\x27\x0c\x7c\x29\x65\xea\xff\xe5\x29\x84\xc5\x0d\xc4\xd0\x93\xa0\x36\x5d\x92\xdf\xf9\x32\x5c\x42\xe2\x4b\x3d\x16\xff\xfe\x61\xbe\x2a\x81\xfd\x09\x00\x00\xff\xff\x94\x6c\xa7\x53\x08\x04\x00\x00"

func echo_ifaceTmplBytes() ([]byte, error) {
	return bindataRead(
		_echo_ifaceTmpl,
		"echo_iface.tmpl",
	)
}

func echo_ifaceTmpl() (*asset, error) {
	bytes, err := echo_ifaceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "echo_iface.tmpl", size: 1032, mode: os.FileMode(0644), modTime: time.Unix(1584106466, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x37, 0x4a, 0x24, 0x80, 0x2, 0x1, 0x7a, 0xe7, 0xc7, 0x69, 0xfe, 0xfe, 0x94, 0xe9, 0xc8, 0x41, 0x3, 0xe2, 0x10, 0x67, 0x2, 0x85, 0x15, 0x9d, 0x7d, 0x29, 0xd9, 0x20, 0x63, 0x29, 0x93, 0x80}}
	return a, nil
}

var _gin_ifaceTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\xc1\x6e\xdb\x30\x0c\xbd\xeb\x2b\x08\x9f\xec\xa1\x95\xef\x03\x76\x28\xd2\x6c\xdd\x61\x6d\xb1\x66\x1f\xa0\xda\xb4\x2c\x34\xa6\x0c\x49\xee\x66\x08\xfa\xf7\x41\xb2\xd2\x38\x45\xd2\x9e\x4c\x53\xe4\x7b\x8f\x8f\xac\x6b\xd8\xe8\x16\x41\x22\xa1\x11\x0e\x5b\x78\x9e\x41\xea\xeb\x41\x19\x0e\xb7\x0f\x70\xff\xb0\x83\xed\xed\xcf\x1d\x67\x6c\x14\xcd\x8b\x90\x08\xde\x03\x7f\x7c\x91\xf7\x62\x40\x08\x81\x31\x35\x8c\xda\x38\x28\x19\x00\x40\x21\x95\xeb\xa7\x67\xde\xe8\xa1\x96\x8a\xae\xa5\x26\xd5\xc4\xa8\x60\x15\x63\xde\xab\x0e\x48\xbb\xed\x30\xba\xf9\xc9\x19\xe0\x1b\x3d\x0c\x48\x0e\x42\xa8\x6b\xf0\xfe\xf0\x1f\x82\xf7\x48\x6d\x08\xcc\xcd\x63\xa4\xe4\xbb\x79\xc4\x48\x19\x02\x28\x72\x68\x3a\xd1\x20\xf8\x73\x88\xbd\x50\xb4\xe0\x2d\xe1\x68\xf4\xab\x6a\x11\x7a\x41\xed\x1e\x8d\x85\x26\xa5\x3b\x6d\x40\x2a\x4a\xaa\x23\x71\x4c\x86\x50\x56\x31\xc9\xef\x72\x6d\xca\xb2\x83\x16\xef\x8d\x20\x89\xc0\xbf\x2b\xdc\xb7\x36\x8a\xbc\x3c\x4f\xc4\x7d\x37\xd3\x81\xeb\x17\xba\x5e\xb7\xcb\x34\xe5\x97\xc8\xb7\xd1\xe4\xf0\x9f\xab\xbc\xc7\xbd\xc5\xa5\xf9\xb3\xc2\x2c\x69\xf9\x06\xc6\xea\x1a\x7e\xa3\x54\xd6\xa1\x39\xf1\xeb\x09\xcd\xab\x20\x07\x26\x3f\xbe\x73\xd3\xe6\x67\xa7\x93\x1d\xdd\x44\xcd\x87\x38\x25\x42\x52\xb2\x25\xa9\x08\xaf\xc0\x9e\xe2\x55\x67\x97\xf2\xc3\xe8\x69\xcc\xa6\x18\x3d\x45\x15\x5f\xbf\x01\x2e\xf9\xb2\xf0\x7e\x89\x42\x28\x4e\x2c\x58\x95\x1e\x07\xbe\xb4\xf0\x6c\xf8\x64\x71\xb5\xe1\x05\x21\x19\x3f\xa8\xb6\xdd\xe3\x5f\x61\xd0\x46\x44\xcb\xd7\x5b\x67\x47\x3a\xfe\xc7\x62\xb9\x2a\xe6\x9c\x57\x6f\x3e\x67\x92\x37\x2f\x53\x8b\x05\x45\x9d\x8e\x0e\x66\xba\x8b\x87\x72\xe7\xdc\x78\x43\x33\xf0\x18\x2c\xdb\x3d\x71\x25\x1f\x5e\x72\xe4\x58\x12\x42\x71\x05\x31\xf5\x28\x5c\x9f\x7e\x92\xf8\xf5\x75\x9c\xb1\x8d\xdf\xd0\x5c\x7e\xde\xb5\x3e\xa3\xc0\xfe\x07\x00\x00\xff\xff\x20\xfd\x7a\xfb\x10\x04\x00\x00"

func gin_ifaceTmplBytes() ([]byte, error) {
	return bindataRead(
		_gin_ifaceTmpl,
		"gin_iface.tmpl",
	)
}

func gin_ifaceTmpl() (*asset, error) {
	bytes, err := gin_ifaceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gin_iface.tmpl", size: 1040, mode: os.FileMode(0664), modTime: time.Unix(1584082042, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xfe, 0x89, 0x2a, 0xf1, 0x17, 0x71, 0xe9, 0xe0, 0x19, 0xf6, 0x3, 0x38, 0xb3, 0xb2, 0xb1, 0xf7, 0x1a, 0x7c, 0x40, 0x27, 0xfc, 0x42, 0x80, 0xf6, 0x4b, 0xbd, 0x4e, 0x7c, 0xe0, 0x35, 0x59, 0xa7}}
	return a, nil
}

var _httprouter_ifaceTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\xc1\x6e\xdb\x30\x0c\x86\xef\x7a\x0a\xc2\xd8\x21\x19\x32\xf9\x19\x86\xa4\x6b\x76\x58\x13\xa4\x01\x76\x56\x6d\xc6\x56\x6b\x4b\x1a\x45\x0f\x30\x08\xbd\xfb\xe0\xd8\xd9\xd2\x00\x2d\xb0\xc2\x27\x99\xa6\xfc\xfd\x3f\x7f\xc2\x79\x0e\x6b\x5f\x22\x54\xe8\x90\x0c\x63\x09\x4f\x3d\x54\xfe\x4b\x6b\x49\xc3\x66\x07\x0f\xbb\x23\xdc\x6d\xbe\x1f\xb5\x52\xc1\x14\x2f\xa6\x42\x10\x01\xbd\x7f\xa9\x1e\x4c\x8b\x90\x92\x52\xb6\x0d\x9e\x18\x16\x0a\x00\x20\x73\xc8\x79\xcd\x1c\x32\x35\xd6\x95\xe5\xba\x7b\xd2\x85\x6f\xf3\xe7\xae\xb1\xe8\x62\x51\xb7\xb6\x1c\x2f\x91\xef\x18\x29\x53\x4b\xa5\x44\xec\x09\x9c\xe7\xbb\x36\x70\xff\xc8\x04\x7a\xed\xdb\x16\x1d\x43\x4a\x79\x0e\x22\x97\x3a\x25\x11\x74\x65\x4a\x8a\xfb\x30\xb8\xd1\xc7\x3e\xe0\xe0\x26\x25\xb0\x8e\x91\x4e\xa6\x40\x10\x25\x42\xc6\x55\x08\xfa\x9b\xc5\xa6\x8c\xc3\x77\x6f\x4b\x0c\x5e\x6f\x64\xce\xfe\x45\xf4\x0f\xe4\xda\x97\xa3\xc0\x62\x70\xad\x0f\x18\x83\x77\x11\x7f\x92\x65\xa4\x15\x7c\x9e\xde\xfe\xea\x30\xf2\x0a\xfe\x4d\xa6\xf7\x86\x4c\x1b\x97\x13\x0a\x9b\x88\xa3\xd4\x3c\xd8\x4b\x10\x97\x33\x29\x95\xe7\x70\xc0\xca\x46\x46\x7a\x95\xcc\x23\xd2\x6f\xe3\x18\x68\x6a\xde\xe4\x16\xa7\x36\xfb\x2b\x19\x75\xea\x5c\xf1\x2e\x6e\x41\xa3\xc9\xc9\xd6\xe1\x7c\xac\x20\xbe\xa6\x2f\xdf\x5b\xc6\x96\x39\x7c\x75\x3d\xe8\xe1\x61\xcc\x64\x5a\x07\xe9\xad\x71\x65\x83\x8b\x4c\xe4\xaa\x9b\x52\xb6\x82\x4c\xe4\xd9\x5b\xb7\x37\x5c\xc3\x27\x7d\x4f\xbe\x0b\xa0\x87\xea\xdc\x8d\xfa\x26\xdf\xe5\x75\xf6\x7f\xb9\xe7\x78\xc7\x6b\xf7\xc8\xff\x0f\x55\x6f\xd1\xf6\xdd\xac\x34\x1f\xe7\xc4\x6d\xb0\x41\xc6\x19\x81\x5b\x34\xe5\x9c\xe3\x1a\x2e\xea\x19\x79\xbb\xc0\xd6\xbb\x38\x23\x71\xed\x9d\xc3\x62\xce\x9d\x1c\xc9\x14\x1f\x58\xc9\xcd\xff\x9f\xd4\x9f\x00\x00\x00\xff\xff\xc5\xf9\xba\x0c\xce\x05\x00\x00"

func httprouter_ifaceTmplBytes() ([]byte, error) {
	return bindataRead(
		_httprouter_ifaceTmpl,
		"httprouter_iface.tmpl",
	)
}

func httprouter_ifaceTmpl() (*asset, error) {
	bytes, err := httprouter_ifaceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "httprouter_iface.tmpl", size: 1486, mode: os.FileMode(0664), modTime: time.Unix(1584082042, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x40, 0x4e, 0xdc, 0x83, 0x3b, 0xe7, 0x4c, 0x9a, 0x7, 0x4d, 0xc3, 0x91, 0x9b, 0x6e, 0x77, 0xfa, 0x30, 0xfa, 0xeb, 0x42, 0xe3, 0x19, 0x8a, 0x7d, 0x9c, 0xbe, 0x61, 0xae, 0xbe, 0x5a, 0xa8, 0x7c}}
	return a, nil
}

var _iris_ifaceTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xc1\x6e\xdb\x30\x0c\x86\xef\x7a\x0a\xc2\x27\x67\xe8\x24\x6c\xc7\x01\x3b\x04\x69\xb7\xee\xb0\x36\x58\xb3\x07\x50\x6d\xc6\x11\x62\x4b\x82\x44\x67\x33\x04\xbd\xfb\x20\xcb\x6e\x9d\xa2\xe9\x7a\xb2\x25\x52\x3f\x7f\x7e\xa4\x10\xb0\x31\x35\x42\x83\x1a\x9d\x24\xac\xe1\x71\x80\xc6\x7c\xec\x94\xe3\x70\x7d\x0f\x77\xf7\x3b\xb8\xb9\xfe\xb1\xe3\x8c\x59\x59\x1d\x65\x83\x10\x02\xf0\xed\xb1\xb9\x93\x1d\x42\x8c\x8c\xa9\xce\x1a\x47\x50\x32\x00\x80\xa2\x51\x74\xe8\x1f\x79\x65\x3a\x71\x94\x24\x9d\xf4\x42\x39\xe5\xc5\xe9\xd3\xe7\xe2\xbf\x19\xa2\x32\x9a\xf0\x2f\x15\x6c\xc5\x58\x08\x6a\x0f\xda\xd0\x4d\x67\x69\x78\x20\x07\x7c\x63\xba\x0e\x35\x41\x8c\x42\x40\x08\xf3\x39\xc6\x10\x50\xd7\x31\x32\x1a\x6c\xf2\xc7\x77\x83\xc5\xe4\x2f\x46\x50\x9a\xd0\xed\x65\x85\x10\x5e\x53\x3c\x48\xa5\xb3\x5e\xfe\xb5\xce\x9c\x54\x8d\x70\x90\xba\x6e\xd1\x79\xa8\xc6\xeb\xbd\x71\x90\x4c\x8e\x1d\xa4\xca\xe9\x36\xc6\x72\x05\x93\x63\x7e\x3b\x3d\x60\xb3\x97\x10\x9c\xd4\x0d\x02\xff\xa6\xb0\xad\x7d\x32\x79\xb9\x9f\x24\xfb\xa2\xa7\xb9\xd4\x4f\xa4\x83\xa9\x73\x37\xe5\x5c\x6d\x93\xbf\xab\x10\xb0\xf5\x98\xdf\xbf\x23\x77\x32\x96\xbf\x91\x31\x21\xe0\x17\x36\xca\x13\xba\x33\x6a\x0f\xe8\x4e\x52\x13\xb8\x29\xf8\x82\xa9\x9f\xc2\x64\x32\x94\x7d\xaf\xab\x37\x85\x4a\x69\x2d\x7c\x48\xb9\x7c\x6d\x6d\xab\x2a\x49\xca\xe8\x2b\xf0\xe7\xc2\xab\x57\x67\xf4\xdd\x99\xde\x4e\x8c\x2c\x7c\xf9\x0a\xd2\x5a\xbe\x95\x8e\x86\xb2\x08\x21\x87\x63\x2c\xce\x58\xcc\x79\xcf\x3d\x5f\x9a\xfc\x44\xbe\xf7\xb8\x18\xb5\x4d\xea\xe3\x00\x3a\x55\xd7\x2d\xfe\x91\x0e\x7d\x92\xf4\x7c\x39\xfc\x31\xc3\xf2\xdf\x1e\xcb\x45\x1e\xe7\x7c\xf5\x04\x79\x92\x7f\x02\xe9\x4c\x4f\xe8\x41\xe9\xbd\x49\xf8\x72\xa1\x8b\xab\x72\x4b\x64\xd7\x7a\x00\x9e\x7e\xf2\x70\x67\x10\xd3\xc2\x8d\x08\x9e\xa3\x31\x16\x57\x90\xae\xb6\x92\x0e\xe3\x61\x74\xbc\xdc\x8b\x73\x4e\x7c\xad\x33\xc6\xb7\x1f\x2c\x17\x27\xb2\x7f\x01\x00\x00\xff\xff\xc1\xb0\xde\x9b\x35\x04\x00\x00"

func iris_ifaceTmplBytes() ([]byte, error) {
	return bindataRead(
		_iris_ifaceTmpl,
		"iris_iface.tmpl",
	)
}

func iris_ifaceTmpl() (*asset, error) {
	bytes, err := iris_ifaceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "iris_iface.tmpl", size: 1077, mode: os.FileMode(0644), modTime: time.Unix(1584107777, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x57, 0x3f, 0x30, 0x7c, 0xaf, 0x52, 0x34, 0xd5, 0x2f, 0x48, 0xa5, 0xb9, 0x1, 0x59, 0x45, 0xd0, 0x82, 0x42, 0x2f, 0x9e, 0x50, 0xfd, 0xa0, 0x60, 0x59, 0x4e, 0xe, 0xc6, 0x4e, 0x9d, 0x73, 0x87}}
	return a, nil
}

var _macaron_ifaceTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x54\x4d\x8f\xd3\x30\x10\xbd\xfb\x57\x8c\x7a\x4a\x51\x71\xc4\x15\x89\xc3\xaa\xbb\xb0\x1c\xf6\x43\x6c\x39\x21\x0e\xa6\x99\xa6\x56\x1b\xdb\xb2\xa7\x85\xc8\x9a\xff\x8e\x1c\x27\x4b\xc2\xb6\x2b\x84\xd8\x1b\x27\x7f\x4d\xde\x7b\xf3\xde\x28\x65\x09\x4b\x5b\x21\xd4\x68\xd0\x2b\xc2\x0a\xbe\xb5\x50\xdb\xd7\x8d\xf6\x12\x2e\xef\xe0\xf6\x6e\x05\x57\x97\x1f\x57\x52\x08\xa7\xd6\x3b\x55\x23\xc4\x08\xf2\x7e\x57\xdf\xaa\x06\x81\x59\x08\xdd\x38\xeb\x09\x0a\x01\x00\x30\xab\xad\xdb\xd5\x52\x9b\xb2\x51\x6b\xe5\xad\x91\xc7\x37\x33\x31\x17\x22\x46\xbd\x01\x63\xe9\xaa\x71\xd4\x3e\x90\x07\xb9\xb4\x4d\x83\x86\x80\xb9\x2c\x21\xc6\xe1\xcc\x1c\x23\x9a\x8a\x59\x50\xeb\x12\x9b\x5c\xb5\x0e\x13\x1b\x33\x68\x43\xe8\x37\x6a\x8d\x10\x4f\x21\x6e\x95\x36\xc0\x9c\x84\xa4\xc6\xba\xa3\xf3\xf6\xa8\x2b\x84\xad\x32\xd5\x1e\x7d\x80\x75\x77\xbd\xb1\x1e\x7a\x89\x9d\xf0\x24\x20\x3d\x30\x17\x73\xf8\xf2\x75\x50\x7f\x9d\xbf\x12\x83\xa6\x18\xbd\x32\x35\x82\x7c\xaf\x71\x5f\x85\x24\xf6\x7c\x5f\xbd\x8e\x71\x6f\x03\xd7\x0d\xd2\xd6\x56\xb9\xab\xe2\xd5\xc0\xb6\xb4\x86\xf0\x07\xcd\x63\xc4\x7d\xc0\x0c\xf0\x27\xc5\xbd\xb4\xbc\xb2\x10\x65\x09\x9f\xb0\xd6\x81\xd0\x4f\xfc\x7b\x40\x7f\x54\x86\xc0\xf7\x8f\xbf\xb9\x1b\xfa\x67\xb2\x8f\xd6\x6c\x0e\x66\xfd\x2c\x56\xd1\xc0\xa3\xa2\x9b\xbc\x2e\x20\x4c\x81\xe7\x27\xd3\xfa\xe0\xed\xc1\xc1\x49\x07\xa7\x39\x1e\x02\x8e\x42\xf3\xf6\x40\xe8\x3b\x1f\x1b\x5d\x55\x7b\xfc\xae\x3c\x06\x78\xfb\x0e\x82\x1c\x87\x98\x2b\x32\x4b\x31\x8b\x31\xef\x98\x67\x0b\x48\x4d\x15\x59\xd4\x99\x34\xaf\x89\xdc\x85\x69\x41\xa6\x4d\xb6\xbf\xd7\x93\x51\xf3\x58\x74\xb0\xbf\x2a\x3a\xec\x74\x75\xaf\x68\xdb\x1d\x9e\xcc\x51\xec\x34\x8e\xf3\x64\x29\xe5\x24\xef\x4c\x70\x61\xda\xe2\x6f\xa1\xc6\xd3\xc0\x8b\xb1\x4b\x13\xae\xff\xde\x0c\x17\x03\xc5\xcb\x4e\xe2\xe7\x80\xc5\xd3\x2c\x9e\xff\xab\x9c\xf7\xfa\x45\x7d\xfe\x77\x1e\x0f\x56\x8b\x9f\x01\x00\x00\xff\xff\x57\x26\x85\xc8\x69\x06\x00\x00"

func macaron_ifaceTmplBytes() ([]byte, error) {
	return bindataRead(
		_macaron_ifaceTmpl,
		"macaron_iface.tmpl",
	)
}

func macaron_ifaceTmpl() (*asset, error) {
	bytes, err := macaron_ifaceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "macaron_iface.tmpl", size: 1641, mode: os.FileMode(0664), modTime: time.Unix(1584092222, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x32, 0xf, 0xc5, 0x78, 0xf3, 0xbd, 0xd0, 0x47, 0x53, 0x9b, 0xc3, 0xee, 0x65, 0x41, 0xb7, 0x4a, 0x87, 0xe0, 0x72, 0x1d, 0xb9, 0x64, 0x9c, 0x93, 0x81, 0x3, 0x2f, 0x22, 0x2, 0xb1, 0x78, 0xc1}}
	return a, nil
}

var _mux_ifaceTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x54\x39\x6f\xdb\x3e\x14\xdf\xf9\x29\x1e\x34\x49\x41\xfe\x14\xf0\x1f\x03\x64\x28\x72\x34\x1d\x72\x34\x49\xd1\xa1\xe8\xc0\x58\x4f\x32\x1b\x91\x54\x79\xb8\x16\x08\x7e\xf7\x82\xa4\x6c\xcb\x71\x92\xa1\x5d\xea\x45\x20\xf9\xae\xdf\xf1\x5c\xd7\x70\xa6\x1a\x84\x0e\x25\x6a\x66\xb1\x81\xa7\x11\x3a\xf5\x9f\xe0\x9a\xc2\xf9\x2d\xdc\xdc\x3e\xc2\xc5\xf9\xa7\x47\x4a\xc8\xc0\x16\xcf\xac\x43\xf0\x1e\xe8\xdd\x73\x77\xc3\x04\x42\x08\x84\x70\x31\x28\x6d\xa1\x24\x00\x00\x85\x44\x5b\x2f\xad\x1d\x0a\x92\xcf\x1d\xb7\x4b\xf7\x44\x17\x4a\xd4\xac\xe7\x62\xac\x05\xd7\xf5\xea\xff\xe2\xe0\xb5\x53\x9a\xf7\x3d\xab\x85\x5b\x17\xa4\x22\xa4\xae\xe1\x87\x33\x16\x9c\x41\x68\x95\x86\x38\xd1\x07\x39\x82\x40\xbb\x54\x0d\x58\x05\x1a\x3b\x6e\x2c\x6a\xb2\x62\x1a\x62\xcf\xeb\xf4\x64\xe0\x34\x05\x5f\xed\x6e\x08\xf1\x9e\xb7\x20\x95\xbd\x10\x83\x1d\x1f\xac\x06\x7a\xa6\x84\x40\x69\x21\x84\xba\x06\xef\x37\xe7\x10\xbc\x47\xd9\x84\x40\xec\x38\x44\xb0\xf4\x71\x1c\x30\x82\x0d\x01\xb8\xb4\xa8\x5b\xb6\x40\xf0\xaf\x55\x5c\x32\x2e\x21\x04\x00\x88\xb4\xa6\xd3\xa0\xd5\x8a\x37\x08\x82\x37\x4d\x8f\xbf\x98\x46\x93\xd1\xb8\x75\x62\x20\x36\x8e\x81\x21\x94\x15\x7c\xfb\x2e\xdc\x9a\x5e\x6f\x43\x2f\x9d\x5c\x90\xcd\x38\xde\x6b\x26\x3b\x04\x7a\xc9\xb1\x6f\x4c\x9c\xf3\x6d\x48\x79\x84\x39\xaa\x4d\xb7\xcc\x48\xc6\x53\x46\xd2\xe8\x3d\x9a\x41\x49\x83\x5f\x35\xb7\xa8\x8f\xe1\x68\xba\xfd\xe9\xd0\xd8\xca\x7b\xec\x4d\xc4\xfe\x67\xb9\xd3\xe8\xf9\x1b\x92\xaa\xf7\x93\x6c\x7b\xd4\x3e\xa0\x5e\x31\x69\xb7\x9a\xbe\x20\xde\x4c\xcf\x56\x25\xe6\x5a\x27\x17\xef\xd6\x29\x35\x1c\x45\x2e\xef\x95\x4b\x73\x99\xfd\x7a\xd5\xab\xfa\x7d\xd4\xca\x0d\x99\x3c\xd0\x29\x11\x4e\x4e\x41\xd3\x3b\x66\x97\x77\x1a\x5b\xbe\x2e\x0b\xef\x73\x58\x08\x45\x45\x1f\xdc\x53\x8e\x2b\x77\x34\xed\xe7\xce\xd4\x7b\xc7\x2d\x51\xab\x68\xf3\x45\xba\x8a\xf6\xc8\x25\x92\x66\x73\xe7\x9c\x9c\x82\xa1\x73\xcb\x90\x5d\x3f\xfa\xc5\x60\x39\x0b\xa6\x94\x56\xdb\xf6\x64\xea\xb2\xa5\x37\xe5\x18\xe0\xb2\x55\x69\x95\x72\xbf\x37\x3d\x16\x77\x29\x2e\xdf\x6c\xa9\xf6\x88\xa2\x57\x4c\x36\x7d\x32\x6c\xe2\x28\x52\x16\x42\x71\x9c\xc7\x9d\xfb\xa6\xa2\x69\x96\xf8\x9b\x96\x33\x25\xec\xea\x46\x66\x0f\xd9\xba\x52\xc6\x86\xb0\xcb\x8d\xe7\x9c\x98\x1e\x8a\x8d\xd7\x52\xe6\x8a\xf5\xbc\xc1\xcf\x0e\xf5\x08\x34\x7e\x38\x9a\x79\xf2\x74\x55\x7a\xcf\x65\xdb\x33\x7b\x10\xba\xab\xf6\x52\xd5\xbf\x02\x3a\xfb\x8f\x8a\xea\xfc\x43\x28\xe7\x6b\x1a\xc8\xef\x00\x00\x00\xff\xff\x24\x46\x70\x02\x15\x06\x00\x00"

func mux_ifaceTmplBytes() ([]byte, error) {
	return bindataRead(
		_mux_ifaceTmpl,
		"mux_iface.tmpl",
	)
}

func mux_ifaceTmpl() (*asset, error) {
	bytes, err := mux_ifaceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mux_iface.tmpl", size: 1557, mode: os.FileMode(0664), modTime: time.Unix(1584082042, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1, 0x36, 0x24, 0x3e, 0x4a, 0x8e, 0x7c, 0xe6, 0x8a, 0xa2, 0x6c, 0x68, 0x99, 0x34, 0xec, 0xd4, 0xee, 0x11, 0x7a, 0xa0, 0xe6, 0xb6, 0x71, 0xb2, 0x65, 0x8f, 0x63, 0x7c, 0xab, 0x10, 0xa5, 0x4a}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"chi_iface.tmpl":        chi_ifaceTmpl,
	"echo_iface.tmpl":       echo_ifaceTmpl,
	"gin_iface.tmpl":        gin_ifaceTmpl,
	"httprouter_iface.tmpl": httprouter_ifaceTmpl,
	"iris_iface.tmpl":       iris_ifaceTmpl,
	"macaron_iface.tmpl":    macaron_ifaceTmpl,
	"mux_iface.tmpl":        mux_ifaceTmpl,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"chi_iface.tmpl":        {chi_ifaceTmpl, map[string]*bintree{}},
	"echo_iface.tmpl":       {echo_ifaceTmpl, map[string]*bintree{}},
	"gin_iface.tmpl":        {gin_ifaceTmpl, map[string]*bintree{}},
	"httprouter_iface.tmpl": {httprouter_ifaceTmpl, map[string]*bintree{}},
	"iris_iface.tmpl":       {iris_ifaceTmpl, map[string]*bintree{}},
	"macaron_iface.tmpl":    {macaron_ifaceTmpl, map[string]*bintree{}},
	"mux_iface.tmpl":        {mux_ifaceTmpl, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
