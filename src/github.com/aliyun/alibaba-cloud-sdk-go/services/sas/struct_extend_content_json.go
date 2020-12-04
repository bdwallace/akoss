package sas

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

// ExtendContentJson is a nested struct in sas response
type ExtendContentJson struct {
	Owasp         string      `json:"Owasp" xml:"Owasp"`
	Title         string      `json:"Title" xml:"Title"`
	Effect        string      `json:"Effect" xml:"Effect"`
	VulType       string      `json:"VulType" xml:"VulType"`
	EmgProof      string      `json:"EmgProof" xml:"EmgProof"`
	Reason        string      `json:"Reason" xml:"Reason"`
	Os            string      `json:"Os" xml:"Os"`
	Wasc          string      `json:"Wasc" xml:"Wasc"`
	Tag           string      `json:"Tag" xml:"Tag"`
	AliasName     string      `json:"AliasName" xml:"AliasName"`
	Description   string      `json:"Description" xml:"Description"`
	Solution      string      `json:"Solution" xml:"Solution"`
	Status        int         `json:"Status" xml:"Status"`
	AbsolutePath  string      `json:"AbsolutePath" xml:"AbsolutePath"`
	Ip            string      `json:"Ip" xml:"Ip"`
	Proof         string      `json:"Proof" xml:"Proof"`
	Cwe           string      `json:"Cwe" xml:"Cwe"`
	LastTs        int64       `json:"LastTs" xml:"LastTs"`
	OsRelease     string      `json:"OsRelease" xml:"OsRelease"`
	Target        string      `json:"Target" xml:"Target"`
	Reference     string      `json:"Reference" xml:"Reference"`
	PrimaryId     int64       `json:"PrimaryId" xml:"PrimaryId"`
	CveList       []string    `json:"cveList" xml:"cveList"`
	Necessity     Necessity   `json:"Necessity" xml:"Necessity"`
	RpmEntityList []RpmEntity `json:"RpmEntityList" xml:"RpmEntityList"`
}
