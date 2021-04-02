package ffmpeg

import (
	"fmt"
	"io"
	"os/exec"
	"runtime"

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

	if runtime.GOOS == "windows" {
		err := command.Run()

		return err
	} else {
		ptyFile, err := pty.Start(command)

		if err != nil {
			return err
		}

		progress := &Progress{}

		_, err = io.Copy(progress, ptyFile)

		return err
	}
}

func (progress *Progress) Write(p []byte) (int, error) {
	n := len(p)

	err := re.MatchToTarget(string(p), progress)

	fmt.Printf("Processed kB: %d\n", progress.Processed)

	return n, err
}
