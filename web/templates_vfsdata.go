// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package web

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Templates statically implements the virtual filesystem provided to vfsgen.
var Templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2021, 2, 15, 13, 43, 55, 575838238, time.UTC),
		},
		"/admin": &vfsgen۰DirInfo{
			name:    "admin",
			modTime: time.Date(2021, 2, 15, 13, 43, 55, 578838236, time.UTC),
		},
		"/admin/allow-list-remove-confirm.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "allow-list-remove-confirm.tmpl",
			modTime:          time.Date(2021, 2, 15, 17, 42, 2, 266073953, time.UTC),
			uncompressedSize: 413,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x8c\x90\xc1\x6a\xf3\x30\x10\x06\xef\x7a\x8a\x65\xef\x7f\x42\x6e\x3f\xc1\x36\x84\xb6\x87\xd2\x42\x43\x1a\xe8\xd9\xb1\xd6\x78\x41\x5a\x19\x45\x4e\x30\x62\xdf\xbd\xc8\x4d\x4a\xa1\x97\xde\x47\xf3\x8d\x36\x67\xb0\xd4\xb3\x10\x60\xe2\xe4\x08\x41\x35\x67\xde\xfc\x17\xc0\x9d\xf5\x2c\x3b\xe7\xc2\xf5\x95\xcf\xe9\x40\x3e\x5c\xe8\x21\x48\xcf\xd1\x1f\x17\xb6\xa0\x40\x62\x41\xd5\xfc\x10\x75\x41\x12\x49\x2a\x2a\x53\x0d\x1b\x60\x5b\xe3\x95\x5c\x17\x3c\x61\xf3\x07\xf9\xc7\x8d\x55\xad\xd6\xc3\xa6\x31\xa6\x1a\x23\x2d\x96\x0b\x45\xee\xe7\x22\x59\x3d\x49\x8a\xf3\x6a\x3f\x9d\x5e\x68\x5e\x1d\xa8\x2f\xf0\x18\xa9\x31\xc6\x54\x7d\x88\x7e\xe1\xbb\x2f\x21\x42\xdb\x25\x0e\x52\x63\xce\x53\x74\xc7\x00\xd8\x96\xf9\x6d\x5b\xf6\xff\x39\x3e\xa7\x6d\x5c\x0a\x50\x15\xc1\x53\x1a\x82\xad\x71\xff\xf6\x7e\xc4\xc6\x00\x54\x2c\xe3\x94\x20\xcd\x23\xd5\x38\xb0\xb5\x24\x08\xd2\x7a\xaa\x91\x2d\xc2\xa5\x75\x13\xd5\xdf\x4d\xcf\x8f\xaa\xbf\x5e\x9d\xa7\x93\xe7\x74\x67\xf1\x7e\x85\xdb\x8f\xcb\x6c\x63\xaa\x75\x09\x6f\x8c\xc9\x99\xc4\xaa\x9a\xcf\x00\x00\x00\xff\xff\x2d\xb2\xea\x04\x9d\x01\x00\x00"),
		},
		"/admin/allow-list.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "allow-list.tmpl",
			modTime:          time.Date(2021, 2, 15, 17, 14, 16, 230167877, time.UTC),
			uncompressedSize: 742,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x7c\x52\x4d\x8b\xdb\x30\x10\xbd\xfb\x57\x4c\xe7\xd4\x1e\x64\x93\xee\xa5\x04\xd9\x10\xda\x1e\x4a\x0b\x5d\xb6\x81\x1e\x8b\x62\x8d\xd7\xa2\xfa\x30\xb2\x94\x34\x08\xfd\xf7\x62\xc5\xa6\x5b\xd8\xdd\x93\xed\x79\xf3\x3e\xfc\xa4\x94\x40\xd2\xa0\x2c\x01\x06\x15\x34\x21\xe4\x9c\x92\xda\x7d\xb0\x80\x07\x69\x94\x3d\x68\xed\x2e\xdf\xd4\x1c\x8e\x05\x5e\x50\x20\x2b\x21\xe7\xea\x09\xb7\x77\x36\x90\x0d\x0b\xbb\xe2\x52\x9d\xa1\xd7\x62\x9e\x5b\x9c\xc4\x23\xb1\x91\x84\x24\x8f\x5d\x05\xc0\xc7\x1d\x28\xd9\xe2\x85\x74\xef\x0c\x61\xf7\xbc\xd7\xcf\x15\xce\x99\x37\xe3\xae\xab\x78\x23\xd5\xb9\xfb\x4f\xd9\xbb\x4b\x51\x04\x78\x3a\xed\x9d\x66\x46\xb2\xdd\xfb\x0d\x9b\x8a\x9d\xd8\x84\x3f\xba\x68\xc3\xe6\x3a\x69\xc0\x7f\x43\xa8\xcb\x73\xb1\x9c\x56\x72\xd4\x85\x1d\x46\x5a\xd6\x56\xc9\x94\xbc\xb0\x8f\x04\xf5\x67\x1b\xbc\xa2\x39\xe7\x32\x06\xe0\x5a\x75\x29\xd5\xf7\xf1\xf4\x95\xae\xf5\x03\x0d\x39\xc3\x5b\x2e\x60\xf4\x34\xb4\x98\x52\xf4\xfa\xe8\x00\xc5\xf2\xa7\xfb\x92\x88\x69\x35\x87\xbd\x27\xe3\xce\xb4\xef\x9d\x1d\x94\x37\x08\xa8\x24\x42\xfd\xe5\x53\xce\xd8\x3d\x14\x8c\x37\xa2\x7b\x07\xbc\xd1\x6a\x8b\x40\x56\xae\xbe\xbc\x89\x7a\x8d\x7b\x2b\x69\x79\xbd\x7d\x8f\x77\x2f\xd4\x7b\x90\xf2\x56\xed\xdd\xca\x1c\x9c\x37\xb7\xa2\xa4\x64\x64\x83\xbf\x22\x88\x3e\x28\x67\x5f\x0b\x2e\x8a\x0c\x82\xa1\x30\x3a\xd9\xe2\xfd\xf7\x1f\xc7\xb5\x24\x00\xae\xec\x14\x03\x84\xeb\x44\x2d\x06\xfa\x13\x10\xac\x30\xd4\xe2\x14\x4f\xbf\x7e\xd3\xf5\xf9\xc5\x39\x9e\x8c\x0a\x08\x67\xa1\x23\x2d\xde\x2f\xc7\xdf\x4e\xb8\x59\xc2\x6f\x57\x04\xf8\x1b\xc6\xa0\xf1\xee\x02\x8c\x75\xd5\x56\xd4\xdf\x00\x00\x00\xff\xff\xa6\x94\x91\xce\xe6\x02\x00\x00"),
		},
		"/admin/dashboard.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "dashboard.tmpl",
			modTime:          time.Date(2021, 2, 15, 13, 43, 32, 430852769, time.UTC),
			uncompressedSize: 408,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x6c\x8f\x41\x6b\xc4\x20\x10\x85\xef\xf9\x15\x53\xef\x26\xa4\xa7\x1e\x5c\xa1\x6c\x7f\x41\x29\xf4\x6c\xe3\x74\x33\x60\xc6\x60\xdc\x4d\x40\xfc\xef\x25\x36\x96\x5d\xe8\x69\xe4\x7d\xbe\xf7\x66\x52\x02\x8b\xdf\xc4\x08\x22\x52\x74\x28\x20\xe7\x94\xa8\x7f\x61\x10\xaf\x76\x22\x7e\x33\xcb\xf8\xe5\x4d\xb0\x1f\x05\xef\x14\x90\x2d\xe4\xdc\xdc\x79\x07\xcf\x11\x39\xee\xee\x46\x59\xba\xc1\xe0\xcc\xb2\x9c\xc4\x6c\x2e\x28\x47\x34\x16\x83\xd0\x0d\x80\x1a\x7b\x20\x7b\x12\x2b\xba\xc1\x4f\x28\xf4\xff\x5d\x9f\x07\xce\x59\x75\x63\xaf\x1b\xd5\x59\xba\xe9\x87\xe4\xe0\xd7\x92\x08\x70\xaf\x46\xdc\xa2\xdc\x96\x4a\xe6\x07\xfd\x79\x73\xa2\xb4\x07\xef\xa7\xb3\xbf\x72\xac\xfd\xb3\x3b\x36\x78\xff\x23\xd0\x96\xb9\x6f\x30\x1f\x69\x57\xf7\xfb\x48\x29\x18\xbe\x20\xb4\x67\x47\xc8\x71\xc9\xb9\xc8\xa0\x1c\xe9\x94\xda\xdd\xe2\xa8\x7e\x45\xb6\x07\x57\x5d\x0d\xa8\xe7\x94\x01\xea\x49\x4a\xe8\x82\x5f\x41\x4a\xdd\x1c\x8e\x9f\x00\x00\x00\xff\xff\xc9\xcb\xf6\x0b\x98\x01\x00\x00"),
		},
		"/auth": &vfsgen۰DirInfo{
			name:    "auth",
			modTime: time.Date(2021, 2, 10, 16, 34, 42, 221227039, time.UTC),
		},
		"/auth/fallback_sign_in.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "fallback_sign_in.tmpl",
			modTime:          time.Date(2021, 2, 10, 16, 34, 42, 221227039, time.UTC),
			uncompressedSize: 445,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x74\x8f\xcb\x6a\xeb\x30\x10\x86\xf7\x7e\x8a\x39\xb3\x17\xc6\x67\x55\x82\x6d\xe8\xa6\xdb\x16\x1a\xe8\x5a\xb1\x26\xb1\xa8\x2e\x46\x1a\xc5\x2d\x42\xef\x5e\xec\xd8\x90\x42\xba\xd2\xe5\xfb\xe7\x9b\x99\x9c\x41\xd1\x59\x3b\x02\x64\xcd\x86\x10\x4a\xc9\x59\x37\x4f\x0e\xf0\x39\xf1\xf8\x22\x8d\x39\xc9\xe1\xf3\xb8\xc2\x85\x01\x39\x05\xa5\x54\x77\x95\x83\x77\x4c\x8e\x97\xda\xaa\x55\xfa\x0a\x83\x91\x31\x76\x38\xc9\x0b\x89\x91\xa4\xa2\x80\x7d\x05\xd0\x8e\x0d\x68\xd5\xe1\x4c\x66\xf0\x96\xb0\x7f\xd4\xe9\x63\x83\xa5\xb4\xf5\xd8\xf4\x55\x5b\x2b\x7d\xed\x7f\x79\x83\x9f\x57\x1f\xc0\xfd\xef\xe0\x8d\xb0\x4a\x34\xff\x77\x76\xf6\xc1\x82\x25\x1e\xbd\xea\xf0\xed\xf5\xfd\x88\x20\x07\xd6\xde\x75\x39\xa7\x60\x8e\x1e\x50\x26\x1e\x0f\xe7\xad\xf3\x21\xea\x8b\xd3\x0e\x4b\xb9\x09\x56\x89\x76\x53\x62\xe0\xef\x89\x3a\x64\xfa\x62\x04\x27\x2d\x75\x98\xe2\xb6\xd4\x83\xdc\x24\x63\x9c\x7d\x50\x7b\x76\x79\xff\x95\x8d\xe9\x64\x35\xef\x23\xd7\xcb\xcc\xfb\xfd\xb6\xf7\x7a\x40\xfb\x4f\x08\xa8\x83\x9f\x41\x88\xbe\xca\x99\x9c\x2a\xe5\x27\x00\x00\xff\xff\x62\x75\x88\x18\xbd\x01\x00\x00"),
		},
		"/base.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "base.tmpl",
			modTime:          time.Date(2021, 2, 15, 13, 43, 32, 444852763, time.UTC),
			uncompressedSize: 1153,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xbc\x94\xc1\x6e\xdb\x30\x0c\x86\xef\x79\x0a\x56\xd8\x71\x8e\xdb\x5b\xe1\xca\x06\xba\x75\xd8\x06\x0c\xed\xb0\xf4\xb2\x53\xa1\xd8\x4c\x4c\x94\x96\x02\x8b\x76\x56\x18\x7e\xf7\x41\xb2\xd3\x2c\x45\x81\xa2\x97\x9d\x6c\x8a\xff\xff\x59\xa4\x28\xeb\xb3\x9b\xbb\xcf\xf7\xbf\x7f\x7e\x81\x5a\x1a\x2e\x16\xfa\xf0\x40\x53\x15\x0b\x00\x7d\x96\x24\x50\xba\x66\xc7\x28\x08\x6b\xe3\x11\x04\x9b\x1d\x1b\x41\x48\x92\xa8\x68\x50\x0c\x94\xb5\x69\x3d\x4a\xae\x3a\xd9\x24\x97\x0a\xd2\x63\xca\x9a\x06\x73\xd5\x13\xee\x77\xae\x15\x05\xa5\xb3\x82\x56\x72\xb5\xa7\x4a\xea\xbc\xc2\x9e\x4a\x4c\x62\xf0\x11\xc8\x92\x90\xe1\xc4\x97\x86\x31\xbf\x58\x9e\x1f\x50\x5e\x9e\x18\x8b\xb0\xbb\xa1\x27\x4f\x6b\x62\x92\xa7\xac\xa6\xaa\x42\x7b\xe5\x76\xa6\x0c\xe1\xf9\xd5\xa8\xd3\x49\x19\x3c\x4c\xf6\x11\xea\x16\x37\xb9\x4a\x8d\xf7\x28\x7e\x4a\x2e\x4b\xef\x15\xb4\xc8\xb9\x8a\xb1\xaf\x11\x45\x45\x8b\x90\x30\x16\xc3\xb0\x66\x57\x3e\x82\x8a\xa1\x82\xe5\x38\x7e\x75\xc9\x6a\xf5\x09\x7e\x39\xd7\xc0\x0a\xdb\x1e\xdb\x61\x40\x5b\x8d\xa3\x4e\x27\xcf\x42\xa7\x53\xd3\xf4\xda\x55\x4f\x01\xf6\x4c\xc1\x3f\xd2\x9a\x48\x99\x3d\x8b\xf0\xa9\x8a\xfa\xa0\x0a\xba\x0f\x9d\xc7\x16\xb2\x1c\xc8\x3f\xb0\xdb\x6e\xb1\x7a\x20\x3b\x8e\x73\x96\x36\x10\x05\xf3\x02\x80\xde\x15\xf3\x1b\xc0\x37\x64\x76\x07\xc4\xf2\xd6\x34\x38\x8e\x67\x07\x5d\xfa\x2c\xd4\x1d\x1f\x3d\x9a\xa9\xd0\x66\xee\xcc\x30\x74\x2d\xdf\x3b\x50\x16\xf7\x3e\x73\x3d\xb6\xe1\xa8\xd4\x38\xaa\x62\x18\xe8\xe2\xd2\x82\xba\xc5\xbd\xbf\x8f\x9d\x08\xe5\x9a\x42\xa7\x4c\x6f\xd1\x4c\xd5\x90\xcd\x2a\xe3\xeb\xb5\x33\x6d\x75\xca\x33\xfd\x75\x48\xdf\xfc\x93\x7d\x1f\xd7\x30\xbb\x7d\xc2\xe4\xe5\xf5\x1d\x47\xfc\x75\x10\xfd\x20\x2f\xef\xde\x7b\x27\x75\xb6\x31\xcc\x6b\x53\x3e\x66\xec\xb6\xae\x93\x53\x7c\x27\xf5\x8a\xb6\xf6\x2e\xae\xbf\xc4\xea\xf4\xd0\xec\x61\x40\xf6\x78\x3c\xb7\xff\x7e\x08\x27\x85\x78\xda\x5a\xb2\xd9\xc6\xb5\xcd\xab\xd5\x7c\xb7\x6f\x15\x13\x67\x37\xac\xc5\xd9\x7d\x31\xc4\x30\x4f\xfb\x7c\xbf\x15\x2c\xe1\x38\xf0\xcf\x26\x9d\x4e\xd7\x43\xa7\xd3\x9f\xe6\x6f\x00\x00\x00\xff\xff\x96\x49\x08\x96\x81\x04\x00\x00"),
		},
		"/error.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "error.tmpl",
			modTime:          time.Date(2021, 2, 15, 13, 43, 55, 575838238, time.UTC),
			uncompressedSize: 365,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x5c\xd0\xc1\x4e\xf4\x20\x10\x07\xf0\x7b\x9f\x62\xc2\x77\xf9\x3c\xd0\xa6\x1e\x0d\xe5\xb0\x66\x9f\xc0\x27\x98\xc2\xac\xc5\x6d\xa1\x19\xb0\xa6\x21\xbc\xbb\x51\xb6\x59\xf5\x34\xc9\x1f\xf8\xfd\xc9\xe4\x6c\xe9\xe2\x3c\x81\x48\x2e\xcd\x24\x4a\x61\xf2\x96\x18\x24\x9c\x99\x03\x43\xce\xed\x4b\xc2\xf4\x1e\x9f\x83\xa5\x52\x72\x26\x6f\x4b\x69\xee\xef\x4c\xf0\x89\x7c\x12\xa5\x34\xca\xba\x0d\xcc\x8c\x31\x0e\x62\xc5\x57\x92\x13\xa1\x25\x16\xba\x01\x50\x53\xaf\xab\xf8\xef\x0f\xa9\xe2\x82\xf3\xac\x41\xde\xbb\x4a\x51\x5d\x4d\x55\x37\xf5\xba\x51\x9d\x75\x9b\xfe\xe5\x73\xf8\xa8\xee\x8f\xcc\x84\x59\xc6\x45\xf6\x8f\xdf\x27\x00\x6a\x05\x67\x07\x41\xcc\xa7\x60\x77\xa1\x73\x6e\xcf\xcc\x5f\xf8\x7a\x5c\xa8\x13\x40\x21\x4c\x4c\x97\x41\xbc\xe1\x86\xd1\xb0\x5b\xd3\xd3\xe4\x62\x0a\xbc\xb7\x23\x9a\xeb\xff\x07\x71\xb4\x8c\xc9\xc3\x98\xbc\x5c\xd9\x2d\xc8\xbb\xd0\x27\x34\x57\xd5\xe1\x8d\xac\xf6\xf1\xe3\x3a\x6e\x5b\xfb\x0c\x00\x00\xff\xff\x9c\xcf\x6d\x54\x6d\x01\x00\x00"),
		},
		"/landing": &vfsgen۰DirInfo{
			name:    "landing",
			modTime: time.Date(2021, 2, 12, 9, 56, 41, 54426042, time.UTC),
		},
		"/landing/about.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "about.tmpl",
			modTime:          time.Date(2021, 2, 12, 9, 56, 41, 54426042, time.UTC),
			uncompressedSize: 135,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x4c\xcb\xb1\xaa\x02\x31\x10\x46\xe1\x3e\x4f\xf1\xdf\xed\x43\xd8\x7e\x08\xdc\xde\xd2\x17\x88\xce\xb8\x09\xc8\x44\x64\xd4\x62\x98\x77\x97\xd5\xc6\xea\x34\xdf\x71\x07\xcb\x65\xa8\x60\xb1\x61\x57\x59\x10\x71\x68\xca\x43\x37\x64\xfc\x9f\xe6\xc3\xdc\x21\xca\x88\x48\x3f\xf8\x3c\xd5\x44\x6d\xe7\x89\x78\x3c\x6b\x02\xa8\xaf\xf5\xd8\x05\x6d\xbf\x70\x6b\x9b\x50\xe9\x6b\x4d\x54\x3e\xe0\x1b\xd0\x5f\xce\x28\xf7\xf9\x42\xce\x35\xb9\x8b\x72\xc4\x3b\x00\x00\xff\xff\xe2\x37\xb4\x29\x87\x00\x00\x00"),
		},
		"/landing/index.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "index.tmpl",
			modTime:          time.Date(2021, 2, 12, 9, 56, 41, 54426042, time.UTC),
			uncompressedSize: 4104,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x94\x57\x4b\x6f\x1c\xb9\x11\xbe\xeb\x57\x54\x74\x5e\xb5\xe0\x5b\x0e\x5a\x01\x42\x56\x8b\x35\x60\x6b\xb3\x2b\xdb\x39\x73\xc8\xd2\x4c\x05\x7c\xb4\xc8\xe2\x24\x80\xa0\xff\x1e\x54\xf1\x31\x3d\xb1\x25\x60\x4f\xb6\xba\xd9\xac\xc7\xf7\xa8\x9a\x97\x17\x70\xf8\x44\x11\xe1\x92\x89\x3d\x5e\xc2\xeb\xeb\xcb\x0b\x7d\xf8\x7b\x84\xcb\x4f\x26\x3a\x8a\xfb\x2f\xfa\x5c\x1e\x03\x46\x07\xaf\xaf\x17\x9b\x8f\x6c\x8a\x8c\x91\xe5\xb3\x8b\x1b\x47\x47\xb0\xde\x94\xf2\xb3\x3e\x37\x14\x31\x43\xf8\xef\x95\xa9\x9c\x2e\x6f\x2f\x00\x6e\x0e\x1f\x6e\xff\xef\xf6\x7f\xa1\xb7\x29\xc8\xfd\x37\xd7\x87\x0f\xb7\x17\x72\x8a\xc2\x1e\x4a\xb6\x3f\x5f\x5e\x9b\x52\x90\xcb\x35\x85\xfd\x35\x63\xe1\xab\x03\xe6\x40\xb8\xac\x71\x7f\xd9\x8e\xae\xb7\x9f\x52\xc6\x00\xb4\x96\x1a\xc0\x25\x9f\x32\x14\x62\x30\x01\xf9\x27\xb0\x29\x16\xb4\x8c\x5c\x33\x18\x47\x2b\x15\x4b\x71\x0f\xe8\x89\x7f\x82\x82\x0e\x5c\x02\xa4\x5a\x42\x72\xc0\x18\xd6\x94\x81\xa2\x25\x47\xae\x46\x86\xca\xe0\xcd\x2e\x65\x04\xe4\x76\x35\x42\x30\xfb\x68\x2e\x00\x00\x8c\xa7\xe7\x6a\x16\x78\x44\x07\x4f\xa6\x5a\xda\xd5\x02\x5c\xf3\x4a\x05\x28\x02\x56\x08\x04\x3b\xda\x61\x74\x35\x2c\x70\x07\x26\xdb\x0a\xb6\xe6\x52\x0b\x1c\x89\x0d\x4a\x7a\xfb\x2a\x77\xd6\x4c\x05\xf2\x21\x45\x5b\x0b\x18\x8c\x68\x22\x1c\xd1\x6b\xa2\x0b\xdc\x49\x24\x64\x8d\x3a\x2e\x04\x8c\x14\xe0\xc9\x58\xf2\x54\xa8\xc0\x3e\x9b\x23\x39\x03\x11\x9f\xab\x5e\x7c\x34\xde\x53\x59\xe0\x4b\x2b\x2b\x65\x4b\x92\x93\x4f\xbb\x94\x99\x0a\xa0\xc7\x80\x91\x6b\x80\x48\xbb\x03\x30\x7a\x5f\x0b\x84\xe4\xb1\x30\xe1\x02\x8f\xbd\x89\x1a\x15\x0b\xc3\xea\x8d\xc5\x6c\x58\x8b\xdb\x63\x61\x53\xa0\xfd\x1d\x56\xcc\x8e\x50\x32\x0d\xc8\xef\x35\x1d\x5c\xa5\x02\x9c\xa9\xb0\x94\x04\x25\x79\x4f\x96\xb8\x3a\x8a\x2d\x8f\x81\xdd\xa2\x71\xbf\x7a\xce\x64\x51\x1b\x4a\xc6\x56\x4f\x05\x62\x8d\x56\xb0\x5b\xe0\x77\x47\x69\xd3\x82\xde\xc6\x71\x03\x04\x53\x4a\xc7\xe7\x58\xfd\x5a\xd9\x30\x0a\x26\x9b\x03\xfa\x81\x4d\x21\x24\x97\xe0\xb9\x52\xd1\xa8\x9b\x82\x7e\xc3\xe8\x32\x66\xea\xf8\x37\xf8\xa5\x7c\xd6\x9e\xf8\x37\xb9\xb7\xc0\x83\x22\xd1\x80\x16\x6e\xd5\x02\xcf\xd5\x04\x58\xd1\x7b\x8c\x8c\x45\x5e\x47\xb4\x10\x4d\x68\xd5\xfe\x21\xcd\xc9\x24\xfc\x28\x67\x49\x27\x29\xb4\xf2\x02\x5f\x94\x9d\xc2\xcd\x89\xe3\x13\xd6\x3d\x19\x86\x23\x1d\x4d\x10\xf6\x30\x98\xba\xaf\xb8\xdc\x5c\xaf\xb7\x4d\x21\x8f\x18\x20\x56\xef\x0d\xac\x07\x93\x91\xb3\x01\x47\x26\x9c\x1a\x11\xa9\xf8\xd6\x28\x8d\x14\x52\xde\x91\xb6\x63\xf6\x46\x9f\x77\x62\x4a\x6e\x27\x5c\x17\xf8\x18\x21\xe5\x68\x32\x6a\x7d\x5a\xc9\x91\x8e\x98\xb3\x69\xac\x2b\x66\x4f\xcc\xc2\xc5\x7b\xa1\x80\x36\x6f\xd0\xd5\xd6\x00\x25\x59\x12\x5c\x0d\xa7\xe7\x8a\x0b\xdc\x55\xcb\x29\x77\x26\x7f\xd7\xbf\x05\xfe\x99\x32\x33\xb1\xa2\x51\x8a\x01\x72\x1a\xb3\x9d\x6f\xaa\x0c\x70\x14\x1a\xef\xaa\xaf\xa1\x97\xb3\xf3\xe2\x37\x3c\x04\xa8\x6d\x5e\xe0\xf7\x96\xb8\x47\xcb\x75\x43\x9d\xf7\xf8\xbe\xc0\x57\x56\xf5\x35\x3d\xf6\x6b\x8f\xc9\x57\x5e\x8d\x90\x0a\x2d\x46\x53\x16\xf8\x15\x73\x57\x98\x76\x0f\xeb\xc4\x6a\xcd\xc8\xd4\x95\xb7\xc0\x1d\x77\xd0\x47\xd7\x36\xa2\x91\x70\xf7\x3c\xe3\x9f\x28\xaa\x94\x2b\xb5\xac\x18\x1d\x95\x82\x0b\x7c\x16\x2d\x9d\xe4\x01\xe4\xa0\xa0\x30\xb9\xdf\x4e\x11\x0e\x83\xcd\x02\xb6\xbe\x1a\x19\x35\xed\xa1\x83\xb5\xfa\x23\x45\x93\x37\xf4\x79\x18\x37\x2a\xaf\xc0\x1b\xbb\xc9\x55\xf1\x59\xe0\x6b\x8e\x02\xc4\xa9\x0d\xed\x94\x37\x29\xa3\x10\x2c\x45\x69\xbc\xd9\x91\x98\xc2\x00\x5f\x0d\xd1\x58\xed\x40\x47\xa0\xe6\xee\xae\xae\xb1\xaf\xdb\xd8\xe6\x7f\xdd\xa8\xc8\x01\x45\xc6\xac\xde\xfa\x9b\xb1\x70\x90\xcb\x65\x5c\x08\x70\x8c\x42\x71\xcb\x35\x14\xd6\xa6\x29\xb1\xce\x89\x7a\x46\x52\x8d\x39\x88\x2a\x50\xcd\x42\xd2\x99\xc9\x2c\x70\x67\xe1\x09\x25\x11\x97\x44\xbc\xd8\x4f\x9c\xa9\x5a\xd5\x35\x6f\x18\x1a\x52\xcf\xba\x1f\xb6\xdb\x60\x8c\x8c\x7b\xcc\xcd\xcc\x3b\xe1\xc7\x67\xc6\x02\x4f\xb9\x37\x19\x74\x3c\xa5\x20\xf1\x17\xb5\x8e\xc9\xf8\x82\x41\xb2\xe1\x94\x45\x1a\x6a\xc2\xcf\x55\xe8\xeb\x60\x4d\x99\x4d\x63\x9b\x86\xfd\xac\xba\xe9\xd2\x9a\xe7\x1d\x75\xba\x7a\x63\x29\x92\x69\x0e\x70\x44\x3f\xbc\x6c\xb0\xb0\x8f\xad\xde\xd7\xf9\x58\xab\xa6\xb8\x9d\x73\x0b\xfc\xd9\x86\x9a\x86\x0d\x46\xbb\x3b\xe6\x9c\x60\x3d\x55\xde\xa0\xf8\x77\x2d\x9c\x16\xf8\x43\xea\x19\xd7\x93\xeb\xfd\x44\xdf\x70\x3b\x4d\xae\x49\x56\x61\xd0\x6c\x9c\x9e\xae\x6d\x70\x1d\x31\x62\x34\xbc\xe1\x4d\x84\x80\xc2\xb4\x93\xbb\x62\x5d\xe0\x97\x4a\x12\xa8\xd3\x43\x2b\xe8\xa6\xab\x7d\x3d\x77\xcf\xe6\x00\x55\xdf\x0c\x04\x5b\x33\x35\xff\x36\x31\xf7\x3a\x60\x64\x5a\x3c\x4d\x1f\xe8\x23\x6c\x81\x3f\xdf\x33\xf9\xfb\x37\xd9\xb7\x19\x71\x0b\x7c\x43\x0f\x4f\x99\xe2\x9e\x24\x37\x2c\xad\x60\x49\x34\xd8\x94\x85\x27\x9a\x43\x4b\x7d\x5c\x01\xc8\x24\x1e\xfa\xa0\xe7\x94\x63\x0d\xff\x96\xff\x5a\x73\x2d\x73\x9b\xd9\x78\xc0\xbd\x10\x74\xda\xe3\xa9\xab\x8a\x61\xb7\xd4\x1e\x4e\xcc\xc2\xa2\xc7\xdc\xe0\xeb\xd0\x4e\x26\x34\x6f\xdd\xaa\x45\x17\x83\x4d\xde\x4d\xff\xb4\x8f\x54\x0a\x05\x95\xed\x96\xe3\xa7\x9a\x2b\x77\x6f\x3f\xe9\xa4\x59\xd4\xc9\x19\xc6\x16\xf3\xde\x52\xb2\xc0\xc7\x36\x41\xfe\xb2\x79\x49\x1f\xa3\x6d\xd9\x21\xff\xe5\xdd\x4c\x83\xde\xe7\x34\xbc\xc4\xd8\x86\x37\x6f\x56\x36\x67\x56\x5d\x2c\xeb\xf7\x9b\x90\x74\xa6\xd5\x79\x56\x9d\xa7\x1d\xe6\x24\x6c\x9e\x5b\x69\x9b\xef\xf7\x45\x90\xf3\xc4\x83\xa0\x4d\xb5\x6f\x32\x7f\x81\xc7\x0d\x8e\xcd\xf4\xe6\x6a\x04\x6b\x4e\x14\x37\xdc\xf6\x98\x54\xa2\x9a\x75\x37\x9c\x14\x17\x78\x10\xd6\x29\xbb\x9a\xc7\xd6\x62\x69\x25\x9e\x0b\x67\x63\x74\xdf\x48\x4e\xc3\xf3\x81\x8a\x9f\x73\xf2\xa9\x16\xab\x46\x70\x6c\x5c\x99\x2e\x37\x0e\x0c\x92\x9d\x08\xfb\xcd\x64\xd2\x7d\x56\xf8\xd1\xec\x55\xa0\x9a\xd4\x16\xc7\x30\x6d\xd1\x68\x3e\x24\x6f\xc9\x4d\xdf\xfa\xf6\x03\x3d\x8c\x26\x9d\x8d\xdd\xaf\xdc\x1f\x6b\x79\xef\x2d\x12\x0b\x7c\x6a\xd0\xa8\x4b\x4c\x7a\x9d\x36\xc6\xc7\xef\x65\x33\xfa\x43\x6f\x6b\xa0\xd1\xe8\x63\xd4\x58\x26\xb2\xd0\xbf\x6f\x16\x5f\xda\xaf\x91\x01\xf5\x9c\x36\x22\xd4\x4e\x5a\xe9\xf4\xee\x00\x36\x9b\x32\xed\x74\x18\xf5\x5c\x25\xfa\x92\xb3\xc0\xaf\x9b\xb8\xcd\xd2\xdf\xd5\x5f\xc7\x0b\x2b\xe9\x2f\x2c\x8a\x0b\x7c\x36\x1e\x4b\x35\x4e\x2c\x29\x60\x51\x07\x3a\xcf\x52\x02\xf6\x7d\x4f\x91\x11\xae\x4d\xe3\xbd\xaf\x1a\xf6\x47\x8b\x54\xdb\x89\x24\x82\x76\x6c\x0a\x06\x45\x60\x14\x87\x4b\xf5\x68\xfd\xc7\xc1\xfd\x9b\x2e\xae\x5d\x5a\xe0\x97\xe1\x43\x83\xbd\x03\x79\x5d\x0b\x4f\x1d\x5f\xe0\xe3\x94\x46\xb3\xdc\x30\x2b\xfd\xce\xed\x94\x35\x93\x55\xb3\xed\xa7\xc1\xb6\xdd\x0b\x4e\xcb\xed\x02\xff\x98\x53\xbd\xaf\x02\x83\x26\xb2\x08\xf8\x1f\xb0\xab\x29\xe2\xe6\xda\xd1\x71\xfc\x03\x37\x7f\xbb\xba\x82\xeb\x9c\xfe\x03\x57\x57\xb7\x17\x2f\x2f\x18\xdd\xeb\xeb\xff\x02\x00\x00\xff\xff\x8f\x92\xaf\x72\x08\x10\x00\x00"),
		},
		"/news": &vfsgen۰DirInfo{
			name:    "news",
			modTime: time.Date(2021, 2, 15, 13, 43, 55, 578838236, time.UTC),
		},
		"/news/overview.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "overview.tmpl",
			modTime:          time.Date(2021, 2, 16, 9, 31, 25, 618043975, time.UTC),
			uncompressedSize: 386,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x54\x90\x41\x6b\xc3\x30\x0c\x85\xef\xfa\x15\x9a\xe8\x71\x6e\x49\x4f\xa3\x38\x86\xfd\x81\xb2\x43\x61\x67\x13\xab\x8d\xc1\xb1\x43\xec\x35\x05\xe1\xff\x3e\x52\x52\x48\x4f\x86\xef\xf9\x3d\x3d\x49\x04\x1d\x5f\x7d\x64\xa4\xe2\x4b\x60\xc2\x5a\x45\x7c\xf3\x15\x91\xce\x3c\xe7\xcb\x13\x2e\x0c\x39\x3a\xac\x15\x36\x8e\x2e\xc5\xc2\xb1\x2c\x1e\xd0\xce\xdf\xb1\x0b\x36\xe7\x96\x46\x7b\x63\xd5\xb3\x75\x3c\x91\x01\x44\xdd\x37\xe8\x5d\x4b\x33\x87\x2e\x0d\x4c\x66\x3b\xe1\x77\x85\xb5\xea\x43\xdf\x18\xd0\x07\xe7\xef\xe6\x2d\x6f\x4a\x33\xbd\x93\x2e\x05\x35\x38\xd5\x1c\xc9\x80\xc8\x64\xe3\x8d\x71\xe7\xa3\xe3\xc7\x27\xee\x38\xf0\xc0\xb1\xe0\xa9\xc5\xfd\x77\x08\x3f\x29\x97\x5c\x2b\x00\x68\x8b\xfd\xc4\xd7\x96\x44\xfe\xa6\x70\x49\x48\x91\xe7\x7c\x1a\x53\x2e\x84\xe4\x1d\xad\x19\xb5\xae\xb5\x8f\x46\xe4\x15\xb7\x3f\xdb\x81\x9f\x25\x8f\x4b\x49\x6b\x00\xf4\xb8\xd5\x2f\xfc\x28\x8b\x3e\x1a\x00\x11\x8e\x6e\xb9\xca\xba\xcc\xf3\x41\xfd\xa1\x14\x1e\xa6\x34\xa3\x52\xe6\xf5\xe7\x3f\x00\x00\xff\xff\xf8\x06\x06\xca\x82\x01\x00\x00"),
		},
		"/news/post.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "post.tmpl",
			modTime:          time.Date(2021, 2, 15, 13, 43, 55, 578838236, time.UTC),
			uncompressedSize: 384,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x74\xce\xb1\x4e\xc5\x20\x14\xc6\xf1\xbd\x4f\x71\x3c\x93\x0e\xb4\xa9\x93\xb9\xe1\x32\xb9\xb8\x54\x87\xbe\x00\x29\xe7\x5a\x92\x5e\x68\x00\x4b\x13\x72\xde\xdd\xb4\xea\xd5\x0e\x4e\x90\x7f\xf8\x7e\xa1\x14\x30\x74\xb1\x8e\x00\x93\x4d\x13\x21\x30\x97\x62\xdb\x27\x07\xd8\x51\x8e\xfd\x1e\x99\x4f\x50\x4a\xfd\xe6\x63\xaa\x3b\x7d\xa5\xed\x0d\x90\x33\xc0\x5c\xfd\x11\x06\xef\x12\xb9\xb4\x19\x95\x34\x76\x81\x61\xd2\x31\x9e\x71\xd6\xef\x24\x46\xd2\x86\x02\xaa\x0a\x40\x8e\xad\x3a\x72\xb2\x19\x5b\x55\xc9\xc6\xd8\x45\x1d\xa6\xc1\x67\x3c\x96\xc1\x4f\xe2\x6a\x44\xfb\xb8\xf5\xf9\xe6\xf4\xb4\xa6\xcd\x99\x7f\x99\xfd\x00\x79\x27\x04\x34\xc1\x67\x10\x42\x55\x52\xc3\x18\xe8\x72\xc6\x52\x3e\xc2\xd4\x7b\x40\x47\x39\x9e\xfc\x42\x61\xb1\x94\x91\x19\xc1\x9a\x33\xde\x82\x7a\xfd\xbe\xc9\x46\xff\x3b\x9f\x7d\x4c\x08\x68\x0d\xc2\xbd\x75\x03\x7c\x7d\xe9\xe5\xf9\xe1\x87\x73\xb4\x26\x54\x1d\xad\xa9\xae\xeb\x5d\x2a\x85\x9c\x61\xfe\x0c\x00\x00\xff\xff\xec\x11\x39\x89\x80\x01\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/admin"].(os.FileInfo),
		fs["/auth"].(os.FileInfo),
		fs["/base.tmpl"].(os.FileInfo),
		fs["/error.tmpl"].(os.FileInfo),
		fs["/landing"].(os.FileInfo),
		fs["/news"].(os.FileInfo),
	}
	fs["/admin"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/admin/allow-list-remove-confirm.tmpl"].(os.FileInfo),
		fs["/admin/allow-list.tmpl"].(os.FileInfo),
		fs["/admin/dashboard.tmpl"].(os.FileInfo),
	}
	fs["/auth"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/auth/fallback_sign_in.tmpl"].(os.FileInfo),
	}
	fs["/landing"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/landing/about.tmpl"].(os.FileInfo),
		fs["/landing/index.tmpl"].(os.FileInfo),
	}
	fs["/news"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/news/overview.tmpl"].(os.FileInfo),
		fs["/news/post.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
