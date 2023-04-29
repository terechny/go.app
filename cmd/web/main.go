package main

import (
	"fmt"
	"net/http"

	"post.loc/app/cmd/web/handler"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/post", handler.Post)
	mux.HandleFunc("/list", handler.List)
	mux.HandleFunc("/create", handler.Create)
	mux.HandleFunc("/store", handler.Store)
	mux.HandleFunc("/redact", handler.Redact)
	mux.HandleFunc("/update", handler.Update)
	mux.HandleFunc("/delete", handler.Delete)

	fileServer := http.FileServer(http.Dir("../../ui/image/"))
	mux.Handle("/image/", http.StripPrefix("/image", fileServer))

	fmt.Println("Started on http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)

	fmt.Println(err)

}
