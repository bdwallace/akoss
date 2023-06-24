package aws

import (
	"context"
	"fmt"
	"library/common"
	"models"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/kafka"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/s3"
)

/*
	基础类
		包含 Region session 资源链接对象 资源监控信息等内容
*/

type BaseAws struct {
	GroupNames  []string
	GroupIds    []string
	ToPort      int
	CidrIp      string
	Description string
	//Level       int
	Project int
	Alias   string
	Region  string

	Session        *session.Session
	SourceInstance *Instance
	SourceInfo     *ResourceInfo
}

/*
	初始化session
		依据 Alias Region 信息创建aws session并储存到 基础对象的seesion中
*/
func (c *BaseAws) InitSession() (err error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		// Specify profile to load for the session's config
		Profile: c.Alias,

		// Provide SDK Config options, such as Region.
		Config: aws.Config{
			Region: aws.String(c.Region),
		},

		// Force enable Shared Config support
		SharedConfigState: session.SharedConfigEnable,
	})
	c.SourceInstance = new(Instance)
	//c.Session = new(session.Session)
	c.Session = sess

	return
}

/*
	获取所有merchant 安全组
*/
func (c *BaseAws) GetAllSgroupForMerchant() (result ec2.DescribeSecurityGroupsOutput, err error) {

	// 获取所有安全组
	groups, err := c.SourceInstance.Ec2.DescribeSecurityGroups(nil)
	if err != nil{
		return result, err
	}

	for _, group := range groups.SecurityGroups {
		groupName := *group.GroupName

		//  添加过滤 只追加 merchant 安全组
		if strings.Index(groupName, "merchant") >= 0 {
			result.SecurityGroups = append(result.SecurityGroups, group)
		}
	}

	return
}


func (c *BaseAws) GetAllSgroupForGraylog() (result ec2.DescribeSecurityGroupsOutput, err error) {

	c.Region = "ap-southeast-1"
	_ = c.InitSession()
	c.SourceInstance.Ec2 = ec2.New(c.Session)
	// 获取所有安全组
	groups, err := c.SourceInstance.Ec2.DescribeSecurityGroups(nil)
	if err != nil{
		return result, err
	}

	for _, group := range groups.SecurityGroups {
		groupName := *group.GroupName

		//  添加过滤 只追加 merchant 安全组
		if strings.Index(groupName, "Graylog") >= 0 || strings.Index(groupName, "graylog") >= 0{
			result.SecurityGroups = append(result.SecurityGroups, group)
		}
	}
	return
}

/*
	获取所有安全组
*/
func (c *BaseAws) GetAllSgroup() (result ec2.DescribeSecurityGroupsOutput, err error) {

	// 获取所有安全组
	groups, err := c.SourceInstance.Ec2.DescribeSecurityGroups(nil)
	for _, group := range groups.SecurityGroups {
		result.SecurityGroups = append(result.SecurityGroups, group)
	}

	return
}

/*
	获取指定多个安全组 groupIds
*/
func (c *BaseAws) GetSgroup() (iprangstr []*ec2.IpPermission, err error) {
	result, err := c.SourceInstance.Ec2.DescribeSecurityGroups(&ec2.DescribeSecurityGroupsInput{
		GroupIds: aws.StringSlice(c.GroupIds),
	})

	iprangstr = result.SecurityGroups[0].IpPermissions
	return

}

/*
	???
*/
func (c *BaseAws) IngressSgroup() (result *ec2.AuthorizeSecurityGroupIngressOutput, err error) {
	ToPort := int64(c.ToPort)
	input := &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: aws.String(c.GroupIds[0]),
		IpPermissions: []*ec2.IpPermission{
			{
				FromPort:   aws.Int64(ToPort),
				IpProtocol: aws.String("tcp"),
				IpRanges: []*ec2.IpRange{
					{
						CidrIp:      aws.String(c.CidrIp),
						Description: aws.String(c.Description),
					},
				},
				ToPort: aws.Int64(ToPort),
			},
		},
	}

	result, err = c.SourceInstance.Ec2.AuthorizeSecurityGroupIngress(input)
	if err != nil{
		if strings.Index(err.Error(),"The maximum number of rules") > 0{
			fmt.Printf("error: IngressSgroup.Ec2.AuthorizeSecurityGroupIngress  ::  %s\n",err.Error())
			err = fmt.Errorf( "安全组规则数量达到上限,删除部分规则再添加新规则")
		} else if strings.Index(err.Error(),"already exists") > 0 {
			fmt.Printf("error: IngressSgroup.Ec2.AuthorizeSecurityGroupIngress  ::  %s\n",err.Error())
			err = fmt.Errorf( "该安全组规则已存在")
		}else{
			fmt.Printf("error: IngressSgroup.Ec2.AuthorizeSecurityGroupIngress  ::  %s\n",err.Error())
			err = fmt.Errorf( "添加安全组失败")
		}
	}

	return
}

/*
	???
*/
func (c *BaseAws) DeleteSgroup(iprangstr []*ec2.IpPermission) (result *ec2.RevokeSecurityGroupIngressOutput, err error) {
	input := &ec2.RevokeSecurityGroupIngressInput{
		GroupId:       aws.String(c.GroupIds[0]),
		IpPermissions: iprangstr,
	}

	// result, err = svc.AuthorizeSecurityGroupIngress(input.SetIpPermissions(iprangstr))
	result, err = c.SourceInstance.Ec2.RevokeSecurityGroupIngress(input)

	return
}

