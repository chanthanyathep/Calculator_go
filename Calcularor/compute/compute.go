package compute

import (
	"errors"
	"go/operators"
	"strconv"
	"strings"
	"text/scanner"
)

func InToPost(in string) string {
	var s scanner.Scanner
	s.Init(strings.NewReader(in))
	Stack := NewStringStack()
	var postfix string
	sqrtStatus := 0

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		if s.TokenText() == "√" {
			sqrtStatus++
			continue
		}
		if isOperand(tok) {
			postfix += "(" + s.TokenText() + ")"
		} else if isOperator(s.TokenText()) {
			if Stack.Pos == -1 || s.TokenText() == "(" {
				Stack.Push(s.TokenText())
			} else if operators.Ops[s.TokenText()].Precedence > operators.Ops[Stack.SafeTop()].Precedence {
				Stack.Push(s.TokenText())
			} else {
				for t := Stack.Pos; t != -1; t-- {
					if operators.Ops[s.TokenText()].Precedence <= operators.Ops[Stack.SafeTop()].Precedence {
						postfix += Stack.SafeTop()
						Stack.SafePop()
					}
					Stack.Push(s.TokenText())
				}
			}
		} else if s.TokenText() == ")" {
			for t := Stack.Pos; t != -1; t-- {
				if Stack.SafeTop() != "(" {
					postfix += Stack.SafeTop()
					Stack.Pop()
				}
			}
			if sqrtStatus != 0 {
				postfix += "√"
				sqrtStatus--
			}
			Stack.Pop()
		}
	}

	for t := Stack.Pos; t != -1; t-- {
		postfix += Stack.SafeTop()
		Stack.Pop()
	}

	return postfix
}

func Evaluate(postfix string) float64 {
	var s scanner.Scanner
	s.Init(strings.NewReader(postfix))
	Stack := NewFloatStack()

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		if s.TokenText() == "(" || s.TokenText() == ")" {
			continue
		} else if isOperand(tok) {
			Stack.Push(safeParseFloat(s.TokenText()))
		} else if isOperator(s.TokenText()) {
			if s.TokenText() != "√" {
				var m = make([]float64, 2)
				for i := 1; i >= 0; i-- {
					m[i] = Stack.SafePop()
				}
				Stack.Push(operators.Ops[s.TokenText()].Operation(m))
			} else {
				var m = make([]float64, 2)
				m[0] = Stack.SafePop()
				Stack.Push(operators.Ops[s.TokenText()].Operation(m))
			}
		}

	}

	return Stack.SafeTop()
}

func isOperand(tok int32) bool {
	return tok == -3 || tok == -4
}

func isOperator(lit string) bool {
	return operators.IsOperator(lit)
}

func parseFloat(lit string) (float64, error) {
	f, err := strconv.ParseFloat(lit, 64)
	if err != nil {
		return 0, errors.New("Cannot parse recognized float: " + lit)
	}
	return f, nil
}

func safeParseFloat(lit string) float64 {
	f, err := parseFloat(lit)
	if err != nil {
		return 0
	}
	return f
}
