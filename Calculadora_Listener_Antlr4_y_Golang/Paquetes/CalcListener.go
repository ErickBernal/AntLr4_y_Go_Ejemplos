package Paquetes

import (
	"fmt"
	"pruebaGo/parser"
	"strconv"
)

/*
	NOTA:
		El listener esta a la espera de cualquier cambio,
		por ello ofrece metodos de Enter() y Exit(), respectivos
		para cada produccion, que se accionan al momento de alguna modificacion
		(esto ocurre por cualquier reconocimiento de un terminal o no terminal).
*/
type CalcListener struct {
	*parser.BaseCalcListener
	Stack []int
}

func (x *CalcListener) Push(i int) {
	x.Stack = append(x.Stack, i)
}

func (x *CalcListener) Pop() int {
	if len(x.Stack) < 1 {
		panic("< empty stack >")
	}
	result := x.Stack[len(x.Stack)-1]
	x.Stack = x.Stack[:len(x.Stack)-1]
	return result
}

func (x *CalcListener) ExitSum(c *parser.SumContext) {
	right, left := x.Pop(), x.Pop()
	fmt.Println("+ <left:", left, " , right:", right, ">")
	x.Push(left + right)
}

func (x *CalcListener) ExitMul(c *parser.MulContext) {
	right, left := x.Pop(), x.Pop()
	fmt.Println("* <left:", left, " , right:", right, ">")

	x.Push(left * right)
}

func (x *CalcListener) ExitDigit(c *parser.DigitContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("<dig:", i, ">")
	x.Push(i)
}

func (x *CalcListener) ExitPassE(c *parser.PassEContext) {
	f := x.Pop()
	fmt.Println("	<Par: ", f, " , text: ", c.GetText(), " >")
	x.Push(f)
}