/*
	???
*/
func InitCredentials() {
	homedir := common.HomeDir()
	aws_dir := fmt.Sprintf("%s/.aws", homedir)
	err := common.Mkdir(aws_dir)
	if err != nil {
		fmt.Printf("error:  新建目录%s错误:%s\n", aws_dir, err)
		return
	}

	credentialsPath := fmt.Sprintf("%s/credentials", aws_dir)

	// system, err := models.GetSystemById(1)
	projects, err := models.GetProjectAll()
	var credentials string
	for _, project := range projects {
		// if( err == nil && system.AwsKeyId != "" && system.AwsKeySecret != "" ) {
		credentials = fmt.Sprintf("%s[%s]\naws_access_key_id=%s\naws_secret_access_key=%s\n\n", credentials, project.Alias, project.AwsKeyId, project.AwsKeySecret)
		// }
	}

	//加入亚马逊云s3备份需要的key
	credentials = fmt.Sprintf("%s[%s]\naws_access_key_id=%s\naws_secret_access_key=%s\n\n", credentials, beego.AppConfig.String("backupAlias"), beego.AppConfig.String("backupAwsKeyId"), beego.AppConfig.String("backupAwsKeySecret"))

	err = common.FileWriter(credentialsPath, credentials)
	if err != nil {
		beego.Info("--------写文件出错:", credentialsPath)
	} else {
		beego.Info("-----写入文件：", credentialsPath)
	}
}

/*
   资源检测
   以下方法提供了 AWS 基础服务检测 用户数据统计等功能
*/

/*
   定义用户类
   包含 资源检测 所需字段
*/

/*
	保存各组件链接实例对象
*/
type Instance struct {
	Iam         *iam.IAM
	Vpc         *ec2.Vpc
	Ec2         *ec2.EC2
	Volume      *ec2.Volume
	Rds         *rds.RDS
	Elasticache *elasticache.ElastiCache
	ES          *elasticsearchservice.ElasticsearchService
	Kafka       *kafka.Kafka
	ELB         *elb.ELB
}

/*
	系统资源类
	保存各组件资源信息对象指针
*/
type ResourceInfo struct {
	Users               []*UserInfo
	AllSecurityGroups   *AllSecurityGroupPair
	VPCInfo             *VPCsInfo
	EC2Info             *EC2Info
	Subnets             *SubnetsInfo
	RDSCluster          *RDSInfo
	ElastiCacheClusters *ElastiCacheInfo
	ESCluster           *ESInfo
	KafkaCluster        *KafkaCluster

	//VolumeCluster  			[]*VolumeInstance
}

/*
	用户信息
*/
type UserInfo struct {
	UserARN         string
	UserID          string
	GroupID         string
	GroupARN        string
	GroupName       string
	UserAccessKeyID string

	//  需要展示字段
	UserName string
}

/*
	安全组 id name 信息
*/
type SecurityGroupPair struct {
	GroupID   string
	GroupName string
}

/*
	保存安全组信息对儿
*/
type AllSecurityGroupPair struct {
	// 可以优化为map 储存安全组信息对儿
	//AllSecurityGroupIDNamePair  map[string]string
	AllSecurityGroupIDNamePair []*SecurityGroupPair
	OwnerID                    string
}

/*
type InstanceStatus struct {
	Name 		string
	code 		string
}
*/

/*
	vpc 实例资源信息类
*/
type VPCInstanceInfo struct {

	//  需要展示字段
	ID        string
	Name      string
	Status    string
	CidrBlock string
}

/*
	系统所有vpc资源信息
*/
type VPCsInfo struct {
	OwnerID string

	VPCs []*VPCInstanceInfo
}

/*
	安全组信息 id name
*/
type SourceSecurityGroup struct {
	SecurityGroupID   string
	SecurityGroupName string
}

/*
	ec2 实例信息
*/
type EC2InstanceInfo struct {
	InstanceID string
	//VPCID      				string

	//  需要展示字段
	InstanceName   string
	PrivateIP      string
	PublicIP       string
	InstanceType   string // 实例类型  配置
	SubnetID       string
	SubnetName     string
	SecurityGroups []*SourceSecurityGroup
	Status         string
}

/*
	???
*/
type EC2Reservation struct {
	OwnerID string
	//OwnerName	string
	//  需要展示字段
	Instances []*EC2InstanceInfo
}

/*
	所有 ec2 系统资源
*/
type EC2Info struct {
	EC2Reservations []*EC2Reservation
}

/*
	子网详情
*/
type Subnet struct {
	Name      string
	ID        string
	Status    string
	CidrBlock string
}

/*
	子网信息
*/
type SubnetsInfo struct {
	Subnets []*Subnet
}

/*
type VolumeInstance struct {
	VolumeID string
	Status   string
	Size     int64
}
*/

/*
	rds 实例详情
*/
type RDSInstanceInfo struct {
	InstanceArn string

	//  需要展示字段
	InstanceIdentifier string
	EndpointAddr       string //  链接地址
	InstanceType       string // 实例类型  配置
	SecurityGroups     []*SourceSecurityGroup
	InstanceStatus     string
}

/*
	rds 资源信息
*/
type RDSInfo struct {
	UserArn      string
	RDSInstances []*RDSInstanceInfo
}

/*
	Cache 实例详情
		不一定是redis  需要教研clusterEngine
*/
type CacheInstanceInfo struct {
	ClusterEngine  string
	CacheClusterID string
	//  需要展示字段
	//   cluster
	CacheClusterName    string
	CacheClusterType    string
	CacheSecurityGroups []*SourceSecurityGroup
	CacheClusterStatus  string
}

/*
	Cache 资源信息
*/
type ElastiCacheInfo struct {
	//UserArn 				string
	//UserName 				string

	InstancesInfo []*CacheInstanceInfo
}

/*
	ES domain 详情
*/
type ESDomainInfo struct {
	VPCID    string
	DomainID string

	//  需要展示字段
	//DomainARN       		string
	EndpointsVPC          string // ES链接地址
	DomainName            string
	DomainStatus          string
	DomainType            string //  aws接口返回 []* string 多个实例类型？
	ClusterSecurityGroups []*SourceSecurityGroup
}

