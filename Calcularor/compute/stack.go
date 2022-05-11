package compute

import (
	"errors"
)

type StringStack struct {
	Info []string
	Pos  int
}

func NewStringStack() *StringStack {
	return &StringStack{
		Info: []string{},
		Pos:  -1,
	}
}

type FloatStack struct {
	Info []float64
	Pos  int
}

func NewFloatStack() *FloatStack {
	return &FloatStack{
		Info: []float64{},
		Pos:  -1,
	}
}

func (s *StringStack) Push(a string) {
	s.Pos++
	if s.Pos < len(s.Info) {
		s.Info[s.Pos] = a
	} else {
		s.Info = append(s.Info, a)
	}
}

func (s *StringStack) Pop() (string, error) {
	ret, err := s.Top()
	if err != nil {
		return "", errors.New("Can't pop; stack is empty!")
	}
	s.Pos--
	return ret, nil
}

func (s *StringStack) SafePop() string {
	ret, _ := s.Pop()
	return ret
}

func (s *StringStack) Top() (string, error) {
	if s.Pos < 0 {
		return "", errors.New("No elements in stack!")
	}
	return s.Info[s.Pos], nil
}

func (s *StringStack) SafeTop() string {
	ret, _ := s.Top()
	return ret
}

func (s *FloatStack) Push(a float64) {
	s.Pos++
	if s.Pos < len(s.Info) {
		s.Info[s.Pos] = a
	} else {
		s.Info = append(s.Info, a)
	}
}

func (s *FloatStack) Pop() (float64, error) {
	ret, err := s.Top()
	if err != nil {
		return 0, errors.New("Can't pop; stack is empty!")
	}
	s.Pos--
	return ret, nil
}

func (s *FloatStack) SafePop() float64 {
	ret, _ := s.Pop()
	return ret
}

func (s *FloatStack) Top() (float64, error) {
	if s.Pos < 0 {
		return 0, errors.New("No elements in stack!")
	}
	return s.Info[s.Pos], nil
}

func (s *FloatStack) SafeTop() float64 {
	ret, _ := s.Top()
	return ret
}
