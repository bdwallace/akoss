package ecs

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

// ResourceTypeCount is a nested struct in ecs response
type ResourceTypeCount struct {
	Instance         int `json:"Instance" xml:"Instance"`
	Disk             int `json:"Disk" xml:"Disk"`
	Volume           int `json:"Volume" xml:"Volume"`
	Image            int `json:"Image" xml:"Image"`
	Snapshot         int `json:"Snapshot" xml:"Snapshot"`
	Securitygroup    int `json:"Securitygroup" xml:"Securitygroup"`
	LaunchTemplate   int `json:"LaunchTemplate" xml:"LaunchTemplate"`
	Eni              int `json:"Eni" xml:"Eni"`
	Ddh              int `json:"Ddh" xml:"Ddh"`
	KeyPair          int `json:"KeyPair" xml:"KeyPair"`
	SnapshotPolicy   int `json:"SnapshotPolicy" xml:"SnapshotPolicy"`
	ReservedInstance int `json:"ReservedInstance" xml:"ReservedInstance"`
}
