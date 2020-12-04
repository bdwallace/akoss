package yundun_ds

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

// Rule is a nested struct in yundun_ds response
type Rule struct {
	Name          string `json:"Name" xml:"Name"`
	Status        int    `json:"Status" xml:"Status"`
	RiskLevelName string `json:"RiskLevelName" xml:"RiskLevelName"`
	RuleName      string `json:"RuleName" xml:"RuleName"`
	DepartName    string `json:"DepartName" xml:"DepartName"`
	Id            int64  `json:"Id" xml:"Id"`
	CustomType    int    `json:"CustomType" xml:"CustomType"`
	DisplayName   string `json:"DisplayName" xml:"DisplayName"`
	GmtCreate     int64  `json:"GmtCreate" xml:"GmtCreate"`
	CategoryName  string `json:"CategoryName" xml:"CategoryName"`
	Content       string `json:"Content" xml:"Content"`
	LoginName     string `json:"LoginName" xml:"LoginName"`
	Count         int64  `json:"Count" xml:"Count"`
	GmtModified   int64  `json:"GmtModified" xml:"GmtModified"`
	Category      int    `json:"Category" xml:"Category"`
	RiskLevelId   int64  `json:"RiskLevelId" xml:"RiskLevelId"`
	UserId        int64  `json:"UserId" xml:"UserId"`
	Description   string `json:"Description" xml:"Description"`
}
