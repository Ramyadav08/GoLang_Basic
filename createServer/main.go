package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w,"Welcome to my server......❤️❤️❤️❤️❤️❤️❤️❤️")
	})
	http.ListenAndServe(":5000", nil)
}