/*
	Elasticsearch 资源信息
*/
type ESInfo struct {

	///  未找到用户关联字段
	UserArn  string
	UserName string

	DomianNames []string
	DomainInfo  []*ESDomainInfo
}

/*
	Kafka集群详情
*/
type KafkaClusterInfo struct {

	//  需要展示字段
	//  Cluster
	ClusterArn            string
	ClusterAddr           string ///  未找到该字段
	ClusterName           string
	ClusterType           string
	ClusterSecurityGroups []*SourceSecurityGroup

	ClusterStatus string
}

/*
	所有 kafka 资源信息
*/
type KafkaCluster struct {
	UserArn  string
	UserName string
	//nodeInfo 					[]*KafkaNode

	Cluster []*KafkaClusterInfo
}

/*
	recover
		处理 runtime err
*/
func recoverCatch() {
	err := recover()
	switch err.(type) {
	case runtime.Error:
		fmt.Println("runtime error:", err)
	}
}

/*
	默认错误输出
*/
//func defaultFailOnError(err error) {
//	if aerr, ok := err.(awserr.Error); ok {
//		switch aerr.Code() {
//		default:
//			fmt.Println("22222",aerr.Error())
//		}
//	} else {
//		fmt.Println("1111",err.Error())
//	}
//}

///////  users
/*
   	获取所有 aws 用户\组信息
		传入 []UserInfo类型空切片
   		传出 传出users []UserInfo定义的属性
*/
//func (c * BaseAws)GetUsersInfo(users *[]UserInfo)(*[]UserInfo,error){
func (c *BaseAws) GetUsersInfo(users *[]UserInfo) error {

	// get IAM obj
	c.SourceInstance.Iam = iam.New(c.Session)

	// get usersInfo
	userInput := &iam.ListUsersInput{}
	userRet, err := c.SourceInstance.Iam.ListUsers(userInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return err
	}

	// get groupsInfo
	groupInput := &iam.ListGroupsInput{}
	groupRet, err := c.SourceInstance.Iam.ListGroups(groupInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return err
	}

	var user UserInfo

	//  获取userInfo  groupInfo 到 SourceInfo.Users
	for _, u := range userRet.Users {

		user.UserARN = *u.Arn
		user.UserID = *u.UserId
		user.UserName = *u.UserName

		*users = append(*users, user)
	}

	for i, g := range groupRet.Groups {

		(*users)[i].GroupARN = *g.Arn
		(*users)[i].GroupID = *g.GroupId
		(*users)[i].GroupName = *g.GroupName
	}

	//  test  output
	if users == nil {
		fmt.Println("aws.go  func GetUsersInfo users is nil")
		err := fmt.Errorf("aws.go  func GetUsersInfo users is nil")
		return err
	}
	/*
		for _, v := range *users {
			fmt.Println("userId:", v.UserID, "  userName:", v.UserName, "   groupName:", v.GroupName)
		}
	*/
	return nil
}

/////// securityGroups

/*
	获取所有安全组
		传出 保存在 obj.SourceInfo.AllSecurityGroups
*/
func (c *BaseAws) GetAllSecurityGourpSave2Obj() {

	allSecurityGroups, err := c.GetAllSgroup()
	if err != nil {
		fmt.Println("error: GetAllSecurityGourpSave2Obj,  ", err)
		return
	}

	//c.SourceInfo = new(ResourceInfo)
	c.SourceInfo.AllSecurityGroups = new(AllSecurityGroupPair)

	for _, v := range allSecurityGroups.SecurityGroups {
		tempSecurityPair := new(SecurityGroupPair)

		tempSecurityPair.GroupName = *v.GroupName
		tempSecurityPair.GroupID = *v.GroupId

		c.SourceInfo.AllSecurityGroups.AllSecurityGroupIDNamePair = append(c.SourceInfo.AllSecurityGroups.AllSecurityGroupIDNamePair, tempSecurityPair)
	}
}

///////  VPC

/*
	获取所有vpc资源信息并解析
		传出 obj.SourceInfo.vpc...
*/

func (c *BaseAws) GetVPCAndAnalytical() error {

	//  获取所有vpc信息
	vpcInfo, err := c.GetVPCInfo()
	if err != nil {
		fmt.Println("error:  GetVPCAndAnalytical  ", err)
		return err
	}

	//  解析信息 并 存储  obj.SourceInfo.vpc...
	c.VPCAnalyticalData(vpcInfo)

	/*
		// out put
			for i, v := range c.SourceInfo.VPCInfo.VPCs{
				//fmt.Println("vpc_",i,"： vpcid: ",v,"     ownerid: ",v.OwnerID,"    status: ",v.Status)
				fmt.Println("vpc_",i)
				fmt.Println(v)
			}
	*/

	return err
}

/*
	基本调用流程
		...

*/

/*
	解析vpc详情
		传入 vpcInfo 指针
		传出 将解析的数据存放 obj.SourceInfo.vpc...
*/

func (c *BaseAws) VPCAnalyticalData(vpcInfo *ec2.DescribeVpcsOutput) {

	//  recover
	defer recoverCatch()

	//c.SourceInfo = new(ResourceInfo)
	c.SourceInfo.VPCInfo = new(VPCsInfo)

	for _, v := range vpcInfo.Vpcs {
		tempVpc := new(VPCInstanceInfo)

		tempVpc.ID = *v.VpcId
		tempVpc.Status = *v.State
		tempVpc.CidrBlock = *v.CidrBlock

		for _, n := range v.Tags {
			if 0 == strings.Compare(*n.Key, "Name") {
				tempVpc.Name = *n.Value
			}
		}
		c.SourceInfo.VPCInfo.VPCs = append(c.SourceInfo.VPCInfo.VPCs, tempVpc)
	}

}

