// Simple command line program to crop images
package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/oliamb/cutter"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Return the cropped Image
func crop(imageData []byte, imageType string, cWidth int, cHeight int) image.Image {

	// Check if we can handle this image
	if imageType != ".jpg" && imageType != ".png" {
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

	return croppedImg
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
	for i, inputFilePath := range flag.Args() {
		imageData, err := ioutil.ReadFile(inputFilePath)
		if err != nil {
			log.Fatal("Cannot open file", err)
		}

		// Open the file specified as the first argument
		imageType := filepath.Ext(inputFilePath)
		croppedImg := crop(imageData, imageType, *cWidth, *cHeight)

		// Write the cropped image into a file
		croppedFileName := fmt.Sprintf("cropped_%d%s", i, imageType)
		croppedFile, err := os.Create(croppedFileName)
		if err != nil {
			log.Fatal("Could not create file for cropped image")
		}
		defer croppedFile.Close()

		// Now we encode the cropped image data using the appropriate
		// encoder and save it to the above file
		switch imageType {
		case ".png":
			err = png.Encode(croppedFile, croppedImg)
		case ".jpg":
			err = jpeg.Encode(croppedFile, croppedImg, &jpeg.Options{})
		}

		if err != nil {
			log.Fatal("Error saving")
		}
	}
}
