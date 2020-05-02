//go:generate broccoli -src=assets -o fonts

package fonts

import (
	"bufio"
	"io/ioutil"
	"path"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
)

// SetFontFace loads and set an existing font face stored in the assets folder
func SetFontFace(dc *gg.Context, font string, opts *truetype.Options) error {
	b, _ := br.Open(path.Join("assets", font))
	r := bufio.NewReader(b)

	fontBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	f, err := truetype.Parse(fontBytes)
	if err != nil {
		return err
	}

	ff := truetype.NewFace(f, opts)

	dc.SetFontFace(ff)

	return nil
}
