package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website! on "+hostname+" with tag "+strings.ToUpper(getTag())+"\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func getTag() string {
	r := os.Getenv("FOO")
	if r == "" {
		return "default"
	}
	return r
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
