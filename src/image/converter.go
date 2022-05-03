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

func (converter Converter) JPEGToPNG() error {
	files, err := ioutil.ReadDir(converter.Directory)
	if err != nil {
		return err
	}

	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() {
			directory := filepath.Join(converter.Directory, fileName)
			childConverter := Converter{Directory: directory}
			childConverter.JPEGToPNG()
			continue
		}
		if filepath.Ext(fileName) == ".jpeg" || filepath.Ext(fileName) == ".jpg" {
			if err := converter.convert(file); err != nil {
				return err
			}
		}
	}
	return nil
}

func (converter Converter) convert(fileInfo fs.FileInfo) error {
	fileName := fileInfo.Name()
	file, err := os.Open(filepath.Join(converter.Directory, fileName))
	defer file.Close()
	if err != nil {
		return err
	}

	output, err := os.Create(filepath.Join(converter.Directory, fileName) + ".png")
	defer output.Close()
	if err != nil {
		return err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	err = png.Encode(output, img)
	if err != nil {
		return err
	}

	return nil
}
