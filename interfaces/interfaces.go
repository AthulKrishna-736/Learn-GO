package main

import (
	"fmt"
)

// type InterfaceName interface {
//     Method1() ReturnType
//     Method2(ParameterType) ReturnType
// ... more methods
// }

type MathOperation interface {
	Add(a int, b int) int
	Sub(a int, b int) int
	Mul(a int, b int) int
}

type Calc struct{}

func (c Calc) Add(a int, b int) int {
	return a + b
}

func (c Calc) Sub(a int, b int) int {
	return a - b
}

func (c Calc) Mul(a int, b int) int {
	return a * b
}

var b MathOperation

func main() {
	var calc MathOperation = Calc{}
	fmt.Println("res: ", calc.Add(1, 2), calc.Mul(1, 2), calc.Sub(1, 2))
}
