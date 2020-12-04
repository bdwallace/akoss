package domain

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

// Domain is a nested struct in domain response
type Domain struct {
	RegistrationDate         string                           `json:"RegistrationDate" xml:"RegistrationDate"`
	RegistrationDateLong     int64                            `json:"RegistrationDateLong" xml:"RegistrationDateLong"`
	ZhRegistrantOrganization string                           `json:"ZhRegistrantOrganization" xml:"ZhRegistrantOrganization"`
	DomainStatus             string                           `json:"DomainStatus" xml:"DomainStatus"`
	Email                    string                           `json:"Email" xml:"Email"`
	DomainType               string                           `json:"DomainType" xml:"DomainType"`
	Remark                   string                           `json:"Remark" xml:"Remark"`
	DomainName               string                           `json:"DomainName" xml:"DomainName"`
	ProductId                string                           `json:"ProductId" xml:"ProductId"`
	ExpirationDateStatus     string                           `json:"ExpirationDateStatus" xml:"ExpirationDateStatus"`
	ExpirationDateLong       int64                            `json:"ExpirationDateLong" xml:"ExpirationDateLong"`
	RegistrantType           string                           `json:"RegistrantType" xml:"RegistrantType"`
	ExpirationDate           string                           `json:"ExpirationDate" xml:"ExpirationDate"`
	Premium                  bool                             `json:"Premium" xml:"Premium"`
	DomainAuditStatus        string                           `json:"DomainAuditStatus" xml:"DomainAuditStatus"`
	DomainGroupName          string                           `json:"DomainGroupName" xml:"DomainGroupName"`
	InstanceId               string                           `json:"InstanceId" xml:"InstanceId"`
	ExpirationCurrDateDiff   int                              `json:"ExpirationCurrDateDiff" xml:"ExpirationCurrDateDiff"`
	RegistrantOrganization   string                           `json:"RegistrantOrganization" xml:"RegistrantOrganization"`
	DomainGroupId            string                           `json:"DomainGroupId" xml:"DomainGroupId"`
	DnsList                  DnsListInQueryAdvancedDomainList `json:"DnsList" xml:"DnsList"`
}
