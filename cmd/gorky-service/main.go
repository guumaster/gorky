package main

import (
	"flag"
	"log"
	"time"

	"github.com/guumaster/gorky/pkg/gorky"
	"github.com/guumaster/gorky/pkg/service"
)

func main() {
	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()

	action := *svcFlag
	hasAction := len(action) > 0

	repeatAfter := time.Hour * 12

	cfg := gorky.MakeConfig()

	if action == "config" {
		err := gorky.CreateConfig(cfg)
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	runner, err := service.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if hasAction {
		err = runner.ManageService(action)
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	runner.RepeatAfter(repeatAfter)
}
