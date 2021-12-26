package main

import (
	"fmt"
	"strconv"

	"github.com/discord/fasthttp"
)

var widthDefault string = "100"

func Handler(ctx *fasthttp.RequestCtx) {
	ImageToAsciiHandler(ctx)
}

func ImageToAsciiHandler(ctx *fasthttp.RequestCtx) {
	fileHeader, err := ctx.FormFile("image")

	if err != nil {
		fmt.Println(err)
		ctx.Error("Bad request form file", fasthttp.StatusBadRequest)
		return
	}
	
	if fileHeader == nil {
		fmt.Println("fileHeader is nil")
		ctx.Error("No image found", fasthttp.StatusBadRequest)
		return
	}
	
	fmt.Println("image found")
	
	mimeType := fileHeader.Header.Get("Content-Type")

	fmt.Println(mimeType)

	if mimeType != "image/jpeg" && mimeType != "image/png" {
		ctx.Error("Mime Type invalid", fasthttp.StatusBadRequest)
		return
	}

	file, err := fileHeader.Open()

	if err != nil {
		fmt.Println(err)
		ctx.Error("Internal error", fasthttp.StatusInternalServerError)
		return
	}

	width := string(ctx.FormValue("width"))

	if width == "" {
		width = widthDefault
	}

	parseWidth, err := strconv.ParseInt(width, 10, 32)

	if err != nil {
		fmt.Println(err)
		ctx.Error("Width invalid parsing", fasthttp.StatusBadRequest)
		return
	}

	invert := string(ctx.FormValue("invert"))
	baseInvert := false

	if invert == "true" {
		baseInvert = true
	}

	ctx.Write(ImageToAscii(file, int(parseWidth), baseInvert))
}