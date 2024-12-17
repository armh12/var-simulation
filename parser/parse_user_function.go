package parser

import (
	"fmt"
	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/vm"
	"math"
)

func GetEnvironmentForParse() map[string]interface{} {
	return map[string]interface{}{
		// base functions
		"sqrt": math.Sqrt,
		"log":  math.Log,
		"log2": math.Log2,
		// trigonometric and reverse trigonometric functions
		"sin":  math.Sin,
		"cos":  math.Cos,
		"tan":  math.Tan,
		"asin": math.Asin,
		"acos": math.Acos,
		"atan": math.Atan,
		// hyperbolic trigonometric functions
		"sinh": math.Sinh,
		"cosh": math.Cosh,
		"tanh": math.Tanh,
		// exponential
		"exp":  math.Exp,
		"exp2": math.Exp2,
	}
}

func ParseUserInput(userInput string) (*vm.Program, error) {
	compiledUserInput, err := expr.Compile(userInput)
	if err != nil {
		return nil, fmt.Errorf("error compiling user input: %v", err)
	}
	return compiledUserInput, nil
}
