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

// Project is a nested struct in yundun_ds response
type Project struct {
	Name            string `json:"Name" xml:"Name"`
	TotalCount      int64  `json:"TotalCount" xml:"TotalCount"`
	LastCount       int64  `json:"LastCount" xml:"LastCount"`
	LoginName       string `json:"LoginName" xml:"LoginName"`
	TopicCount      string `json:"TopicCount" xml:"TopicCount"`
	Id              int64  `json:"Id" xml:"Id"`
	CreationTime    int64  `json:"CreationTime" xml:"CreationTime"`
	TopicBlobCount  int64  `json:"TopicBlobCount" xml:"TopicBlobCount"`
	UserId          int64  `json:"UserId" xml:"UserId"`
	Description     string `json:"Description" xml:"Description"`
	TopicTupleCount string `json:"TopicTupleCount" xml:"TopicTupleCount"`
	DisplayName     string `json:"DisplayName" xml:"DisplayName"`
}
