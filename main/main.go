package main

import (
	"github.com/horvatic/vaticwella/imageHttp"
	"net/http"
)

func main() {
	http.HandleFunc("/ss", imageHttp.SuperSampleService)
	http.ListenAndServe(":8080", nil)
}
