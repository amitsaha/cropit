package crop_test

import (
	"bufio"
	"fmt"
	"github.com/amitsaha/cropit/crop"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func ExampleCrop() {
	// Test image to crop
	inputFilePath := "../test_images/cat1.jpg"
	imageData, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		log.Fatal("Cannot open file", err)
	}

	// Write the cropped image to a file cropped_cat1.jpg
	croppedFileDir, fileName := filepath.Split(inputFilePath)
	croppedFileName := fmt.Sprintf("%scropped_%s", croppedFileDir, fileName)
	croppedFile, err := os.Create(croppedFileName)
	if err != nil {
		log.Fatal("Could not create file for cropped image")
	}
	defer croppedFile.Close()

	croppedFileWriter := bufio.NewWriter(croppedFile)
	// Crop width and height
	cWidth := 5000
	cHeight := 4000
	crop.Crop(imageData, cWidth, cHeight, croppedFileWriter)
	croppedFileWriter.Flush()
}
