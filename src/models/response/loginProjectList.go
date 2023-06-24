package response

import (
	"time"
)

type ProjectListForLoginResponse struct {
	ProjectNameList         []string    `json:"project_name_list"`
	ProjectList				[]*Project	`json:"project_list"`
}

type Project struct {
	Id           int   		`json:"id"`
	Name         string 	`json:"name"`
	Alias        string 	`json:"alias"`
	Nacos1       string 	`json:"nacos_1"`
	Nacos2       string 	`json:"nacos_2"`
	Nacos3       string 	`json:"nacos_3"`
	AwsKeyId     string 	`json:"aws_key_id"`
	AwsKeySecret string 	`json:"aws_key_secret"`
	AwsRegion    string 	`json:"aws_region"`

	CreatedAt time.Time 	`json:"created_at"`
	UpdatedAt time.Time		`json:"updated_at"`
}