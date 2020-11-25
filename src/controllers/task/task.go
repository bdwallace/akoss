package taskcontrollers

import (
	"controllers"
	"encoding/json"
	"fmt"
	"models"
)

type TaskController struct {
	controllers.BaseController
}


///////////
/*
	task  curd
 */


// @Title 添加 task
// @Description 添加 task

// @Success 0 {id} int
// @Failure 1 添加 task 失败
// @Failure 2 User not found
// @router /task/ [post]
func (c *TaskController) AddTask() {

	//beego.Info(string(c.Ctx.Input.RequestBody))
	task := new(models.Task)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, task)
	if err != nil {
		c.SetJson(1, nil, "数据格式解析错误")
		return
	}

	var id int64
	if task.Id != 0 {
		if id, err = models.UpdateTaskById(task);err != nil{
			c.SetJson(1, nil, "更新 task by id 失败")
			return
		}
	}

	if task.Id == 0{
		if id, err = models.AddTask(task); err != nil{
			c.SetJson(1, nil, "添加 task by id 失败")
			return
		}
	}

	c.SetJson(0, id, "")
	return
}



// @Title 获取 task by Id
// @Description 获取 task by Id
// @Param   id       query     int  	true       "task id"
// @Success 0 {object} models.Task
// @Failure 1 获取 task by Id 失败
// @Failure 2 User not found
// @router /task/id/ [get]
func (c *TaskController) GetTaskById() {

	taskId, _ := c.GetInt("id")
	task, err := models.GetTaskById(taskId)
	if err != nil{
		c.SetJson(1, nil, "获取 task by Id 失败")
		return
	}

	c.SetJson(0, task, "")
	return
}




// @Title 获取 tasks by service id
// @Description get tasks by service id
// @Param   service_id      query     int  	  true      "service id"
// @Success 0 {object} models.Task
// @Failure 1 获取 task by Id 失败
// @Failure 2 User not found
// @router /task/serviceId/ [get]
func (c *TaskController) GetTaskByServiceId() {

	serviceId, err :=  c.GetInt("service_id")
	if err != nil{
		c.SetJson(1, err, "获取 service id 失败")
		return
	}

	resTask, err := models.GetTaskByServiceId(serviceId)
	if err != nil{
		c.SetJson(1, err,"获取 host by service id失败")
		return
	}

	c.SetJson(0,resTask,"")
	return
}



// @Title 获取所有 task
// @Description get all the task
// @Success 0 {object} models.Task
// @Failure 1 获取所有 task 失败
// @Failure 2 User not found
// @router /task/list/ [get]
func (c *TaskController)GetAllTask() {

	resTask, err := models.GetAllTask()

	if err != nil {
		fmt.Println("error: GetAllTask()", err)
		c.SetJson(1, err, "获取所有 task 失败")
		return
	}

	c.SetJson(0, resTask, "")
	return

}



// @Title delete task by task id
// @Description delete task by task id
// @Param   id      query     int 		true       "task id"
// @Success 0 {object} models.Task
// @Failure 1 删除 task by task id 失败
// @Failure 2 User not found
// @router /task/id/ [delete]
func  (c *TaskController)DeleteTask(id int) () {

	resId, err := models.DeleteTask(id)
	if err != nil{
		c.SetJson(1, err, "删除 task by id 失败")
		return
	}

	c.SetJson(0, resId, "")
	return

}



