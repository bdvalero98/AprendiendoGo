package main

import "AprendiendoGo/teclado"

func main() {
	/*estado, texto := variables.ConviertoaTexto(3529)
	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "linux." || os == "OS X." {
		fmt.Println("Esto no es Windows es: ", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es darwin")
	default:
		fmt.Printf("%s \n", os)
	}

	number, texto := ejercicios.ConvertiraEntero("349082")
	fmt.Println(number)
	fmt.Println(texto)*/

	teclado.IngresoNumeros()
}
