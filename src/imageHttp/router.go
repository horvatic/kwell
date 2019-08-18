package imageHttp

import (
	"net/http"
)

func StartService() {
	http.HandleFunc("/ss", superSampleService)
	http.HandleFunc("/mf", medianFilterService)
	http.HandleFunc("/bf", binaryFilterService)
	http.ListenAndServe(":8080", nil)
}
