package main

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Struktur model untuk tabel "users"
type User struct {
	ID    uint   `gorm:"primaryKey"`
	Nama  string `gorm:"not null"`
	Age   int    `gorm:"not null"`
	Phone string
}

func main() {
	// Membuat context dengan timeout
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Konfigurasi koneksi PostgreSQL
	dsn := "host=localhost user=postgres password=edan dbname=edan port=5555 sslmode=disable"

	// Membuka koneksi ke PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal(err)
		}
		sqlDB.Close()
	}()

	// Membuat tabel "users" jika belum ada
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}

	// Query SELECT menggunakan ORM dan context
	var users []User
	err = db.WithContext(ctx).Find(&users).Error
	if err != nil {
		log.Fatal(err)
	}

	// Menampilkan hasil query
	for _, user := range users {
		fmt.Printf("ID: %d, Nama: %s, Age: %d, Phone: %s\n", user.ID, user.Nama, user.Age, user.Phone)
	}
}
