package db

// Copyright [2022] [totoroyyw]

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"Store-Man/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type FragmentObj struct {
	gorm.Model
	FileName            string // original name
	Path                string // /aaa/bbb/
	Node                string // ip:port
	FragmentIndexOfFile uint32 // 1
}

// TODO: which is key?

// TODO: how about pg?

var (
	Config config.CONF
	db     *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open(Config.SqliteName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&FragmentObj{})
}

func DBGetFileName(path string) []string {
	var filenames []string
	db.Where("path = ?", path).Find(&filenames)
	return filenames
}
