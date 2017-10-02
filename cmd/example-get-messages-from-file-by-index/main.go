package main

import (
	"flag"
	"io"
	"log"
	"runtime/debug"

	"github.com/pkg/errors"

	"github.com/BCM-ENERGY-team/go-eccodes"
)

func main() {
	filename := flag.String("file", "", "file path, e.g. /tmp/ARPEGE_0.1_SP1_00H12H_201709290000.grib2")

	flag.Parse()

	filter := map[string]interface{}{
		"level": int64(10),
	}

	index, err := codes.NewIndex(nil, *filename, "r", filter)
	if err != nil {
		log.Fatalf("failed to create index for file: %s", err.Error())
	}
	defer index.Close()

	n := 0
	for {
		msg, err := next(index)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("failed to get message (#%d) from index: %s", n, err.Error())
		}

		log.Printf("============= BEGIN MESSAGE N%d ==========\n", n)

		shortName, _ := msg.Parameters()["shortName"]
		name, _ := msg.Parameters()["name"]

		log.Printf("Variable = [%s](%s)\n", shortName, name)

		log.Printf("============= END MESSAGE N%d ============\n\n", n)
		n++

		debug.FreeOSMemory()
	}
}

func next(index codes.Index) (codes.Message, error) {
	ms, err := index.Next(nil)
	if err != nil {
		return nil, err
	}
	defer ms.Close()

	msg, err := ms.Message()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get message from message stub")
	}

	return msg, nil
}
