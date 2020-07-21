package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Platform struct {
	Id        int               `orm:"column(id);pk;auto"`
	Name      string            `orm:"column(name);size(100);unique"`          // 平台名称
	ChannelId string            `orm:"column(channel_id);size(100);"`          // 平台名称
	Value 	  string 			`orm:"column(value);type(text)"`                      // 配置项， 以 k -> v 形式储存json

	CreatedAt              time.Time `orm:"column(created_at);type(datetime);auto_now_add;"`
	UpdatedAt              time.Time `orm:"column(updated_at);type(datetime);auto_now;"`

	Project		*Project		`orm:"rel(fk)"`
	Services	[]*Service  	`orm:"reverse(many)"`
	Domains 	[]*Domain		`orm:"reverse(many)"`
}


//  platformConf
type PlatformValue struct {
	Key 	string	`json:"Key"`
	Value 	string	`json:"Value"`
}

// func (c *Platform) TableName() string {
//     // db table name
//     return "t_platform"
// }

func init(){
	// orm.RegisterModel(new(Platform))
	orm.RegisterModelWithPrefix("t_",new(Platform))
}



/*
	创建 platform 中 porject_id
*/
func (c * SqlClass)PlatformCreateForeignKeyProject() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf( "alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);",platformTableName,platformProjectForeignKeyName,projectId,projectTableName,primaryKey)

	// 创建 domain  platform 外键约束
	err := c.CreateForeignKey(addForeignSqlStr)
	if err != nil{
		if strings.Index(err.Error(), errDuplicate) != -1 {
			beego.Info("key is duplicate")
		} else {
			beego.Error(addForeignSqlStr, " : ", err.Error())
		}
	} else {
		beego.Info("key is create")
	}
}






////////  platform curd

func GetAllPlatform(projectId int) (p []*Platform, err error) {

	o := orm.NewOrm()
	if _, err = o.QueryTable(platformTableName).Filter("project_id",projectId).RelatedSel().All(&p);err != nil{
		return nil,err
	}
	return p,err
}



func SearchPlatform(searchText string, projectId int)(p []*Platform, err error){

	cond := orm.NewCondition()
	// search := cond.And("project_id",projectId).And("name__icontains",searchText).Or("value__icontains",searchText).And("project_id",projectId).Or("channel_id__icontains",searchText).And("project_id",projectId)
	or := cond.Or("name__icontains",searchText).Or("value__icontains",searchText).Or("channel_id__icontains", searchText)
	search := cond.And("project_id", projectId).AndCond(or)

	o := orm.NewOrm()
	qs := o.QueryTable(platformTableName)
	if _, err := qs.SetCond(search).RelatedSel().All(&p); err != nil{
		return nil, err
	}

	return p, nil
}



func GetPlatformById(id int) (p *Platform, err error) {

	o := orm.NewOrm()
	p = &Platform{}
	err = o.QueryTable(platformTableName).Filter("id", id).RelatedSel().One(p)
	if err != nil {
		return nil, err
	}
	return
}

func GetPlatformByProjectId(projectId int) (p []*Platform, err error) {

	o := orm.NewOrm()
	_ ,err = o.QueryTable(platformTableName).Filter("project_id",projectId).RelatedSel().All(&p)
	if err != nil {
		return nil, err
	}
	return p,err
}


func GetPlatformByProjectForName(projectId int) (list *orm.ParamsList, err error) {
	list = new(orm.ParamsList)
	o := orm.NewOrm()
	_ ,err = o.QueryTable(platformTableName).Filter("project_id",projectId).GroupBy("name").OrderBy("name").ValuesFlat(list, "name")
	return
}


func GetPlatformAndDomainRelated(platform *Platform) ( *Platform,  error) {

	o := orm.NewOrm()
	if _, err := o.LoadRelated(platform,"Domains"); err != nil{
		return platform, err
	}

	return platform,nil
}



func AddPlatform(p *Platform) (id int64, err error) {

	o := orm.NewOrm()

	id, err = o.Insert(p)
	if err != nil{
		return
	}
	return
}


func UpdatePlatformById(p *Platform)(id int64, err error){

	o := orm.NewOrm()

	if id, err = o.Update(p); err != nil{
		return
	}
	return
}


