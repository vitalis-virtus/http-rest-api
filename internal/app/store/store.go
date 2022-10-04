package store

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	config *Config
	DB     *sql.DB
}

func New(config *Config) *Store {

	return &Store{
		config: config,
	}
}

// Open func for store initalizing, connecting to BD
func (s *Store) Open() error {
	db, err := sql.Open("mysql", s.config.DatabaseURL)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.DB = db

	fmt.Println("Connected to DB")

	return nil
}

// Close disconnecting for BD
func (s *Store) Close() {
	s.DB.Close()
}
