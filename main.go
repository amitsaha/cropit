package main

import (
	"fmt"
	"os"
	"github.com/oliamb/cutter"
	"image"
	"image/jpeg"
	"log"
)

func main() {

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Cannot open file", err)
	}
	defer f.Close()
	
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal("Cannot decode image:", err)
	}

	cImg, err := cutter.Crop(img, cutter.Config{
		Height:  1000, 
		Width:   1000, 
		Mode:    cutter.TopLeft,
		Anchor:  image.Point{60, 10},
		Options: 0,
	})

	if err != nil {
		log.Fatal("Cannot crop image:", err)
	}

	// Write the cropped image into a file
	cropped_f, err := os.Create("cropped.jpg")
	if err != nil {
		log.Fatal("Could not save cropped image")
	}
	defer cropped_f.Close()
	err = jpeg.Encode(cropped_f, cImg, &jpeg.Options{})
	if err != nil {
		log.Fatal("Error saving")
	}
	fmt.Println("Saved cropped image as cropped.jpg")
}
