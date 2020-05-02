package gorky

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"net/http"
)

func imgFromUnsplash(collectionID string, s *Size) (*BackgroundImage, error) {
	if s == nil {
		s = &Size{defaultHeight, defaultWidth}
	}

	url := fmt.Sprintf("https://source.unsplash.com/collection/%s/%s", collectionID, s)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var bufferRead bytes.Buffer

	// Get BackgroundImage data
	r := io.TeeReader(res.Body, &bufferRead)

	data, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}

	// Get BackgroundImage Size
	c, _, err := image.DecodeConfig(bytes.NewReader(bufferRead.Bytes()))
	if err != nil {
		return nil, err
	}

	return &BackgroundImage{
		Image: data,
		Size:  Size{c.Width, c.Height},
	}, nil
}
