/* Simple command line program using github.com/amitsaha/cropit/crop to crop images

Usage:

$ go run main.go --height=5000 --width=7000 <path to>/cat1.jpg  <path to>cat2.png

The cropped images will be placed in the same directory as the original images with
the file names being ``cropped_<original_file_name>.<original_extension>``
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/amitsaha/cropit/crop"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func cropper(inputFilePath string, cWidth int, cHeight int, done chan string) {

	imageData, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		log.Fatal("Cannot open file", err)
	}

	croppedFileDir, fileName := filepath.Split(inputFilePath)
	croppedFileName := fmt.Sprintf("%scropped_%s", croppedFileDir, fileName)
	croppedFile, err := os.Create(croppedFileName)
	if err != nil {
		log.Fatal("Could not create file for cropped image")
	}
	defer croppedFile.Close()

	croppedFileWriter := bufio.NewWriter(croppedFile)
	crop.Crop(imageData, cWidth, cHeight, croppedFileWriter)
	croppedFileWriter.Flush()

	// We are done with this file
	done <- inputFilePath
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

	numImages := len(flag.Args())
	if numImages == 0 {
		log.Fatal("Must specify at least 1 image to crop")
	}

	// Channel to synchronize with the go routines
	done := make(chan string)
	// Loop over each file , crop and save the cropped image.
	for _, inputFilePath := range flag.Args() {
		go cropper(inputFilePath, *cWidth, *cHeight, done)
	}

	// Wait for all the images to be cropped
	for i := 0; i < numImages; i++ {
		fmt.Printf("Cropped %s\n", <-done)
	}

}
