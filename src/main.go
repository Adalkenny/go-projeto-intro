package main

import (
	"fmt"
	"log"

	"github.com/Adalkenny/go-projeto-intro/src/modulo"
)

func main() {

	log.SetPrefix("Mantenha: ")
	log.SetFlags(0)
	message, erro := modulo.Comprimentar("Adalberto S. M. Axelrod")
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(message)
}
