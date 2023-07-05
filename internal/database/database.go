package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type IDatabase interface {
}

type database struct {
	db *sqlx.DB
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

func New(db *sqlx.DB) *database {
	return &database{db: db}
}

func Connect(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DbName, cfg.Password))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to DB :)")
	return db, nil
}

/*
 * TODO:
 * init database struct
 * connect to the postgres
 * use GORM
 */
