package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Deploy struct {
	Id     int    `orm:"column(id);auto"`
	Action int16  `orm:"column(action)"`
	Status int16  `orm:"column(status)" description:"0未上线1状态未定义,2审核失败,3上线完成,4上线失败"`
	Title  string `orm:"column(title);size(100);null"`
	IsRun  int    `orm:"column(is_run);null"`
	Class  string `orm:"column(class)"`
	Count  int    `orm:"column(count);default(0)"`

	CreatedAt time.Time `orm:"column(created_at);type(datetime);auto_now_add;"`
	UpdatedAt time.Time `orm:"column(updated_at);type(datetime);auto_now;"`

	Project *Project `orm:"rel(fk)"`
	User    *User    `orm:"rel(fk)"`
	Tasks   []*Task  `orm:"reverse(many)"`
}

// func (t *Deploy) TableName() string {
// 	return "deploy"
// }

func init() {
	orm.RegisterModelWithPrefix("t_", new(Deploy))
}

/*
	添加 user id 外键约束
*/
func (c *SqlClass) DeployCreateForeignKeyUser() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", deployTableName, deployUserForeignKeyName, userId, userTableName, primaryKey)

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
	创建 project 外键约束
*/
func (c *SqlClass) DeployCreateForeignKeyProject() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", deployTableName, deployProjectForeignKeyName, projectId, projectTableName, primaryKey)

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

func AddDeploy(d *Deploy) (id int64, err error) {
	o := orm.NewOrm()
	if id, err = o.Insert(d); err == nil {
		return id, nil
	}

	return -1, err
}

func UpdateDeployAndAddTaskRelated(d *Deploy, tasks []*Task) (deployId int64, err error) {

	o := orm.NewOrm()
	o.Begin()
	deployId, err = o.Update(d)
	if err != nil {
		o.Rollback()
		return -1, err
	}

	_, err = o.InsertMulti(len(tasks), tasks)
	if err != nil {
		o.Rollback()
		return -1, err
	}
	o.Commit()

	return deployId, err
}

func GetDeployById(id int) (d *Deploy, err error) {
	o := orm.NewOrm()
	d = &Deploy{}

	err = o.QueryTable(deployTableName).Filter("id", id).RelatedSel().One(d)
	if err != nil {
		fmt.Println("id: ", id, "    error: ", err)
		return nil, err
	}

	return
}

func GetDeployByProjectId(projectId int) (t []*Deploy, err error) {

	o := orm.NewOrm()
	_, err = o.QueryTable(deployTableName).Filter("project_id", projectId).OrderBy("-id").RelatedSel().All(&t)
	if err != nil {
		return nil, err
	}
	return t, err

}

func GetDeployByProjectIdUserId(projectId int, userId int, start int, length int) (count int64, t []*Deploy, err error) {
	o := orm.NewOrm()
	// count, _ = o.QueryTable(crontabLogTableName).Count()
	// o.QueryTable(crontabLogTableName).Limit(length, start).OrderBy("-id").All(&m)
	if userId == 0 {
		count, _ = o.QueryTable(deployTableName).Filter("project_id", projectId).Count()
		_, err = o.QueryTable(deployTableName).Filter("project_id", projectId).Limit(length, start).OrderBy("-id").RelatedSel("user").All(&t)
	} else {
		count, _ = o.QueryTable(deployTableName).Filter("project_id", projectId).Filter("user_id", userId).Count()
		_, err = o.QueryTable(deployTableName).Filter("project_id", projectId).Filter("user_id", userId).Limit(length, start).OrderBy("-id").RelatedSel("user").All(&t)
	}
	if err != nil {
		return 0, nil, err
	}
	return

}

func GetDeployByServiceId(serviceId int) (d []*Deploy, err error) {

	o := orm.NewOrm()
	_, err = o.QueryTable(deployTableName).Filter("service_id", serviceId).RelatedSel().All(&d)
	if err != nil {
		return nil, err
	}
	if _, err = o.LoadRelated(d, "Tasks"); err != nil {
		return nil, err
	}
	return d, err
}

func GetAllDeploy() (d []*Deploy, err error) {

	o := orm.NewOrm()
	if _, err = o.QueryTable(deployTableName).All(&d); err != nil {
		return nil, err
	}
	return d, nil

}

func UpdateDeployById(d *Deploy) (id int64, err error) {
	o := orm.NewOrm()

	if id, err = o.Update(d); err == nil {
		return -1, err
	}
	return id, nil
}

func DeleteDeploy(id int) (num int64, err error) {

	o := orm.NewOrm()
	if num, err = o.Delete(&Deploy{Id: id}); err == nil {
		return num, nil
	}
	return -1, err
}

/*
	通过 project id 查询 deploy 以及关联表数据
*/
func GetDeployRelated(deploys *Deploy) error {

	o := orm.NewOrm()

	_, err := o.LoadRelated(deploys, "Tasks")
	if err != nil {
		return err
	}

	for _, task := range deploys.Tasks {
		_, err = o.LoadRelated(task, "Project")
		if err != nil {
			return err
		}

		_, err = o.LoadRelated(task, "Service")
		if err != nil {
			return err
		}

		_, err = o.LoadRelated(task.Service, "Hosts")
		if err != nil {
			return err
		}

		_, err = o.LoadRelated(task.Service, "Domains")
		if err != nil {
			return err
		}

		_, err = o.LoadRelated(task.Service, "Confs")
		if err != nil {
			return err
		}

		_, err = o.LoadRelated(task.Service, "Platforms")
		if err != nil {
			return err
		}

		_, err = o.LoadRelated(task.Service, "Project")
		if err != nil {
			return err
		}

		_, err = o.LoadRelated(task, "Hosts")
		if err != nil {
			return err
		}

	}

	return nil

}
