package main

import (
	"fmt"
	"github.com/urusofam/urlShortener/internal/config"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("./config/local.yaml")

	if err != nil {
		log.Fatalf("error loading config: %s", err)
	}

	fmt.Printf("%+v\n", cfg)
}
