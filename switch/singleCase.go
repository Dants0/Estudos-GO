package main

import "fmt"

func main() {

	var x int = 7

	switch x {
	case 1:
		fmt.Println("Segunda")
	case 2:
		fmt.Println("Terça")
	case 3:
		fmt.Println("Quarta")
	case 4:
		fmt.Println("Quinta")
	case 5:
		fmt.Println("Sexta")
	case 6:
		fmt.Println("Sábado")
	case 7:
		fmt.Println("Domingo")

	default:
		fmt.Println("Unknown day")
	}

}
