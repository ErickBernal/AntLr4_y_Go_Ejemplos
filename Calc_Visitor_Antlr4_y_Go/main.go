package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"main/Paquetes"
	"main/parser"
)

//---------------------------------------
func main() {
	fmt.Println("------   Prueba de paquetes:   ---------")
	s := Paquetes.Prueba{}
	s.Saludo()
	//-----------------    uso de antlr4  ------------------------------------------------
	fmt.Println("------   Parser:   ---------")
	exp := "3*(32+4)"
	input := antlr.NewInputStream(exp)
	lexer := parser.NewCalcLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCalcParser(stream)
	p.BuildParseTrees = true
	tree := p.Expr()
	var visitor = Paquetes.Visitor{}
	var result = visitor.Visitor(tree)
	fmt.Println(exp, " = ", result)
}
