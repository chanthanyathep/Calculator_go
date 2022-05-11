package operators

import (
	"math"
)

var (
	add = &Operator{
		Name:       "+",
		Precedence: 1,
		Operation: func(args []float64) float64 {
			return args[0] + args[1]
		},
	}
	sub = &Operator{
		Name:       "-",
		Precedence: 1,
		Operation: func(args []float64) float64 {
			return args[0] - args[1]
		},
	}
	mul = &Operator{
		Name:       "*",
		Precedence: 2,
		Operation: func(args []float64) float64 {
			return args[0] * args[1]
		},
	}
	div = &Operator{
		Name:       "/",
		Precedence: 2,
		Operation: func(args []float64) float64 {
			return args[0] / args[1]
		},
	}
	sqrt = &Operator{
		Name:       "√",
		Precedence: 0,
		Operation: func(args []float64) float64 {
			return math.Sqrt(args[0])
		},
	}
	pow = &Operator{
		Name:       "^",
		Precedence: 3,
		Operation: func(args []float64) float64 {
			return math.Pow(args[0], args[1])
		},
	}
	open = &Operator{
		Name:       "(",
		Precedence: 0,
		Operation: func(args []float64) float64 {
			return 0
		},
	}
)

func init() {
	Register(add)
	Register(sub)
	Register(mul)
	Register(div)
	Register(sqrt)
	Register(pow)
	Register(open)
}
