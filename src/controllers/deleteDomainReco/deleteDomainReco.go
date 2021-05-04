package deleteDomainRecoControllers

import (
	"controllers"
	"fmt"
	"models"
)

type DeleteDomainRecoController struct {
	controllers.BaseController
}



/*
// @Title 获取所有 delete domain list
// @Description get all delete domain list
// @Success 0 {object} models.DeleteDomainReco
// @Failure 1 获取所有 delete domain reco 失败
// @Failure 2 User not found
// @router /domainReco [get]
func (c *DeleteDomainRecoController)GetAllDomainReco() {

}
*/


// @Title search and get all delete domain
// @Description search get all delete domain
// @Success 0 {object} models.DeleteDomainReco
// @Failure 1 search get all delete domain reco 失败
// @Failure 2 User not found
// @router /domainReco [get]
func (c *DeleteDomainRecoController)GetDomainReco() {

	searchText := c.GetString("search_text")
	class := c.GetString("class")
	projectId, err := c.GetInt("project_id")
	if err != nil{
		c.SetJson(1,err,"获取 project_id 失败")
		return
	}

	page, _ := c.GetInt("page", 0)
	length, _ := c.GetInt("length", 20)
	start := 0
	var total int64

	if page > 0 {
		start = (page - 1) * length
	}
	resDomainReco := make([]*models.DeleteDomainReco, 0)
	total, resDomainReco, err = models.SearchDomainRecords(projectId,start,length, class, searchText)
	if err != nil{
		c.SetJson(1,err,"搜索 domain reco 匹配内容 失败")
		return
	}

	fmt.Println("total: ",total)
	fmt.Println("len: ",len(resDomainReco))
	fmt.Println("page: ",page)
	c.SetJson(0, map[string]interface{}{"total": total, "currentPage": page, "table_data": resDomainReco}, "")
	return

}


// @Title Get DomainReco By ProjectId For Class
// @Description Get DomainReco By ProjectId For Class
// @Success 0 {object} models.DeleteDomainReco
// @Failure 1 Get DomainReco By ProjectId For Class 失败
// @Failure 2 User not found
// @router /domainReco/class [get]
func (c *DeleteDomainRecoController)GetDomainRecoByProjectForClass(){

	projectId, err :=  c.GetInt("project_id")
	if err != nil{
		c.SetJson(1, err, "获取project_id 失败")
		return
	}

	data, err := models.GetDomainRecoByProjectForClass(projectId)
	if err != nil{
		c.SetJson(1, err,"获取 domainReco class by project id失败")
		return
	}

	c.SetJson(0, data, "")
	return
}