/*
   	获取所有 VPCs id status
		传出 所有vpc资源信息
*/
func (c *BaseAws) GetVPCInfo() (*ec2.DescribeVpcsOutput, error) {

	// get aws session
	c.SourceInstance.Ec2 = ec2.New(c.Session)

	// get VPCs info
	input := &ec2.DescribeVpcsInput{}
	//  获取所有vpcInfo
	result, err := c.SourceInstance.Ec2.DescribeVpcs(input)

	if err != nil {
		fmt.Println("error: Ec2.DescribeVpcs  ", err)
		return nil, err
	}
	return result, err
}

////////  EC2

/*
	匹配subnetId -> subnetName
*/
func (c *BaseAws) GetEc2CollocationSubnet() error {

	// get
	err := c.GetEC2AndAnalytical()
	if err != nil {
		newErr := fmt.Errorf("获取ec2Info失败： ")
		return newErr
	}

	err = c.GetSubnetAndAnalytical()
	if err != nil {
		newErr := fmt.Errorf("获取EC2 subnet失败")
		return newErr
	}

	// 匹配 subnetId -> subnetName
	subnetName := make(map[string]string, 0)
	for _, s := range c.SourceInfo.Subnets.Subnets {

		id := s.ID
		name := s.Name
		subnetName[id] = name
	}

	for _, s := range c.SourceInfo.EC2Info.EC2Reservations {
		for _, v := range s.Instances {

			id := v.SubnetID

			if _, ok := subnetName[id]; ok {
				v.SubnetName = subnetName[id]
			}
		}
	}

	return nil
}

/*
	获取所有ec2实例信息， 解析ec2资源详情
		输出 保存 obj.Sourceinfo.ec2...
*/
func (c *BaseAws) GetEC2AndAnalytical() error {

	// get ec2 info
	ec2Info, err := c.GetEC2sInfo()
	if err != nil {
		fmt.Println("error:  GetEC2AndAnalytical  c.GetEC2sInfo  ", err)
		return err
	}

	// analytical data
	c.EC2AnalyticalData(ec2Info)

	return nil
}

/*
  	解析 EC2sInstanceInfo
		输出 保存 obj.Sourceinfo.ec2...
*/
func (c *BaseAws) EC2AnalyticalData(ec2Info *ec2.DescribeInstancesOutput) {

	//  注册recover
	defer recoverCatch()

	//c.SourceInfo = new(ResourceInfo)
	c.SourceInfo.EC2Info = new(EC2Info)

	// 解析
	for i, v := range ec2Info.Reservations {

		tempReser := new(EC2Reservation)

		tempReser.OwnerID = *v.OwnerId

		c.SourceInfo.EC2Info.EC2Reservations = append(c.SourceInfo.EC2Info.EC2Reservations, tempReser)
		for j, d := range v.Instances {

			tempInstance := new(EC2InstanceInfo)

			if d.InstanceId != nil {
				tempInstance.InstanceID = *d.InstanceId
			}
			if d.PrivateIpAddress != nil {
				tempInstance.PrivateIP = *d.PrivateIpAddress
			}
			if d.PublicIpAddress != nil {
				tempInstance.PublicIP = *d.PublicIpAddress
			}
			if d.InstanceType != nil {
				tempInstance.InstanceType = *d.InstanceType
			}
			if d.SubnetId != nil {
				tempInstance.SubnetID = *d.SubnetId
			}
			if d.State.Name != nil {
				tempInstance.Status = *(*d.State).Name
			}

			for _, name := range d.Tags {
				if 0 == strings.Compare(*name.Key, "Name") {
					tempInstance.InstanceName = *name.Value
				}
			}

			c.SourceInfo.EC2Info.EC2Reservations[i].Instances = append(c.SourceInfo.EC2Info.EC2Reservations[i].Instances, tempInstance)

			for _, sGroup := range d.SecurityGroups {
				tempSecurityGroups := new(SourceSecurityGroup)

				tempSecurityGroups.SecurityGroupID = *sGroup.GroupId
				tempSecurityGroups.SecurityGroupName = *sGroup.GroupName

				c.SourceInfo.EC2Info.EC2Reservations[i].Instances[j].SecurityGroups = append(c.SourceInfo.EC2Info.EC2Reservations[i].Instances[j].SecurityGroups, tempSecurityGroups)
			}
		}
	}

}

/*
	获取所有 EC2 instance
		输出	 保存 obj.SourceInfo.ec2...
*/

func (c *BaseAws) GetEC2sInfo() (*ec2.DescribeInstancesOutput, error) {

	c.SourceInstance.Ec2 = ec2.New(c.Session)

	// get all ec2 instance info
	input := &ec2.DescribeInstancesInput{}
	result, err := c.SourceInstance.Ec2.DescribeInstances(input)
	if err != nil {
		fmt.Println("error: Ec2.DescribeInstances")
		return nil, err
	}

	return result, nil
}

//////// Subnet
/*
	获取subnet信息并解析
		获取所有sebnet信息
		传出 解析并保存 obj.SourceInfo.subnet...
*/
func (c *BaseAws) GetSubnetAndAnalytical() error {

	res, err := c.GetSubnetInfo()
	if err != nil {
		fmt.Println("error: GetSubnetAndAnalytical  ", err)
		return err
	}

	//  analytical data
	c.SubnetAnalyticalData(res)
	/*
		for i, v := range c.SourceInfo.Subnets.Subnets{
			fmt.Println("sub_",i,"   name:",v.Name,"  ID:",v.ID,"   status:",v.Status,"    cidr:",v.CidrBlock)
		}
	*/
	return nil
}

