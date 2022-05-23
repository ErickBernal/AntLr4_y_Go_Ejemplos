package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"pruebaGo/Paquetes"
	"pruebaGo/parser"
)

func main() {
	//----------------- prueba de uso de structs en paquetes --------
	fmt.Println("------   paquetes:   ---------")
	paquetes := Paquetes.Clase{}
	paquetes.Hola()
	//-----------------    uso de antlr4  ------------------------------------------------
	fmt.Println("------   Parser:   ---------")
	//setup the input
	input := "3*(32+4)"
	is := antlr.NewInputStream(input)

	//create the Lexer
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	//create the parser(p) with string(stream)
	p := parser.NewCalcParser(stream)
	//finaly parse the expression
	//p.L()   =  primer pruduccion que yo utilice (puede ser cualquiera)
	var listener Paquetes.CalcListener
	fmt.Println("<input: \"", input, "\">")
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.L())
	fmt.Println("<Result: \"", listener.Pop(), "\" >")
	fmt.Println("----------------------fin-----------------")
}
