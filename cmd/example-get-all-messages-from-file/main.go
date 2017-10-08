package main

/*
#cgo LDFLAGS: -leccodes -leccodes_memfs -lpng -laec -ljasper -lopenjp2 -lz -lm
*/
import "C"

import (
	"flag"
	"io"
	"log"
	"runtime/debug"
	"time"

	"github.com/amsokol/go-errors"

	"github.com/amsokol/go-eccodes"
	cio "github.com/amsokol/go-eccodes/io"
)


func main() {
	filename := flag.String("file", "", "io path, e.g. /tmp/ARPEGE_0.1_SP1_00H12H_201709290000.grib2")

	flag.Parse()

	f, err := cio.OpenFile(*filename, "r")
	if err != nil {
		log.Fatalf("failed to open file on file system: %s", err.Error())
	}
	defer f.Close()

	file, err := codes.OpenFile(f)
	if err != nil {
		log.Fatalf("failed to open file: %s", err.Error())
	}
	defer file.Close()

	n := 0
	for {
		err = process(file, n)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("failed to get message (#%d) from index: %s", n, err.Error())
		}
		n++
	}
}

func process(file codes.File, n int) error {
	start := time.Now()

	msg, err := file.Next()
	if err != nil {
		return err
	}
	defer msg.Close()

	log.Printf("============= BEGIN MESSAGE N%d ==========\n", n)

	shortName, err := msg.GetString("shortName")
	if err != nil {
		return errors.Wrap(err, "failed to get 'shortName' value")
	}
	name, err := msg.GetString("name")
	if err != nil {
		return errors.Wrap(err, "failed to get 'name' value")
	}

	log.Printf("Variable = [%s](%s)\n", shortName, name)

	// just to measure timing
	_, _, _, err = msg.Data()
	if err != nil {
		return errors.Wrap(err, "failed to get data (latitudes, longitudes, values)")
	}

	log.Printf("elapsed=%.0f ms", time.Since(start).Seconds()*1000)
	log.Printf("============= END MESSAGE N%d ============\n\n", n)

	debug.FreeOSMemory()

	return nil
}
