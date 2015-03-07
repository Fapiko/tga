package tga

import (
	"encoding/binary"
	"log"
	"os"
)

const (
	GRAYSCALE = 1
	RGB       = 3
	RGBA      = 4
)

type Color struct {
	// We're storing the color data in this order rather than the more canonical Red Blue Green Alpha as this is how
	// the order in which the data is written out so it simplifies the reading and writing process
	blueGreenRedAlpha [3]byte
}

type Image struct {
	width  int
	height int
	format byte
	data   []byte
}

type Header struct {
	idLength        byte
	colormapType    byte
	datatypeCode    byte
	colormapOrigin  int16
	colormapLength  int16
	colormapDepth   byte
	xOrigin         int16
	yOrigin         int16
	width           int16
	height          int16
	bitsPerPixel    byte
	imageDescriptor byte
}

func NewColor(red byte, blue byte, green byte, alpha byte) *Color {
	color := new(Color)

	color.blueGreenRedAlpha = [3]byte{blue, green, red}

	return color
}

func NewImage(width int, height int, format byte) *Image {
	// Figure out how large the data array needs to be. Format indicates the number of bytes per pixel
	numBytes := width * height * int(format)

	image := &Image{
		width:  width,
		height: height,
		format: format,
		data:   make([]byte, numBytes),
	}

	return image
}

// Set a single pixel to a specific color
// TODO: Add validation & error handling
func (image *Image) Set(x int, y int, color *Color) {
	// Calculate where we need to start copying the bytes in
	startLocation := (x + y*image.width) * int(image.format)

	log.Println(startLocation)

	copy(image.data[startLocation:startLocation+int(image.format)], color.blueGreenRedAlpha[0:3])
}

// Flip the image vertically
func (image *Image) FlipVertically() {

}

// Write out the TGA file
func (image *Image) WriteFile(location string) {
	header := &Header{
		bitsPerPixel:    image.format << 3,
		width:           int16(image.width),
		height:          int16(image.height),
		datatypeCode:    2,    //image.format,
		imageDescriptor: 0x20, // Flag that sets the image origin to the top left
	}

	log.Println(header)

	fileHandle, err := os.Create(location)
	if err != nil {
		panic(err)
	}

	// Write out header
	if err = binary.Write(fileHandle, binary.LittleEndian, header); err != nil {
		panic(err)
	}

	if err = binary.Write(fileHandle, binary.LittleEndian, &image.data); err != nil {
		panic(err)
	}

	developerArea := [4]byte{}
	if err = binary.Write(fileHandle, binary.LittleEndian, developerArea); err != nil {
		panic(err)
	}

	extensionArea := [4]byte{}
	if err = binary.Write(fileHandle, binary.LittleEndian, extensionArea); err != nil {
		panic(err)
	}

	footer := []byte("TRUEVISION-XFILE.\x00")
	if err = binary.Write(fileHandle, binary.LittleEndian, footer); err != nil {
		panic(err)
	}

	fileHandle.Close()
	// Data type code
	//	header[2] = 1

	//	developer

}
