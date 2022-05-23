package main

import (
	"main/paquetes"
	"main/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	prog := `
	a = true
	b = 0
	if (a) {
		println("a is true!")
	}
	while (b < 5) {
		println(b)
		b = b + 2
	}`
	input := antlr.NewInputStream(prog)
	lexer := parser.NewControlLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewControlParser(tokens)
	p.BuildParseTrees = true
	tree := p.Prog()                                                  //se asigna la produccion inicial
	eval := paquetes.Visitor{Memory: make(map[string]paquetes.Value)} //instancia del mapa con make
	eval.Visit(tree)                                                  //.visit(), es el que nosotros construimos

	prueba := paquetes.Prueba{}
	prueba.Saludo()

}
