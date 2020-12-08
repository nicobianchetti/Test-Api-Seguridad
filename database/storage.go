package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db     *gorm.DB
	once   sync.Once
	hostPg = os.Getenv("DB_HOST_PG")
	// portPg, _ = strconv.Atoi(os.Getenv("DB_PORT_PG"))
	portPg   = os.Getenv("DB_PORT_PG")
	userPg   = os.Getenv("DB_USER_PG")
	passPg   = os.Getenv("DB_PASS_PG")
	dbNamePg = os.Getenv("DB_NAME_PG")
)

// Driver of storage
type Driver string

// Drivers
const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// New create the connection with db
func New(d Driver) {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		NewPostgresDB()
	}
}

//NewPostgresDB created new connection in postgres
func NewPostgresDB() {
	//Singleton
	once.Do(func() {
		var err error

		host := "localhost"
		port := "5432"
		user := "postgres"
		pass := "postgres"
		dbName := "api_seguridad"

		// fmt.Printf("PArametros : host %s port %s user %s pass %s dbName %s \n", hostPg, portPg, userPg, passPg, dbNamePg)

		// host := hostPg
		// port := portPg
		// user := userPg
		// pass := passPg
		// dbName := dbNamePg

		db, err = gorm.Open("postgres", "postgres://"+user+":"+pass+"@"+host+":"+port+"/"+dbName+"?sslmode=disable")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("Conectado a postgres")
	})
}

func newMySQLDB() {
	// once.Do(func() {
	// 	var err error
	// 	db, err = gorm.Open("mysql", "edteam:edteam@tcp(localhost:7531)/godb?parseTime=true")
	// 	if err != nil {
	// 		log.Fatalf("can't open db: %v", err)
	// 	}

	// 	fmt.Println("conectado a mySQL")
	// })
}

// DB return a unique instance of db
func DB() *gorm.DB {
	return db
}
