package codes

import (
	"runtime"

	"github.com/BCM-ENERGY-team/go-eccodes/log"
	"github.com/BCM-ENERGY-team/go-eccodes/native"
	"github.com/pkg/errors"
)

type File interface {
	isOpen() bool
	Native() native.CFILE
	Close() error
}

type file struct {
	filename string
	mode     string
	file     native.CFILE
}

func OpenFile(filename string, mode string) (File, error) {
	fp, err := native.Cfopen(filename, mode)
	if err != nil {
		return nil, errors.Wrapf(err, "failed open file '%s' mode '%s'", filename, mode)
	}

	f := &file{filename: filename, mode: mode, file: fp}
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
	err := native.Cfclose(f.file)
	if err != nil {
		return errors.Wrapf(err, "failed to close '%s' mode '%s'", f.filename, f.mode)
	}
	return nil
}

func fileFinalizer(f *file) {
	if f.isOpen() {
		log.LogMemoryLeak.Printf("file '%s' mode '%s' is not closed", f.filename, f.mode)
		f.Close()
	}
}
