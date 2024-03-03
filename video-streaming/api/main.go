package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

var (
	err    error
	webcam *gocv.VideoCapture
)

func main() {
	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer webcam.Close()
	img := gocv.NewMat() // これどういう意味だろう
	defer img.Close()

	window := gocv.NewWindow("Video Capture")

	for {
		ok := webcam.Read(&img)
		if !ok {
			fmt.Printf("Device closed %v\n", ok)
			return
		}

		window.IMShow(img)
		window.WaitKey(1)
	}
}
