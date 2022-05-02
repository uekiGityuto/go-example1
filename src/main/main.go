package main

import (
	"flag"
	"fmt"
	"github.com/uekiGityuto/go-example/src/image"
)

func main() {
	var (
		directory = flag.String("directory", "resources", "target directory")
	)
	flag.Parse()
	fmt.Println(image.JpegToPng(*directory))
}
