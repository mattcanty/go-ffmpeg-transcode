// +build !windows

package ffmpeg

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/creack/pty"
	"github.com/oriser/regroup"
)

type Progress struct {
	Processed int64 `regroup:"processed"`
}

var re = regroup.MustCompile(`(?m)(?P<processed>\d+)kB`)

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

	ptyFile, err := pty.Start(command)

	if err != nil {
		return err
	}

	progress := &Progress{}

	_, err = io.Copy(progress, ptyFile)

	return err

}

func (progress *Progress) Write(p []byte) (int, error) {
	n := len(p)

	err := re.MatchToTarget(string(p), progress)

	fmt.Printf("Processed kB: %d\n", progress.Processed)

	return n, err
}
