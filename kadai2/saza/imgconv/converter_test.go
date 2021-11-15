package imgconv_test

import (
	"testing"

	"github.com/saza-ku/gopherdojo-studyroom/kadai2/saza/imgconv"
)

func TestChangeExt(t *testing.T) {
	cases := []struct {name string; inputName string; inputExt imgconv.ExportFileType; expected string}{
		{name: "jpeg to png", inputName: "hoge.jpeg", inputExt: imgconv.ExportPngType, expected: "hoge.png"},
		{name: "png to gif", inputName: "hoge.png", inputExt: imgconv.ExportGifType, expected: "hoge.gif"},
		{name: "gif to jpg", inputName: "hoge.gif", inputExt: imgconv.ExportJpegType, expected: "hoge.jpg"},
		{name: "jpeg to jpg", inputName: "hoge.jpeg", inputExt: imgconv.ExportJpegType, expected: "hoge.jpg"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := imgconv.ExportChangeExt(c.inputName, c.inputExt); c.expected != actual {
				t.Errorf(
					"want changeExt(%s, %s) = %s, got %s",
					c.inputName, c.inputExt, c.expected, actual)}})
	}
}
