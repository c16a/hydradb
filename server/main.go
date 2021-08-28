package main

import (
	"github.com/c16a/hydradb/internal"
	"log"
)

func main() {
	config := &internal.Config{
		Storage: &internal.Storage{
			Directory: "/tmp/hydradb",
		},
	}
	db, err := internal.InitDb(config)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Validate()
	if err != nil {
		log.Fatalln(err)
	}
}
