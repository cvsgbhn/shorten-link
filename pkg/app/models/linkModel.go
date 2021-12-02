package models

import (
	"shorten-link/pkg/db"

	"log"
	"time"
	"context"
	_ "github.com/lib/pq"
	"github.com/Masterminds/squirrel"
	"github.com/go-redis/redis"
)

var ctx = context.Background()

type LinkInfo struct {
	Id int
	Hash string
	OriginalUrl string
	CreationDate time.Time
	ExpirationDate time.Time
}

func RedisGetByFullLink(fullLink string, rdClient *redis.Client) string {
	log.Println("RedisGetByFullLink // linkModel")
	result := ""
	result, err := rdClient.Get(fullLink).Result()
	if err != nil {
		log.Println(err)
	}

	log.Println("redis res-t: ", result)
	return result
}

func RedisGetByShortenedLink(shortLink string, rdClient *redis.Client) string {
	var result string
	result, err := rdClient.Get(shortLink).Result()
	if err != nil {
		log.Println(err)
	}

	return result
}

func RedisAddLink(newLink *LinkInfo, rdClients []*redis.Client) {
	err := rdClients[0].Set(newLink.OriginalUrl, newLink.Hash, 0).Err()
    if err != nil {
        panic(err)
    }
	err = rdClients[1].Set(newLink.Hash, newLink.OriginalUrl, 0).Err()
    if err != nil {
        panic(err)
    }
}

/*
Selects all not expired shortened links for a given full one
*/
func GetByFullLink(fullLink string, postgresDB *db.DB) ([]*LinkInfo) {

	links := make([]*LinkInfo, 0)
	currentDate := time.Now()

	//postgresDB := db.NewDB(db.BuildConfig())

	sql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(
			"id",
			"hash",
			"original_url",
			"creation_date",
			"expiration_date",
		).
		From("links").
		Where(squirrel.Eq{"original_url": fullLink}).
		Where("expiration_date>=?", currentDate)

		rows, err := sql.RunWith(postgresDB).Query()
		if err != nil {
			log.Println(err)
			return nil
		}
		defer rows.Close()
	
		for rows.Next() {
			var row LinkInfo
	
			err = rows.Scan(
				&row.Id,
				&row.Hash,
				&row.OriginalUrl,
				&row.CreationDate,
				&row.ExpirationDate,
			)
			links = append(links, &row)
		}
	
	//defer postgresDB.Close()

	return links
}

/*
Selects all not expired shortened links matching a given one
*/
func GetByShortenedLink(shortLink string, postgresDB *db.DB) ([]*LinkInfo) {

	links := make([]*LinkInfo, 0)
	currentDate := time.Now()

	//postgresDB := db.NewDB(db.BuildConfig())

	sql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(
			"id",
			"hash",
			"original_url",
			"creation_date",
			"expiration_date",
		).
		From("links").
		Where(squirrel.Eq{"hash": shortLink}).
		Where("expiration_date>=?", currentDate)

		rows, err := sql.RunWith(postgresDB).Query()
		if err != nil {
			log.Println(err)
			return nil
		}
		defer rows.Close()
	
		for rows.Next() {
			var row LinkInfo
	
			err = rows.Scan(
				&row.Id,
				&row.Hash,
				&row.OriginalUrl,
				&row.CreationDate,
				&row.ExpirationDate,
			)
			links = append(links, &row)
		}
	
	//defer postgresDB.Close()

	return links
}

/*
Inserts a new link
*/
func AddLink(newLink *LinkInfo) (id int, err error) {

	postgresDB := db.NewDB(db.BuildConfig())
	if err != nil {
		log.Panic(err)
	}

	sql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert("links").
		Columns(
			"hash",
			"original_url",
			"creation_date",
			"expiration_date",
		).
		Values(
			newLink.Hash,
			newLink.OriginalUrl,
			newLink.CreationDate,
			newLink.ExpirationDate,
		).
		Suffix("RETURNING \"id\"")

	err = sql.RunWith(postgresDB).QueryRow().Scan(&id)

	//defer postgresDB.Close()

	return id, err
}
