package main

import (
	"log"

	"github.com/guumaster/gorky/pkg/gorky"
	"github.com/guumaster/gorky/pkg/xdg"
)

func main() {
	dirs, err := xdg.New()
	if err != nil {
		log.Fatal(err)
	}

	err = gorky.Run(dirs)
	if err != nil {
		log.Fatal(err)
	}
}
