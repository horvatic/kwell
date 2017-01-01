package imageHttp

import (
	"net/http"
)

func SetRoutes() {
	http.HandleFunc("/ss", superSampleService)
	http.HandleFunc("/mf", medianFilterService)
	http.ListenAndServe(":8080", nil)
}
