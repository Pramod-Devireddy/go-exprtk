package exprtk

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestNewExprtk(t *testing.T) {
	exprtkObj := NewExprtk()

	if reflect.TypeOf(exprtkObj).String() != "exprtk.GoExprtk" {
		t.Error("NewExprtk returned incorrect type")
	}
}

func TestCompileExpression(t *testing.T) {
	exprtkObj := NewExprtk()

	exprtkObj.SetExpression("x.1 + y.1")

	exprtkObj.AddDoubleVariable("x.1")
	exprtkObj.AddDoubleVariable("y.1")

	err := exprtkObj.CompileExpression()
	if err.Error() != "failed to compile the expression" {
		t.Error("negative case failed")
	}

	exprtkObj.SetExpression("x + y")

	exprtkObj.AddDoubleVariable("x")
	exprtkObj.AddDoubleVariable("y")

	err = exprtkObj.CompileExpression()
	if err != nil {
		t.Error("failed to compile the expression")
	}
}

func TestDoubleVariables(t *testing.T) {
	exprtkObj := NewExprtk()

	exprtkObj.SetExpression("(x + 2)*(y-2)")

	exprtkObj.AddDoubleVariable("x")
	exprtkObj.AddDoubleVariable("y")

	err := exprtkObj.CompileExpression()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	exprtkObj.SetDoubleVariableValue("x", 18)
	exprtkObj.SetDoubleVariableValue("y", 32)

	if exprtkObj.GetEvaluatedValue() != 600 {
		t.Error("Incorrect Value")
	}

}

func TestMixedVariables(t *testing.T) {
	var eqn string
	eqn = "if (eqn == 'avg') avg(x); "
	eqn += "else if (eqn == 'max') max(x); "
	eqn += "else if (eqn == 'min') min(x); "
	eqn += "else if (eqn == 'sum') sum(x); "
	eqn += "else 0; "

	var eqnStr string
	var array []float64 = []float64{1, 2, 3, -4.3, 10, -6.5, 7, 8, -1.3}

	exprtkObj := NewExprtk()
	exprtkObj.SetExpression(eqn)
	exprtkObj.AddStringVariable("eqn")
	exprtkObj.AddVectorVariable("x")
	exprtkObj.CompileExpression()
	exprtkObj.SetVectorVariableValue("x", array)

	eqnStr = "avg"
	exprtkObj.SetStringVariableValue("eqn", eqnStr)
	if math.Round(exprtkObj.GetEvaluatedValue()*10)/10 != 2.1 {
		t.Error("Incorrect Value")
	}

	eqnStr = "max"
	exprtkObj.SetStringVariableValue("eqn", eqnStr)
	if exprtkObj.GetEvaluatedValue() != 10 {
		t.Error("Incorrect Value")
	}

	eqnStr = "min"
	exprtkObj.SetStringVariableValue("eqn", eqnStr)
	if exprtkObj.GetEvaluatedValue() != -6.5 {
		t.Error("Incorrect Value")
	}

	eqnStr = "sum"
	exprtkObj.SetStringVariableValue("eqn", eqnStr)
	if exprtkObj.GetEvaluatedValue() != 18.9 {
		t.Error("Incorrect Value")
	}

	eqnStr = "xyz"
	exprtkObj.SetStringVariableValue("eqn", eqnStr)
	if exprtkObj.GetEvaluatedValue() != 0.0 {
		t.Error("Incorrect Value")
	}
}
