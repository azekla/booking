package main

import (
	"fmt"
	"github.com/azekla/booking/pkg/handlers"
	"net/http"
)

const portNumber = ":7182"

// the main function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println(err)
	}

}
