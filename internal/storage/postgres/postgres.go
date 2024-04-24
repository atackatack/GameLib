package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type Storage struct {
	DataBase *sqlx.DB
	Config   StorageConfig
}

func NewStorage(cfg StorageConfig) (Storage, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Name, cfg.DBName, cfg.Password, cfg.SSLMode))

	if err != nil {
		log.Fatalln("error connect to DB")
		return Storage{}, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("error ping to DB %s %s:%s \n", err, cfg.Host, cfg.Port)
		return Storage{}, err
	}

	return Storage{
		DataBase: db,
		Config:   cfg,
	}, nil
}
