package imageHttp

import (
	"net/http"

	"github.com/horvatic/vaticwella/src/imageProcessing"
)

func superSampleService(res http.ResponseWriter, req *http.Request) {
	imageService(res, req, imageProcessing.SuperSampling)
}

func medianFilterService(res http.ResponseWriter, req *http.Request) {
	imageService(res, req, imageProcessing.MedianFilter)
}

func binaryFilterService(res http.ResponseWriter, req *http.Request) {
	imageService(res, req, imageProcessing.BinaryFilter)
}
