package main

import (
	"fmt"
	"log"
	"net/http"

	githubhook "gopkg.in/rjz/githubhook.v0"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	secret := []byte("don't tell!")
	hook, err := githubhook.Parse(secret, r)
	if err != nil {
		return
	}

	fmt.Println(hook.Event)
	fmt.Println(hook.Id)

	fmt.Fprintf(w, "Hi there")
}

func hello() {
	log.Println("hello")
}
