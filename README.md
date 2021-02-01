# HTTP Server

Serves HTTP requests with the contents of the file system rooted. It is used to maintain my personal blog [http://accu.cc](http://accu.cc), and it works well.

Note that HTTPs service is provided by [https://github.com/mohanson/https_proxy](https://github.com/mohanson/https_proxy).

```text
Usage of http_server:
  -d string
        root directory (default ".")
  -l string
        listen address (default "127.0.0.1:8080")
  -r404 string
        page uri for 404
```
