<p align="center">
  <h1 align="center">Go-ExprTk</h1>
  <p align="center">Blazing fast mathematical expression evaluation for Go, powered by C++ ExprTk</p>
</p>

<p align="center">
  <a href="https://pkg.go.dev/github.com/Pramod-Devireddy/go-exprtk">
    <img src="https://pkg.go.dev/badge/github.com/Pramod-Devireddy/go-exprtk.svg" alt="Go Reference">
  </a>
  <a href="https://github.com/Pramod-Devireddy/go-exprtk/actions/workflows/ci.yml">
    <img src="https://github.com/Pramod-Devireddy/go-exprtk/actions/workflows/ci.yml/badge.svg" alt="CI">
  </a>
  <a href="https://codecov.io/gh/Pramod-Devireddy/go-exprtk">
    <img src="https://codecov.io/gh/Pramod-Devireddy/go-exprtk/branch/main/graph/badge.svg?token=9pNExL9yZt"/>
  </a>
  <a href="https://goreportcard.com/report/github.com/Pramod-Devireddy/go-exprtk">
    <img src="https://goreportcard.com/badge/github.com/Pramod-Devireddy/go-exprtk" alt="Go Report Card">
  </a>
  <a href="https://github.com/Pramod-Devireddy/go-exprtk/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/Pramod-Devireddy/go-exprtk" alt="License">
  </a>
</p>

---

