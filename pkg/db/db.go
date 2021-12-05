package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
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

type DBClient interface{
	GetByFullLink()
	GetByShortenedLink()
}

func RedisClient(db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "db-redis:6379",
		Password: "",
		DB:       db,
	})

	fmt.Println("redis client: ", client)
	return client
}

func RedisTwoTables() []*redis.Client {
	return []*redis.Client{RedisClient(0), RedisClient(1)}
}

func NewDB(config *Config) (*DB) {
	db, err := sql.Open("postgres", config.BuildDsn())
	if err != nil {
		return nil
	}

	if err = db.Ping(); err != nil {
		return nil
	}

	return &DB{
		DB:     db,
		config: config,
	}
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
		"host=%s port=%s user=%s sslmode=%s",
		dc.Host,
		dc.Port,
		dc.User,
		dc.SSLMode,
	)
}

func (db *DB) Close() error {
	return db.DB.Close()
}

func BuildConfig() *Config {
	var config *Config
	file, err := os.Open("config.postgres.json")
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
