package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

// table
type DomainBackup struct {
	Id              int    `orm:"column(id);pk;auto"`
	DomainName      string `orm:"column(domain_name);size(100)"`
	InstanceEndTime string `orm:"column(instance_end_time);size(100)"`
	DnsServer       string `orm:"column(dns_server);size(100)"`
	DomainSource    string `orm:"column(domain_source);size(100)"`
	DomainInfo      string `orm:"column(domain_info);type(text)"`
	DomainRecord    string `orm:"column(domain_record);type(text)"`
}

type TempDomainBackup struct {
	Id              int    `db:"id"`
	DomainName      string `db:"domain_name"`
	InstanceEndTime string `db:"instance_end_time"`
	DnsServer       string `db:"dns_server"`
	DomainSource    string `db:"domain_source"`
	DomainInfo      string `db:"domain_info"`
	DomainRecord    string `db:"domain_record"`
}

type DomainBackupRecord struct {
	Value      string `json:"Value" xml:"Value"`
	TTL        int64  `json:"TTL" xml:"TTL"`
	Remark     string `json:"Remark" xml:"Remark"`
	DomainName string `json:"DomainName" xml:"DomainName"`
	RR         string `json:"RR" xml:"RR"`
	Priority   int64  `json:"Priority" xml:"Priority"`
	RecordId   string `json:"RecordId" xml:"RecordId"`
	Status     string `json:"Status" xml:"Status"`
	Locked     bool   `json:"Locked" xml:"Locked"`
	Weight     int    `json:"Weight" xml:"Weight"`
	Line       string `json:"Line" xml:"Line"`
	Type       string `json:"Type" xml:"Type"`
}

var DBStr string
var MysqlDB *sql.DB

func init() {
	orm.RegisterModelWithPrefix("t_", new(DomainBackup))
}

func ConnDB() {
	MysqlDB, err := sql.Open("mysql", DBStr)
	if err != nil {
		fmt.Printf("error:  MysqlDB.init is failed  ::  %s\n", err.Error())
		return
	}

	MysqlDB.SetMaxOpenConns(100)
	MysqlDB.SetMaxIdleConns(20)
	MysqlDB.SetConnMaxLifetime(100 * time.Second)

	if err = MysqlDB.Ping(); err != nil {
		panic("数据库连接失败： " + err.Error())
	}
	return
}

func GetBackupById(id int) (*DomainBackup, error) {

	o := orm.NewOrm()
	domainBackup := &DomainBackup{Id: id}
	if err := o.Read(domainBackup); err != nil {
		return nil, err
	}
	return domainBackup, nil
}

func GetBackupByDomainName(domainName string) (*DomainBackup, error) {
	o := orm.NewOrm()

	r := new(DomainBackup)
	if err := o.QueryTable(DomainBackupTableName).Filter("domain_name", domainName).One(r); err != nil {

		return nil, err
	}

	return r, nil
}

func GetBackupDomainByPageBreak(start, length int) ([]*DomainBackup, error) {

	o := orm.NewOrm()
	resDomainBackup := make([]*DomainBackup, 0, 60)
	_, err := o.QueryTable(DomainBackupTableName).Limit(length, start).OrderBy("id").All(&resDomainBackup)
	if err != nil {
		return nil, err
	}

	return resDomainBackup, nil
}

func AddOrUpdateBackupDomain(domainBackup *DomainBackup) (err error) {

	o := orm.NewOrm()
	r := new(DomainBackup)
	err = o.QueryTable(DomainBackupTableName).Filter("domain_name", domainBackup.DomainName).One(r)

	if r.Id > 0 {
		// update
		domainBackup.Id = r.Id
		if _, err = o.Update(domainBackup); err != nil {
			return
		}
	} else if err.Error() == "<QuerySeter> no row found" {
		// add
		_, err = o.Insert(domainBackup)
		if err != nil {
			return
		}
	} else {
		err = fmt.Errorf("error:  Add or update domain backup is failed!  domian: %s\n", domainBackup.DomainName)
		return
	}

	return nil
}

func CreateTableForBackupTheDomainWithAccKeyAndDate(accessKeyId string) (tableName string, err error) {

	date := time.Now().Format("2006-01-02")
	tableName = fmt.Sprintf("t_domain_backup_%s_%s", accessKeyId, date)

	createSql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` ("+
		"`id` int(11) NOT NULL AUTO_INCREMENT, "+
		"`domain_name` varchar(100) NOT NULL DEFAULT '', "+
		"`instance_end_time` varchar(100) NOT NULL DEFAULT '', "+
		"`dns_server` varchar(100) NOT NULL DEFAULT '', "+
		"`domain_info` longtext NOT NULL, "+
		"`domain_record` longtext NOT NULL, "+
		"`domain_source` varchar(100) NOT NULL DEFAULT '', "+
		"PRIMARY KEY (`id`)"+
		") ENGINE=InnoDB AUTO_INCREMENT=102 DEFAULT CHARSET=utf8;",
		tableName,
	)

	o := orm.NewOrm()
	_, err = o.Raw(createSql).Exec()
	if err != nil {
		return "", err
	}

	return tableName, nil
}

func TruncateBackupTable(tableName string) (err error) {

	o := orm.NewOrm()
	trunSql := fmt.Sprintf("TRUNCATE table `%s`", tableName)
	if _, err = o.Raw(trunSql).Exec(); err != nil {

		return err
	}
	return
}

func ExportAliDomainDataToBackup() (allData []*DomainBackup, err error) {

	o := orm.NewOrm()
	_, err = o.QueryTable(DomainBackupTableName).All(&allData)
	return
}

func InsertTempDomainBackup(tableName string, domain *TempDomainBackup) (err error) {
	o := orm.NewOrm()
	insertSql := fmt.Sprintf("INSERT INTO `%s` (domain_name,instance_end_time,dns_server,domain_info,domain_record,domain_source) VALUES('%s','%s','%s','%s','%s','%s');",
		tableName,
		domain.DomainName,
		domain.InstanceEndTime,
		domain.DnsServer,
		domain.DomainInfo,
		domain.DomainRecord,
		domain.DomainSource)

	_, err = o.Raw(insertSql).Exec()
	if err != nil {
		return err
	}

	return
}

func InsertTempDomainBackupFormBackup(tempTableName string) (err error) {

	o := orm.NewOrm()
	insertSql := fmt.Sprintf("INSERT INTO `%s` SELECT * FROM `%s`", tempTableName, DomainBackupTableName)
	_, err = o.Raw(insertSql).Exec()
	return
}

func CountTempDomainBackup(tableName string) (count int, err error) {
	o := orm.NewOrm()
	contSql := fmt.Sprintf("SELECT COUNT(*) FROM `%s`", tableName)
	var maps []orm.Params
	_, err = o.Raw(contSql).Values(&maps)
	if err != nil {
		return 0, err
	}
	couStr := maps[0]["COUNT(*)"].(string)
	count, err = strconv.Atoi(couStr)
	if err != nil {
		return 0, err
	}
	return
}

func DropTableBackupAccAndDate(tableName string) (err error) {
	o := orm.NewOrm()
	dropSql := fmt.Sprintf("DROP TABLE `%s`", tableName)
	_, err = o.Raw(dropSql).Exec()
	if err != nil {
		return err
	}
	return
}
