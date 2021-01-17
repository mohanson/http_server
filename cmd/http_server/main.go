package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/mohanson/doa"
)

var (
	flListen = flag.String("l", "127.0.0.1:8080", "listen address")
	flRoot   = flag.String("d", ".", "root directory")
	flR404   = flag.String("r404", "", "page uri for 404")
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

func aop404(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(path.Join(*flRoot, r.URL.Path)); os.IsNotExist(err) {
			if *flR404 != "" {
				http.Redirect(w, r, *flR404, 301)
				return
			}
		}
		handler.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()
	log.Println("root", *flRoot)
	var handler http.Handler
	handler = http.FileServer(http.Dir(*flRoot))
	handler = aop404(handler)
	handler = aopLog(handler)
	handler = aopBanMethods(handler)
	http.Handle("/", handler)
	log.Println("listen and serve on", *flListen)
	doa.Try1(http.ListenAndServe(*flListen, nil))
}
