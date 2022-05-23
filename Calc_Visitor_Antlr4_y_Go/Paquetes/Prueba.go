package Paquetes

import "fmt"

type Prueba struct {
}

func (x *Prueba) Saludo() {
	fmt.Println("<saludo>")
}