package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, this is my go app deployed with Jenkins!")
	})

	fmt.Println("Server running on port 8088...")
	http.ListenAndServe(":8088", nil)
}
