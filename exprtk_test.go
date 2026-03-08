package exprtk

import (
	"fmt"
	"math"
	"reflect"
	"sync"
	"testing"
)

func TestNewExprtk(t *testing.T) {
	exprtkObj := NewExprtk()
	defer exprtkObj.Delete()

	if reflect.TypeOf(exprtkObj).String() != "exprtk.GoExprtk" {
		t.Error("NewExprtk returned incorrect type")
	}
}

func TestCompileExpression(t *testing.T) {
	// Negative case: completely invalid expression syntax
	exprtkObj := NewExprtk()
	exprtkObj.SetExpression("@@@ invalid +++ expression")
	err := exprtkObj.CompileExpression()
	if err == nil {
		t.Error("expected compilation to fail for invalid expression")
	}
	exprtkObj.Delete()

	// Positive case: valid expression
	exprtkObj2 := NewExprtk()
	defer exprtkObj2.Delete()
	exprtkObj2.SetExpression("x + y")
	exprtkObj2.AddDoubleVariable("x")
	exprtkObj2.AddDoubleVariable("y")
	err = exprtkObj2.CompileExpression()
	if err != nil {
		t.Error("failed to compile the expression")
	}
}

func TestDoubleVariables(t *testing.T) {
	exprtkObj := NewExprtk()
	defer exprtkObj.Delete()

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
	defer exprtkObj.Delete()

	exprtkObj.SetExpression(eqn)
	exprtkObj.AddStringVariable("eqn")
	exprtkObj.AddVectorVariable("x")
	// Set vector values before compilation so ExprTk knows the correct size
	exprtkObj.SetVectorVariableValue("x", array)
	err := exprtkObj.CompileExpression()
	if err != nil {
		t.Fatal(err)
	}

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

// --- New Tests ---

func TestEval(t *testing.T) {
	result, err := Eval("(x + 2) * (y - 2)", map[string]float64{"x": 18, "y": 32})
	if err != nil {
		t.Fatalf("Eval failed: %v", err)
	}
	if result != 600 {
		t.Errorf("expected 600, got %f", result)
	}
}

func TestEvalInvalidExpression(t *testing.T) {
	_, err := Eval("invalid @@@ expression", map[string]float64{"x": 1})
	if err == nil {
		t.Error("expected error for invalid expression")
	}
}

func TestEvalNoVariables(t *testing.T) {
	result, err := Eval("2 + 3 * 4", map[string]float64{})
	if err != nil {
		t.Fatalf("Eval failed: %v", err)
	}
	if result != 14 {
		t.Errorf("expected 14, got %f", result)
	}
}

func TestTrigonometry(t *testing.T) {
	result, err := Eval("sin(x)^2 + cos(x)^2", map[string]float64{"x": 1.5})
	if err != nil {
		t.Fatalf("Eval failed: %v", err)
	}
	if math.Abs(result-1.0) > 1e-10 {
		t.Errorf("expected 1.0, got %f", result)
	}
}

func TestMathFunctions(t *testing.T) {
	tests := []struct {
		name   string
		expr   string
		vars   map[string]float64
		expect float64
	}{
		{"abs", "abs(x)", map[string]float64{"x": -5}, 5},
		{"sqrt", "sqrt(x)", map[string]float64{"x": 16}, 4},
		{"ceil", "ceil(x)", map[string]float64{"x": 2.3}, 3},
		{"floor", "floor(x)", map[string]float64{"x": 2.7}, 2},
		{"exp", "exp(x)", map[string]float64{"x": 0}, 1},
		{"log", "log(x)", map[string]float64{"x": 1}, 0},
		{"pow", "x^y", map[string]float64{"x": 2, "y": 10}, 1024},
		{"modulo", "x % y", map[string]float64{"x": 10, "y": 3}, 1},
		{"min", "min(x, y)", map[string]float64{"x": 3, "y": 7}, 3},
		{"max", "max(x, y)", map[string]float64{"x": 3, "y": 7}, 7},
		{"clamp", "clamp(0, x, 10)", map[string]float64{"x": 15}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Eval(tt.expr, tt.vars)
			if err != nil {
				t.Fatalf("Eval failed: %v", err)
			}
			if math.Abs(result-tt.expect) > 1e-10 {
				t.Errorf("expected %f, got %f", tt.expect, result)
			}
		})
	}
}

func TestConstants(t *testing.T) {
	result, err := Eval("pi", map[string]float64{})
	if err != nil {
		t.Fatalf("Eval failed: %v", err)
	}
	if math.Abs(result-math.Pi) > 1e-10 {
		t.Errorf("expected pi=%f, got %f", math.Pi, result)
	}
}

