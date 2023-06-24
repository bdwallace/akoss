package ccc

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

// AgentDevice is a nested struct in ccc response
type AgentDevice struct {
	InstanceId     string `json:"InstanceId" xml:"InstanceId"`
	RamId          int64  `json:"RamId" xml:"RamId"`
	LoginTime      int64  `json:"LoginTime" xml:"LoginTime"`
	ClientIp       string `json:"ClientIp" xml:"ClientIp"`
	ClientPort     string `json:"ClientPort" xml:"ClientPort"`
	BrowserVersion string `json:"BrowserVersion" xml:"BrowserVersion"`
	IsLogin        int    `json:"IsLogin" xml:"IsLogin"`
	Remark         string `json:"Remark" xml:"Remark"`
}
