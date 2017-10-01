package codes

import (
	"log"
	"os"
)

var logMemoryLeak = log.New(os.Stderr, "MEMORY LEAK", log.LstdFlags)
