package main

import (
	"net/http"
	"github.com/amitsaha/cropit/crop"
	"log"
	"strconv"
	"io/ioutil"
)
// cropImage expectes the following HTTP FORM data:

// image: The image must use this form key name
// w: Desired width of the image
// h: Desired height of the image
func cropImage(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Cannot read image data", 400)
	}
	cWidth := r.FormValue("w")
	cHeight := r.FormValue("h")

	if cWidth == "" || cHeight == "" {
		http.Error(w, "Must supply both width and height", 400)
	} else {
		originalImage, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, "Cannot read image data", 400)
		}
		cWidth, err := strconv.Atoi(cWidth)
		if err != nil {
			http.Error(w, "Bad value for crop width supplied", 400)
		}
		cHeight, err := strconv.Atoi(cHeight)
		if err != nil {
			http.Error(w, "Bad value for crop height supplied", 400)
		}
		log.Printf("Recieved %s, Desired Width and height: %d, %d\n", header.Filename, cHeight, cWidth)

		// Call the crop.Crop() function passing the Response Writer object
		// to it
		// XXX: Add proper error handling via recover()
		crop.Crop(originalImage, cWidth, cHeight, w)
	}
}
func main() {
	http.HandleFunc("/", cropImage)
	http.ListenAndServe(":9090", nil)
}
