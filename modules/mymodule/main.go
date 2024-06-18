package main

import (
	"fmt"
	"mymodule/mypackage"
	"mymodule/mymathpackage"
	"github.com/spf13/cobra"
)

func main() {

	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			mymathpackage.RandomNumber()
			mypackage.CreatePerson("Guilherme", 24)
			mypackage.CreatePerson("Fulano", 24)
			mypackage.CreatePerson("Beltrano", 24)
		},
	}
	fmt.Println("Chamando cmd.execute()")
	cmd.Execute()
}
