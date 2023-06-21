package main

import {
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
}

func main() {
	http.HandleFunc("/", HandleGetvideos)
	http.HandleFunc("/update", HandleUpdateVideos)
	http.ListenAndServe(":80", nil)
}