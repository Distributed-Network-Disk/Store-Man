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

#[derive(Clone)]
pub struct File {
    pub id: i32,
    pub name: String,
    pub path: String,
    pub attribute: String,
    pub time: i64,
    pub is_folder: bool,
    pub file_type: String,
    pub file_size: i32,
    pub fraction_num: i32,
}

impl File {
    pub fn init(
        id: i32,
        name: String,
        path: String,
        attribute: String,
        time: i64,
        is_folder: bool,
        file_type: String,
        file_size: i32,
        fraction_num: i32,
    ) -> Self {
        File {
            id,
            name,
            path,
            attribute,
            time,
            is_folder,
            file_type,
            file_size,
            fraction_num,
        }
    }
}
