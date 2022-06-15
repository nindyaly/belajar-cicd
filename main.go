package main

import (
	"log"
	"net/http"
)

import (
	"fmt"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("listening on port 3000...")
	http.ListenAndServe(":3000", nil)
	fmt.Println("Hello World!")
}
