package image_test_test

import (
	"github.com/uekiGityuto/go-example1/src/image"
	"testing"
)

func TestJPEGToPNG_正常系(t *testing.T) {
	converter := image.Converter{Directory: "testdata"}
	if err := converter.JPEGToPNG(); err != nil {
		t.Errorf("画像変換失敗。%s", err)
	}
}

func TestJPEGToPNG_異常系_存在しないディレクトリ(t *testing.T) {
	converter := image.Converter{Directory: "illegal_directory"}
	if err := converter.JPEGToPNG(); err == nil {
		t.Errorf("存在しないディレクトリを指定してもエラーにならない。")
	}
}
