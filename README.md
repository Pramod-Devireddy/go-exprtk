<p align="center">
    <h1 align="center"> Go Mathematical Expression Toolkit</h1>
</p>

<p align="center">
  <a href="https://godoc.org/github.com/Pramod-Devireddy/go-exprtk">
  	<img src="https://godoc.org/github.com/Pramod-Devireddy/go-exprtk?status.svg" alt="Go Doc Reference">
  </a>
  <a href="https://travis-ci.com/Pramod-Devireddy/go-exprtk">
  <img src="https://www.travis-ci.com/Pramod-Devireddy/go-exprtk.svg?branch=main" alt="Build Status">
  </a>
  <a href="https://codecov.io/gh/Pramod-Devireddy/go-exprtk">
  	<img src="https://codecov.io/gh/Pramod-Devireddy/go-exprtk/branch/main/graph/badge.svg?token=9pNExL9yZt"/>
  </a>
  <a href="https://goreportcard.com/report/github.com/Pramod-Devireddy/go-exprtk">
  	<img src="https://goreportcard.com/badge/github.com/Pramod-Devireddy/go-exprtk" alt="Go Report Card">
  </a>
</p>



The Go Expression Toolkit (Go-ExprTk) is a wrapper library based on C++ Mathematical Expression Toolkit Library ([ExprTk](http://www.partow.net/programming/exprtk/)).

It is a simple to use, easy to integrate and extremely efficient run-time mathematical expression parser and evaluation engine.

Go-ExprTk supports numerous forms of functional, logical and vector processing semantics and is very easily extendible.

## Installation

```bash
go get github.com/Pramod-Devireddy/go-exprtk
```

## Examples

```Go
package main

import (
	"fmt"

	"github.com/Pramod-Devireddy/go-exprtk"
)

func main() {
	exprtkObj := exprtk.NewExprtk()
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

	fmt.Println(exprtkObj.GetEvaluatedValue())
}
```

Check out more [Examples](https://github.com/Pramod-Devireddy/go-exprtk/tree/main/examples) 


## ⚡ Features

The Go-ExprTk library has the following capabilities:

* Mathematical operators (+, -, *, /, %, ^)

* Functions (min, max, avg, sum, abs, ceil, floor, round, roundn, exp, log, log10, logn, pow, root, sqrt, clamp, inrange, swap)

* Trigonometry (sin, cos, tan, acos, asin, atan, atan2, cosh, cot, csc, sec, sinh, tanh, d2r, r2d, d2g, g2d, hyp)

* Equalities & Inequalities (=, ==, <>, !=, <, <=, >, >=)

* Assignment (:=, +=, -=, *=, /=, %=)

* Logical operators (and, nand, nor, not, or, xor, xnor, mand, mor)

* Control structures (if-then-else, ternary conditional, switch case, return-statement)

* Loop structures (while loop, for loop, repeat until loop, break, continue)

* Optimization of expressions (constant folding, strength reduction, operator coupling, special functions and dead code elimination)

* String operations (equalities, inequalities, logical operators, concatenation and sub-ranges)

* Expression local variables, vectors and strings

* User defined variables, vectors, strings, constants and function support

* Multivariate function composition

* Multiple sequence point and sub expression support

* Numeric integration and differentiation

* Vector Processing: BLAS-L1 (axpy, axpby, axpb), all/any-true/false, count, rotate-left/right, shift-left/right, sort, nth_element, iota, sum, kahan-sum, dot-product, copy


## ❤️ Credits

This module could not be possible without the [ExprTk](http://www.partow.net/programming/exprtk/) library by [Arash Partow](https://github.com/ArashPartow) and the idea of creating the wrapper module by [Narayana Rao G S](https://twitter.com/narayanraogs)


Published under MIT License