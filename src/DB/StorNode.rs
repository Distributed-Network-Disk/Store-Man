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
pub struct StorNode {
    pub id: i32,
    pub ip: String,
    pub port: i32,
    pub is_online: bool,
    pub all_size: i32,
    pub left_size: i32,
}

impl DeviceItem {
    pub fn init(
        id: i32,
        ip: String,
        port: i32,
        is_online: bool,
        all_size: i32,
        left_size: i32,
    ) -> Self {
        DeviceItem {
            id: id,
            ip: ip,
            port: port,
            is_online: is_online,
            all_size: all_size,
            left_size: left_size,
        }
    }
}
