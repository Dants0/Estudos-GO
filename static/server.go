package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":1718", "http service address")

func main() {
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":3000", nil)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("Listen And Serve: ", err)
	}
}