/*
	解析subnet详情
		传出 解析并保存 obj.SourceInfo.subnet...
*/
func (c *BaseAws) SubnetAnalyticalData(SubnetInfo *ec2.DescribeSubnetsOutput) {

	//  注册recover
	defer recoverCatch()

	//c.SourceInfo = new(ResourceInfo)
	c.SourceInfo.Subnets = new(SubnetsInfo)

	for _, v := range SubnetInfo.Subnets {
		tempSubnet := new(Subnet)

		tempSubnet.ID = *v.SubnetId
		tempSubnet.CidrBlock = *v.CidrBlock
		tempSubnet.Status = *v.State
		for _, n := range v.Tags {
			if 0 == strings.Compare(*n.Key, "Name") {
				tempSubnet.Name = *n.Value
			}
		}

		c.SourceInfo.Subnets.Subnets = append(c.SourceInfo.Subnets.Subnets, tempSubnet)
	}
	/*
		for i, v := range c.SourceInfo.Subnets.Subnets{
			fmt.Println("sub_",i,"   name:",v.Name,"  ID:",v.ID,"   status:",v.Status,"    cidr:",v.CidrBlock)
		}
	*/
}

/*
	获取所有subnet Info
		输出 subnetInfo point
*/
func (c *BaseAws) GetSubnetInfo() (*ec2.DescribeSubnetsOutput, error) {

	c.SourceInstance.Ec2 = ec2.New(c.Session)

	// 获取所有 subnetInfo
	intput := &ec2.DescribeSubnetsInput{}
	result, err := c.SourceInstance.Ec2.DescribeSubnets(intput)
	if err != nil {
		fmt.Println("error: Ec2.DescribeSubnets  ", err)
		return nil, err
	}

	return result, nil
}

/////// volume

/*
   获取所有 volume info
*/
func (c *BaseAws) GetVolumeInfo(ec2s *[]EC2Info) error {

	/*
		...
	*/

	return nil
}

///////  RDS
/*
	RDS err输出
*/
func RDSFailOnError(err error) {
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case rds.ErrCodeDBInstanceNotFoundFault:
			fmt.Println(rds.ErrCodeDBInstanceNotFoundFault, aerr.Error())
		default:
			fmt.Println(aerr.Error())
		}
	} else {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
	}
}

/*
	获取RDS信息并解析
		获取所有rds信息
		传出 解析并保存 obj.SourceInfo.rds...
*/

func (c *BaseAws) GetRDSAndAnalytical() error {

	res, err := c.GetRDSInfo()
	if err != nil {
		fmt.Println("error:GetRDSAndAnalytical   ", err)
		return err
	}

	c.SourceInstance.Ec2 = ec2.New(c.Session)

	c.GetAllSecurityGourpSave2Obj()

	// analytical data
	c.RDSAnalyticalData(res)

	return nil
}

/*
	解析RDS详情
		传出 解析并保存 obj.SourceInfo.RDS...
*/
func (c *BaseAws) RDSAnalyticalData(RDSRes *rds.DescribeDBInstancesOutput) {

	//  注册recover
	defer recoverCatch()

	//c.SourceInfo = new(ResourceInfo)
	c.SourceInfo.RDSCluster = new(RDSInfo)

	// analytical data
	for i, v := range RDSRes.DBInstances {
		tempRDSInstance := new(RDSInstanceInfo)

		tempRDSInstance.InstanceIdentifier = *v.DBInstanceIdentifier
		tempRDSInstance.EndpointAddr = *v.Endpoint.Address
		tempRDSInstance.InstanceType = *v.DBInstanceClass
		tempRDSInstance.InstanceStatus = *v.DBInstanceStatus

		c.SourceInfo.RDSCluster.RDSInstances = append(c.SourceInfo.RDSCluster.RDSInstances, tempRDSInstance)

		for _, s := range v.VpcSecurityGroups {
			tempSecurityGroup := new(SourceSecurityGroup)

			tempSecurityGroup.SecurityGroupID = *s.VpcSecurityGroupId
			//  根据 安全组id 查询匹配安全组名称
			for _, allsg := range c.SourceInfo.AllSecurityGroups.AllSecurityGroupIDNamePair {
				if 0 == strings.Compare(tempSecurityGroup.SecurityGroupID, allsg.GroupID) {
					tempSecurityGroup.SecurityGroupName = allsg.GroupName
				}
			}
			c.SourceInfo.RDSCluster.RDSInstances[i].SecurityGroups = append(c.SourceInfo.RDSCluster.RDSInstances[i].SecurityGroups, tempSecurityGroup)
		}
	}

}

/*
	获取所有RDS Info
		输出 RDSInfo point
*/
//func (c *BaseAws) GetRDSInfo(RDS *[]RDSInstance) error {
func (c *BaseAws) GetRDSInfo() (*rds.DescribeDBInstancesOutput, error) {

	c.SourceInstance = new(Instance)
	c.SourceInstance.Rds = rds.New(c.Session)

	/*
		// get rds cluster
		clusterInput := &rds.DescribeDBClustersInput{}
		clusterRes, cErr := c.SourceInstance.Rds.DescribeDBClusters(clusterInput)
		if cErr != nil {
			if aerr, ok := cErr.(awserr.Error); ok {
				switch aerr.Code() {
				case rds.ErrCodeDBClusterNotFoundFault:
					fmt.Println(rds.ErrCodeDBClusterNotFoundFault, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(cErr.Error())
			}
			return cErr
		}
	*/

	// get rds instance info
	instanceInput := &rds.DescribeDBInstancesInput{}
	instanceRes, err := c.SourceInstance.Rds.DescribeDBInstances(instanceInput)
	if err != nil {
		RDSFailOnError(err)
		return nil, err
	}

	if instanceRes == nil {
		err := fmt.Errorf("err:   clusterRes or instanceRes is nil")
		fmt.Println(err)
		return nil, err
	}

	return instanceRes, nil
}

