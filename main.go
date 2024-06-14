package main

import (
	"fmt"
	"golang-dev-logic-challenge-dvanauken/routes"
)

//func main() {
//	http.HandleFunc("/analyze", controllers.AnalysisHandler) // Use AnalysisHandler from the controllers package
//
//	fmt.Println("Starting server on port 8080")
//	http.ListenAndServe(":8080", nil) // Start the server on port 8080
//}

func main() {
	router := routes.SetupRouter()
	fmt.Println("Starting server on port 8080")
	router.Run(":8080")
}
