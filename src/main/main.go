package main

import (
	"flag"
	"github.com/uekiGityuto/go-example/src/image"
)

func main() {
	var (
		directory = flag.String("directory", "resources", "target directory")
	)
	flag.Parse()
	converter := image.Converter{Directory: *directory}
	converter.JpegToPng()
}
