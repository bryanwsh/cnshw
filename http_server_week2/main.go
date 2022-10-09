package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"github.com/golang/glog"
)

func answer4(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	io.WriteString(res, "ok")
}

func answer1(res http.ResponseWriter, req *http.Request) {
	glog.V(2).Info("answer1 run")
	for k, v := range req.Header {
		res.Header().Set(k, v[0])
	}
	io.WriteString(res, "answer1")
}

func answer2(res http.ResponseWriter, req *http.Request) {
	os.Setenv("VERSION", "123456")
	version := os.Getenv("VERSION")
	io.WriteString(res, version)
}

func answer3(res http.ResponseWriter, req *http.Request) {
	req_ip := req.RemoteAddr
	statusCode := 203
	res.WriteHeader(statusCode)
	glog.V(2).Info("Client IP:", req_ip, ", Response Code:", statusCode,".")
	io.WriteString(res, "answer3")
}

func main() {
	// flag.Parse()
	flag.Set("v", "4")
	glog.Info("HTTP Server Start")
	// handler
	http.HandleFunc("/answer1", answer1)
	http.HandleFunc("/answer2", answer2)
	http.HandleFunc("/answer3", answer3)
	http.HandleFunc("/healthz", answer4)
	// listen 80
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
