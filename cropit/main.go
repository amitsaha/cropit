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

type Result struct {
	FilePath string
	Error    string
}

func cropper(inputFilePath string, cWidth int, cHeight int, done chan Result) {

	// Recover from any panic here so that one bad image doesn't bring the entire
	// program down
	defer func() {
		if err := recover(); err != nil {
			done <- Result{FilePath: inputFilePath, Error: fmt.Sprintf("%s", err)}
		}
	}()

	imageData, _ := ioutil.ReadFile(inputFilePath)
	croppedFileDir, fileName := filepath.Split(inputFilePath)
	croppedFileName := fmt.Sprintf("%scropped_%s", croppedFileDir, fileName)
	croppedFile, _ := os.Create(croppedFileName)
	defer croppedFile.Close()

	croppedFileWriter := bufio.NewWriter(croppedFile)
	crop.Crop(imageData, cWidth, cHeight, croppedFileWriter)
	croppedFileWriter.Flush()

	// We are done with this file
	done <- Result{FilePath: inputFilePath, Error: ""}
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
	done := make(chan Result)
	// Loop over each file , crop and save the cropped image.
	for _, inputFilePath := range flag.Args() {
		go cropper(inputFilePath, *cWidth, *cHeight, done)
	}

	// Wait for all the images to be cropped
	for i := 0; i < numImages; i++ {
		result := <-done
		if result.Error == "" {
			fmt.Printf("Cropped %s\n", result.FilePath)
		} else {
			fmt.Printf("Error cropping %s, %s\n", result.FilePath, result.Error)
		}
	}
}
