package rsa

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"strings"
)

type RSA struct {
	PublicKey 	string
	PrivateKey 	string
}
func GenRsaKey(bits int) (*RSA, error){

	rsaObj := new(RSA)
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	priBuf := new(bytes.Buffer)
	err = pem.Encode(priBuf, block)
	if err != nil {
		return nil, err
	}

	//rsaObj.PrivateKey = base64.StdEncoding.EncodeToString(priBuf.String())
	rsaObj.PrivateKey = removeHeadAndEnd(priBuf.String())

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	block = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: derPkix,
	}

	pubBuf := new(bytes.Buffer)
	err = pem.Encode(pubBuf, block)
	if err != nil {
		return nil, err
	}

	//rsaObj.PublicKey = base64.StdEncoding.EncodeToString(pubBuf.String())
	rsaObj.PublicKey = removeHeadAndEnd(pubBuf.String())

	return rsaObj,nil
}

func removeHeadAndEnd(secretKey string)string{

	FirstWrap := strings.Index(secretKey,"\n")
	LastWrap := strings.LastIndex(secretKey,"\n")
	tmp1 := secretKey[FirstWrap:LastWrap]
	LastWrap2 := strings.LastIndex(tmp1,"\n")
	return secretKey[FirstWrap+1:LastWrap2]
}
