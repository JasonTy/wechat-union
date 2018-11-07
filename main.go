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
	var sessionKey = "LXFJEfsRiNLfk2q4z3ajyQ=="
	var encryptedData = "AZMiC/4nuJCP/BAqGZ9WHGCzdg8Nva/3gl54y/242nDCCumn0H/DZPvhNURwG0mFbRVrTDwPS/NbG4a/NUuypuLZs5BnLgSA6MPzAJKJyc4V7a+bvT+b90UWJYvHYBsnxrxDeYzX7CyZ0fQJizQOoB0C8O2mvobAPvzddpju4yOxBiULwDBGV0ldTNPND8+PsXNbLnPTEHh9rdvx/umfp1ngT2RqZhttBmp/0xdhVM2c7BzAU66raezJajB5XEzfAwd5ppy1ZIS/Fg9oymPfMzjLMv72464WthXMszqW68YRvu/bdL6XcnSdRiHIlb2EPN+Hd3CiyM7cFe42lWXkXAxCh/D4Uv+IMcDf2nWftpKIasmKyozskjst7YLu4Vas4Lj+uduIkv472kI8O/HcGsCIJcWQvUgvgmuixRVVxMzKOYVf8PykxnapNw2MQCtx3l8HJ8AIZgo8am4Vi4DcOX/S7RKhUXSdXMUeu/QX1zA="
	var iv = "QFZw1/47zE2aHmU1UNuHPw=="
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


