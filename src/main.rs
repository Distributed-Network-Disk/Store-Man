use actix_web::{get, post, App, HttpRequest, HttpResponse, HttpServer, Responder, Result};

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

#[post("/listfile")]
async fn listfile(params: web::Json<listfile_param>) -> web::Json<ret_string> {
    let result: String = GetFileList::execute(params.whose.clone(), params.QueryPath.clone());
    web::Json(ret_string { result })
}

#[actix_rt::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(move || App::new().service(listfile))
        .bind("127.0.0.1:8000")?
        .run()
        .await
}
