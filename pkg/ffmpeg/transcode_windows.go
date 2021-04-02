// +build windows

package ffmpeg

import (
	"os/exec"
)

func Transcode(inFilePath string, outFilePath string) error {
	return exec.Command(
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
