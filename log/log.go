package log

import (
	"log"
	"os"
)

var LogMemoryLeak = log.New(os.Stderr, "MEMORY LEAK ", log.LstdFlags)
