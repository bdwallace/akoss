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

// Result is a nested struct in opensearch response
type Result struct {
	Config             map[string]interface{}   `json:"config" xml:"config"`
	Traffic            int                      `json:"traffic" xml:"traffic"`
	Created            int                      `json:"created" xml:"created"`
	ReceivedCount      int                      `json:"receivedCount" xml:"receivedCount"`
	End                int                      `json:"end" xml:"end"`
	IndustryName       string                   `json:"industryName" xml:"industryName"`
	Start              int                      `json:"start" xml:"start"`
	AnalyzeStatus      string                   `json:"analyzeStatus" xml:"analyzeStatus"`
	Domain             string                   `json:"domain" xml:"domain"`
	Name               string                   `json:"name" xml:"name"`
	Index              int                      `json:"index" xml:"index"`
	CreateTime         string                   `json:"createTime" xml:"createTime"`
	Type               string                   `json:"type" xml:"type"`
	Updated            int                      `json:"updated" xml:"updated"`
	Content            string                   `json:"content" xml:"content"`
	RegionId           string                   `json:"regionId" xml:"regionId"`
	Status             int                      `json:"status" xml:"status"`
	Version            int64                    `json:"version" xml:"version"`
	Online             bool                     `json:"online" xml:"online"`
	Description        string                   `json:"description" xml:"description"`
	SundialId          string                   `json:"sundialId" xml:"sundialId"`
	Params             map[string]interface{}   `json:"params" xml:"params"`
	AppQuery           string                   `json:"appQuery" xml:"appQuery"`
	ModifyTime         string                   `json:"modifyTime" xml:"modifyTime"`
	DataCollectionType string                   `json:"dataCollectionType" xml:"dataCollectionType"`
	Id                 string                   `json:"id" xml:"id"`
	Active             bool                     `json:"active" xml:"active"`
	Processors         []map[string]interface{} `json:"processors" xml:"processors"`
	Indexes            []string                 `json:"indexes" xml:"indexes"`
	Values             []string                 `json:"values" xml:"values"`
	ReceivedSample     []ReceivedSampleItem     `json:"receivedSample" xml:"receivedSample"`
	Meta               []MetaItem               `json:"meta" xml:"meta"`
}
