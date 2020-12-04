package cassandra

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

// Cluster is a nested struct in cassandra response
type Cluster struct {
	ClusterId         string                `json:"ClusterId" xml:"ClusterId"`
	PayType           string                `json:"PayType" xml:"PayType"`
	MaintainStartTime string                `json:"MaintainStartTime" xml:"MaintainStartTime"`
	AutoRenewPeriod   int                   `json:"AutoRenewPeriod" xml:"AutoRenewPeriod"`
	MajorVersion      string                `json:"MajorVersion" xml:"MajorVersion"`
	CreatedTime       string                `json:"CreatedTime" xml:"CreatedTime"`
	IsLatestVersion   bool                  `json:"IsLatestVersion" xml:"IsLatestVersion"`
	AutoRenewal       bool                  `json:"AutoRenewal" xml:"AutoRenewal"`
	MinorVersion      string                `json:"MinorVersion" xml:"MinorVersion"`
	MaintainEndTime   string                `json:"MaintainEndTime" xml:"MaintainEndTime"`
	ExpireTime        string                `json:"ExpireTime" xml:"ExpireTime"`
	ClusterName       string                `json:"ClusterName" xml:"ClusterName"`
	DataCenterCount   int                   `json:"DataCenterCount" xml:"DataCenterCount"`
	Status            string                `json:"Status" xml:"Status"`
	LockMode          string                `json:"LockMode" xml:"LockMode"`
	Tags              TagsInDescribeCluster `json:"Tags" xml:"Tags"`
}
