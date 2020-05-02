package gorky

import (
	"fmt"
	"image"
	"log"

	"github.com/OpenPeeDeeP/xdg"
)

// https://source.unsplash.com/collection/{collection_id}/1600x900

// Future
// https://wall.alphacoders.com/api.php?lang=Spanish

const defaultWidth = 1600
const defaultHeight = 900

// Size stores with and height of an image
type Size struct {
	W, H int
}

func (s Size) String() string {
	return fmt.Sprintf("%dx%d", s.W, s.H)
}

type BackgroundImage struct {
	image.Image
	Size
}

func Run(dirs *xdg.XDG) error {
	dataDir := dirs.DataHome()

	log.Println("Downloading new image")

	img, err := imgFromUnsplash("10041812", &Size{W: 1600, H: 900})
	if err != nil {
		return err
	}

	log.Println("Changing background")

	newBack, err := setNewBackground(img, dataDir)
	if err != nil {
		return err
	}

	log.Println("Cleaning old backgrounds")

	err = cleanOldBackgrounds(dataDir, newBack)
	if err != nil {
		return err
	}

	return nil
}
