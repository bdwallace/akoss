package qiniu

import (
	"fmt"
	"io/ioutil"
	"library/common"
	"net/http"
	"strings"
)

func Balance(requstUri, cookie string) (*float64, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", requstUri, nil)
	if err != nil {
		fmt.Println("error: HttpAuth http.NewRequest ", err)
		return nil, err
	}
	req.Header.Add("cookie", cookie)

	httpResp, err := client.Do(req)
	if err != nil {
		fmt.Println("error: HttpAuth client.Do  ", err)
		return nil, err
	}
	defer httpResp.Body.Close()

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		fmt.Println("error: HttpAuth ioutil.ReadAll  ", err)
		return nil, err
	}

	strBody := string(body)
	var balanceStr string
	if index := strings.Index(strBody, "cash_balance"); index > 0 {
		str1 := strBody[index+19:]
		index2 := strings.Index(str1, "\",")
		balanceStr = str1[:index2]
	}

	balanceFloat := common.StrToFloat(strings.Replace(balanceStr, ",", "", -1))

	return &balanceFloat, err

}
