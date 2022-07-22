package main

import (
	"github.com/A-ndrey/gonew/builder"
	"log"
)

func main() {
	b := builder.CreateFromUserChoices()
	err := b.Build()
	if err != nil {
		log.Fatalln(err)
	}
}
