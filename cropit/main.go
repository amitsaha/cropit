// Simple command line program using github.com/amitsaha/cropit/crop to crop images
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"github.com/amitsaha/cropit/crop"
)

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
		
		croppedFileDir, fileName := filepath.Split(inputFilePath)
		croppedFileName := fmt.Sprintf("%scropped_%s", croppedFileDir, fileName)
		croppedFile, err := os.Create(croppedFileName)
		if err != nil {
			log.Fatal("Could not create file for cropped image")
		}
		defer croppedFile.Close()
		
		croppedFileWriter := bufio.NewWriter(croppedFile)
		crop.Crop(imageData, *cWidth, *cHeight, croppedFileWriter)
		croppedFileWriter.Flush()
	}
}
