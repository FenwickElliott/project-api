package main

import (
	"fmt"

	project "github.com/fenwickelliott/project-models"
)

func main() {
	thing := project.Thing{
		Declaration: "I be a API",
	}

	fmt.Println(thing.Declaration)
}
