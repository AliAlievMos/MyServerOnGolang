package dbconnection

import (
	"database/sql"
	"fmt"
	"log"

	"MyServer/clients"

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

	rows := chkid(sender.Id, db)
	if rows == nil {
		return 1
	}

	for rows.Next() {
		var currency int
		if err := rows.Scan(&currency); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s на счету\n", string(rune(currency)))
		if transaction > currency {
			return 2
		}

	}

	rows1 := chkid(recipient.Id, db)
	if rows1 == nil {
		return 3
	}

	return 0
}

func chkid(id int, db *sql.DB) *sql.Rows {
	query := `SELECT currency FROM clients WHERE id = $1`

	rows, err := db.Query(query, id)
	if err != nil {
		fmt.Printf("%s ошибка! \n", err)
		return nil
	}
	return rows
}

func transaction(db *sql.DB, sender clients.Client, recipient clients.Client, transaction int) {
	sender.Currency -= transaction
	recipient.Currency += transaction
}
