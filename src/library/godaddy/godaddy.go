package godaddy

import (
	"encoding/json"
	"fmt"
	gopubssh "library/ssh"
	"models"
	"models/response"
	"strconv"
)

type GoDaddy struct {
	Key		string
	Secret 	string
}
func AnalysisGodaddyRespSlice (godaddyByte string)(jsonGodaddy []*response.RespGodaddy,err error){
	jsonGodaddy = make([]*response.RespGodaddy,10)

	if err = json.Unmarshal([]byte(godaddyByte), &jsonGodaddy); err != nil{
		fmt.Println("error：  AnalysisGodaddyRespSlice.json.Unmarshal is failed  ::  ", err)
		return nil, err
	}
	return jsonGodaddy,nil
}

func AnalysisGodaddayToDataBase (jsonGodaddy []*response.RespGodaddy, godaddy *GoDaddy)(failedDomains []*models.Godaddy){

	for _, item := range jsonGodaddy {
		tmp := new(models.Godaddy)
		tmp.Domain = item.Domain
		tmp.DomainId = strconv.Itoa(item.DomainId)
		tmp.Expires = item.Expires
		tmp.Status = item.Status

		GetGoDaddyDomainInfo(item.Domain,godaddy)
		err := models.AddOrUpdateGodaddy(tmp)
		if err != nil{
			failedDomains = append(failedDomains, tmp)
		}
	}

	return
}

func GetGoDaddyDomainInfo(domain string, godaddy *GoDaddy){

	//cmd := fmt.Sprintf("curl -X GET -H \"Authorization: sso-key %s:%s\" \"https://api.godaddy.com/v1/domains?limit=500\"",godaddy.Key,godaddy.Secret)
	//cmd := fmt.Sprintf("curl -X GET -H \"Authorization: sso-key %s:%s\" \"https://api.godaddy.com/v1/domains?limit=500\"",godaddy.Key,godaddy.Secret)

	//cmd := fmt.Sprintf("curl -X GET -H \"Authorization: sso-key %s:%s\" \"https://api.ote-godaddy.com/v1/domains/%s\" -H \"accept: application/json\" ",godaddy.Key, godaddy.Secret, domain)
	cmd := fmt.Sprintf("curl -X \"GET\" \"https://api.ote-godaddy.com/v1/domains/%s\" -H \"accept: application/json\"  -H \"Authorization: sso-key %s:%s\"",domain, godaddy.Key, godaddy.Secret)
	res := gopubssh.LocalExec(cmd)
	fmt.Printf("res: %s\n",res.Result)

}