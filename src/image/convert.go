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

func JpegToPng(directory string) []string {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}

	var fileNames []string
	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() {
			fileNames = append(fileNames, JpegToPng(filepath.Join(directory, fileName))...)
			continue
		}
		if filepath.Ext(fileName) == ".jpeg" || filepath.Ext(fileName) == ".jpg" {
			convert(directory, file)
			fileNames = append(fileNames, fileName)
		}
	}
	return fileNames
}

func convert(directory string, fileInfo fs.FileInfo) {
	fileName := fileInfo.Name()
	file, err := os.Open(filepath.Join(directory, fileName))
	defer file.Close()
	if err != nil {
		panic(err)
	}

	output, err := os.Create(filepath.Join(directory, fileName) + ".png")
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
