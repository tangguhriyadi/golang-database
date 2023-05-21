package db

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func TestEmpty(t *testing.T) {
	connStr := "user=postgres dbname=edan password=edan port=5555 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	// rows, err := db.Query("SELECT * from users")
	// if err != nil {
	// 	panic(err)
	// }

	// defer rows.Close()

	// for rows.Next() {
	// 	var id int
	// 	var nama string
	// 	var age int
	// 	var phone string
	// 	err = rows.Scan(&id, &nama, &age, &phone)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Printf("ID: %d, Nama: %s, Age: %d, Phone: %s\n", id, nama, age, phone)
	// }

	// if err = rows.Err(); err != nil {
	// 	panic(err)
	// }

}
