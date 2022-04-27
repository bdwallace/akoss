package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Service struct {
	Id            int    `orm:"column(id);pk;auto"`
	Name          string `orm:"column(name);size(100);null"`  // 服务名称
	Alias         string `orm:"column(alias);size(100);null"` // 服务名称
	Port          string `orm:"column(port);size(100);null"`
	ImagePath     string `orm:"cloumn(image_path);size(100);null"`
	Value         string `orm:"column(value);type(text)"` // 配置项， 以 k -> v 形式储存json
	KillTime      string `orm:"column(kill_time);size(100);"`
	ReleaseTo     string `orm:"column(release_to);size(200);null"`
	Tag           string `orm:"column(tag);size(100);null"`      // 当前tag
	LastTag       string `orm:"column(last_tag);size(100);null"` // 上次服务tag
	LogKeyword    string `orm:"column(log_keyworld);size(100);null"`
	BlackList     string `orm:"column(black_list);type(text);null"`
	DenyUserAgent string `orm:"column(deny_user_agent);type(text);null"`
	UseNacos      string `orm:"column(use_nacos);size(100);"`
	DockerNetwork string `orm:"column(docker_network);size(100);"`


	CreatedAt       time.Time `orm:"column(created_at);type(datetime);auto_now_add;"`
	UpdatedAt       time.Time `orm:"column(updated_at);type(datetime);auto_now;"`
	LastTagUpdateAt time.Time `orm:"column(last_tag_update_at);type(datetime);null"`

	DockerName string `orm:"column(docker_name);size(100);null"`
	Class      string `orm:"column(class)"`
	Health     string `orm:"column(health)"` // 服务健康监测字段
	IsDel      int    `orm:"column(is_del);default(0)"`

	Project   *Project    `orm:"rel(fk)"`
	Platforms []*Platform `orm:"rel(m2m)"`
	Confs     []*Conf     `orm:"rel(m2m)"` //  rel_through(/src/models.Conf)
	Hosts     []*Host     `orm:"rel(m2m)"`
	Domains   []*Domain   `orm:"rel(m2m)"`
}

// class:0表示java应用，1表示前端应用

type ServiceValue struct {
	HostId int    `json:"HostId"`
	Value  string `json:"Value"`
}

func init() {
	orm.RegisterModelWithPrefix("t_", new(Service))
}

/*func (s *Service) TableUnique() [][]string {
	return [][]string{
		[]string{"Project"},
	}
}*/

/*
	创建 service 中 project_id
*/
func (c *SqlClass) ServiceCreateForeignKeyProject() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", serviceTableName, serviceProjectForeignKeyName, projectId, projectTableName, primaryKey)

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

////////   多对多表外键 创建

/*
	service  <----->  conf
*/
/*
	创建 t_services_t_conf  conf外键约束
*/
func (c *SqlClass) ServiceConfCreateForeignKeyForConf() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", serviceConfTableName, serviceConfForConfForeignKeyName, serviceConfForConfId, confTableName, primaryKey)

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
	创建 t_services_t_conf   service外键约束
*/
func (c *SqlClass) ServiceConfCreateForeignKeyForService() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", serviceConfTableName, serviceConfForServiceForeignKeyName, serviceConfForServiceId, serviceTableName, primaryKey)

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
	service  <---->  domain
*/

// 创建 t_services_t_Domain  conf外键约束
func (c *SqlClass) ServiceDomainCreateForeignKeyForDomain() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", serviceDomainTableName, serviceDomainForServiceForeignKeyName, serviceDomainForDomainId, domainTableName, primaryKey)

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

// 创建 t_services_t_Domain  service外键约束
func (c *SqlClass) ServiceDomainCreateForeignKeyForService() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", serviceDomainTableName, serviceDomainForDomainForeignKeyName, serviceDomainForServiceId, serviceTableName, primaryKey)

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

//////   service  <----->  host

// 创建 t_services_t_Host  host外键约束
func (c *SqlClass) ServiceHostCreateForeignKeyForHost() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", serviceHostTableName, serviceHostForHostForeignKeyName, serviceHostForHostId, hostTableName, primaryKey)

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

