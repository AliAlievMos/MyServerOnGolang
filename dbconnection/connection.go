package dbconnection

import (
	"MyServer/clients"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = 1234
	dbname   = "postgres"
)

func Connect() *sql.DB {
	info := fmt.Sprintf("host=%s port=%d user=%s "+"password=%d dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", info)
	if err != nil {
		fmt.Println("не удалось открыть бд")
		panic(err)
	} else {
		fmt.Println("подключено!")

	}
	return db

}
func Clouse(db *sql.DB) {
	err := db.Close()
	if err != nil {
		fmt.Println("не удалось закрыть бд")
		panic(err)
	} else {
		fmt.Println("отключено!")
	}
}

func Checks(db *sql.DB, sender clients.Client, recipient clients.Client, transaction int) int8 {

	fmt.Printf("%d переводит %d %d \n", sender.Id, recipient.Id, transaction)
	bl, currency := chkId(sender.Id, db)
	if bl == false {
		return 1
	}

	sender.Currency = currency
	if sender.Currency < transaction {
		return 2
	}

	bl, currency1 := chkId(recipient.Id, db)
	if bl == false {
		return 3
	}

	recipient.Currency = currency1
	transactionDo(db, sender, recipient, transaction)

	return 0
}

func chkId(id int, db *sql.DB) (bool, int) {

	query := `SELECT currency FROM clients WHERE id = $1`
	var currency int
	bl := db.QueryRow(query, id).Scan(&currency)
	//row := db.QueryRow(query, id)
	fmt.Println("name")
	fmt.Println(currency)
	if bl == sql.ErrNoRows {
		return false, currency
	} else {
		return true, currency
	}
}

func transactionDo(db *sql.DB, sender clients.Client, recipient clients.Client, transaction int) {
	sender.Currency -= transaction
	recipient.Currency += transaction
	fmt.Println("transaction")
	query := `UPDATE clients SET currency = $1 WHERE id = $2;`
	_, err := db.Query(query, sender.Currency, sender.Id)
	if err != nil {
		fmt.Printf("%s ошибка! \n", err)
	}
	_, err1 := db.Query(query, recipient.Currency, recipient.Id)
	if err1 != nil {
		fmt.Printf("%s ошибка! \n", err1)
	}
}