///////  ElastiCache
/*
	获取所有ElastiCache信息并解析
		获取所有ElastiCache信息
		传出 解析并保存 obj.SourceInfo.ElastiCache...
*/

func (c *BaseAws) GetElastiCacheAndAnalytical() error {

	desCacheCluster, err := c.GetElastiCacheInfo()

	if err != nil {
		fmt.Println("error: GetElastiCacheInfo  ", err)
		return err
	}

	//  获取所有安全组
	c.SourceInstance.Ec2 = ec2.New(c.Session)
	c.GetAllSecurityGourpSave2Obj()

	c.ElastiCacheAnalyticalData(desCacheCluster)
	/*
		for _, v := range c.SourceInfo.ElastiCacheClusters.InstancesInfo{
			fmt.Println("name: ",v.CacheClusterName,"    engine:",v.ClusterEngine,"    status:",v.CacheClusterStatus,"     type:",v.CacheClusterType)
			for _, s := range v.CacheSecurityGroups{
				fmt.Println("       sgroup id: ",s.SecurityGroupID,"   name:",s.SecurityGroupName)
			}
		}

	*/
	return nil
}

/*
	ElastiCache err 输出
*/
func ElastiCacheFailOnError(err error) {

	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case elasticache.ErrCodeCacheClusterNotFoundFault:
			fmt.Println(elasticache.ErrCodeCacheClusterNotFoundFault, aerr.Error())
		case elasticache.ErrCodeInvalidParameterValueException:
			fmt.Println(elasticache.ErrCodeInvalidParameterValueException, aerr.Error())
		case elasticache.ErrCodeInvalidParameterCombinationException:
			fmt.Println(elasticache.ErrCodeInvalidParameterCombinationException, aerr.Error())
		default:
			fmt.Println(aerr.Error())
		}
	} else {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
	}
}

/*
	解析 ElastiCache详情
		传出 解析并保存 obj.SourceInfo.ElastiCache...
*/
func (c *BaseAws) ElastiCacheAnalyticalData(desElestiCache *elasticache.DescribeCacheClustersOutput) {

	//  注册recover
	defer recoverCatch()

	//c.SourceInfo = new(ResourceInfo)
	c.SourceInfo.ElastiCacheClusters = new(ElastiCacheInfo)

	/*
		   	拿到的是每个节点的信息
			需要在controller中添加相同集群多节点校验
			组合成为一个集群
	*/
	for i, v := range desElestiCache.CacheClusters {
		tempCacheCluster := new(CacheInstanceInfo)

		tempCacheCluster.ClusterEngine = *v.Engine
		//  后续添加 engine 校验

		tempCacheCluster.CacheClusterName = *v.ReplicationGroupId
		tempCacheCluster.CacheClusterType = *v.CacheNodeType
		tempCacheCluster.CacheClusterStatus = *v.CacheClusterStatus

		c.SourceInfo.ElastiCacheClusters.InstancesInfo = append(c.SourceInfo.ElastiCacheClusters.InstancesInfo, tempCacheCluster)

		// security group
		for _, s := range v.SecurityGroups {
			tempSecurityGroup := new(SourceSecurityGroup)

			//  没有 securtiy group name 字段   只有id 和 status
			tempSecurityGroup.SecurityGroupID = *s.SecurityGroupId

			for _, allsg := range c.SourceInfo.AllSecurityGroups.AllSecurityGroupIDNamePair {
				if 0 == strings.Compare(tempSecurityGroup.SecurityGroupID, allsg.GroupID) {
					tempSecurityGroup.SecurityGroupName = allsg.GroupName
				}
			}

			c.SourceInfo.ElastiCacheClusters.InstancesInfo[i].CacheSecurityGroups = append(c.SourceInfo.ElastiCacheClusters.InstancesInfo[i].CacheSecurityGroups, tempSecurityGroup)
		}
	}

}

/*
	获取所有ElastiCache Info
		输出 ElastiCacheInfo point
*/
func (c *BaseAws) GetElastiCacheInfo() (*elasticache.DescribeCacheClustersOutput, error) {

	c.SourceInstance.Elasticache = elasticache.New(c.Session)

	input := &elasticache.DescribeCacheClustersInput{}
	desCacheCluster, err := c.SourceInstance.Elasticache.DescribeCacheClusters(input)
	if err != nil {
		ElastiCacheFailOnError(err)
		return nil, err
	}

	if desCacheCluster == nil {
		err = fmt.Errorf("err:  result is nil")
		fmt.Println("error: GetElastCacheInfo   ", err)
		return nil, err
	}

	return desCacheCluster, nil

}

///// ES
/*
	获取所有ES信息并解析
		获取所有ES信息
		传出 解析并保存 obj.SourceInfo.ES...
*/

func (c *BaseAws) GetESAndAnalytical() error {

	domainRes, domainStatus, err := c.GetESInfo()
	if err != nil {
		fmt.Println("error: GetEsInfo  ", err)
		return err
	}

	// 获取所有安全组
	c.SourceInstance.Ec2 = ec2.New(c.Session)
	c.GetAllSecurityGourpSave2Obj()

	//  analytical data
	c.ESAnalyticalData(domainRes, domainStatus)
	/*
		// out put
		for _, v := range c.SourceInfo.ESCluster.DomainInfo{
			fmt.Println("ES name:",v.DomainName,"  type:",v.DomainType," endponit:",v.EndpointsVPC)
			for _, s := range v.ClusterSecurityGroups{
				fmt.Println("     sgroup:",s.SecurityGroupID,"   name:",s.SecurityGroupName)
			}
		}
	*/
	return nil
}

