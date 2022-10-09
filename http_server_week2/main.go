package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)


func healthz(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "ok")
}

func main() {
	fmt.Print("HTTP Server Start")
	//handle
	http.HandleFunc("/healthz", healthz)
	// listen 80
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
