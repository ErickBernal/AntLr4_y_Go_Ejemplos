package Paquetes

import "fmt"

type Clase struct {
	Dato int
}

func (x *Clase) Hola() {
	x.Dato = 2
	fmt.Println("la clase dice hola", x.Dato)
}
