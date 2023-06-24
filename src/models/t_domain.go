package models

import (
	"fmt"
	"library/common"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Domain struct {
	Id      int    `orm:"column(id);auto"`
	Name    string `orm:"column(name);size(200)"`
	Domain  string `orm:"column(domain);size(200)"`
	Port    string `orm:"column(port);size(200)"`
	Class   string `orm:"column(class);size(200)"`
	Conf    string `orm:"column(conf);type(text)"`
	Conf80  string `orm:"column(conf_80);type(text)"`
	Crt     string `orm:"column(crt);type(text)"`
	Key     string `orm:"column(key);type(text)"`
	Monitor int16  `orm:"column(monitor)"`
	Quicken int    `orm:"column(quicken);default(0)"`

	CreatedAt time.Time `orm:"column(created_at);type(datetime);auto_now_add;"`
	UpdatedAt time.Time `orm:"column(updated_at);type(datetime);auto_now;"`

	Project   *Project    `orm:"rel(fk)"`
	Platforms []*Platform `orm:"rel(m2m)"`
	Services  []*Service  `orm:"reverse(many)"`
}

// func (c *Domain) TableName() string {
//     // db table name
//     return domainTableName
// }

func init() {
	// orm.RegisterModel(new(Domain))
	orm.RegisterModelWithPrefix("t_", new(Domain))
}

/*
	创建 domain 中 project_id
*/
func (c *SqlClass) DomainCreateForeignKeyProject() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", domainTableName, domainProjectForeignKeyName, projectId, projectTableName, primaryKey)

	// 创建 domain  project 外键约束
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

// 创建 t_Domain_t_platform  domain外键约束
func (c *SqlClass) DomainPlatformCreateForeignKeyForDomain() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", domainPlatformTableName, domainPlatformForDomainForeignKeyName, domainPlatformForDomainId, domainTableName, primaryKey)

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

// 创建 t_Domain_t_platform  platform外键约束
func (c *SqlClass) DomainPlatformCreateForeignKeyForPlatform() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", domainPlatformTableName, domainPlatformForPlatformForeignKeyName, domainPlatformForPlatformId, platformTableName, primaryKey)

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
	创建 domain 中 service_id
*/
//func (c * SqlClass)DomainCreateForeignKeyService() {
//
//	// 拼接两张表外键约束sql
//	addForeignSqlStr := fmt.Sprintf( "alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);",domainTableName,domainServiceForeignKeyName,serviceId,serviceTableName,primaryKey)
//
//	// 创建 domain  service 外键约束
//	if err := c.CreateForeignKey(addForeignSqlStr); err != nil{
//		errStr := DefaultFailOnError("domain service_id foreign key sql exec",err)
//		fmt.Println(errStr)
//		return
//	}else {
//		fmt.Println("domain tables service_id foreign key relationship created successfully")
//	}
//}

/*
	创建 domain 中 domain_id
*/
// func (c * SqlClass)DomainCreateForeignKeyPlatform() {

// 	// 拼接两张表外键约束sql
// 	addForeignSqlStr := fmt.Sprintf( "alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);",domainTableName,domainPlatformForeignKeyName,domainId,domainTableName,primaryKey)

// 	// 创建 domain  domain 外键约束
// 	if err := c.CreateForeignKey(addForeignSqlStr); err != nil{
// 		errStr := DefaultFailOnError("domain domain_id foreign key sql exec",err)
// 		fmt.Println(errStr)
// 		return
// 	}else {
// 		fmt.Println("domain tables domain_id foreign key relationship created successfully")
// 	}
// }

////////  domain curd

func GetAllDomain(projectId, start, length int) (count int64, d []*Domain, err error) {

	o := orm.NewOrm()

	if count, err = o.QueryTable(domainTableName).Filter("project_id", projectId).Count(); err != nil {
		return 0, nil, err
	}

	if _, err = o.QueryTable(domainTableName).Filter("project_id", projectId).Limit(length, start).RelatedSel().All(&d); err != nil {
		return 0, nil, err
	}

	return
}

func GetDomainById(id int) (d *Domain, err error) {

	o := orm.NewOrm()
	d = &Domain{}
	err = o.QueryTable(domainTableName).Filter("id", id).RelatedSel().One(d)

	return
}

func GetDomainByDomain(domain *Domain) (d *Domain, err error) {

	o := orm.NewOrm()
	d = &Domain{}
	err = o.QueryTable(domainTableName).Filter("domain", domain.Domain).One(d)
	return
}

func SearchDomain(projectId, start, length int, class, platforms, services, quicken, searchText string) (total int, d []*Domain, err error) {
	base := fmt.Sprintf("SELECT `t_domain`.`id`, `t_domain`.`name`, `t_domain`.`domain`, `t_domain`.`class`, `t_domain`.`monitor`, `t_domain`.`quicken`,`t_domain`.`port`, `t_domain`.`crt`, `t_domain`.`key` FROM `%s` `t_domain`", domainTableName)

	where := "WHERE"
	if projectId != 0 {
		where = fmt.Sprintf("%s `t_domain`.`project_id` = %d", where, projectId)
	}
	if platforms != "NOT_FILTER" {
		base = fmt.Sprintf("%s INNER JOIN `%s` `t_domain_platform` ON `t_domain`.`id`=`t_domain_platform`.`%s`", base, domainPlatformTableName, domainPlatformForDomainId)
		base = fmt.Sprintf("%s INNER JOIN `%s` `t_platform` ON `t_platform`.`id`=`t_domain_platform`.`%s`", base, platformTableName, domainPlatformForPlatformId)

		where = fmt.Sprintf("%s AND `t_platform`.Name in (%s)", where, common.AddStringListSingle(platforms))
	}
	if services != "NOT_FILTER" {
		base = fmt.Sprintf("%s INNER JOIN `%s` `t_service_domain` ON `t_domain`.`id`=`t_service_domain`.`%s`", base, serviceDomainTableName, domainPlatformForDomainId)
		base = fmt.Sprintf("%s INNER JOIN `%s` `t_service` ON `t_service`.`id`=`t_service_domain`.`%s`", base, serviceTableName, serviceDomainForServiceId)

		where = fmt.Sprintf("%s AND `t_service`.`Name` in (%s)", where, common.AddStringListSingle(services))
	}
	if class != "NOT_FILTER" {
		where = fmt.Sprintf("%s AND `t_domain`.`class` in (%s)", where, common.AddStringListSingle(class))
	}
	if quicken != "NOT_FILTER" {
		where = fmt.Sprintf("%s AND `t_domain`.`quicken` in (%s)", where, common.AddStringListSingle(quicken))
	}

	if searchText != "" {
		// where = fmt.Sprintf("%s AND (`t_domain`.`name` LIKE '%%%s%%' OR `t_domain`.`domain` LIKE '%%%s%%' OR `t_service`.`name` LIKE '%%%s%%' OR `t_platform`.`name` LIKE '%%%s%%')", where, searchText, searchText, searchText, searchText)
		where = fmt.Sprintf("%s AND (`t_domain`.`name` LIKE '%%%s%%' OR `t_domain`.`domain` LIKE '%%%s%%')", where, searchText, searchText)
	}

	sql := fmt.Sprintf("%s %s", base, where)

	o := orm.NewOrm()

	var count []orm.Params
	o.Raw("SELECT count(id) FROM (" + sql + ") as temp").Values(&count)
	if len(count) > 0 {
		total = common.GetInt(count[0]["count(id)"])
	}

	_, err = o.Raw(sql+" LIMIT ?,?", start, length).QueryRows(&d)

	return
}

func GetDomainByProjectId(projectId int) (d []*Domain, err error) {

	o := orm.NewOrm()
	_, err = o.QueryTable(domainTableName).Filter("project_id", projectId).RelatedSel().All(&d)
	if err != nil {
		return nil, err
	}
	return d, err
}

func GetDomainByProjectForClass(projectId int) (list *orm.ParamsList, err error) {
	list = new(orm.ParamsList)
	o := orm.NewOrm()
	_, err = o.QueryTable(domainTableName).Filter("project_id", projectId).GroupBy("class").OrderBy("class").ValuesFlat(list, "class")
	return
}

func GetDomainByProjectForQuicken(projectId int) (d []*Domain, err error) {
	o := orm.NewOrm()
	//_ ,err = o.QueryTable(domainTableName).Filter("project_id",projectId).GroupBy("quicken").OrderBy("quicken").ValuesFlat(list, "quicken")
	_, err = o.QueryTable(domainTableName).Filter("project_id", projectId).Filter("quicken", 1).RelatedSel().All(&d)
	return
}

func GetDomainByServiceId(serviceId int) (d []*Domain, err error) {

	o := orm.NewOrm()
	_, err = o.QueryTable(domainTableName).Filter("service_id", serviceId).RelatedSel().All(&d)
	if err != nil {
		return nil, err
	}
	return d, err
}

func GetDomainMonitorAll() (m []Domain, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(domainTableName).Filter("Monitor", 1).All(&m, "Id", "class", "Domain", "Port")
	return
}

func AddDomain(d []*Domain) (id int64, err error) {

	o := orm.NewOrm()

	id, err = o.InsertMulti(len(d), d)
	if err != nil {
		return
	}
	return
}

func UpdateDomainById(d *Domain) (id int64, err error) {

	o := orm.NewOrm()

	if id, err = o.Update(d); err != nil {
		return
	}
	return
}

func DeleteDomainById(id int) (num int64, err error) {

	o := orm.NewOrm()
	if num, err = o.Delete(&Domain{Id: id}); err == nil {
		return num, nil
	}
	return -1, err
}

//////////////
/*
	多对多操作
*/

/*
	通过 domain 查询 domain 以及关联表数据
*/
func GetDomainRelated(domain *Domain) (*Domain, error) {

	o := orm.NewOrm()

	if _, err := o.LoadRelated(domain, "Services"); err != nil {
		return domain, err
	}

	if _, err := o.LoadRelated(domain, "Platforms"); err != nil {
		return domain, err
	}

	return domain, nil
}

/*
	添加 domain 以及关联表 数据
*/

func AddDomainAndRelated(domain *Domain) error {

	o := orm.NewOrm()

	// 校验 platform 数据 是否为 nil
	if len(domain.Platforms) != 0 {
		// 建立 platform 关系, 插入domain_platform表 platform 数据
		m2mForDomainAndPlatform := o.QueryM2M(domain, "Platforms")
		for _, p := range domain.Platforms {
			_, err := m2mForDomainAndPlatform.Add(p)
			if err != nil {
				fmt.Printf("error:  m2mForServiceAndHost.Add(domain.Plaftforms): %s   %s", p.Name, err)
				return err
			}
		}

	}

	// 校验 service 数据 是否为 nil
	// 在域名保存时不修改与service的关系
	//if len(domain.Services) != 0{
	// 建立 service 关系, 插入domain_service表 service 数据
	//	m2mForDomainAndService := o.QueryM2M(domain,"Services")
	//	_, err := m2mForDomainAndService.Add(domain.Services)
	//	if err != nil{
	//		fmt.Println("error:  m2mForServiceAndHost.Add(domain.Domains):  ",err)
	//		return err
	//	}
	//}

	return nil
}

/*
	更新 domain 以及关联表 数据
*/
func UpdateDomainAndRelated(domain *Domain) (err error) {

	o := orm.NewOrm()

	tempDomain := new(Domain)
	tempDomain.Id = domain.Id
	err = o.QueryTable(domainTableName).Filter("id", tempDomain.Id).RelatedSel().One(tempDomain)
	if err != nil {
		err = fmt.Errorf("error: get domain by domain id fail %s\n", err)
		return
	}

	if _, err := o.LoadRelated(tempDomain, "Platforms"); err != nil {
		return err
	}

	if _, err := o.LoadRelated(tempDomain, "Services"); err != nil {
		return err
	}

	if len(tempDomain.Services) != 0 {
		if tempDomain.Class != domain.Class {
			err = fmt.Errorf("error: The domain name is already bound to other services, please delete the domain name and service association first\n")
			return
		}
	}
	_, err = o.Update(domain)
	if err != nil {
		err = fmt.Errorf("error: update domain fail  %s\n", err)
		return
	}

	if len(tempDomain.Platforms) != 0 && len(domain.Platforms) != 0 {
		if tempDomain.Platforms[0].Id != domain.Platforms[0].Id {
			err = fmt.Errorf("error: The domain name is already bound to other services, please delete the domain name and service association first\n")
			return
		}
	}

	// 加载到domain加载的Platforms关系，直接全部清除
	m2mForDomainAndPlatforms := o.QueryM2M(domain, "Platforms")
	if _, err = m2mForDomainAndPlatforms.Clear(); err != nil {
		return
	}

	// // 加载到domain加载的Services关系，直接全部清除
	// 在域名保存时不修改与service的关系
	m2mForDomainAndServices := o.QueryM2M(domain, "Services")
	if _, err = m2mForDomainAndServices.Clear(); err != nil {
		return
	}

	// 不存在关联关系,需要建立关联关系
	if err := AddDomainAndRelated(domain); err != nil {
		fmt.Println("error:  add related ", err)
		return err
	}

	return nil
}

/*
	删除 platform 与 service 关联关系
*/
func DeletePlatformAndDomainRelatedByPlatformId(platform *Platform) (err error) {

	o := orm.NewOrm()

	// 获取 platfrom 关联的 services， 直接全部清除
	m2mForPlatformAndDomains := o.QueryM2M(platform, "Domains")
	if _, err = m2mForPlatformAndDomains.Clear(); err != nil {
		return
	}

	return
}
