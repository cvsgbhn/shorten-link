package logic

import (
	"shorten-link/pkg/app/models"

	"encoding/hex"
	"crypto/md5"
	"fmt"
)

func byteToString62(initialText []byte) string {}

func ShortenLink(longLink string) models.LinkInfo {
	data := []byte(longLink)
	hashedLink := md5.Sum(data)

	var newLink models.LinkInfo
	//newLink.Hash = string(hashedLink[:])[:5]
	newLink.Hash = (hex.EncodeToString(hashedLink[:]))[0:5]

	fmt.Println(hashedLink[:])
	fmt.Println(newLink.Hash)

	return newLink
}