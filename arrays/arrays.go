package main

import "fmt"

func main(){
	var array = [...]string{"G1", "G2", "G3", "G4"}

	using_inferred := [...]string{"test1", "test2", "test3", "test4"}
	
	array[1]="G6"
	using_inferred[1]="teste99"

	//inicializando o array com valores pre definidos em posições pre definidas do array
	array22 := [...]int{3:20, 4:2}


	//usa-se a funçõa len para informar qual o tamanho do array
	tamanho_do_array := [...]string{"BMW", "FERRARI", "PORSCHE"}

	fmt.Println(len(tamanho_do_array))


	fmt.Println(array)
	fmt.Println(using_inferred)
	fmt.Println(array22)
}