// +build windows

package ffmpeg

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func Transcode(inFilePath string, outFilePath string) error {
	command := exec.Command(
		"ffmpeg",
		"-y",
		"-loglevel",
		"quiet",
		"-stats",
		"-i",

		inFilePath,
		outFilePath,
	)

	out, _ := command.CombinedOutput()

	if !strings.HasPrefix(string(out), "size=") {
		return errors.New(
			fmt.Sprintf("Failed to transcode with commend '%s'", command.Args),
		)
	}

	return nil
}
