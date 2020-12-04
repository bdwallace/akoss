package tencent

import (
	"io/ioutil"
	"library/common"
	"net/http"

	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	// "encoding/json"
	"fmt"
	"time"
	// "github.com/astaxie/beego/httplib"
)

type Response struct {
	Response Data `json: "Response"`
}

type Data struct {
	Balance int64 `json: "Balance"`
}

func sha256hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

func hmacsha256(s, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return string(hashed.Sum(nil))
}

func Balance(access_key_id, access_secret string) (balance float64, err error) {
	secretId := access_key_id
	secretKey := access_secret
	host := "billing.tencentcloudapi.com"
	algorithm := "TC3-HMAC-SHA256"
	service := "billing"
	version := "2018-07-09"
	action := "DescribeAccountBalance"
	// region := "ap-guangzhou"
	var timestamp int64 = time.Now().Unix()
	// var timestamp int64 = 1551113065

	// step 1: build canonical request string
	httpRequestMethod := "POST"
	canonicalURI := "/"
	canonicalQueryString := ""
	canonicalHeaders := "content-type:application/json; charset=utf-8\n" + "host:" + host + "\n"
	signedHeaders := "content-type;host"
	payload := "{}"
	hashedRequestPayload := sha256hex(payload)
	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		httpRequestMethod,
		canonicalURI,
		canonicalQueryString,
		canonicalHeaders,
		signedHeaders,
		hashedRequestPayload)
	// fmt.Println(canonicalRequest)

	// step 2: build string to sign
	date := time.Unix(timestamp, 0).UTC().Format("2006-01-02")
	credentialScope := fmt.Sprintf("%s/%s/tc3_request", date, service)
	hashedCanonicalRequest := sha256hex(canonicalRequest)
	string2sign := fmt.Sprintf("%s\n%d\n%s\n%s",
		algorithm,
		timestamp,
		credentialScope,
		hashedCanonicalRequest)
	// fmt.Println(string2sign)

	// step 3: sign string
	secretDate := hmacsha256(date, "TC3"+secretKey)
	secretService := hmacsha256(service, secretDate)
	secretSigning := hmacsha256("tc3_request", secretService)
	signature := hex.EncodeToString([]byte(hmacsha256(string2sign, secretSigning)))
	// fmt.Println(signature)

	// step 4: build authorization
	authorization := fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		algorithm,
		secretId,
		credentialScope,
		signedHeaders,
		signature)
	// fmt.Println(authorization)

	// curl := fmt.Sprintf(`curl -X POST https://%s\
	//  -H "Authorization: %s"\
	//  -H "Content-Type: application/json; charset=utf-8"\
	//  -H "Host: %s" -H "X-TC-Action: %s"\
	//  -H "X-TC-Timestamp: %d"\
	//  -H "X-TC-Version: %s"\
	//  -d '%s'`, host, authorization, host, action, timestamp, version, payload)
	// fmt.Println(curl)

	var head map[string]string
	head = make(map[string]string)
	head["Authorization"] = authorization
	head["Content-Type"] = "application/json; charset=utf-8"
	head["Host"] = host
	head["X-TC-Action"] = action
	head["X-TC-Timestamp"] = fmt.Sprintf("%d", timestamp)
	head["X-TC-Version"] = version

	body := []byte(payload)

	var resp *http.Response
	resp, err = common.HttpPost(fmt.Sprintf("https://%s/", host), head, body)
	if err != nil {
		return
	}

	var response Response
	resp_body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(resp_body, &response)
	if err != nil {
		return
	}

	balance = float64(response.Response.Balance) / 100.00

	return
}
