package main

import (
	"log"

	"github.com/amsokol/go-errors"
)

func main() {
	log.Print(errors.NewM("message 1"))
	log.Print(errors.NewMf("message 2: %s", "param"))
	log.Print(errors.NewMff(errors.Fields{"f1": "val1", "f2": 15}, "message 3: %s", "param"))
}
