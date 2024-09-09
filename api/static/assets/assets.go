// Copyright 2016 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package assets

import (
	"fmt"
	"github.com/xxl6097/go-http/api/static/assets/html"
	"io/fs"
	"net/http"
)

var (
	// read-only filesystem created by "embed" for embedded files
	content    fs.FS
	FileSystem http.FileSystem
	// if prefix is not empty, we get file content from disk
	prefixPath string
)

func init() {
	fmt.Println("加载冬桃资源")
	Register(html.AssetsHml)
}

// Load if path is empty, load assets in memory
// or set FileSystem using disk files
func Load(path string) {
	prefixPath = path
	if prefixPath != "" {
		FileSystem = http.Dir(prefixPath)
	} else {
		FileSystem = http.FS(content)
	}
}

func RegisterPath(fileSystem fs.FS, path string) {
	subFs, err := fs.Sub(fileSystem, path)
	if err == nil {
		content = subFs
	}
}

func Register(fileSystem fs.FS) {
	subFs, err := fs.Sub(fileSystem, ".")
	if err == nil {
		content = subFs
	}
}