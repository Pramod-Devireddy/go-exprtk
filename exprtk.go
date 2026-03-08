package exprtk

// #cgo CXXFLAGS: -flto -fuse-linker-plugin -std=c++11
// #cgo LDFLAGS: -L.
// #include <stdlib.h>
// #include "exprtkwrapper.h"
import "C"
import (
	"errors"
	"unsafe"
)

// GoExprtk is the main struct that wraps the C++ ExprTk library.
type GoExprtk struct {
	exprtk C.exprtkWrapper
}

// NewExprtk creates a new ExprTk instance.
func NewExprtk() GoExprtk {
	var obj GoExprtk
	obj.exprtk = C.exprtkWrapperInit()
	return obj
}

// SetExpression sets the mathematical expression string to be evaluated.
func (obj GoExprtk) SetExpression(expr string) {
	cstr := C.CString(expr)
	defer C.free(unsafe.Pointer(cstr))
	C.setExpressionString(obj.exprtk, cstr)
}

// AddDoubleVariable registers a numeric (double) variable for use in expressions.
func (obj GoExprtk) AddDoubleVariable(x string) {
	cstr := C.CString(x)
	defer C.free(unsafe.Pointer(cstr))
	C.addDoubleVariable(obj.exprtk, cstr)
}

// AddStringVariable registers a string variable for use in expressions.
func (obj GoExprtk) AddStringVariable(x string) {
	cstr := C.CString(x)
	defer C.free(unsafe.Pointer(cstr))
	C.addStringVariable(obj.exprtk, cstr)
}

// AddVectorVariable registers a vector variable for use in expressions.
func (obj GoExprtk) AddVectorVariable(x string) {
	cstr := C.CString(x)
	defer C.free(unsafe.Pointer(cstr))
	C.addVectorVariable(obj.exprtk, cstr)
}

// SetDoubleVariableValue sets the value of a previously registered double variable.
func (obj GoExprtk) SetDoubleVariableValue(varName string, val float64) {
	cstr := C.CString(varName)
	defer C.free(unsafe.Pointer(cstr))
	C.setDoubleVariableValue(obj.exprtk, cstr, C.double(val))
}

// SetStringVariableValue sets the value of a previously registered string variable.
func (obj GoExprtk) SetStringVariableValue(varName string, val string) {
	cName := C.CString(varName)
	defer C.free(unsafe.Pointer(cName))
	cVal := C.CString(val)
	defer C.free(unsafe.Pointer(cVal))
	C.setStringVariableValue(obj.exprtk, cName, cVal)
}

// SetVectorVariableValue sets the value of a previously registered vector variable.
func (obj GoExprtk) SetVectorVariableValue(varName string, val []float64) {
	arr := make([]C.double, len(val))
	for i := 0; i < len(val); i++ {
		arr[i] = C.double(val[i])
	}
	firstValue := &(arr[0])
	var arrayLength C.int = C.int(len(arr))
	cstr := C.CString(varName)
	defer C.free(unsafe.Pointer(cstr))
	C.setVectorVariableValue(obj.exprtk, cstr, firstValue, arrayLength)
}

// CompileExpression compiles the expression after all variables have been registered.
// Returns an error if the expression is invalid or cannot be compiled.
func (obj GoExprtk) CompileExpression() error {
	value := C.compileExpression(obj.exprtk)
	if value == 0 {
		return errors.New("failed to compile the expression")
	}
	return nil
}

// GetEvaluatedValue evaluates the compiled expression and returns the result.
func (obj GoExprtk) GetEvaluatedValue() float64 {
	return float64(C.getEvaluatedValue(obj.exprtk))
}

// Delete destroys the ExprTk instance and releases all associated memory.
// Always call this when done, preferably with defer.
func (obj GoExprtk) Delete() {
	C.deleteExprtk(obj.exprtk)
}

// Eval is a convenience function that compiles and evaluates a mathematical expression
// with the given variables in a single call.
//
// Example:
//
//	result, err := exprtk.Eval("(x + 2) * (y - 2)", map[string]float64{"x": 18, "y": 32})
func Eval(expression string, vars map[string]float64) (float64, error) {
	obj := NewExprtk()
	defer obj.Delete()

	obj.SetExpression(expression)

	for name := range vars {
		obj.AddDoubleVariable(name)
	}

	if err := obj.CompileExpression(); err != nil {
		return 0, err
	}

	for name, val := range vars {
		obj.SetDoubleVariableValue(name, val)
	}

	return obj.GetEvaluatedValue(), nil
}
