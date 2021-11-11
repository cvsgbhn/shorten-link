package models

import (
	"shorten-link/pkg/db"

	"log"
	"time"
	_ "github.com/lib/pq"
	"github.com/Masterminds/squirrel"
)

type LinkInfo struct {
	Id int
	Hash string
	OriginalUrl string
	CreationDate time.Time
	ExpirationDate time.Time
}

/*
Selects all not expired shortened links for a given full one
*/
func GetByFullLink(fullLink string) ([]*LinkInfo) {

	links := make([]*LinkInfo, 0)
	currentDate := time.Now()

	postgresDB, err := db.NewDB(db.BuildConfig())
	if err != nil {
		log.Panic(err)
	}

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
	
	defer postgresDB.Close()

	return links
}

/*
Selects all not expired shortened links matching a given one
*/
func GetByShortenedLink(shortLink string) ([]*LinkInfo) {

	links := make([]*LinkInfo, 0)
	currentDate := time.Now()

	postgresDB, err := db.NewDB(db.BuildConfig())
	if err != nil {
		log.Panic(err)
	}

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
	
	defer postgresDB.Close()

	return links
}

/*
Inserts a new link
*/
func AddLink(newLink *LinkInfo) (id int, err error) {

	postgresDB, err := db.NewDB(db.BuildConfig())
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

	defer postgresDB.Close()

	return id, err
}
