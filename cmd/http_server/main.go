package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/mohanson/doa"
)

var (
	flListen = flag.String("l", "127.0.0.1:8080", "listen address")
	flRoot   = flag.String("d", ".", "root directory")
)

func aopBanMethods(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.ServeHTTP(w, r)
		case http.MethodHead:
			handler.ServeHTTP(w, r)
		case http.MethodPost:
			w.WriteHeader(http.StatusMethodNotAllowed)
		case http.MethodPut:
			w.WriteHeader(http.StatusMethodNotAllowed)
		case http.MethodPatch:
			w.WriteHeader(http.StatusMethodNotAllowed)
		case http.MethodDelete:
			w.WriteHeader(http.StatusMethodNotAllowed)
		case http.MethodConnect:
			handler.ServeHTTP(w, r)
		case http.MethodOptions:
			handler.ServeHTTP(w, r)
		case http.MethodTrace:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

func aopLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.RequestURI())
		handler.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()
	log.Println("root", *flRoot)
	var handler http.Handler
	handler = http.FileServer(http.Dir(*flRoot))
	handler = aopLog(handler)
	handler = aopBanMethods(handler)
	http.Handle("/", handler)
	log.Println("listen and serve on", *flListen)
	doa.Try1(http.ListenAndServe(*flListen, nil))
}
