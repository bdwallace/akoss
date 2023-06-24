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

// QuaraFile is a nested struct in sas response
type QuaraFile struct {
	Name         string `json:"Name" xml:"Name"`
	Path         string `json:"Path" xml:"Path"`
	Id           int    `json:"Id" xml:"Id"`
	InfoType     string `json:"InfoType" xml:"InfoType"`
	ModifyTime   string `json:"ModifyTime" xml:"ModifyTime"`
	InternetIp   string `json:"InternetIp" xml:"InternetIp"`
	Ip           string `json:"Ip" xml:"Ip"`
	InstanceName string `json:"InstanceName" xml:"InstanceName"`
	EventName    string `json:"EventName" xml:"EventName"`
	Uuid         string `json:"Uuid" xml:"Uuid"`
	Value        string `json:"Value" xml:"Value"`
	NameDisplay  string `json:"NameDisplay" xml:"NameDisplay"`
	EventType    string `json:"EventType" xml:"EventType"`
	ValueDisplay string `json:"ValueDisplay" xml:"ValueDisplay"`
	Tag          string `json:"Tag" xml:"Tag"`
	Md5          string `json:"Md5" xml:"Md5"`
	Status       string `json:"Status" xml:"Status"`
	Type         string `json:"Type" xml:"Type"`
}