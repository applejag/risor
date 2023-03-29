package vm

import (
	"github.com/cloudcmds/tamarin/internal/compiler"
	"github.com/cloudcmds/tamarin/object"
)

const DefaultFrameLocals = 4

type Frame struct {
	fn             *object.CompiledFunction
	scope          *compiler.Scope
	returnAddr     int
	localsCount    int
	locals         [DefaultFrameLocals]object.Object
	extendedLocals []object.Object
}

func (f *Frame) Init(fn *object.CompiledFunction, returnAddr int, localsCount int) {
	f.fn = fn
	if fn != nil {
		f.scope = fn.Scope().(*compiler.Scope)
	} else {
		f.scope = nil
	}
	f.returnAddr = returnAddr
	f.localsCount = localsCount
	if localsCount > DefaultFrameLocals {
		f.extendedLocals = make([]object.Object, localsCount)
	}
}

func (f *Frame) InitWithLocals(fn *object.CompiledFunction, returnAddr int, locals []object.Object) {
	count := len(locals)
	f.Init(fn, returnAddr, count)
	if count > DefaultFrameLocals {
		copy(f.extendedLocals, locals)
	} else {
		copy(f.locals[:], locals)
	}
}

func (f *Frame) Locals() []object.Object {
	if f.localsCount > DefaultFrameLocals {
		return f.extendedLocals
	}
	return f.locals[:f.localsCount]
}

func (f *Frame) Function() *object.CompiledFunction {
	return f.fn
}

func (f *Frame) Scope() *compiler.Scope {
	return f.scope
}
