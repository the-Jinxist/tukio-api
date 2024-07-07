package config

import (
	"log"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var db *sqlx.DB

func GetDB() *sqlx.DB {
	return db
}

func InitDB() {
	conn := connectToDb()
	if conn == nil {
		log.Panic("can't connect to db")
	}

	db = conn
}

func connectToDb() *sqlx.DB {
	counts := 0

	dsn := viper.GetString("DSN")
	log.Println("dsn ", dsn)
	for {
		connextion, err := openDB(dsn)
		if err != nil {
			log.Printf("postgress not ready")

		} else {
			log.Println("postgress is ready")
			return connextion

		}

		if counts > 10 {
			return nil
		}

		log.Println("Backing off, waiting for db to be ready")
		time.Sleep(time.Second * time.Duration(counts))
		counts++
		continue
	}
}

func openDB(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
