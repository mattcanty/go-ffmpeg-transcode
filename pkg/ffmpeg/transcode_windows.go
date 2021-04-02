// +build windows

package ffmpeg

import (
	"os/exec"
)

func Transcode(inFilePath string, outFilePath string) error {
	return exec.Command(
		"cmd",
		"/C",
		"ffmpeg",
		"-y",
		"-loglevel",
		"quiet",
		"-stats",
		"-i",

		inFilePath,
		outFilePath,
	).Run()
}
