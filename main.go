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

func main() {

	// Setup the flags
	cHeight := flag.Int("height", 0, "Crop Height")
	cWidth := flag.Int("width", 0, "Crop Width")
	flag.Parse()

	//XX: Check flags, file name specified

	// Open the file specified as the first argument
	inputFilePath := flag.Args()[0]
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
		Height:  *cHeight,
		Width:   *cWidth,
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
	switch filepath.Ext(inputFileName) {
	case ".png":
		err = png.Encode(croppedFile, croppedImg)
	case ".jpg":
		err = jpeg.Encode(croppedFile, croppedImg, &jpeg.Options{})
	default:
		err = errors.New("Unsupported format: " + filepath.Ext(inputFileName))
	}

	if err != nil {
		log.Fatal("Error saving")
	}
	fmt.Printf("Saved cropped image as %s\n", croppedFilePath)
}
