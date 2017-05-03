package main

import (
	"flag"
	"net/http"
)

func main() {
	port := flag.String("p", "8080", "Port to host files on.")
	dir := flag.String("d", "./", "Directory in which to root server.")
	flag.Parse()
	panic(http.ListenAndServe(":"+*port, http.FileServer(http.Dir(*dir))))
}
