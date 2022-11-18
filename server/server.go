package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"MyServer/clients"
	"MyServer/dbconnection"
)

func Server() {
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
		var sender = clients.Client{Id: idSender, Currency: trans}
		var recipient = clients.Client{Id: idRec, Currency: trans}

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
