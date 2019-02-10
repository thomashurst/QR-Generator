package main

import (
	"encoding/base64"
	"fmt"
	"image/color"
	"log"

	"github.com/skip2/go-qrcode"
)

var (
	imageData    []byte
	imageSize    = 256 // height/width in px
	err          error
	dataToEncode string
)

func main() {
	fmt.Printf("Please paste URL in: ")
	fmt.Scanln(&dataToEncode) // Enter URL to QR
	imageData, err = qrcode.Encode(dataToEncode, qrcode.Medium, imageSize)
	if err != nil {
		log.Fatal("Error generating QR code. ", err)
	}
	// Take QR parameters and convert in base64 encoded string
	encodedImageData := base64.StdEncoding.EncodeToString(imageData)
	// Code to have a live refresh of QR code
	imgTag := "<img src=\"data:image/png;base64" + encodedImageData + "\" />"
	fmt.Println(imgTag)
	// Generate QR
	err = qrcode.WriteColorFile(dataToEncode, qrcode.Medium, imageSize, color.White, color.RGBA{0xE2, 0x4F, 0x63, 0xff}, "image/qr.png")
	if err != nil {
		log.Fatal("Error generating QR code to file. ", err)
	}
}
