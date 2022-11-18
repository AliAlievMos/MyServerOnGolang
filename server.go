package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	//"net/url"

	"MyServer/clients"
	"MyServer/dbconnection"
	//_ "github.com/lib/pq"
)

func main() {

	http.HandleFunc("/transaction/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.URL.Path)
		str := strings.Split(r.URL.Path[13:], "/")
		idSender, err := strconv.Atoi(str[0])
		if err != nil {
			return
		}
		idRec, err := strconv.Atoi(str[2])
		if err != nil {
			return
		}
		trans, err := strconv.Atoi(str[1])
		if err != nil {
			return
		}

		var sender = clients.Client{idSender, "", 0}
		var recipient = clients.Client{idRec, "", 0}

		db := dbconnection.Connect()
		result := dbconnection.Checks(db, sender, recipient, trans)
		if result == 1 {
			fmt.Fprintf(w, "такого клиента нет")
		} else if result == 2 {
			fmt.Fprintf(w, "недостаточно средств")
		} else if result == 3 {
			fmt.Fprintf(w, "такого получателя нет")
		} else {
			fmt.Fprintf(w, "успешно")
		}
		defer dbconnection.Clouse(db)
	})
	http.ListenAndServe(":80", nil)
}
