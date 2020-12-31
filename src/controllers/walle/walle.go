package wallecontrollers

import (
	"controllers"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"library/common"
	"library/components"
	"models"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type WalleController struct {
	controllers.BaseController
}

// @Title 获取 tag list for ImagePath
// @Description 根据 Service.ImagePath 查询所有 tag list 信息
// @Param   image_path query     string true         "service image_path"
// @Success 0 {list} {["tag", ...]}
// @Failure 1 获取 失败
// @Failure 2 User not found
// @router /walle/taglist [get]
func (c *WalleController) GetTagList() {

	imagePath := c.GetString("image_path")

	repoApi := beego.AppConfig.String("repoApi")
	// tag_url := fmt.Sprintf("%s/api/repositories/%s/tags", repoApi, imagePath)
	repo_project := strings.Split(imagePath, "/")[0]
	repo_repository := strings.Split(imagePath, "/")[1]
	tag_url := fmt.Sprintf("%s/api/v2.0/projects/%s/repositories/%s/artifacts", repoApi, repo_project, repo_repository)
	tag_url = fmt.Sprintf("%s?with_tag=true&with_scan_overview=true&with_label=true&page_size=220&page=1", tag_url)
	//beego.Info(tag_url)

	req := httplib.Get(tag_url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	bytes, err := req.Bytes()
	if err != nil {
		c.SetJson(1, nil, "获取tag错误—"+err.Error())
		return
	}

	type Tags struct {
		Name string
	}

	type Artifacts struct {
		Tags []Tags
	}

	var data []Artifacts
	json.Unmarshal(bytes, &data)

	var reList []string
	for _, s := range data {
		for _, ss := range s.Tags {
			if ss.Name != "" {
				reList = append(reList, ss.Name)
			}
		}
	}

	c.SetJson(0, reList, "")
	return

}

// @Title 获取 获取nacos的下线列表
// @Description 根据 Project.Nacos 查询它当套环境的下线主机
// @Param   projects_id      query     int   		true         "projects id"
// @Success 0 {list} {["host:port", ...]}
// @Failure 1 获取 失败
// @Failure 2 User not found
// @router /walle/lineget [get]
func (c *WalleController) LineGet() {

	levelId, _ := c.GetInt("project_id")

	project, _ := models.GetProjectById(levelId)
	url := "http://" + project.Nacos + "/nacos/v1/cs/configs"

	ipList, err := components.LineGet(url)
	if err != nil {
		c.SetJson(1, err.Error(), "访问nacos的下线列表失败")
		return
	}

	c.SetJson(0, ipList, "")
	return
}

// @Title 获取 获取nacos的下线列表
// @Description 根据 Project.Nacos 查询它当套环境的下线主机
// @Param   projects_id      query     int   		true         "projects id"
// @Success 0 {list} {["host:port", ...]}
// @Failure 1 获取 失败
// @Failure 2 User not found
// @router /walle/linechange [get]
func (c *WalleController) LineChange() {

	nacos := c.GetString("nacos")
	hostPort := c.GetString("host_port")
	line := c.GetString("line")

	url := "http://" + nacos + "/nacos/v1/cs/configs"
	lineData, err := components.LineGet(url)
	if err != nil {
		c.SetJson(1, err.Error(), "获取nacos接口数据返回失败")
		return
	}

	lineList := strings.Split(lineData, ",")
	if line == "on" {
		for common.InListIndex(hostPort, lineList) != -1 {
			i := common.InListIndex(hostPort, lineList)
			lineList = append(lineList[:i], lineList[i+1:]...)
		}
	}

	if line == "off" {
		if common.InListIndex(hostPort, lineList) == -1 {
			lineList = append(lineList, hostPort)
		}
	}

	reqPost := httplib.Post(url)
	reqPost.Param("dataId", "grayReleaseConfig.properties")
	reqPost.Param("group", "DEFAULT_GROUP")
	reqPost.Param("type", "properties")
	content1 := "grayRelease.enable=true\r\n"
	content2 := "grayRelease.grayLogEnabled=true\r\n"
	content3 := "grayRelease.grayClientIpList=\r\n"
	content4 := "grayRelease.grayServerList="
	content := content1 + content2 + content3 + content4 + strings.Trim(strings.Join(lineList, ","), ",")
	// content := content1 + content2 + content3 + content4 + strings.Trim(lineListStr, ",")
	reqPost.Param("content", content)
	_, err = reqPost.String()
	if err != nil {
		c.SetJson(1, err.Error(), "向nacos发送数据失败")
		return
	}
	//beego.Info(fmt.Sprintf("%s %s %s", hostPort, line, resultPost))

	c.SetJson(0, "", "")
	return
}
