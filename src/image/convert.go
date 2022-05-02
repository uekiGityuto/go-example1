package image

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Converter struct {
	Directory string
}

func (converter Converter) JpegToPng() {
	files, err := ioutil.ReadDir(converter.Directory)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() {
			directory := filepath.Join(converter.Directory, fileName)
			childConverter := Converter{Directory: directory}
			childConverter.JpegToPng()
			continue
		}
		if filepath.Ext(fileName) == ".jpeg" || filepath.Ext(fileName) == ".jpg" {
			converter.convert(file)
		}
	}
}

func (converter Converter) convert(fileInfo fs.FileInfo) {
	fileName := fileInfo.Name()
	file, err := os.Open(filepath.Join(converter.Directory, fileName))
	defer file.Close()
	if err != nil {
		panic(err)
	}

	output, err := os.Create(filepath.Join(converter.Directory, fileName) + ".png")
	defer output.Close()
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	err = png.Encode(output, img)
	if err != nil {
		panic(err)
	}
}
