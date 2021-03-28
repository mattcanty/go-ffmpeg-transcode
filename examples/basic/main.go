package main

import (
	"log"

	"github.com/matt.canty/go-ffmpeg-transcode/pkg/ffmpeg"
)

func main() {
	inFilePath := "input.webm"
	outFilePath := "output.mp3"

	err := ffmpeg.Transcode(inFilePath, outFilePath)
	if err != nil {
		log.Fatal(err)
	}
}
