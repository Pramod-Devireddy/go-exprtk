package exprtk

// #cgo CXXFLAGS: -std=c++11
// #cgo LDFLAGS: -L.
// #include "exprtkwrapper.h"
import "C"
import (
	"errors"
)

// GoExprtk ...Exprtk Structure
type GoExprtk struct {
	exprtk C.exprtkWrapper
}

// NewExprtk ... Creates a new object
func NewExprtk() GoExprtk {
	var obj GoExprtk
	obj.exprtk = C.exprtkWrapperInit()
	return obj
}

// SetExpression ... Sets an Expression
func (obj GoExprtk) SetExpression(expr string) {
	C.setExpressionString(obj.exprtk, C.CString(expr))
}

// AddDoubleVariable ... Adds variable to the expression
func (obj GoExprtk) AddDoubleVariable(x string) {
	C.addDoubleVariable(obj.exprtk, C.CString(x))
}

// AddStringVariable ... Adds variable to the expression
func (obj GoExprtk) AddStringVariable(x string) {
	C.addStringVariable(obj.exprtk, C.CString(x))
}

// AddVectorVariable ... Adds variable to the expression
func (obj GoExprtk) AddVectorVariable(x string) {
	C.addVectorVariable(obj.exprtk, C.CString(x))
}

// SetDoubleVariableValue ... Sets value to the variable
func (obj GoExprtk) SetDoubleVariableValue(varName string, val float64) {
	C.setDoubleVariableValue(obj.exprtk, C.CString(varName), C.double(val))
}

// SetStringVariableValue ... Sets value to the variable
func (obj GoExprtk) SetStringVariableValue(varName string, val string) {
	C.setStringVariableValue(obj.exprtk, C.CString(varName), C.CString(val))
}

// SetVectorVariableValue ... Sets value to the variable
func (obj GoExprtk) SetVectorVariableValue(varName string, val []float64) {
	arr := make([]C.double, 0)
	for i := 0; i < len(val); i++ {
		arr = append(arr, C.double(val[i]))
	}
	firstValue := &(arr[0])
	var arrayLength C.int = C.int(len(arr))
	C.setVectorVariableValue(obj.exprtk, C.CString(varName), firstValue, arrayLength)
}

// CompileExpression ... Compiles the Expression
func (obj GoExprtk) CompileExpression() error {
	value := C.compileExpression(obj.exprtk)
	if value == 0 {
		return errors.New("failed to compile the expression")
	}
	return nil
}

// GetEvaluatedValue ... Returns the evaluated value
func (obj GoExprtk) GetEvaluatedValue() float64 {
	return float64(C.getEvaluatedValue(obj.exprtk))
}

// Delete ... Destroys the created object and releases the memory
func (obj GoExprtk) Delete() {
	C.deleteExprtk(obj.exprtk)
}
