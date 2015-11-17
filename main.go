// Simple command line program to crop images
package main

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"flag"
	"github.com/oliamb/cutter"
)

func cropAndSave(inputFilePath string, cWidth int, cHeight int) () {

	// Check if we can handle this image
	fileExtension := filepath.Ext(inputFilePath)
	if fileExtension != ".jpg" && fileExtension != ".png" {
		log.Fatal("Cannot handle image of this type")
	}
	
	inputFileDir, inputFileName := filepath.Split(inputFilePath)
	f, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal("Cannot open file", err)
	}
	defer f.Close()

	// We first decode the image
	img, _, err := image.Decode(f)
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

	// Write the cropped image into a file
	croppedFileName := "cropped_" + inputFileName
	croppedFilePath := filepath.Join(inputFileDir, croppedFileName)
	croppedFile, err := os.Create(croppedFilePath)
	if err != nil {
		log.Fatal("Could not save cropped image")
	}
	defer croppedFile.Close()

	// Now we encode the cropped image data using the appropriate
	// encoder and save it to the above file
	switch fileExtension {
	case ".png":
		err = png.Encode(croppedFile, croppedImg)
	case ".jpg":
		err = jpeg.Encode(croppedFile, croppedImg, &jpeg.Options{})
	}

	if err != nil {
		log.Fatal("Error saving")
	}
	fmt.Printf("Saved cropped image as %s\n", croppedFilePath)
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
	
	// Open the file specified as the first argument
	inputFilePath := flag.Args()[0]
	cropAndSave(inputFilePath, *cWidth, *cHeight)
}
