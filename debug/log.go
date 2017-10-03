package debug

import (
	"log"
	"os"
)

var MemoryLeakLogger = log.New(os.Stderr, "MEMORY LEAK ", log.LstdFlags)
