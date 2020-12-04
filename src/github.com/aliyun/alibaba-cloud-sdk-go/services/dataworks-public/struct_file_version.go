package dataworks_public

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

// FileVersion is a nested struct in dataworks_public response
type FileVersion struct {
	FileVersion         int    `json:"FileVersion" xml:"FileVersion"`
	FileContent         string `json:"FileContent" xml:"FileContent"`
	CommitTime          int64  `json:"CommitTime" xml:"CommitTime"`
	CommitUser          string `json:"CommitUser" xml:"CommitUser"`
	FileName            string `json:"FileName" xml:"FileName"`
	Status              string `json:"Status" xml:"Status"`
	ChangeType          string `json:"ChangeType" xml:"ChangeType"`
	IsCurrentProd       bool   `json:"IsCurrentProd" xml:"IsCurrentProd"`
	NodeId              int64  `json:"NodeId" xml:"NodeId"`
	Comment             string `json:"Comment" xml:"Comment"`
	NodeContent         string `json:"NodeContent" xml:"NodeContent"`
	FilePropertyContent string `json:"FilePropertyContent" xml:"FilePropertyContent"`
	UseType             string `json:"UseType" xml:"UseType"`
}
