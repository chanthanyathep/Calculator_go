package operators

var Ops = map[string]*Operator{}

type Operator struct {
	Name       string
	Precedence int
	Operation  func(args []float64) float64
}

func Register(op *Operator) {
	Ops[op.Name] = op
}

func IsOperator(str string) bool {
	_, exist := Ops[str]
	return exist
}
