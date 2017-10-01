package codes

import (
	"runtime"

	"github.com/BCM-ENERGY-team/go-eccodes/native"
)

type Context interface {
	IsOpen() bool
	Native() native.Ccodes_context
	Close()
}

type context struct {
	context native.Ccodes_context
}

func DefaultContext() Context {
	ctx := native.Ccodes_context_get_default()
	runtime.SetFinalizer(ctx, contextFinalizer)
	return &context{context: ctx}
}

func (c *context) IsOpen() bool {
	return c.context != nil
}

func (c *context) Native() native.Ccodes_context {
	return c.context
}

func (c *context) Close() {
	defer func() { c.context = nil }()
	native.Ccodes_context_delete(c.context)
}

func contextFinalizer(c *context) {
	if c.IsOpen() {
		logMemoryLeak.Print("context is not closed")
		c.Close()
	}
}
