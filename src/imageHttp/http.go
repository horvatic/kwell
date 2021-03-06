package imageHttp

import (
	"bytes"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type imageProcessor func(source image.Image, searchArea int, runInParallel bool) image.Image

func imageService(res http.ResponseWriter, req *http.Request, process imageProcessor) {
	defer req.Body.Close()
	body, readErr := ioutil.ReadAll(req.Body)
	if readErr != nil {
		fmt.Fprintf(res, "error reading: %s!", readErr)
		return
	}

	indexHead := bytes.Index(body, []byte("\r\n\r\n")) + 4
	bodyDecode := bytes.NewReader(body[indexHead:])
	img, _, imageErr := image.Decode(bodyDecode)
	if imageErr != nil {
		bodyDecode := bytes.NewReader(body)
		img, _, imageErr = image.Decode(bodyDecode)
		if imageErr != nil {
			fmt.Fprintf(res, "error image: %s!%s", imageErr, body)
			return
		}
	}

	processImage := process(img, getSearchAreaSize(req.URL), getrunInParallel(req.URL))

	res.Header().Set("Content-Type", "image/png")
	png.Encode(res, processImage)
}

func getrunInParallel(u *url.URL) bool {
	q := u.Query()
	runInParallel, errors := strconv.ParseBool(q.Get("runInParallel"))
	if errors != nil {
		return false
	}
	return runInParallel
}

func getSearchAreaSize(u *url.URL) int {
	q := u.Query()
	area, errors := strconv.Atoi(q.Get("sa"))
	if errors != nil {
		return 0
	}
	return area
}
