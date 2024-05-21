package main

import (
	"app-commnad-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("StartUP")

	app := app.Gerar()
	// erro := app.Run(os.Args) Outra maneira de fazer

	// if erro != nil {
	// 	log.Fatal(erro)
	// }

	if erro := app.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
