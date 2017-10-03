package io

import (
	"runtime"
	"strings"

	"github.com/BCM-ENERGY-team/go-eccodes/debug"
	"github.com/BCM-ENERGY-team/go-eccodes/native"
)

type File interface {
	isOpen() bool
	Native() native.CFILE
	Close() error
}

type file struct {
	debugID string
	file     native.CFILE
}

func OpenFile(path string, mode string) (File, error) {
	fp, err := native.Cfopen(path, mode)
	if err != nil {
		return nil, err
	}

	f := &file{file: fp, debugID: strings.Join([]string{"file=", path, "', mode='", mode, "'"}, "")}
	runtime.SetFinalizer(f, fileFinalizer)
	return f, nil
}

func (f *file) isOpen() bool {
	return f.file != nil
}

func (f *file) Native() native.CFILE {
	return f.file
}

func (f *file) Close() error {
	defer func() { f.file = nil }()
	return native.Cfclose(f.file)
}

func fileFinalizer(f *file) {
	if f.isOpen() {
		debug.MemoryLeakLogger.Printf("'%s' is not closed", f.debugID)
		f.Close()
	}
}
