package cusanalytic_sc_online

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

// VisitorItem is a nested struct in cusanalytic_sc_online response
type VisitorItem struct {
	Img           string `json:"Img" xml:"Img"`
	PkId          string `json:"PkId" xml:"PkId"`
	LatelyTime    int64  `json:"LatelyTime" xml:"LatelyTime"`
	EarliestPlace string `json:"EarliestPlace" xml:"EarliestPlace"`
	UkId          string `json:"UkId" xml:"UkId"`
	Gender        string `json:"Gender" xml:"Gender"`
	EarliestTime  int64  `json:"EarliestTime" xml:"EarliestTime"`
	LatelyPlace   string `json:"LatelyPlace" xml:"LatelyPlace"`
	Age           int64  `json:"Age" xml:"Age"`
	StoreId       int64  `json:"StoreId" xml:"StoreId"`
	EnterCount    int64  `json:"EnterCount" xml:"EnterCount"`
}
