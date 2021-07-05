package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	//Greet(os.Stdout, "Elodie")
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "worldddd")
}
