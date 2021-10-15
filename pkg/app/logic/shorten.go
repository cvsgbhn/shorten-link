package logic

import (
	"shorten-link/pkg/app/models"

	//"encoding/hex"
	"crypto/md5"
	"time"
	"fmt"
)

func base62Convert(initNum []byte) string {
	alphabet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	result := ""
	for _, i := range(initNum) {
		for i > 0 {
			r := i % 62
			i /= 62
			result = string(alphabet[r]) + result
		}
	}
	return result
}

func ShortenLink(longLink string) models.LinkInfo {
	data := []byte(longLink)
	hashedLink := md5.Sum(data)

	var newLink models.LinkInfo
	newLink.OriginalUrl = longLink
	newLink.CreationDate = time.Now()
	newLink.ExpirationDate = (newLink.CreationDate).AddDate(0, 1, 0)
	for i := 0; i <= len(hashedLink[:])- 4; i++ {
		tempHash := base62Convert(hashedLink[:])[i:i+4]
		checkHash := models.GetByShortenedLink(tempHash)
		if len(checkHash) == 0 {
			newLink.Hash = tempHash
			models.AddLink(&newLink)
			return newLink	
		}
	}

	fmt.Println(hashedLink[:])
	fmt.Println(newLink.Hash)

	return newLink
}