/*
	解析 ES 详情
		传出 解析并保存 obj.SourceInfo.ES...
*/
func (c *BaseAws) ESAnalyticalData(decESDomain *elasticsearchservice.DescribeElasticsearchDomainsOutput, desDomainStatus *elasticsearchservice.DescribeReservedElasticsearchInstancesOutput) {

	//  注册recover
	defer recoverCatch()

	// new SourceInfo 多层 obj
	//c.SourceInfo = new(ResourceInfo)
	c.SourceInfo.ESCluster = new(ESInfo)

	//  需要修改  返回 err
	if c.SourceInfo.ESCluster == nil || c.SourceInfo == nil {
		fmt.Println("  c.SourceInfo.ESCluster == nil")
		return
	}

	//  创建配置对象指针获取数据，	将指针append SourceInfo
	for _, v := range decESDomain.DomainStatusList {
		tempDomain := new(ESDomainInfo)

		// get ES info
		tempDomain.DomainName = *v.DomainName
		tempDomain.DomainID = *v.DomainId
		tempDomain.DomainType = *v.ElasticsearchClusterConfig.InstanceType
		tempDomain.EndpointsVPC = *v.Endpoints["vpc"]

		c.SourceInfo.ESCluster.DomainInfo = append(c.SourceInfo.ESCluster.DomainInfo, tempDomain)
		// get SecrityGroups
		for j, sGroup := range v.VPCOptions.SecurityGroupIds {
			tempSecurityGroup := new(SourceSecurityGroup)

			tempSecurityGroup.SecurityGroupID = *sGroup

			for _, allsg := range c.SourceInfo.AllSecurityGroups.AllSecurityGroupIDNamePair {
				if 0 == strings.Compare(tempSecurityGroup.SecurityGroupID, allsg.GroupID) {
					tempSecurityGroup.SecurityGroupName = allsg.GroupName
				}
			}

			c.SourceInfo.ESCluster.DomainInfo[j].ClusterSecurityGroups = append(c.SourceInfo.ESCluster.DomainInfo[j].ClusterSecurityGroups, tempSecurityGroup)
		}

	}

}

/*
	获取所有 ES Info
		输出 ES Info point
*/
func (c *BaseAws) GetESInfo() (*elasticsearchservice.DescribeElasticsearchDomainsOutput, *elasticsearchservice.DescribeReservedElasticsearchInstancesOutput, error) {

	c.SourceInstance.ES = elasticsearchservice.New(c.Session)

	// 获取所有es DomainName
	input := &elasticsearchservice.ListDomainNamesInput{}
	domainNamesRes, err := c.SourceInstance.ES.ListDomainNames(input)
	if err != nil {
		fmt.Println("error: ES.ListDomainNames  ", err)
		return nil, nil, err
	}

	if domainNamesRes == nil {
		err = fmt.Errorf("error: GetESInfo  result is nil")
		fmt.Println(err)
		return nil, nil, err
	}

	names := make([]*string, 0)
	for _, v := range domainNamesRes.DomainNames {
		tempName := new(string)
		tempName = v.DomainName
		names = append(names, tempName)
	}

	//  获取ES集群信息 传入DomainName
	desDomainInput := elasticsearchservice.DescribeElasticsearchDomainsInput{
		DomainNames: names,
	}

	desDomansOutPut, err := c.SourceInstance.ES.DescribeElasticsearchDomains(&desDomainInput)
	if err != nil {
		fmt.Println("error: ES.DescribeElasticsearchDomains  ", err)
		return nil, nil, err
	}

	if desDomansOutPut == nil {
		err = fmt.Errorf("error: GetESInfo desDomansOutPut  is nil")
		fmt.Println(err)
		return nil, nil, err
	}

	//  get ES cluster instance info
	domainStatusInput := &elasticsearchservice.DescribeReservedElasticsearchInstancesInput{}
	desdomainStatusRes, err := c.SourceInstance.ES.DescribeReservedElasticsearchInstances(domainStatusInput)
	if err != nil {
		fmt.Println("error: ES.DescribeReservedElasticsearchInstances  ", err)
		return nil, nil, err
	}

	return desDomansOutPut, desdomainStatusRes, nil
}

///////////// kafka
/*
	获取所有 Kafka 信息并解析
		获取所有 Kafka 信息
		传出 解析并保存 obj.SourceInfo.Kafka...
*/
func (c *BaseAws) GetKafkaAndAnalytical() error {

	listCluster, err := c.GetKafkaInfo()
	//_, err := baseAws.GetKafkaInfo()
	if err != nil {
		fmt.Println("error:  GetKafkaInfo  ", err)
		return err
	}

	// get all SecurityGroups
	c.SourceInstance.Ec2 = ec2.New(c.Session)
	c.GetAllSecurityGourpSave2Obj()

	// analyrical data
	c.KafkaAnalyticalData(listCluster)

	/*
		// out put
		for _, v := range c.SourceInfo.KafkaCluster.Cluster{
			fmt.Println("Arn:",v.ClusterArn,"        Name:",v.ClusterName,"       status:",v.ClusterStatus,"        type",v.ClusterType)
			for _, s := range v.ClusterSecurityGroups{
				fmt.Println("      sgroup:",s.SecurityGroupID)
			}
		}
	*/

	return nil
}

/*
	解析 Kafka 详情
		传出 解析并保存 obj.SourceInfo.Kafka...
*/

