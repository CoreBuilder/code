package main

import (
	"net/http"

	"github.com/CoreBuilder/go-starcraftservice/controllers"
)

// http://localhost:3000/races/
func main() {

	controllers.RegisterControllers()

	http.ListenAndServe(":3000", nil)

	/*  Model
		r := models.Race{
	 	ID:       1,
	 	Name:     "Protos",
	 	Religion: "xel'naga",
	 	}

		fmt.Println(r)
	*/
}
