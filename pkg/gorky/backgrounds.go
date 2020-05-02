package gorky

import (
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/reujab/wallpaper"

	"github.com/guumaster/gorky/fonts"
)

func setNewBackground(img *BackgroundImage, dataDir string) (string, error) {
	dc := gg.NewContext(img.W, img.H)
	dc.DrawImage(img, 0, 0)

	err := addWatermark(dc, img.Size)
	if err != nil {
		return "", err
	}

	newBack, err := ioutil.TempFile(dataDir, "gorky_*.png")
	if err != nil {
		return "", err
	}

	log.Printf("Saving background to %s\n", newBack.Name())

	err = imgio.Save(newBack.Name(), dc.Image(), imgio.PNGEncoder())
	if err != nil {
		return "", err
	}

	err = wallpaper.SetFromFile(newBack.Name())
	if err != nil {
		return "", err
	}

	return newBack.Name(), nil
}

func addWatermark(dc *gg.Context, s Size) error {
	dc.SetRGBA(1, 1, 1, .3)

	err := fonts.SetFontFace(dc, "Ubuntu-Medium.ttf", &truetype.Options{Size: 15})
	if err != nil {
		return err
	}

	text := "Gorky wallpaper @Unsplash"
	w, h := dc.MeasureString(text)

	margin := 15.0
	x := float64(s.W) - w - (margin * 3.)
	y := float64(s.H) - h - (margin * 2.)

	dc.SetColor(color.RGBA{A: 120})

	w += margin * 2.
	h += margin * 1.
	dc.DrawRoundedRectangle(x, y, w, h, margin/2)
	dc.Fill()

	dc.SetRGBA(1, 1, 1, .7)

	dc.DrawStringAnchored(text, x+margin, y+(margin*0.8), 0, .5)

	return nil
}

func cleanOldBackgrounds(root string, current string) error {
	ext := regexp.MustCompile(".png$")

	return filepath.Walk(root, func(p string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if p == current || f.IsDir() {
			return nil
		}

		isImg := ext.MatchString(f.Name())

		if isImg {
			return os.Remove(p)
		}
		return nil
	})
}
