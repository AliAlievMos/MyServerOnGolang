package main

import (
	"fmt"
	"net/http"

	"MyServer/clients"
	"MyServer/dbconnection"
	//_ "github.com/lib/pq"
)

func main() {

	//var sender = clients.Client{2, "Tanya", 12000}
	//var recipient = clients.Client{1, "Ali", 12000}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Helo World!")
	})
	http.HandleFunc("/transaction/", func(w http.ResponseWriter, r *http.Request) {
		var sender = clients.Client{2, "Tanya", 0}
		var recipient = clients.Client{1, "Ali", 0}

		db := dbconnection.Connect()
		result := dbconnection.Checks(db, sender, recipient, 1)
		if result == 1 {
			fmt.Fprintf(w, "такого клиента нет")
		} else if result == 2 {
			fmt.Fprintf(w, "недостаточно средств")
		} else if result == 3 {
			fmt.Fprintf(w, "такого получателя нет")
		} else {

		}
		dbconnection.Clouse(db)
	})
	http.ListenAndServe(":80", nil)
}
