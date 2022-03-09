use actix_web::{get, post, App, HttpRequest, HttpResponse, HttpServer, Responder, Result};

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