func (c *BaseAws) KafkaAnalyticalData(listCluster *kafka.ListClustersOutput) {

	//  注册recover
	defer recoverCatch()

	//c.SourceInfo = new(ResourceInfo)
	c.SourceInfo.KafkaCluster = new(KafkaCluster)

	if c.SourceInfo.KafkaCluster == nil || c.SourceInfo == nil {
		fmt.Println("error: KafkaAnalyticalData  c.SourceInfo.KafkaCluster == nil || c.SourceInfo == nil ")
		return
	}

	// analytical data
	for _, v := range listCluster.ClusterInfoList {
		tempCluster := new(KafkaClusterInfo)

		// get base info
		tempCluster.ClusterName = *v.ClusterName
		tempCluster.ClusterArn = *v.ClusterArn
		tempCluster.ClusterType = *v.BrokerNodeGroupInfo.InstanceType
		tempCluster.ClusterStatus = *v.State

		c.SourceInfo.KafkaCluster.Cluster = append(c.SourceInfo.KafkaCluster.Cluster, tempCluster)

		// get SecrityGroups
		for j, sGroupID := range v.BrokerNodeGroupInfo.SecurityGroups {
			tempSecurityGroup := new(SourceSecurityGroup)
			tempSecurityGroup.SecurityGroupID = *sGroupID

			// get securityGroupName
			for _, allsg := range c.SourceInfo.AllSecurityGroups.AllSecurityGroupIDNamePair {
				if 0 == strings.Compare(tempSecurityGroup.SecurityGroupID, allsg.GroupID) {
					tempSecurityGroup.SecurityGroupName = allsg.GroupName
				}
			}

			c.SourceInfo.KafkaCluster.Cluster[j].ClusterSecurityGroups = append(c.SourceInfo.KafkaCluster.Cluster[j].ClusterSecurityGroups, tempSecurityGroup)
		}
	}
}

/*
	获取所有 Kafka Info
		输出 KafkaInfo point
*/
func (c *BaseAws) GetKafkaInfo() (*kafka.ListClustersOutput, error) {

	c.SourceInstance.Kafka = kafka.New(c.Session)

	listClustersInput := &kafka.ListClustersInput{}
	listClustersRes, lErr := c.SourceInstance.Kafka.ListClusters(listClustersInput)
	if lErr != nil {
		fmt.Println("error: Kafka.ListClusters ", lErr)
		return nil, lErr
	}

	if listClustersRes == nil {
		err := fmt.Errorf("err:  listClustersRes is nil")
		fmt.Println(err)
		return nil, err
	}

	return listClustersRes, nil
}

/////  ELB

func (c *BaseAws) GetELBAndAnalytical() error {

	return nil
}

func (c *BaseAws) ELBAnalyticalData(listCluster *kafka.ListClustersOutput) {

}

func (c *BaseAws) GetELBInfo() (*elb.DescribeLoadBalancersOutput, error) {

	c.SourceInstance.ELB = elb.New(c.Session)

	fmt.Println("describeLoadBalancers")
	desELBRes, err := c.SourceInstance.ELB.DescribeLoadBalancers(nil)
	if err != nil {
		fmt.Println("error: ELB.DescribeLoadBalancers  ", err)
		return nil, err
	}

	if desELBRes == nil {
		err := fmt.Errorf("error: GetELBInfo listClustersRes is nil")
		fmt.Println(err)
		return nil, err
	}

	// ???  LoadBalancer is nil
	fmt.Println(desELBRes.LoadBalancerDescriptions)

	return desELBRes, nil
}

func (c *BaseAws) GetAllRegion() (awsregion []string, err error) {
	// return []string{"us-west-1", "ap-east-1", "us-east-2", "us-east-1", "us-west-2", "ap-southeast-1"}, nil

	awsregion = strings.Split(beego.AppConfig.String("awsregion"), ",")
	return
}

/*
	获取所有 单个EC2 instance 的状态
		输出	 保存 obj.SourceInfo.ec2...
*/
func (c *BaseAws) GetEC2Status(instanceId string) (*ec2.DescribeInstancesOutput, error) {

	c.SourceInstance.Ec2 = ec2.New(c.Session)

	// get all ec2 instance info
	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceId),
		},
	}
	result, err := c.SourceInstance.Ec2.DescribeInstances(input)
	if err != nil {
		fmt.Println("error: Ec2.DescribeInstances  ", err)
		return nil, err
	}

	return result, nil
}

/*
	停止 单个EC2 instance
		输出	 保存 obj.SourceInfo.ec2...
*/
func (c *BaseAws) StopEC2Status(instanceId string) (*ec2.StopInstancesOutput, error) {
	c.SourceInstance.Ec2 = ec2.New(c.Session)
	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceId),
		},
	}

	result, err := c.SourceInstance.Ec2.StopInstances(input)
	if err != nil {
		if aErr, ok := err.(awserr.Error); ok {
			switch aErr.Code() {
			default:
				fmt.Println(aErr.Error())
			}
		} else {

			fmt.Println(err.Error())
		}
		return nil, err
	}

	return result, err
}

/*
	启动 单个EC2 instance
		输出	 保存 obj.SourceInfo.ec2...
*/
func (c *BaseAws) StartEC2Status(instanceId string) (*ec2.StartInstancesOutput, error) {
	c.SourceInstance.Ec2 = ec2.New(c.Session)
	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceId),
		},
	}

	result, err := c.SourceInstance.Ec2.StartInstances(input)
	if err != nil {
		if aErr, ok := err.(awserr.Error); ok {
			switch aErr.Code() {
			default:
				fmt.Println(aErr.Error())
			}
		} else {

			fmt.Println(err.Error())
		}
		return nil, err
	}

	return result, err
}

func (c *BaseAws) UploadS3(path, s3Path string) (err error) {

	service := s3.New(c.Session)

	fp, err := os.Open(path)
	if err != nil {
		return
	}
	defer fp.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()

	_, err = service.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String("akgo"),
		Key:    aws.String(s3Path),
		Body:   fp,
	})

	return
}
