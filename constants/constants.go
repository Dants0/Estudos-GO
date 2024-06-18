package main

import "fmt"

// valores constantes são imutáveis e apenas para leitura, impossível de alterar.
const MYNAME = "Guilherme"

var name = "dantas"

func main() {

	// MYNAME = "Gui" não funciona

	name = "teste" // por não ser constante a variavel se alterou quando lhe foi atribuida um novo valor

	fmt.Println(MYNAME, name)
}