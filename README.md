# Go FFMPEG Transcode

```shell
go get github.com/mattcanty/go-ffmpeg-transcode
```

See the [basic example][1]:

```golang
package main

import (
	"log"

	"github.com/matt.canty/go-ffmpeg-transcode/pkg/ffmpeg"
)

func main() {
	err := ffmpeg.Transcode("input.webm", "output.mp3")

	if err != nil {
		log.Fatal(err)
	}
}

```

[1]: ./examples/basic/main.go
