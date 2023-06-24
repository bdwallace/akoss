package components

import (
	"models"
	"strings"

	"github.com/astaxie/beego/httplib"
)

/*
	执行 docker ps 查看容器和上下线状态
*/
func LineGet(s * models.Service, url string) (lineStr string, err error) {
	req := httplib.Get(url)
	req.Param("dataId", "grayReleaseConfig.properties")
	req.Param("group", "DEFAULT_GROUP")

	if s.NacosUserName != "" && s.NacosPwd != ""{
		req.Param("username", s.NacosUserName)
		req.Param("password", s.NacosPwd)
	}

	result, err := req.String()
	//beego.Info(fmt.Sprintf("curl: %s", url))

	resultList := strings.Split(result, "=")
	lineStr = resultList[len(resultList)-1]
	//beego.Info(fmt.Sprintf("lineStr: %s", lineStr))
	// ipList = strings.Split(ipListStr, ",")
	// beego.Info(fmt.Sprintf("level:%s %v", levelId, ipList))

	return
}