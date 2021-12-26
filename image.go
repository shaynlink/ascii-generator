package main

import (
	"bytes"
	"image"
	"io"
	"log"
	"reflect"
	"image/color"

	_ "image/jpeg"
	_ "image/png"

	"github.com/nfnt/resize"
)

var patterns string = "@%#*+=-:. "

func ImageToAscii(reader io.Reader, width int, invert bool) []byte {
	img, _, err := image.Decode(reader)

	if err != nil {
		panic(err)
	}

	log.Printf("Image type: %T", img)

	imgBound := img.Bounds()

	newImage := resize.Resize(uint(width), uint((imgBound.Max.Y*width*10)/(imgBound.Max.X*16)), img, resize.Lanczos3)

	grayImg := image.NewGray(newImage.Bounds())
	for y := newImage.Bounds().Min.Y; y < newImage.Bounds().Max.Y; y++ {
		for x := newImage.Bounds().Min.X; x < newImage.Bounds().Max.X; x++ {
			grayImg.Set(x, y, newImage.At(x, y))
		}
	}

	log.Printf("Image type: %T", grayImg)

	table := []byte(patterns)

	if invert {
		table = reverseArray(table)
	}

	buf := new(bytes.Buffer)

	for y := grayImg.Bounds().Min.Y; y < grayImg.Bounds().Max.Y; y++ {
		for x := grayImg.Bounds().Min.X; x < grayImg.Bounds().Max.X; x++ {
			g := color.GrayModel.Convert(grayImg.At(x, y))
			axeY := int(reflect.ValueOf(g).FieldByName("Y").Uint())
			pos := int((axeY * (len(patterns) - 1)) / 255)
			_ = buf.WriteByte(table[pos])
		}

		_ = buf.WriteByte('\n')
	}

	return buf.Bytes()
}

func reverseArray(arr []byte) []byte{
	for i, j := 0, len(arr)-1; i<j; i, j = i+1, j-1 {
	   arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
 }