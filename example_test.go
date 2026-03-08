package exprtk_test

import (
	"fmt"

	"github.com/Pramod-Devireddy/go-exprtk"
)

func ExampleEval() {
	result, err := exprtk.Eval("(x + 2) * (y - 2)", map[string]float64{"x": 18, "y": 32})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
	// Output: 600
}

func ExampleGoExprtk() {
	exprtkObj := exprtk.NewExprtk()
	defer exprtkObj.Delete()

	exprtkObj.SetExpression("(x + 2) * (y - 2)")

	exprtkObj.AddDoubleVariable("x")
	exprtkObj.AddDoubleVariable("y")

	err := exprtkObj.CompileExpression()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	exprtkObj.SetDoubleVariableValue("x", 18)
	exprtkObj.SetDoubleVariableValue("y", 32)

	fmt.Println(exprtkObj.GetEvaluatedValue())
	// Output: 600
}

func ExampleEval_trigonometry() {
	result, _ := exprtk.Eval("sin(pi / 2)", map[string]float64{})
	fmt.Println(result)
	// Output: 1
}
