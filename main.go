// Simple command line program to crop images
package main

import (
	"bytes"
	"bufio"
	"flag"
	"fmt"
	"github.com/oliamb/cutter"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

// Crop the image and write the cropped image to the provided writer
func crop(imageData []byte, imageType string, cWidth int, cHeight int, w io.Writer)  {

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

func main() {

	// Setup the flags
	cHeight := flag.Int("height", 0, "Crop Height")
	cWidth := flag.Int("width", 0, "Crop Width")
	flag.Parse()

	//Check flags, file name specified
	if *cHeight == 0 {
		log.Fatal("Must specify the Crop Height")
	}

	if *cWidth == 0 {
		log.Fatal("Must specify crop width")
	}

	if len(flag.Args()) == 0 {
		log.Fatal("Must specify at least 1 image to crop")
	}

	// Loop over each file , crop and save the cropped image.
	for _, inputFilePath := range flag.Args() {
		imageData, err := ioutil.ReadFile(inputFilePath)
		if err != nil {
			log.Fatal("Cannot open file", err)
		}
		
		// Write the cropped image into a file
		croppedFileDir, fileName := filepath.Split(inputFilePath)
		croppedFileName := fmt.Sprintf("%scropped_%s", croppedFileDir, fileName)
		croppedFile, err := os.Create(croppedFileName)
		if err != nil {
			log.Fatal("Could not create file for cropped image")
		}
		
		defer croppedFile.Close()

		croppedFileWriter := bufio.NewWriter(croppedFile)
		// Detect file type and call the crop() function
		imageType := http.DetectContentType(imageData)
		crop(imageData, imageType, *cWidth, *cHeight, croppedFileWriter)

		// Flush
		croppedFileWriter.Flush()
	}
}
