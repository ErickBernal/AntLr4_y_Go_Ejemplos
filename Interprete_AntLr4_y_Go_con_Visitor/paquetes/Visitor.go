package paquetes

import (
	"fmt"
	"main/parser"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Visitor struct {
	parser.ControlVisitor
	Memory map[string]Value
}

func (v *Visitor) Visit(tree antlr.ParseTree) Value {
	switch val := tree.(type) {
	case *parser.ProgContext:
		return v.VisitProg(val)
	case *parser.BlockContext:
		return v.VisitBlock(val)
	case *parser.StmtContext:
		return v.VisitStmt(val)
	case *parser.AssignstmtContext:
		return v.VisitAssignstmt(val)
	case *parser.IfstmtContext:
		return v.VisitIfstmt(val)
	case *parser.PrintlnstmtContext:
		return v.VisitPrintlnstmt(val)
	case *parser.WhilestmtContext:
		return v.VisitWhilestmt(val)
	case *parser.PrintstmtContext:
		return v.VisitPrintstmt(val)
	case *parser.OpExprContext:
		return v.VisitOpExpr(val)
	case *parser.IntExprContext:
		return v.VisitIntExpr(val)
	case *parser.IdExprContext:
		return v.VisitIdExpr(val)
	case *parser.StrExprContext:
		return v.VisitStrExpr(val)
	case *parser.BoolExprContext:
		return v.VisitBoolExpr(val)
	default:
		panic("Unknown context " + val.GetText())
	}

}

func (v *Visitor) VisitProg(ctx *parser.ProgContext) Value {
	return v.Visit(ctx.Block())
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) Value {
	for i := 0; ctx.Stmt(i) != nil; i++ {
		v.Visit(ctx.Stmt(i))
	}
	return Value{value: true}
}

func (v *Visitor) VisitStmt(ctx *parser.StmtContext) Value {
	if ctx.Assignstmt() != nil {
		return v.Visit(ctx.Assignstmt())
	}
	if ctx.Ifstmt() != nil {
		return v.Visit(ctx.Ifstmt())
	}
	if ctx.Printlnstmt() != nil {
		return v.Visit(ctx.Printlnstmt())
	}
	if ctx.Printstmt() != nil {
		return v.Visit(ctx.Printstmt())
	}
	if ctx.Whilestmt() != nil {
		return v.Visit(ctx.Whilestmt())
	}
	return Value{value: true}
}

func (v *Visitor) VisitAssignstmt(ctx *parser.AssignstmtContext) Value {
	id := ctx.ID().GetText()
	aux := v.Visit(ctx.Expr())
	v.Memory[id] = aux
	return Value{value: true}
}

func (v *Visitor) VisitIfstmt(ctx *parser.IfstmtContext) Value {
	value, ok := v.Visit(ctx.Expr()).value.(bool)
	if ok && value {
		return v.Visit(ctx.Block())
	}
	return Value{value: false}
}

func (v *Visitor) VisitPrintstmt(ctx *parser.PrintstmtContext) Value {
	fmt.Print(v.Visit(ctx.Expr()).value)
	return Value{value: true}
}

func (v *Visitor) VisitPrintlnstmt(ctx *parser.PrintlnstmtContext) Value {
	fmt.Println(v.Visit(ctx.Expr()).value)
	return Value{value: true}
}

func (v *Visitor) VisitWhilestmt(ctx *parser.WhilestmtContext) Value {
	value, ok := v.Visit(ctx.Expr()).value.(bool)
	for ok && value {
		v.Visit(ctx.Block())
		value, ok = v.Visit(ctx.Expr()).value.(bool)
	}
	return Value{value: true}
}

func (v *Visitor) VisitIntExpr(ctx *parser.IntExprContext) Value {
	i, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
	return Value{value: i}
}

func (v *Visitor) VisitStrExpr(ctx *parser.StrExprContext) Value {
	aux := strings.Trim(ctx.GetText(), "\"")
	return Value{value: aux}
}

func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext) Value {
	id := ctx.GetText()
	value, ok := v.Memory[id]
	if ok {
		return value
	} else {
		panic("no such variable: " + id)
	}
	return Value{value: false}
}

func (v *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext) Value {
	value, _ := strconv.ParseBool(ctx.GetText())
	return Value{value: value}
}

func (v *Visitor) VisitOpExpr(ctx *parser.OpExprContext) Value {
	l := v.Visit(ctx.GetLeft()).value.(int64)
	r := v.Visit(ctx.GetRight()).value.(int64)
	op := ctx.GetOp().GetText()
	switch op {
	case "+":
		return Value{value: l + r}
	case "-":
		return Value{value: l - r}
	case "*":
		return Value{value: l * r}
	case "/":
		return Value{value: l / r}
	case "<":
		if l < r {
			return Value{value: true}
		} else {
			return Value{value: false}
		}
	}
	return Value{value: false}
}
