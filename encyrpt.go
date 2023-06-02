package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
)

var (
	libKey = "xidfkx902dfaxd34" // 客户端约定密钥"
)

func test() {

	encrypted := `yLJX1S1dM0jHVaQIjciWuL//iMph2eSIFk35ZJ1RCkuglFvOwUje8+rtbpXE6pgxuLvaMCPPW1E2vVQyIDhALvA+RwHCcuUdRdLUv9ZZC/cMWLgPnh6PVH5voY0GgCVJe9cArCN+E6g2YKnzVS3OChWW+oesM/vOjPva5LZ3XuLr8W/04ZlTmyVNqP9P3c/96Xc05Ud+msEQ9/2L54QPbqbimT7UR+eFSZNvcagvJj2Y2m6TZtMzbE5HMfkruvRv+PJKzGKRPDRXiDkpqTm4U+B8o0Kn3LbDTAKVnvy1W+Q9d2hrvawKu+HzLv4n+qA2nkb4pasUzRu3ykvH32AxYw==`
	//data := "333"
	//e := AesEncryptECB([]byte(data), []byte(libKey))
	//encrypted := base64.StdEncoding.EncodeToString(e)
	//fmt.Println(encrypted)

	d, err := base64.StdEncoding.DecodeString(encrypted)
	fmt.Println(err)
	decrypted := AesDecryptECB(d, []byte(libKey))
	fmt.Println(string(decrypted))
}

func generateToken() {

}

// 加密
func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

// 解密
func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))

	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
