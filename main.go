package main

import (
	"encoding/json"
	"fmt"
	"crypto/cipher"
	"crypto/aes"
	"encoding/base64"
	
)
//Watermark
type Watermark struct {
	Timestamp int `json:"timestamp"`
	Appid string `json:"appid"`
}

//获取用户信息
type User struct {
	OpenId string `json:"openId"`
	NickName string `json:"nickName"`
	Gender int `json:"gender"`
	Language string `json:"language"`
	City string `json:"city"`
	Province string `json:"province"`
	Country string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
	Watermark Watermark `json:"watermark"`
}


// 解密，采用CBC模式
func AesDecrypt(encrypted, key []byte, iv []byte) ([]byte, error) {
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	decrypted := make([]byte, len(encrypted))
	var block cipher.Block
	block, err = aes.NewCipher([]byte(key))
	if err != nil {
		return []byte{}, err
	}
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(decrypted, encrypted)
	return decrypted, nil
}

func main () {
	var sessionKey = ""
	var encryptedData = ""
	var iv = ""
	decoded, _ := base64.StdEncoding.DecodeString(sessionKey)
	encrypteddata, _ := base64.StdEncoding.DecodeString(encryptedData)
	ivdata, _ := base64.StdEncoding.DecodeString(iv)
	origData, err := AesDecrypt(encrypteddata, decoded, ivdata)
	if err != nil {
		panic(err)
	}
	var jsonUsers = []byte(string(origData))
	users := User{}
	errs := json.Unmarshal(jsonUsers, &users)
	if errs != nil {
		fmt.Println(errs)
	}
	fmt.Println(users)
}


