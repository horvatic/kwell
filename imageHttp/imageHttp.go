package imageHttp

import (
	"bytes"
	"fmt"
	"github.com/horvatic/vaticwella/imageProcessing"
	"image/png"
	"io/ioutil"
	"net/http"
)

func SuperSampleService(res http.ResponseWriter, req *http.Request) {
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
	superImage := imageProcessing.SuperSampling(img)

	res.Header().Set("Content-Type", "image/png")
	png.Encode(res, superImage)
}
