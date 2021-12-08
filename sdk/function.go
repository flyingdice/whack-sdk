package sdk

import (
	"github.com/flyingdice/whack-sdk/sdk/exports"
	"github.com/pkg/errors"
	"reflect"
)

// Func represents a generic golang host func that can be called by WASM through runtime translation.
type Func func(...interface{}) (interface{}, error)

// Function represents a generic function that can be imported into a module.
type Function interface {
	Name() string
	NumIn() int
	NumOut() int
	Void() bool
	Func() Func
}

type function struct {
	name   string
	numIn  int
	numOut int
	fn     Func
}

func (f *function) Name() string { return f.name }
func (f *function) NumIn() int   { return f.numIn }
func (f *function) NumOut() int  { return f.numOut }
func (f *function) Void() bool   { return f.numOut == 0 }
func (f *function) Func() Func   { return f.fn }

// NewFunction creates a new function from the given name/golang Func.
func NewFunction(name string, fn exports.Export) *function {
	fnType := reflect.TypeOf(fn)
	return &function{
		name:   name,
		numIn:  fnType.NumIn(),
		numOut: fnType.NumOut(),
		fn: func(args ...interface{}) (interface{}, error) {
			unwrapped, err := unwrap(args)
			if err != nil {
				return nil, errors.Wrap(err, "failed to unwrap function args")
			}
			return fn(unwrapped...)
		},
	}
}

// unwrap takes a variable number of interface{} values and converts them to int32's.
// This is used to unpack function arguments so they can be passed through the SDK FFI
// and be called from WASM.
func unwrap(args ...interface{}) ([]int32, error) {
	retval := make([]int32, len(args))
	for i, arg := range args {
		asserted, ok := arg.(int32)
		if !ok {
			return nil, errors.Errorf("arg %d type assert to int32 failed", i)
		}
		retval[i] = asserted
	}
	return retval, nil
}
