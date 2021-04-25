package main

import (
	"fmt"

	"github.com/Pramod-Devireddy/go-exprtk"
)

func main() {
	// Create a new exprtk instance
	exprtkObj := exprtk.NewExprtk()

	// Set the expression
	exprtkObj.SetExpression("(x + 2)*(y-2)")

	// Add variables of expression
	exprtkObj.AddDoubleVariable("x")
	exprtkObj.AddDoubleVariable("y")

	// Compile the expression after expression and variables declaration
	err := exprtkObj.CompileExpression()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Set values for the variables
	exprtkObj.SetDoubleVariableValue("x", 18)
	exprtkObj.SetDoubleVariableValue("y", 32)

	// Get the evaluated value
	fmt.Println(exprtkObj.GetEvaluatedValue())

	// Modify values for the variables
	exprtkObj.SetDoubleVariableValue("y", 42)

	// Get re-evaluated value
	fmt.Println(exprtkObj.GetEvaluatedValue())
}
