package myoss

type BaseOss struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	Bucket          string
}

func (c *BaseOss) OssInit() {
	// c.Endpoint = "oss-cn-hongkong"
	c.Endpoint = "oss-cn-hongkong.aliyuncs.com"
	c.AccessKeyId = "LTAIyG8LGgEMKHo9"
	c.AccessKeySecret = "8oDMYsnYnGYsSA2ublOAX9OD42vzen"
	c.Bucket = "ak-devops-backup"
}
