package connection

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"sync"
)

var (
	db   *gorm.DB // un puntero a gorm.DB
	once sync.Once
)

type Driver string

// Drivers
const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

func New(d Driver) {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
}

func newPostgresDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open("postgres", "postgres://admin:admin@localhost:5432/crudgo?sslmode=disable")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}
		fmt.Println("conectado a postgres")
	})
}

func newMySQLDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open("mysql", "admin:admin@tcp(localhost:8080)/godb?parseTime=true")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}
		fmt.Println("conectado a mySQL")
	})
}

// DB retorna una unica instance de bd
func DB() *gorm.DB { // retorna un puntero a gorm.DB
	return db
}
