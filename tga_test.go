package tga_test

import (
	"fmt"
	"github.com/fapiko/tga"
	"testing"
)

func TestColor(t *testing.T) {
	white := tga.NewColor(255, 255, 255, 0)

	fmt.Println(white)
}

func TestDrawPixel(t *testing.T) {
	red := tga.NewColor(255, 0, 0, 255)

	//	image := tga.NewImage(100, 100, tga.RGB)
	//	image.Set(52, 41, red)
	image := tga.NewImage(5, 5, tga.RGB)
	image.Set(2, 2, red)
	image.WriteFile("output.tga")

	t.Log(image)
}
