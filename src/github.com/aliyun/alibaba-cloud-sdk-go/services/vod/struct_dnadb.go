package vod

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

// DNADB is a nested struct in vod response
type DNADB struct {
	DBId          string `json:"DBId" xml:"DBId"`
	DBType        string `json:"DBType" xml:"DBType"`
	DBRegion      string `json:"DBRegion" xml:"DBRegion"`
	DBName        string `json:"DBName" xml:"DBName"`
	Status        string `json:"Status" xml:"Status"`
	DBDescription string `json:"DBDescription" xml:"DBDescription"`
}
