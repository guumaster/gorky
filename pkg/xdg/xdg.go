package xdg

import (
	"fmt"
	"os"

	"github.com/OpenPeeDeeP/xdg"
)

func New() (*xdg.XDG, error) {
	d := xdg.New("", "gorky")
	dirs := []string{
		d.CacheHome(),
		d.DataHome(),
		d.ConfigHome(),
	}

	for _, d := range dirs {
		err := os.MkdirAll(d, 0755)
		if err != nil {
			return nil, fmt.Errorf("failed to create XDG folder %q: %w", d, err)
		}
	}

	return d, nil
}
