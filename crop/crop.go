// Crop package exporting the Crop() function used to crop images
package crop

import (
	"bytes"
	"github.com/oliamb/cutter"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"io"
	"net/http"
)

func validateImageType(image_type string) bool {
	recognizedImageTypes := []string{"image/jpeg", "image/png"}
	for _, t := range recognizedImageTypes {
		if image_type == t {
			return true
		}
	}
	return false
}

/*
Crop() takes in four parameters:

imageData : byte[] representing the input image

cWidth    : An int specifying the desired width of the cropped image

cHeight   : An int specifying the desired height of the cropped image

w         : An io.Writer to which to write the encoded ropped image bytes to
*/
func Crop(imageData []byte, cWidth int, cHeight int, w io.Writer)  {

	imageType := http.DetectContentType(imageData)

	// Check if we can handle this image
	if !validateImageType(imageType) {
		log.Fatal("Cannot handle image of this type")
	}

	// We first decode the image
	reader := bytes.NewReader(imageData)
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal("Cannot decode image:", err)
	}

	// Perform the cropping
	croppedImg, err := cutter.Crop(img, cutter.Config{
		Height:  cHeight,
		Width:   cWidth,
		Mode:    cutter.TopLeft,
		Anchor:  image.Point{60, 10},
		Options: 0,
	})

	if err != nil {
		log.Fatal("Cannot crop image:", err)
	}

	// Now we encode the cropped image data using the appropriate
	// encoder and save it to the above file
	switch imageType {
	case "image/png":
		err = png.Encode(w, croppedImg)
	case "image/jpeg":
		err = jpeg.Encode(w, croppedImg, &jpeg.Options{})
	}
	if err != nil {
		log.Fatal("Error saving")
	}
}

