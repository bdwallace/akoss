package wallecontrollers

import (
	"fmt"
	"library/common"
	"os"
	"strings"
)

// @Title 下载站 upload file
// @Description 下载站 upload file
// @Param   name     query    	 int   	 true      "name"
// @Param   file     query		*Task  	 true	   "file"
// @Success 0 {id} int64
// @Failure 1 upload file 失败
// @Failure 2 User not found
// @router /walle/upload [post]
func (c *WalleController) Upload(){

	name := c.GetString("name")
	file, _, err := c.GetFile("file")
	if err != nil {
		c.SetJson(1, nil, "")
		return
	}
	defer file.Close()

	dir := fmt.Sprintf("upload/task/%s", strings.Split(name, ".")[0])
	path := fmt.Sprintf("upload/task/%s/%s", strings.Split(name, ".")[0], name)
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err := common.Mkdir(dir)
		if err != nil {
			c.SetJson(1, nil, "upload 创建目录失败")
			return
		}
	}

	err = c.SaveToFile("file", path)
	if err != nil {
		c.SetJson(1, nil, "保存文件失败")
	} else {
		c.SetJson(0, name, "文件上传成功")
	}

	return
}

