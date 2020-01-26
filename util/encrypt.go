package util

import (
	"crypto"
	"crypto/aes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

// sha256 签名
func Sha256Encode(string string) string {
	sha := sha256.New()
	sha.Write([]byte(string))
	bs := sha.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

// ras sha256 私钥加密
func RsaPrivateSha256Encode(data string, PrivateKey *rsa.PrivateKey) (signer []byte, err error) {
	hash := crypto.Hash.New(crypto.SHA256)
	hash.Write([]byte(data))
	hashed := hash.Sum(nil)
	signer, err = rsa.SignPKCS1v15(rand.Reader, PrivateKey, crypto.SHA256, hashed)
	return
}

// 根据文件名解析出私钥 ,文件必须是 rsa 私钥格式
func ParsePrivateFromFile(file string) (private *rsa.PrivateKey, err error) {
	pemData, err := ioutil.ReadFile(file)
	if err != nil {
		err = fmt.Errorf("ParsePrivateFromFile,ioutil.ReadFile: %s", err)
		return nil, err
	}
	block, _ := pem.Decode(pemData)
	if block == nil {
		err = fmt.Errorf("ParsePrivateFromFile,pem.Decode fail")
		return nil, err
	}
	if got, want := block.Type, "RSA PRIVATE KEY"; got != want {
		err = fmt.Errorf("ParsePrivateFromFile,got: %q, want: %q", got, want)
		return nil, err
	}
	private, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		err = fmt.Errorf("ParsePrivateFromFile,x509.ParsePKCS1PrivateKey: %s", err)
		return nil, err
	}
	return private, nil
}

// RSA 解密
func RsaPrivateDecode(cipherText []byte, file string) ([]byte, error) {
	private, err := ParsePrivateFromFile(file)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, private, cipherText)
}

//AES 128 ECB decode
func Aes128ECBDecode(data, key []byte) []byte {
	cipher, _ := aes.NewCipher([]byte(key))
	decrypted := make([]byte, len(data))
	size := 16

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cipher.Decrypt(decrypted[bs:be], data[bs:be])
	}
	return decrypted
}
