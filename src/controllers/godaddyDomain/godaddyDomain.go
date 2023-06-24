package godaddycontrollers

import (
	"controllers"
	"fmt"
	libGodaddy "library/godaddy"
	gopubssh "library/ssh"
	"models"
)

type GodaddyControllers struct {
	controllers.BaseController
}


// @Title get all domain from godaddy
// @Description get all domain from godaddy
// @Success 0 {object} models.GodaddyDomain
// @Failure 1 获取 godaddy domain 失败
// @Failure 2 User not found
// @router /godaddy [get]
func (c * GodaddyControllers) GetAllDomain(){

	//godaddy := createGodaddyKey()
	godaddy := new(libGodaddy.GoDaddy)
	var searchDomain string
	godaddy.Key  = c.GetString("godaddy_key")
	godaddy.Secret = c.GetString("godaddy_secret")
	searchDomain = c.GetString("search_domain")

	resGodaddy := make([]*models.Godaddy,0)
	resSearchDomains := make([]*models.Godaddy,0)
	var err error

	if godaddy.Key != "" && godaddy.Secret != "" {
		// get remote
		cmd := fmt.Sprintf("curl -X GET -H \"Authorization: sso-key %s:%s\" \"https://api.godaddy.com/v1/domains?limit=500\"",godaddy.Key,godaddy.Secret)
		res := gopubssh.LocalExec(cmd)
		if res.Error != nil{
			c.SetJson(1, res.Error, "获取godaddy域名失败")
			return
		}

		jsonGodaddy, err := libGodaddy.AnalysisGodaddyRespSlice(res.Result)
		if err != nil{
			c.SetJson(1, nil, "解析 godaddy 失败")
			return
		}

		if len(jsonGodaddy) > 0 {
			_ = libGodaddy.AnalysisGodaddayToDataBase(jsonGodaddy,godaddy)
		}
		resGodaddy, err = models.GetAllGodaddyDomains()
		if err != nil{
			c.SetJson(1, nil, "获取 godaddy from database 失败")
			return
		}

	} else if searchDomain != ""{
		// search database
		resSearchDomains, err = models.GetGodaddyDomainByDomainName(searchDomain)
		if err != nil{
			c.SetJson(1, nil, "搜索 godaddy domain 失败")
			return
		}
		resGodaddy = append(resGodaddy,resSearchDomains...)

	} else {
		c.SetJson(1, nil, "请填写 godaddy key / secret 或 domain name")
		return
	}

	//c.SetJson(0, map[string]interface{}{"item_total": resp.Total, "page_num": pageNumber, "domains": resp.Domains}, "")
	c.SetJson(0, map[string]interface{}{"domains": resGodaddy}, "")
	return

}




// @Title Modify domain DNS godaddy
// @Description Modify domain DNS godaddy
// @Success 0 {object} models.GodaddyDomain
// @Failure 1 获取 godaddy domain 失败
// @Failure 2 User not found
// @router /godaddy [post]
func (c * GodaddyControllers) ModifyGodaddyDomainDNS(){

}
