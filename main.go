package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("booted and tooted")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hiit")
		fmt.Fprintf(w, "Jimmys Response")
	})

	http.ListenAndServe(":8081", nil)
}
