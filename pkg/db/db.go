package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	MaxOpenConns    = 10
	MaxIdleConns    = 2
	ConnMaxLifetime = 10 * time.Minute
)

type DB struct {
	*sql.DB
	config *Config
}

func NewDB(config *Config) (*DB, error) {
	db, err := sql.Open("postgres", config.BuildDsn())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{
		DB:     db,
		config: config,
	}, nil
}

type Config struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
	SSLMode  string
}

func (dc *Config) BuildDsn() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dc.Host,
		dc.Port,
		dc.User,
		dc.Pass,
		dc.Database,
		dc.SSLMode,
	)
}

func (db *DB) Close() error {
	return db.DB.Close()
}

func BuildConfig() *Config {
	var config *Config
	file, err := os.Open("config.development.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return config
}
