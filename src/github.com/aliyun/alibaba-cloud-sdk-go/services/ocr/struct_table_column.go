package ocr

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

// TableColumn is a nested struct in ocr response
type TableColumn struct {
	StartColumn int      `json:"StartColumn" xml:"StartColumn"`
	StartRow    int      `json:"StartRow" xml:"StartRow"`
	EndColumn   int      `json:"EndColumn" xml:"EndColumn"`
	EndRow      int      `json:"EndRow" xml:"EndRow"`
	Height      int      `json:"Height" xml:"Height"`
	Width       int      `json:"Width" xml:"Width"`
	Texts       []string `json:"Texts" xml:"Texts"`
}
