package main

import (
	"image"
	"os"
	"testing"
)

func TestCrop(t *testing.T) {
	img := getImage()

	c := Config{
		Width:  512,
		Height: 400,
	}
	r, err := Crop(img, c)
	if err != nil {
		t.Fatal(err)
	}
	if r.Bounds().Dx() != 512 {
		t.Error("Bad width should be 512 but is", r.Bounds().Dx())
	}
	if r.Bounds().Dy() != 400 {
		t.Error("Bad width should be 400 but is", r.Bounds().Dy())
	}
	if r.Bounds().Min.X != 0 {
		t.Error("Invalid Bounds Min X", r.Bounds().Min.X)
	}
	if r.Bounds().Min.Y != 0 {
		t.Error("Invalid Bounds Min Y", r.Bounds().Min.Y)
	}
}
