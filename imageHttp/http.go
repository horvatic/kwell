package imageHttp

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"net/http"
)

type imageProcessor func(source image.Image) image.Image

func imageService(res http.ResponseWriter, req *http.Request, process imageProcessor) {
	defer req.Body.Close()
	body, readErr := ioutil.ReadAll(req.Body)
	if readErr != nil {
		fmt.Fprintf(res, "error reading: %s!", readErr)
		return
	}

	indexHead := bytes.Index(body, []byte("\r\n\r\n")) + 4
	bodyDecode := bytes.NewReader(body[indexHead:])
	img, imageErr := png.Decode(bodyDecode)
	if imageErr != nil {
		fmt.Fprintf(res, "error image: %s!%s", imageErr, body[indexHead:])
		return
	}
	processImage := process(img)

	res.Header().Set("Content-Type", "image/png")
	png.Encode(res, processImage)
}
