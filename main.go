package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"strings"
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
	UnionId string `json:"unionId"`
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
	var sessionKey = "tiihtNczf5v6AKRyjwEUhQ=="
	var encryptedData = "CiyLU1Aw2KjvrjMdj8YKliAjtP4gsMZMQmRzooG2xrDcvSnxIMXFufNstNGTyaGS9uT5geRa0W4oTOb1WT7fJlAC+oNPdbB+3hVbJSRgv+4lGOETKUQz6OYStslQ142dNCuabNPGBzlooOmB231qMM85d2/fV6ChevvXvQP8Hkue1poOFtnEtpyxVLW1zAo6/1Xx1COxFvrc2d7UL/lmHInNlxuacJXwu0fjpXfz/YqYzBIBzD6WUfTIF9GRHpOn/Hz7saL8xz+W//FRAUid1OksQaQx4CMs8LOddcQhULW4ucetDf96JcR3g0gfRK4PC7E/r7Z6xNrXd2UIeorGj5Ef7b1pJAYB6Y5anaHqZ9J6nKEBvB4DnNLIVWSgARns/8wR2SiRS7MNACwTyrGvt9ts8p12PKFdlqYTopNHR1Vf7XjfhQlVsAJdNiKdYmYVoKlaRv85IfVunYzO0IKXsyl7JCUjCpoG20f0a04COwfneQAGGwd5oa+T8yO5hzuyDb/XcxxmK01EpqOyuxINew=="
	var iv = "r7BXXKkLb8qrSNn05n0qiA=="
	defer func () {
		if e := recover(); e != nil {
			spew.Dump(e)
		}
	}()
	decoded, _ := base64.StdEncoding.DecodeString(sessionKey)
	encrypteddata, _ := base64.StdEncoding.DecodeString(encryptedData)
	ivdata, _ := base64.StdEncoding.DecodeString(iv)
	origData, err := AesDecrypt(encrypteddata, decoded, ivdata)
	if err != nil {
		panic(err)
	}

	newStr := strings.Replace(string(origData), "\a", "", -1)
	users := User{}
	errs := json.Unmarshal([]byte(newStr), &users)
	if errs != nil {
		panic(errs)
	}
	spew.Println(users)
}


