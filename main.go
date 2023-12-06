package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok1"))
	})

	httpSrv := http.Server{
		Addr:    fmt.Sprintf("%s:%d", "0.0.0.0", 8080),
		Handler: mux,
	}
	err := httpSrv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
