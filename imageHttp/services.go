package imageHttp

import (
	"github.com/horvatic/vaticwella/imageProcessing"
	"net/http"
)

func superSampleService(res http.ResponseWriter, req *http.Request) {
	imageService(res, req, imageProcessing.SuperSampling)
}

func medianFilterService(res http.ResponseWriter, req *http.Request) {
	imageService(res, req, imageProcessing.MedianFilter)
}
