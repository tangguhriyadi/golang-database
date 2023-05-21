package repository

import (
	"context"
	"errors"
	"fmt"
	"golang-database/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnections() {
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
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatal(err)
	}

	// Query SELECT menggunakan ORM dan context
	var users []entity.User
	err = db.WithContext(ctx).Find(&users).Error
	if err != nil {
		log.Fatal(err)
	}

	// Menampilkan hasil query
	for _, user := range users {
		fmt.Printf("ID: %d, Nama: %s, Age: %d, Phone: %s\n", user.ID, user.Nama, user.Age, user.Phone)
	}
}

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	FindById(ctx context.Context, id int32) (*entity.User, error)
	FindAll(ctx context.Context) ([]entity.User, error)
}

//implementasi

type UserRepositoryImp struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImp{
		db: db,
	}
}

func (implement *UserRepositoryImp) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	result := implement.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (implement *UserRepositoryImp) FindById(ctx context.Context, id int32) (*entity.User, error) {
	var user entity.User
	result := implement.db.WithContext(ctx).First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	} else if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (implement *UserRepositoryImp) FindAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	result := implement.db.WithContext(ctx).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
