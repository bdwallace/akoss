package reid

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

// Action is a nested struct in reid response
type Action struct {
	Status                int                   `json:"Status" xml:"Status"`
	LocationLayerType     string                `json:"LocationLayerType" xml:"LocationLayerType"`
	Score                 float64               `json:"Score" xml:"Score"`
	FacePointNumber       int                   `json:"FacePointNumber" xml:"FacePointNumber"`
	ArriveTimestamp       int64                 `json:"ArriveTimestamp" xml:"ArriveTimestamp"`
	StayValid             bool                  `json:"StayValid" xml:"StayValid"`
	Id                    int64                 `json:"Id" xml:"Id"`
	LeaveTimestamp        int64                 `json:"LeaveTimestamp" xml:"LeaveTimestamp"`
	UkId                  int64                 `json:"UkId" xml:"UkId"`
	InStay                int64                 `json:"InStay" xml:"InStay"`
	GmtCreate             int64                 `json:"GmtCreate" xml:"GmtCreate"`
	SpecialType           string                `json:"SpecialType" xml:"SpecialType"`
	Age                   int                   `json:"Age" xml:"Age"`
	Gender                string                `json:"Gender" xml:"Gender"`
	ImageUrl              string                `json:"ImageUrl" xml:"ImageUrl"`
	GmtModified           int64                 `json:"GmtModified" xml:"GmtModified"`
	ImageType             string                `json:"ImageType" xml:"ImageType"`
	StayPeriod            int                   `json:"StayPeriod" xml:"StayPeriod"`
	ImageObjectKey        string                `json:"ImageObjectKey" xml:"ImageObjectKey"`
	LocationId            int64                 `json:"LocationId" xml:"LocationId"`
	StoreId               int64                 `json:"StoreId" xml:"StoreId"`
	PointInMap            PointInMap            `json:"PointInMap" xml:"PointInMap"`
	ObjectPositionInImage ObjectPositionInImage `json:"ObjectPositionInImage" xml:"ObjectPositionInImage"`
}
