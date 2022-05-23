package Paquetes

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"main/parser"
	"strconv"
)

type Visitor struct {
	parser.CalcVisitor
}

func (v *Visitor) Visitor(tree antlr.ParseTree) float64 {
	switch val := tree.(type) {
	case *parser.OpContext:
		return v.VisitOp(val)
	case *parser.DigitContext:
		return v.VisitDigit(val)
	case *parser.ParenContext:
		return v.VisitParen(val)
	default:
		panic("Unknown context")
	}
}

func (v *Visitor) VisitOp(ctx *parser.OpContext) float64 {
	l := v.Visitor(ctx.GetLeft())
	r := v.Visitor(ctx.GetRight())
	op := ctx.GetOp().GetText()
	fmt.Println("op", op)
	switch op {
	case "+":
		return l + r
	case "-":
		return l - r
	case "*":
		return l * r
	case "/":
		return l / r
	}
	return 0
}

func (v *Visitor) VisitDigit(ctx *parser.DigitContext) float64 {
	fmt.Println("digit", ctx.GetText())
	i1, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return i1
}

func (v *Visitor) VisitParen(ctx *parser.ParenContext) float64 {
	tar := v.Visitor(ctx.Expr())
	fmt.Println("parent", ctx.GetText(), " ,valor:", tar)
	return tar
}
