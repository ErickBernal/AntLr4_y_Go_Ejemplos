package paquetes

import "fmt"

type Prueba struct {
}

func (x *Prueba) Saludo() {
	fmt.Println("<Holi>")
}
