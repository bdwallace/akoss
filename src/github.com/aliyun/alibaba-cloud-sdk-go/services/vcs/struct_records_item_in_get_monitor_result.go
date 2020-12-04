package vcs

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

// RecordsItemInGetMonitorResult is a nested struct in vcs response
type RecordsItemInGetMonitorResult struct {
	RightBottomY  string     `json:"RightBottomY" xml:"RightBottomY"`
	RightBottomX  string     `json:"RightBottomX" xml:"RightBottomX"`
	LeftUpY       string     `json:"LeftUpY" xml:"LeftUpY"`
	LeftUpX       string     `json:"LeftUpX" xml:"LeftUpX"`
	GbId          string     `json:"GbId" xml:"GbId"`
	Score         string     `json:"Score" xml:"Score"`
	PicUrl        string     `json:"PicUrl" xml:"PicUrl"`
	ShotTime      string     `json:"ShotTime" xml:"ShotTime"`
	MonitorPicUrl string     `json:"MonitorPicUrl" xml:"MonitorPicUrl"`
	TargetPicUrl  string     `json:"TargetPicUrl" xml:"TargetPicUrl"`
	TaskId        string     `json:"TaskId" xml:"TaskId"`
	ExtendInfo    ExtendInfo `json:"ExtendInfo" xml:"ExtendInfo"`
}
