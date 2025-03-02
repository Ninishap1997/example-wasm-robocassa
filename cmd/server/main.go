package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "8080", "port to serve on")
	flag.Parse()

	// Сервируем статические файлы из папки static
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Printf("Starting server on port %s", *port)
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
