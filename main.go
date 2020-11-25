package main

import (
	"fmt"

	"github.com/CoreBuilder/go-starcraftservice/models"
)

func main() {

	r := models.Race{
		ID:       1,
		Name:     "Protos",
		Religion: "xel'naga",
	}

	fmt.Println(r)

}
