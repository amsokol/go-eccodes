package codes

import (
	"runtime"

	"github.com/BCM-ENERGY-team/go-eccodes/log"
	"github.com/BCM-ENERGY-team/go-eccodes/native"
)

type Context interface {
	isOpen() bool
	native() native.Ccodes_context
	Close()
}

type context struct {
	defaultContext bool
	context native.Ccodes_context
}

func DefaultContext() Context {
	ctx := native.Ccodes_context_get_default()
	c := &context{defaultContext: true, context: ctx}
	runtime.SetFinalizer(c, contextFinalizer)
	return c
}

func (c *context) isOpen() bool {
	return c.context != nil
}

func (c *context) native() native.Ccodes_context {
	return c.context
}

func (c *context) Close() {
	defer func() { c.context = nil }()
	if !c.defaultContext {
		native.Ccodes_context_delete(c.context)
	}
}

func contextFinalizer(c *context) {
	if c.isOpen() {
		log.LogMemoryLeak.Print("context is not closed")
		c.Close()
	}
}
