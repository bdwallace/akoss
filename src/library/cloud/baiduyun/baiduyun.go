package baiduyun

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"library/common"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	// "fmt"
	// "encoding/hex"
	"crypto/hmac"
	"crypto/sha256"
)

type Response struct {
	CashBalance float64 `json: "cashBalance"`
}

func hmacsha256(s, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return string(hashed.Sum(nil))
}

func Balance(access_key_id, access_secret string) (balance float64, err error) {
	// # 1.AK/SK、host、method、URL绝对路径、querystring
	AK := access_key_id
	SK := access_secret
	host := "billing.baidubce.com"
	method := "POST"
	query := ""
	URI := "/v1/finance/cash/balance"

	// # 2.x-bce-date,要UTC时间
	// x_bce_date := time.Now().Format("2006-01-02T15:04:05Z")
	x_bce_date := time.Now().UTC().Format("2006-01-02T15:04:05Z")

	// # 3.认证字符串前缀
	authStringPrefix := "bce-auth-v1" + "/" + AK + "/" + x_bce_date + "/" + "1800"

	// # 4.生成signingKey
	signingKey := hex.EncodeToString([]byte(hmacsha256(authStringPrefix, SK)))

	// # 5.header和signedHeaders
	header := map[string]string{"Host": host, "content-type": "application/json;charset=utf-8", "x-bce-date": x_bce_date}
	signedHeaders := "content-type;host;x-bce-date"

	// # 6.生成CanonicalRequest
	// #6.1生成CanonicalURI
	// CanonicalURI := url.QueryEscape(URI)
	CanonicalURI := URI
	// #6.2生成CanonicalQueryString
	CanonicalQueryString := query
	// #6.3生成CanonicalHeaders
	var result []string
	for k, v := range header {
		tempStr := url.QueryEscape(strings.ToLower(k)) + ":" + url.QueryEscape(v)
		result = append(result, tempStr)
	}
	sort.Strings(result)
	CanonicalHeaders := strings.Join(result, "\n")
	// #6.4拼接得到CanonicalRequest
	CanonicalRequest := method + "\n" + CanonicalURI + "\n" + CanonicalQueryString + "\n" + CanonicalHeaders

	// # 7.生成Signature
	Signature := hex.EncodeToString([]byte(hmacsha256(CanonicalRequest, signingKey)))

	// # 8.生成Authorization并放到header里
	header["Authorization"] = authStringPrefix + "/" + signedHeaders + "/" + Signature

	// # 9.POST数据
	payload := "{}"
	body := []byte(payload)

	// # 10.请求API
	var resp *http.Response
	resp, err = common.HttpPost(fmt.Sprintf("https://%s%s", host, URI), header, body)
	if err != nil {
		return
	}

	var response Response
	resp_body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(resp_body, &response)
	if err != nil {
		return
	}

	// curl := fmt.Sprintf("curl -X POST https://%s%s \\\n", host, URI)
	// for k, v := range header {
	// 	curl = fmt.Sprintf("%s-H \"%s:%s\" \\\n", curl, k, v)
	// }
	// curl = fmt.Sprintf("%s -d \"%s\"", curl, payload)
	// fmt.Println(curl)

	return response.CashBalance, nil

}
