package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
	"os"
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
	file, err := os.Create("private.pem")
	if err != nil {
		return nil,err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return nil, err
	}

	pri, err := ioutil.ReadFile("private.pem")
	if err != nil{
		return nil, err
	}
	rsaObj.PrivateKey = base64.StdEncoding.EncodeToString(pri)
	if len(rsaObj.PrivateKey) >0 {
		os.Remove("private.pem")
	}

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return nil,err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return nil, err
	}
	pub, err := ioutil.ReadFile("public.pem")
	if err != nil{
		return nil, err
	}
	rsaObj.PublicKey = base64.StdEncoding.EncodeToString(pub)

	if len(rsaObj.PublicKey) >0 {
		os.Remove("public.pem")
	}

	return rsaObj,nil
}