func DeletePlatformById(id int) (num int64, err error) {

	o := orm.NewOrm()
	if num, err = o.Delete(&Platform{Id: id}); err == nil {
		return num, nil
	}
	return -1, err
}

//////////////
/*
	多对多操作
 */

/*
	通过 platform id 查询 platform 以及关联表数据
 */
func GetPlatformRelated(platform *Platform) (*Platform, error){

	o := orm.NewOrm()

	tempPlatform := new(Platform)
	tempPlatform.Id = platform.Id


	if _, err := o.LoadRelated(tempPlatform,"Services"); err != nil{
		return platform, err
	}

	for _, s := range tempPlatform.Services {
		if s.IsDel != 0{
			continue
		}

		if _, err := o.LoadRelated(s,"Project"); err != nil{
			return platform, err
		}
		if _, err := o.LoadRelated(s,"Hosts"); err != nil{
			return platform, err
		}
		// if _, err := o.LoadRelated(s,"Platforms"); err != nil{
		// 	return platform, err
		// }

		platform.Services = append(platform.Services,s)
	}


	// if _, err := o.LoadRelated(platform,"Domains"); err != nil{
	// 	return platform, err
	// }


	return platform, nil
}







///////  多对多操作  platform   <----->   service

/*
	通过 platform id 查询 platfrom 以及关联表数据
*/
func GetPlatformAndServiceRelated(platfrom *Platform) (*Platform, error){

	o := orm.NewOrm()

	if _, err := o.LoadRelated(platfrom,"Services"); err != nil{
		return nil, err
	}

	return platfrom, nil
}



/*
	添加 platform 以及关联表 数据
*/

func AddPlatformRelatedForServices(platform *Platform)error {

	o := orm.NewOrm()

	// 校验 platform.Services 数据 是否为 nil
	if len(platform.Services) != 0{
		// 建立 host 关系, 插入service_host表 host 数据
		m2mForPlatfromAndServices := o.QueryM2M(platform,"Services")
		_, err := m2mForPlatfromAndServices.Add(platform.Services)
		if err != nil{
			fmt.Println("error:  m2mForPlatformAndServices.Add(platform.Services):  ",err)
			return err
		}
	}

	return nil
}


/*
	更新 platform 以及关联表 数据
*/
func UpdatePlatformAndServicesRelated(platform *Platform) (err error) {

	o := orm.NewOrm()

	_, err = o.Update(platform)
	if err != nil {
		err = fmt.Errorf("error: UpdatePlatformAndServicesRelated update platfrom  fail!")
		return
	}


	// 获取 platfrom 关联的 services， 直接全部清除
	m2mForPlatfromAndServices:= o.QueryM2M(platform,"Services")
	if _, err = m2mForPlatfromAndServices.Clear(); err != nil{
		return
	}

	// 不存在关联关系,需要建立关联关系
	if err := AddPlatformRelatedForServices(platform); err != nil{
		fmt.Println("error:  add related ",err)
		return err
	}

	return nil
}


/*
	增加 platform 以及关联表 数据
*/
func AddPlatformAndServiceRelated(platform *Platform) (id int64, err error) {

	o := orm.NewOrm()

	fmt.Println("==============")
	fmt.Printf("platform: %+v\n",platform)
	id, err = o.Insert(platform)
	if err != nil{
		err = fmt.Errorf("error:  AddPlatformAndServiceRelated to insert service fail!")
		return
	}
	// 不存在关联关系,需要建立关联关系
	platform.Id = int(id)
	if err = AddPlatformRelatedForServices(platform); err != nil{
		fmt.Println("error:  AddPlatformRelatedForServices  ",err)
		return
	}
	return
}



/*
	删除 platform 与 service 关联关系
*/
func DeletePlatformAndServiceRelatedByPlatformId(platform *Platform)(err error) {

	o := orm.NewOrm()

	// 获取 platfrom 关联的 services， 直接全部清除
	m2mForPlatformAndServices:= o.QueryM2M(platform,"Services")
	if _, err = m2mForPlatformAndServices.Clear(); err != nil{
		return
	}

	return
}

