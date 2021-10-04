package logic

import (
	"shorten-link/pkg/app/models"

	"encoding/hex"
	"crypto/md5"
	"fmt"
)

func base62Convert(initNum []byte) string {
	alphabet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
}

func ShortenLink(longLink string) models.LinkInfo {
	data := []byte(longLink)
	hashedLink := md5.Sum(data)

	var newLink models.LinkInfo
	newLink.OriginalUrl = longLink
	for i := 0; i <= len(hashedLink[:])- 4; i++ {
		fmt.Println(hex.EncodeToString(hashedLink[:]))
		fmt.Println(hashedLink[:])
		fmt.Println(string(hashedLink[:]))
		tempHash := (hex.EncodeToString(hashedLink[:]))[i:i+4]
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