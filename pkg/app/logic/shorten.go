package logic

import (
	"shorten-link/pkg/app/models"
	"shorten-link/pkg/db"

	"github.com/go-redis/redis"

	//"encoding/hex"
	"crypto/md5"
	"time"
	"fmt"
)

func base63Convert(initNum []byte) string {
	alphabet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

	result := ""
	for _, i := range(initNum) {
		for i > 0 {
			r := i % 63
			i /= 63
			result = string(alphabet[r]) + result
		}
	}
	return result
}

/*
Generates hash for original link and returns its unique part  
*/
func ShortenLink(longLink string, pgDB *db.DB, rdClients []*redis.Client) models.LinkInfo {
	fmt.Println("ShortenLink // shorten.go")
	data := []byte(longLink)
	hashedLink := md5.Sum(data)

	var newLink models.LinkInfo
	newLink.OriginalUrl = longLink
	newLink.CreationDate = time.Now()
	newLink.ExpirationDate = (newLink.CreationDate).AddDate(0, 1, 0)
	for i := 0; i <= len(hashedLink[:])- 10; i++ {
		tempHash := base63Convert(hashedLink[:])[i:i+10]
		if pgDB != nil {
			fmt.Println("ShortenLink // pgDB")
			checkHash := models.GetByShortenedLink(tempHash, pgDB)
			if len(checkHash) == 0 {
				newLink.Hash = tempHash
				models.AddLink(&newLink)
				return newLink	
			}
		} else if rdClients[0] != nil && rdClients[1] != nil {
			fmt.Println("ShortenLink // rdClient not nil")
			checkHash := models.RedisGetByShortenedLink(tempHash, rdClients[1])
			if len(checkHash) == 0 {
				newLink.Hash = tempHash
				models.RedisAddLink(&newLink, rdClients)
				return newLink	
			}
		}
	}

	fmt.Println(hashedLink[:])
	fmt.Println(newLink.Hash)

	return newLink
}