func TestReEvaluation(t *testing.T) {
	exprtkObj := NewExprtk()
	defer exprtkObj.Delete()

	exprtkObj.SetExpression("x * 2")
	exprtkObj.AddDoubleVariable("x")
	if err := exprtkObj.CompileExpression(); err != nil {
		t.Fatal(err)
	}

	exprtkObj.SetDoubleVariableValue("x", 5)
	if exprtkObj.GetEvaluatedValue() != 10 {
		t.Errorf("expected 10, got %f", exprtkObj.GetEvaluatedValue())
	}

	exprtkObj.SetDoubleVariableValue("x", 100)
	if exprtkObj.GetEvaluatedValue() != 200 {
		t.Errorf("expected 200, got %f", exprtkObj.GetEvaluatedValue())
	}
}

func TestControlStructures(t *testing.T) {
	result, err := Eval("if(x > 0, 1, -1)", map[string]float64{"x": 5})
	if err != nil {
		t.Fatalf("Eval failed: %v", err)
	}
	if result != 1 {
		t.Errorf("expected 1, got %f", result)
	}

	result, err = Eval("if(x > 0, 1, -1)", map[string]float64{"x": -5})
	if err != nil {
		t.Fatalf("Eval failed: %v", err)
	}
	if result != -1 {
		t.Errorf("expected -1, got %f", result)
	}
}

func TestMultipleInstances(t *testing.T) {
	obj1 := NewExprtk()
	defer obj1.Delete()
	obj2 := NewExprtk()
	defer obj2.Delete()

	obj1.SetExpression("x + 1")
	obj1.AddDoubleVariable("x")
	obj1.CompileExpression()
	obj1.SetDoubleVariableValue("x", 10)

	obj2.SetExpression("x * 2")
	obj2.AddDoubleVariable("x")
	obj2.CompileExpression()
	obj2.SetDoubleVariableValue("x", 10)

	if obj1.GetEvaluatedValue() != 11 {
		t.Errorf("obj1: expected 11, got %f", obj1.GetEvaluatedValue())
	}
	if obj2.GetEvaluatedValue() != 20 {
		t.Errorf("obj2: expected 20, got %f", obj2.GetEvaluatedValue())
	}
}

func TestConcurrentEval(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val float64) {
			defer wg.Done()
			result, err := Eval("x * x", map[string]float64{"x": val})
			if err != nil {
				t.Errorf("Eval failed for x=%f: %v", val, err)
				return
			}
			expected := val * val
			if math.Abs(result-expected) > 1e-10 {
				t.Errorf("for x=%f, expected %f, got %f", val, expected, result)
			}
		}(float64(i))
	}
	wg.Wait()
}

func TestComplexExpression(t *testing.T) {
	expr := "((x + y) * (x - y)) / (x * y)"
	result, err := Eval(expr, map[string]float64{"x": 10, "y": 5})
	if err != nil {
		t.Fatalf("Eval failed: %v", err)
	}
	expected := ((10.0 + 5.0) * (10.0 - 5.0)) / (10.0 * 5.0)
	if math.Abs(result-expected) > 1e-10 {
		t.Errorf("expected %f, got %f", expected, result)
	}
}

// --- Benchmarks ---

func BenchmarkEval(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Eval("(x + 2) * (y - 2)", map[string]float64{"x": 18, "y": 32})
	}
}

func BenchmarkCompileAndEvaluate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := NewExprtk()
		obj.SetExpression("(x + 2) * (y - 2)")
		obj.AddDoubleVariable("x")
		obj.AddDoubleVariable("y")
		obj.CompileExpression()
		obj.SetDoubleVariableValue("x", 18)
		obj.SetDoubleVariableValue("y", 32)
		obj.GetEvaluatedValue()
		obj.Delete()
	}
}

func BenchmarkEvaluateOnly(b *testing.B) {
	obj := NewExprtk()
	defer obj.Delete()

	obj.SetExpression("(x + 2) * (y - 2)")
	obj.AddDoubleVariable("x")
	obj.AddDoubleVariable("y")
	obj.CompileExpression()
	obj.SetDoubleVariableValue("x", 18)
	obj.SetDoubleVariableValue("y", 32)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		obj.GetEvaluatedValue()
	}
}

func BenchmarkComplexExpression(b *testing.B) {
	expr := "sin(x)^2 + cos(x)^2 + sqrt(abs(y)) * log(z+1)"
	for i := 0; i < b.N; i++ {
		Eval(expr, map[string]float64{"x": 1.5, "y": -4.0, "z": 9.0})
	}
}
