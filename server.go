package main

import (
	"fmt"
	"net/http"

	"MyServer/clients"
)

func main() {

	var ali clients.Client = clients.Client{2, "Tanya", 12000}
	var tanya clients.Client = clients.Client{1, "Ali", 12000}
	fmt.Println(tanya.Currency, tanya.Id)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Helo World!")
		fmt.Fprintf(w, ali.Name, ali.Currency)
	})
	http.HandleFunc("/Ali/1200/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Helo World!")
		ali.Currency = 0
		fmt.Fprintf(w, ali.Name, ali.Currency)
	})
	http.ListenAndServe(":80", nil)
}