// 创建 t_services_t_Host  service外键约束
func (c *SqlClass) ServiceHostCreateForeignKeyForService() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", serviceHostTableName, serviceHostForServiceForeignKeyName, serviceHostForServiceId, serviceTableName, primaryKey)

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

// 创建  t_services_t_platform   platform_id 外键约束
func (c *SqlClass) PlatformServiceCreateForeignKeyForPlatform() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", platformServicesTableName, platformServicesForPlatformForeignKeyName, platformServicesForPlatformId, platformTableName, primaryKey)

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

// 创建  t_services_t_platform   service_id 外键约束
func (c *SqlClass) PlatformServiceCreateForeignKeyForService() {
	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", platformServicesTableName, platformServicesForServiceForeignKeyName, platformServicesForServiceId, serviceTableName, primaryKey)

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

//////  service curd
/*
	查询所有 service
*/
func GetAllServices(projectId int) (s []*Service, err error) {

	o := orm.NewOrm()

	if _, err = o.QueryTable(serviceTableName).Filter("project_id", projectId).Filter("is_del", 0).All(&s); err != nil {
		return nil, err
	}
	return
}

func SearchAllServices(searchText string, projectId int) (s []*Service, err error) {

	cond := orm.NewCondition()
	// search := cond.And("project_id",projectId).And("name__icontains",searchText).Or("value__icontains",searchText).And("project_id",projectId).Or("tag__icontains",searchText).And("project_id",projectId).Or("last_tag__icontains",searchText).And("project_id",projectId).Or("class__icontains",searchText).And("project_id",projectId)
	or := cond.Or("name__icontains", searchText).Or("value__icontains", searchText).Or("tag__icontains", searchText).Or("tag__icontains", searchText).Or("tag__icontains", searchText)
	search := cond.And("project_id", projectId).AndCond(or)

	o := orm.NewOrm()
	qs := o.QueryTable(serviceTableName)

	if _, err := qs.SetCond(search).Filter("is_del", 0).RelatedSel().All(&s); err != nil {
		return nil, err
	}

	return
}

func SearchBackendServices(searchText string, projectId int) (s []*Service, err error) {

	cond := orm.NewCondition()
	// search := cond.And("project_id",projectId).And("name__icontains",searchText).Or("value__icontains",searchText).And("project_id",projectId).Or("tag__icontains",searchText).And("project_id",projectId).Or("last_tag__icontains",searchText).And("project_id",projectId).Or("class__icontains",searchText).And("project_id",projectId)
	or := cond.Or("name__icontains", searchText).Or("value__icontains", searchText).Or("tag__icontains", searchText).Or("tag__icontains", searchText).Or("tag__icontains", searchText)
	search := cond.And("project_id", projectId).AndCond(or)

	o := orm.NewOrm()
	qs := o.QueryTable(serviceTableName)

	if _, err := qs.SetCond(search).Filter("class", "java").Filter("is_del", 0).RelatedSel().All(&s); err != nil {
		return nil, err
	}

	return
}

func SearchFrontendServices(searchText string, projectId int, platformId int) (s []*Service, err error) {

	cond := orm.NewCondition()
	// search := cond.And("project_id",projectId).And("name__icontains",searchText).Or("value__icontains",searchText).And("project_id",projectId).Or("tag__icontains",searchText).And("project_id",projectId).Or("last_tag__icontains",searchText).And("project_id",projectId).Or("class__icontains",searchText).And("project_id",projectId)
	or := cond.Or("name__icontains", searchText).Or("value__icontains", searchText).Or("tag__icontains", searchText).Or("tag__icontains", searchText).Or("tag__icontains", searchText)
	search := cond.And("project_id", projectId).And("platform_id", platformId).AndCond(or)

	o := orm.NewOrm()
	qs := o.QueryTable(serviceTableName)

	if _, err := qs.SetCond(search).Exclude("class", "java").Filter("is_del", 0).RelatedSel().All(&s); err != nil {
		return nil, err
	}

	return
}

//////////////////////////////////////////////////

func GetServiceClassJava() (list *orm.ParamsList, err error) {
	list = new(orm.ParamsList)
	o := orm.NewOrm()
	_, err = o.QueryTable(serviceTableName).Filter("class", "java").Filter("is_del", 0).Filter("project_id__lt",3).GroupBy("name").OrderBy("name").ValuesFlat(list, "name")

	return
}

func GetServiceClassOther() (list *orm.ParamsList, err error) {
	list = new(orm.ParamsList)
	o := orm.NewOrm()
	_, err = o.QueryTable(serviceTableName).Exclude("class", "java").Filter("is_del", 0).Filter("project_id__lt",3).GroupBy("name").OrderBy("name").ValuesFlat(list, "name")
	return
}

//////////////////////////////////////////////////

/*
	通过 id  查询 service
*/
func GetServiceById(id int) (s *Service, err error) {

	o := orm.NewOrm()
	s = &Service{}
	err = o.QueryTable(serviceTableName).Filter("id", id).RelatedSel().One(s)

	return
}

/*
	通过 project_id  获取后端 service
*/
func GetServiceClassOfJavaByProjectId(projectId int) (s []*Service, err error) {

	o := orm.NewOrm()
	_, err = o.QueryTable(serviceTableName).Filter("project_id", projectId).Filter("class", "java").Filter("is_del", 0).RelatedSel().All(&s)
	if err != nil {
		return nil, err
	}
	return s, err
}

func GetServiceByProjectNotJavaForName(projectId int) (list *orm.ParamsList, err error) {
	list = new(orm.ParamsList)
	o := orm.NewOrm()
	sql := fmt.Sprintf("SELECT `t_service`.name FROM `%s` `t_service` INNER JOIN `%s` `t_service_t_domain` ON `t_service`.`id`=`t_service_t_domain`.`%s`", serviceTableName, serviceDomainTableName, serviceDomainForServiceId)
	_, err = o.Raw(sql).ValuesFlat(list, "name")
	return
}

/*
	通过 name 匹配 service
*/
func GetServiceByServiceName(serviceName string) (s []*Service, err error) {

	o := orm.NewOrm()
	_, err = o.QueryTable(serviceTableName).Filter("name", serviceName).Filter("is_del", 0).RelatedSel().All(&s)

	return
}

/*
	通过 class platform  匹配 service
*/
func GetServiceByClassAndPlatform(platform *Platform) (*Platform, error) {

	o := orm.NewOrm()

	if _, err := o.LoadRelated(platform, "Services"); err != nil {
		return nil, err
	}

	return platform, nil
}

func AddService(s *Service) (id int64, err error) {

	o := orm.NewOrm()

	id, err = o.Insert(s)
	if err != nil {
		return
	}
	return
}

func UpdateServiceById(s *Service) (id int64, err error) {

	o := orm.NewOrm()

	if id, err = o.Update(s); err != nil {
		return
	}
	return
}

func GetServiceRelatedForHosts(service *Service) (*Service, error) {

	o := orm.NewOrm()
	if _, err := o.LoadRelated(service, "Hosts"); err != nil {
		return service, err
	}

	return service, nil
}

//////////////
/*
	多对多操作
*/

/*
	通过 service id 查询 service 以及关联表数据
*/
func GetServiceAllRelated(service *Service) (*Service, error) {

	o := orm.NewOrm()

	if _, err := o.LoadRelated(service, "Confs"); err != nil {
		return service, err
	}

	if _, err := o.LoadRelated(service, "Hosts"); err != nil {
		return service, err
	}

	if _, err := o.LoadRelated(service, "Platforms"); err != nil {
		return service, err
	}

	if len(service.Platforms) == 1 {
		if _, err := o.LoadRelated(service.Platforms[0], "Domains"); err != nil {
			return service, err
		}
	}

	if _, err := o.LoadRelated(service, "Domains"); err != nil {
		return service, err
	}

	return service, nil
}

/*
	通过 service id 查询 platform 以及关联表数据
*/
func GetServiceAndPlatformRelated(service *Service) (*Service, error) {
	o := orm.NewOrm()

	if _, err := o.LoadRelated(service, "Platforms"); err != nil {
		return service, err
	}

	return service, nil
}

/*
	添加 service 以及关联表 数据
*/
func UpdateServiceRelated(service *Service) (err error) {

	o := orm.NewOrm()

	if len(service.Confs) != 0 {
		// 建立 host 关系, 插入service_host表 host 数据
		m2mForServiceAndConfs := o.QueryM2M(service, "Confs")
		_, err = m2mForServiceAndConfs.Add(service.Confs)
		if err != nil {
			fmt.Println("error:  m2mForServiceAndConfs.Add(service.Confs):  ", err)
			return err
		}
	}

	if len(service.Hosts) != 0 {
		// 建立 host 关系, 插入service_host表 host 数据
		m2mForServiceAndHost := o.QueryM2M(service, "Hosts")
		_, err = m2mForServiceAndHost.Add(service.Hosts)
		if err != nil {
			fmt.Println("error:  m2mForServiceAndHost.Add(service.Hosts):  ", err)
			return err
		}
	}

	if len(service.Platforms) != 0 {
		// 建立 host 关系, 插入service_host表 host 数据
		m2mForServiceAndPlatform := o.QueryM2M(service, "Platforms")
		_, err = m2mForServiceAndPlatform.Add(service.Platforms)
		if err != nil {
			fmt.Println("error:  m2mForServiceAndHost.Add(service.Plaftforms):  ", err)
			return err
		}
	}

	if len(service.Domains) != 0 {

		// 建立 host 关系, 插入service_host表 host 数据
		m2mForServiceAndDomain := o.QueryM2M(service, "Domains")
		_, err = m2mForServiceAndDomain.Add(service.Domains)
		if err != nil {
			fmt.Println("error:  m2mForServiceAndHost.Add(service.Domains):  ", err)
			return err
		}
	}

	return nil
}

/*
	更新 service 以及关联表 数据
*/
func UpdateServiceAndRelated(service *Service) (err error) {

	o := orm.NewOrm()

	_, err = o.Update(service)
	if err != nil {
		err = fmt.Errorf("insert service fail!")
		return
	}

	// 加载到service加载的hosts关系，直接全部清除
	//if len(service.Hosts) != 0{
	m2mForServiceAndHosts := o.QueryM2M(service, "Hosts")
	if _, err = m2mForServiceAndHosts.Clear(); err != nil {
		return
	}

	m2mForServiceAndConfs := o.QueryM2M(service, "Confs")
	if _, err = m2mForServiceAndConfs.Clear(); err != nil {
		return
	}

	m2mForServiceAndPlatforms := o.QueryM2M(service, "Platforms")
	if _, err = m2mForServiceAndPlatforms.Clear(); err != nil {
		return
	}

	m2mForServiceAndDomains := o.QueryM2M(service, "Domains")
	if _, err = m2mForServiceAndDomains.Clear(); err != nil {
		return
	}

	// 不存在关联关系,需要建立关联关系
	if err := UpdateServiceRelated(service); err != nil {
		fmt.Println("error:  add related ", err)
		return err
	}

	return nil
}

/*
	增加 service 以及关联表 数据
*/
func AddServiceAndRelated(service *Service) (id int64, err error) {

	o := orm.NewOrm()

	id, err = o.Insert(service)
	if err != nil {
		err = fmt.Errorf("insert service fail!")
		return
	}
	// 不存在关联关系,需要建立关联关系
	service.Id = int(id)
	if err = UpdateServiceRelated(service); err != nil {
		fmt.Println("error:  add related ", err)
		return
	}
	return
}

/*
	删除 platform 与 service 关联关系
*/
func DeleteServiceAndAllRelatedByService(service *Service) (err error) {

	o := orm.NewOrm()
	// 获取 platfrom 关联的 services， 直接全部清除

	//if len(service.Hosts) != 0 {
	//}
	m2mForServiceAndHosts := o.QueryM2M(service, "Hosts")
	if _, err = m2mForServiceAndHosts.Clear(); err != nil {
		return
	}

	m2mForServiceAndConfs := o.QueryM2M(service, "Confs")
	if _, err = m2mForServiceAndConfs.Clear(); err != nil {
		return
	}

	m2mForServiceAndPlatforms := o.QueryM2M(service, "Platforms")
	if _, err = m2mForServiceAndPlatforms.Clear(); err != nil {
		return
	}

	m2mForServiceAndDomains := o.QueryM2M(service, "Domains")
	if _, err = m2mForServiceAndDomains.Clear(); err != nil {
		return
	}

	return
}
