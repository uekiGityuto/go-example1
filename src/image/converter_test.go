package image_test

import (
	"github.com/uekiGityuto/go-example/src/image"
	"testing"
)

func TestJPEGToPNG(t *testing.T) {
	converter := image.Converter{Directory: "test"}
	if err := converter.JPEGToPNG(); err != nil {
		t.Errorf("画像変換失敗。%s", err)
	}
}
