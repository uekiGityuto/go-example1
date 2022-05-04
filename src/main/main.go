package main

import (
	"flag"
	"fmt"
	"github.com/uekiGityuto/go-example/src/image"
	"os"
)

// JPEGをPNGに変換する関数
func main() {
	var (
		directory = flag.String("directory", "resources", "target directory")
	)
	flag.Parse()
	converter := image.Converter{Directory: *directory}
	if err := converter.JPEGToPNG(); err != nil {
		fmt.Fprintln(os.Stderr, "Error while converting images:", err)
		os.Exit(1)
	} else {
		fmt.Println("Completed")
	}
}
