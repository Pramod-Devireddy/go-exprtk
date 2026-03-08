package exprtk

import (
	"math"
	"testing"
)

func TestVectorIndexing(t *testing.T) {
	exp := NewExprtk()
	defer exp.Delete()

	exp.SetExpression("v[0] + v[1] + v[2] + v[3]")
	exp.AddVectorVariable("v")
	// Set vector values before compilation so ExprTk knows the size
	exp.SetVectorVariableValue("v", []float64{4.0, 8.0, 12.0, 16.0})

	err := exp.CompileExpression()
	if err != nil {
		t.Fatalf("compile failed: %v", err)
	}

	result := exp.GetEvaluatedValue()
	if result != 40.0 {
		t.Errorf("expected 40.0, got %f", result)
	}
}

func TestVectorIndexSingleElement(t *testing.T) {
	exp := NewExprtk()
	defer exp.Delete()

	exp.SetExpression("v[1]")
	exp.AddVectorVariable("v")
	exp.SetVectorVariableValue("v", []float64{4.0, 8.0, 12.0, 16.0})

	err := exp.CompileExpression()
	if err != nil {
		t.Fatalf("compile failed: %v", err)
	}

	result := exp.GetEvaluatedValue()
	if result != 8.0 {
		t.Errorf("expected 8.0, got %f", result)
	}
}

func TestVectorUpdateSameSize(t *testing.T) {
	exp := NewExprtk()
	defer exp.Delete()

	exp.SetExpression("v[0] + v[1]")
	exp.AddVectorVariable("v")
	exp.SetVectorVariableValue("v", []float64{1.0, 2.0})

	err := exp.CompileExpression()
	if err != nil {
		t.Fatalf("compile failed: %v", err)
	}

	if exp.GetEvaluatedValue() != 3.0 {
		t.Errorf("expected 3.0, got %f", exp.GetEvaluatedValue())
	}

	// Update with same size — no recompilation needed
	exp.SetVectorVariableValue("v", []float64{10.0, 20.0})
	if exp.GetEvaluatedValue() != 30.0 {
		t.Errorf("expected 30.0, got %f", exp.GetEvaluatedValue())
	}
}


func TestVectorWithMathFunctions(t *testing.T) {
	exp := NewExprtk()
	defer exp.Delete()

	exp.SetExpression("avg(v)")
	exp.AddVectorVariable("v")
	exp.SetVectorVariableValue("v", []float64{10.0, 20.0, 30.0})

	err := exp.CompileExpression()
	if err != nil {
		t.Fatalf("compile failed: %v", err)
	}

	result := exp.GetEvaluatedValue()
	if math.Abs(result-20.0) > 1e-10 {
		t.Errorf("expected 20.0, got %f", result)
	}
}
