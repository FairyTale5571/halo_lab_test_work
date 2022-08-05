package database

import (
	"database/sql"

	"github.com/fairytale5571/privat_test/pkg/logger"
	_ "github.com/lib/pq"
)

type Database interface {
	Close() error
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	Prepare(query string) (*sql.Stmt, error)
}

type DB struct {
	db     *sql.DB
	logger *logger.Wrapper
}

func New(uri string) (*DB, error) {
	var err error
	dbConnection := DB{
		logger: logger.New("database"),
	}
	db, err := sql.Open("postgres", uri)
	if err != nil {
		dbConnection.logger.Fatalf("error open database: %v", err)
		return nil, err
	}
	db.SetMaxOpenConns(10)
	dbConnection.db = db

	v, err := dbConnection.Version()
	if err != nil {
		dbConnection.logger.Fatalf("error get version: %v", err)
		return nil, err
	}
	dbConnection.logger.Infof("database version: %v", v)

	return &dbConnection, nil
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) Exec(query string, args ...any) (sql.Result, error) {
	return db.db.Exec(query, args...)
}

func (db *DB) Query(query string, args ...any) (*sql.Rows, error) {
	return db.db.Query(query, args...)
}

func (db *DB) QueryRow(query string, args ...any) *sql.Row {
	return db.db.QueryRow(query, args...)
}

func (db *DB) Prepare(query string) (*sql.Stmt, error) {
	return db.db.Prepare(query)
}

func (db *DB) Version() (string, error) {
	var version string
	err := db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		return "", err
	}
	return version, nil
}
