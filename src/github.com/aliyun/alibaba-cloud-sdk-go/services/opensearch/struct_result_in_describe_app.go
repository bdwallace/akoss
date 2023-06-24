package opensearch

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

// ResultInDescribeApp is a nested struct in opensearch response
type ResultInDescribeApp struct {
	Id               string                 `json:"id" xml:"id"`
	Description      string                 `json:"description" xml:"description"`
	Status           string                 `json:"status" xml:"status"`
	Type             string                 `json:"type" xml:"type"`
	ClusterName      string                 `json:"clusterName" xml:"clusterName"`
	AlgoDeploymentId int                    `json:"algoDeploymentId" xml:"algoDeploymentId"`
	Created          int                    `json:"created" xml:"created"`
	AutoSwitch       bool                   `json:"autoSwitch" xml:"autoSwitch"`
	ProgressPercent  int                    `json:"progressPercent" xml:"progressPercent"`
	Schema           map[string]interface{} `json:"schema" xml:"schema"`
	FetchFields      []string               `json:"fetchFields" xml:"fetchFields"`
	Quota            Quota                  `json:"quota" xml:"quota"`
}
