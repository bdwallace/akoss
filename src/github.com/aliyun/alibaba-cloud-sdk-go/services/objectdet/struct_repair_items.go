package objectdet

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

// RepairItems is a nested struct in objectdet response
type RepairItems struct {
	GarageType           string `json:"GarageType" xml:"GarageType"`
	OeMatch              bool   `json:"OeMatch" xml:"OeMatch"`
	OutStandardPartsId   string `json:"OutStandardPartsId" xml:"OutStandardPartsId"`
	OutStandardPartsName string `json:"OutStandardPartsName" xml:"OutStandardPartsName"`
	PartNameMatch        bool   `json:"PartNameMatch" xml:"PartNameMatch"`
	PartsStdCode         string `json:"PartsStdCode" xml:"PartsStdCode"`
	PartsStdName         string `json:"PartsStdName" xml:"PartsStdName"`
	RelationType         string `json:"RelationType" xml:"RelationType"`
	RepairFee            string `json:"RepairFee" xml:"RepairFee"`
	RepairType           string `json:"RepairType" xml:"RepairType"`
	RepairTypeName       string `json:"RepairTypeName" xml:"RepairTypeName"`
}
