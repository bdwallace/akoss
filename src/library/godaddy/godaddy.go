package godaddy

import (
	"encoding/json"
	"fmt"
	"models"
	"models/response"
	"strconv"
)

func AnalysisGodaddyRespSlice (godaddyByte string)(jsonGodaddy []*response.RespGodaddy,err error){
	jsonGodaddy = make([]*response.RespGodaddy,10)

	if err = json.Unmarshal([]byte(godaddyByte), &jsonGodaddy); err != nil{
		fmt.Println("error：  AnalysisGodaddyRespSlice.json.Unmarshal is failed  ::  ", err)
		return nil, err
	}
	return jsonGodaddy,nil
}

func AnalysisGodaddayToDataBase (jsonGodaddy []*response.RespGodaddy)(failedDomains []*models.Godaddy){

	for _, item := range jsonGodaddy {
		tmp := new(models.Godaddy)
		tmp.Domain = item.Domain
		tmp.DomainId = strconv.Itoa(item.DomainId)
		tmp.Expires = item.Expires
		tmp.Status = item.Status
		err := models.AddOrUpdateGodaddy(tmp)
		if err != nil{
			failedDomains = append(failedDomains, tmp)
		}
	}

	return
}