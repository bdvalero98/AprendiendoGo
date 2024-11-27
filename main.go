package main

import (
	"AprendiendoGo/variables"
	"fmt"
)

func main() {
	estado, texto := variables.ConviertoaTexto(3529)
	fmt.Println(estado)
	fmt.Println(texto)
}
