package alikafka

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// InstanceVO is a nested struct in alikafka response
type InstanceVO struct {
	InstanceId               string                   `json:"InstanceId" xml:"InstanceId"`
	RegionId                 string                   `json:"RegionId" xml:"RegionId"`
	ServiceStatus            int                      `json:"ServiceStatus" xml:"ServiceStatus"`
	VpcId                    string                   `json:"VpcId" xml:"VpcId"`
	VSwitchId                string                   `json:"VSwitchId" xml:"VSwitchId"`
	EndPoint                 string                   `json:"EndPoint" xml:"EndPoint"`
	CreateTime               int64                    `json:"CreateTime" xml:"CreateTime"`
	ExpiredTime              int64                    `json:"ExpiredTime" xml:"ExpiredTime"`
	DeployType               int                      `json:"DeployType" xml:"DeployType"`
	SslEndPoint              string                   `json:"SslEndPoint" xml:"SslEndPoint"`
	Name                     string                   `json:"Name" xml:"Name"`
	IoMax                    int                      `json:"IoMax" xml:"IoMax"`
	EipMax                   int                      `json:"EipMax" xml:"EipMax"`
	DiskType                 int                      `json:"DiskType" xml:"DiskType"`
	DiskSize                 int                      `json:"DiskSize" xml:"DiskSize"`
	MsgRetain                int                      `json:"MsgRetain" xml:"MsgRetain"`
	TopicNumLimit            int                      `json:"TopicNumLimit" xml:"TopicNumLimit"`
	ZoneId                   string                   `json:"ZoneId" xml:"ZoneId"`
	PaidType                 int                      `json:"PaidType" xml:"PaidType"`
	SpecType                 string                   `json:"SpecType" xml:"SpecType"`
	SecurityGroup            string                   `json:"SecurityGroup" xml:"SecurityGroup"`
	UpgradeServiceDetailInfo UpgradeServiceDetailInfo `json:"UpgradeServiceDetailInfo" xml:"UpgradeServiceDetailInfo"`
	Tags                     TagsInGetInstanceList    `json:"Tags" xml:"Tags"`
}
