package main

import (
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"io"
	"net/http"
	"os"
)

func answer4(res http.ResponseWriter, req *http.Request) {
	glog.Info("/healthz has been called")
	res.WriteHeader(200)
	io.WriteString(res, "ok")
}

func answer1(res http.ResponseWriter, req *http.Request) {
	glog.Info("answer1 run")
	for k, v := range req.Header {
		res.Header().Set(k, v[0])
	}
	io.WriteString(res, "answer1")
}

func answer2(res http.ResponseWriter, req *http.Request) {
	version := configs.Version
	io.WriteString(res, version)
}

func answer3(res http.ResponseWriter, req *http.Request) {
	req_ip := req.RemoteAddr
	statusCode := 203
	res.WriteHeader(statusCode)
	glog.Info("Client IP:", req_ip, ", Response Code:", statusCode, ".")
	io.WriteString(res, "answer3")
}

type Configs struct {
	Port    string `json:"port"`
	Version string `json:"version"`
}

var configs Configs

func main() {
	jsonFile, err := os.Open("/app/configs/config.json")
	if err != nil {
		glog.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &configs)
	port := configs.Port
	// flag.Parse()
	flag.Set("v", "4")
	glog.Info("HTTP Server Start")
	// handler
	http.HandleFunc("/answer1", answer1)
	http.HandleFunc("/answer2", answer2)
	http.HandleFunc("/answer3", answer3)
	http.HandleFunc("/healthz", answer4)
	// listen 80
	listenAddress := "0.0.0.0:" + port
	glog.Info(listenAddress)
	listenAndServeErr := http.ListenAndServe(listenAddress, nil)
	if listenAndServeErr != nil {
		glog.Fatal(listenAndServeErr)
	}
}
