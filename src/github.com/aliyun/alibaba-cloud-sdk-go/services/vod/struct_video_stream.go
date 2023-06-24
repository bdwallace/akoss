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

// VideoStream is a nested struct in vod response
type VideoStream struct {
	Index          string `json:"Index" xml:"Index"`
	CodecName      string `json:"CodecName" xml:"CodecName"`
	CodecLongName  string `json:"CodecLongName" xml:"CodecLongName"`
	Profile        string `json:"Profile" xml:"Profile"`
	CodecTimeBase  string `json:"CodecTimeBase" xml:"CodecTimeBase"`
	CodecTagString string `json:"CodecTagString" xml:"CodecTagString"`
	CodecTag       string `json:"CodecTag" xml:"CodecTag"`
	Width          string `json:"Width" xml:"Width"`
	Height         string `json:"Height" xml:"Height"`
	HasBFrames     string `json:"HasBFrames" xml:"HasBFrames"`
	Sar            string `json:"Sar" xml:"Sar"`
	Dar            string `json:"Dar" xml:"Dar"`
	PixFmt         string `json:"PixFmt" xml:"PixFmt"`
	Level          string `json:"Level" xml:"Level"`
	Fps            string `json:"Fps" xml:"Fps"`
	AvgFPS         string `json:"AvgFPS" xml:"AvgFPS"`
	Timebase       string `json:"Timebase" xml:"Timebase"`
	StartTime      string `json:"StartTime" xml:"StartTime"`
	Duration       string `json:"Duration" xml:"Duration"`
	NumFrames      string `json:"NumFrames" xml:"NumFrames"`
	Lang           string `json:"Lang" xml:"Lang"`
	Rotate         string `json:"Rotate" xml:"Rotate"`
	Bitrate        string `json:"Bitrate" xml:"Bitrate"`
}