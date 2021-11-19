package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Task struct {
	Id     int    `orm:"column(id);auto"`
	Action int16  `orm:"column(action);null"`
	Status int16  `orm:"column(status);null"`
	Tag    string `orm:"column(tag);size(800);null"`
	// IsRun          	int       				`orm:"column(is_run);null"`
	Count int    `orm:"column(count);default(0)"`
	Cmd   string `orm:"column(cmd);type(text)"`

	CreatedAt time.Time `orm:"column(created_at);type(datetime);auto_now_add;"`
	UpdatedAt time.Time `orm:"column(updated_at);type(datetime);auto_now;"`

	Project *Project `orm:"rel(fk)"`
	User    *User    `orm:"rel(fk)"`
	Deploy  *Deploy  `orm:"rel(fk)"`
	Service *Service `orm:"rel(fk)"`
	Hosts   []*Host  `orm:"rel(m2m)"`
}

// func (t *Task) TableName() string {
// 	return "task"
// }

func init() {
	orm.RegisterModelWithPrefix("t_", new(Task))
}

/*
	创建 service 中 project_id
*/
func (c *SqlClass) TaskCreateForeignKeyProject() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", taskTableName, taskProjectForeignKeyName, projectId, projectTableName, primaryKey)

	// 创建 service  service_conf 外键约束
	err := c.CreateForeignKey(addForeignSqlStr)
	if err != nil {
		if strings.Index(err.Error(), errDuplicate) != -1 {
			beego.Info("key is duplicate")
		} else {
			beego.Error(addForeignSqlStr, " : ", err.Error())
		}
	} else {
		beego.Info("key is create")
	}
}

/*
	添加 user id 外键约束
*/
func (c *SqlClass) TaskCreateForeignKeyUser() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", taskTableName, taskUserForeignKeyName, userId, userTableName, primaryKey)

	err := c.CreateForeignKey(addForeignSqlStr)
	if err != nil {
		if strings.Index(err.Error(), errDuplicate) != -1 {
			beego.Info("key is duplicate")
		} else {
			beego.Error(addForeignSqlStr, " : ", err.Error())
		}
	} else {
		beego.Info("key is create")
	}
}

/*
	添加 service id 外键约束
*/
func (c *SqlClass) TaskCreateForeignKeyService() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", taskTableName, taskServiceForeignKeyName, serviceId, serviceTableName, primaryKey)

	err := c.CreateForeignKey(addForeignSqlStr)
	if err != nil {
		if strings.Index(err.Error(), errDuplicate) != -1 {
			beego.Info("key is duplicate")
		} else {
			beego.Error(addForeignSqlStr, " : ", err.Error())
		}
	} else {
		beego.Info("key is create")
	}
}

/*
	添加 deploy id 外键约束
*/
func (c *SqlClass) TaskCreateForeignKeyDeploy() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", taskTableName, taskDeployForeignKeyName, deployId, deployTableName, primaryKey)

	err := c.CreateForeignKey(addForeignSqlStr)
	if err != nil {
		if strings.Index(err.Error(), errDuplicate) != -1 {
			beego.Info("key is duplicate")
		} else {
			beego.Error(addForeignSqlStr, " : ", err.Error())
		}
	} else {
		beego.Info("key is create")
	}
}

// 创建 t_task_t_Host  task_id外键约束
func (c *SqlClass) TaskHostCreateForeignKeyForHost() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", taskHostTableName, taskHostForHostForeignKeyName, taskHostForHostId, hostTableName, primaryKey)

	err := c.CreateForeignKey(addForeignSqlStr)
	if err != nil {
		if strings.Index(err.Error(), errDuplicate) != -1 {
			beego.Info("key is duplicate")
		} else {
			beego.Error(addForeignSqlStr, " : ", err.Error())
		}
	} else {
		beego.Info("key is create")
	}
}

// 创建 t_task_t_Host  host_id外键约束
func (c *SqlClass) TaskHostCreateForeignKeyForTask() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", taskHostTableName, taskHostForTaskForeignKeyName, taskHostForTaskId, taskTableName, primaryKey)

	err := c.CreateForeignKey(addForeignSqlStr)
	if err != nil {
		if strings.Index(err.Error(), errDuplicate) != -1 {
			beego.Info("key is duplicate")
		} else {
			beego.Error(addForeignSqlStr, " : ", err.Error())
		}
	} else {
		beego.Info("key is create")
	}
}

/*
	添加 platform id 外键约束
*/
// func (c * SqlClass)TaskCreateForeignKeyPlatform() {

// 	// 拼接两张表外键约束sql
// 	addForeignSqlStr := fmt.Sprintf( "alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);",taskTableName,taskPlatformForeignKeyName,platformId,platformTableName,primaryKey)

// 	if err := c.CreateForeignKey(addForeignSqlStr); err != nil{
// 		errStr := DefaultFailOnError("task platform_id foreign key sql exec",err)
// 		fmt.Println(errStr)
// 		return
// 	}else {
// 		fmt.Println("task tables platform_id foreign key relationship created successfully")
// 	}
// }

////////  task curd

func AddTask(t *Task) (id int64, err error) {
	o := orm.NewOrm()
	if id, err = o.Insert(t); err == nil {
		return id, nil
	}

	return -1, err
}

func GetTaskById(id int) (t *Task, err error) {
	o := orm.NewOrm()
	t = &Task{}

	err = o.QueryTable(taskTableName).Filter("id", id).RelatedSel().One(t)
	// if err := o.Read(t); err != nil {
	// 	fmt.Println("id: ",id,"    error: ",err)
	// 	return nil, err
	// }

	return
}

func GetTaskByServiceId(serviceId int) (t []*Task, err error) {

	o := orm.NewOrm()
	_, err = o.QueryTable(taskTableName).Filter("service_id", serviceId).RelatedSel().All(&t)
	if err != nil {
		return nil, err
	}
	return t, err

}

func GetTaskByDeployId(deployId int) (t []*Task, err error) {

	o := orm.NewOrm()
	_, err = o.QueryTable(taskTableName).Filter("deploy_id", deployId).RelatedSel().All(&t)
	if err != nil {
		return nil, err
	}
	return t, err

}

func GetTaskByDeployIdAndServiceId(deployId int64, serviceId int64) (*Task, error) {

	t := new(Task)

	o := orm.NewOrm()
	err := o.QueryTable(taskTableName).Filter("deploy_id", deployId).Filter("service_id", serviceId).One(t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func GetAllTask() (t []*Task, err error) {

	o := orm.NewOrm()
	if _, err = o.QueryTable(taskTableName).All(&t); err != nil {
		return nil, err
	}
	return t, nil

}

func UpdateTaskById(t *Task) (id int64, err error) {
	o := orm.NewOrm()

	if id, err = o.Update(t); err == nil {
		return -1, err
	}
	return id, nil
}

func DeleteTask(id int) (num int64, err error) {

	o := orm.NewOrm()
	if num, err = o.Delete(&Task{Id: id}); err == nil {
		return num, nil
	}
	return -1, err
}
