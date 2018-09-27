package main

import (
	"fmt"

	project "github.com/fenwickelliott/project/project-models"
)

func main() {
	thing := project.Thing{
		Declaration: "I be a API",
		Number:      13,
	}

	fmt.Println(thing.Declaration)
	fmt.Printf("I am %d minutes old\n", thing.Number)
}
