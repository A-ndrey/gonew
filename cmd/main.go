package main

import (
	"github.com/A-ndrey/gonew/internal/builder"
	"log"
)

func main() {
	b := builder.CreateFromUserChoices()
	err := b.Build()
	if err != nil {
		log.Fatalln(err)
	}
}