Go-ExprTk is a Go wrapper for the [C++ Mathematical Expression Toolkit Library (ExprTk)](http://www.partow.net/programming/exprtk/). It provides a simple API for parsing and evaluating mathematical expressions at runtime with **C++ performance**.

## Why Go-ExprTk?

| Feature | Go-ExprTk | [govaluate](https://github.com/Knetic/govaluate) | [expr](https://github.com/expr-lang/expr) |
|---|---|---|---|
| Speed (eval-only) | ~185ns | ~500ns | ~150ns |
| Speed (compile+eval) | ~330us | ~3us | ~5us |
| Math functions (sin, cos, etc.) | 50+ built-in | Manual | Manual |
| Vector/array operations | Built-in (BLAS-L1) | No | Limited |
| Control structures (if/for/while) | Full support | No | Limited |
| String operations | Yes | Limited | Yes |
| Loop structures | Yes | No | No |

**Best for**: Scientific computing, rule engines, config-driven calculations, financial modeling, and any use case where you need to evaluate user-defined mathematical expressions with maximum performance.

## Installation

```bash
go get github.com/Pramod-Devireddy/go-exprtk
```

**Requirements**: A C++ compiler (GCC, Clang, or MSVC) is needed since Go-ExprTk uses cgo.

## Quick Start

```go
// One-liner evaluation
result, err := exprtk.Eval("(x + 2) * (y - 2)", map[string]float64{"x": 18, "y": 32})
// result = 600
```

## Usage

### Simple Expressions

```go
package main

import (
    "fmt"
    "github.com/Pramod-Devireddy/go-exprtk"
)

func main() {
    exprtkObj := exprtk.NewExprtk()
    defer exprtkObj.Delete()

    exprtkObj.SetExpression("(x + 2) * (y - 2)")

    exprtkObj.AddDoubleVariable("x")
    exprtkObj.AddDoubleVariable("y")

    err := exprtkObj.CompileExpression()
    if err != nil {
        fmt.Println(err)
        return
    }

    exprtkObj.SetDoubleVariableValue("x", 18)
    exprtkObj.SetDoubleVariableValue("y", 32)

    fmt.Println(exprtkObj.GetEvaluatedValue()) // Output: 600

    // Re-evaluate with different values (no recompilation needed)
    exprtkObj.SetDoubleVariableValue("x", 100)
    fmt.Println(exprtkObj.GetEvaluatedValue()) // Output: 3060
}
```

### String and Vector Variables

```go
exprtkObj := exprtk.NewExprtk()
defer exprtkObj.Delete()

eqn := "if (op == 'avg') avg(data); "
eqn += "else if (op == 'sum') sum(data); "
eqn += "else 0;"

exprtkObj.SetExpression(eqn)
exprtkObj.AddStringVariable("op")
exprtkObj.AddVectorVariable("data")

// Set vector values before compilation
exprtkObj.SetVectorVariableValue("data", []float64{1, 2, 3, 4, 5})
exprtkObj.CompileExpression()

exprtkObj.SetStringVariableValue("op", "avg")
fmt.Println(exprtkObj.GetEvaluatedValue()) // Output: 3

exprtkObj.SetStringVariableValue("op", "sum")
fmt.Println(exprtkObj.GetEvaluatedValue()) // Output: 15
```

> **Note**: Vector values should be set before calling `CompileExpression()` to ensure ExprTk allocates the correct buffer size. Double and string values can be set after compilation.

### Convenience Function

For simple one-off evaluations:

```go
result, err := exprtk.Eval("sin(x)^2 + cos(x)^2", map[string]float64{"x": 1.5})
// result = 1.0
```

Check out more [examples](https://github.com/Pramod-Devireddy/go-exprtk/tree/main/examples).

## Performance

Benchmarks on Intel Core i5-8250U @ 1.60GHz:

```
BenchmarkEvaluateOnly-8          6,272,432      185.6 ns/op    0 B/op    0 allocs/op
BenchmarkEval-8                      3,385   330,993 ns/op      0 B/op    0 allocs/op
BenchmarkComplexExpression-8         3,346   373,640 ns/op      0 B/op    0 allocs/op
```

**Key takeaway**: Compile once, evaluate millions of times at ~186ns per evaluation with zero allocations.

## Features

**Mathematical Operators**: `+`, `-`, `*`, `/`, `%`, `^`

**Functions**: min, max, avg, sum, abs, ceil, floor, round, roundn, exp, log, log10, logn, pow, root, sqrt, clamp, inrange, swap

**Trigonometry**: sin, cos, tan, acos, asin, atan, atan2, cosh, cot, csc, sec, sinh, tanh, d2r, r2d, d2g, g2d, hyp

**Equalities & Inequalities**: `=`, `==`, `<>`, `!=`, `<`, `<=`, `>`, `>=`

**Assignment**: `:=`, `+=`, `-=`, `*=`, `/=`, `%=`

**Logical Operators**: and, nand, nor, not, or, xor, xnor, mand, mor

**Control Structures**: if-then-else, ternary conditional, switch case, return-statement

**Loop Structures**: while loop, for loop, repeat until loop, break, continue

**String Operations**: equalities, inequalities, logical operators, concatenation and sub-ranges

**Vector Processing**: BLAS-L1 (axpy, axpby, axpb), all/any-true/false, count, rotate-left/right, shift-left/right, sort, nth_element, iota, sum, kahan-sum, dot-product, copy

**Optimization**: constant folding, strength reduction, operator coupling, special functions, dead code elimination

## API Reference

| Method | Description |
|---|---|
| `NewExprtk()` | Create a new ExprTk instance |
| `SetExpression(expr)` | Set the mathematical expression |
| `AddDoubleVariable(name)` | Register a numeric variable |
| `AddStringVariable(name)` | Register a string variable |
| `AddVectorVariable(name)` | Register a vector variable |
| `CompileExpression()` | Compile the expression (returns error) |
| `SetDoubleVariableValue(name, val)` | Set numeric variable value |
| `SetStringVariableValue(name, val)` | Set string variable value |
| `SetVectorVariableValue(name, val)` | Set vector variable value |
| `GetEvaluatedValue()` | Evaluate and return the result |
| `Delete()` | Free memory (use with defer) |
| `Eval(expr, vars)` | One-liner compile + evaluate |

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## Credits

This module is powered by the [ExprTk](http://www.partow.net/programming/exprtk/) library by [Arash Partow](https://github.com/ArashPartow). The idea of creating the Go wrapper was inspired by [Narayana Rao G S](https://twitter.com/narayanraogs).

## License

Published under [MIT License](LICENSE).
