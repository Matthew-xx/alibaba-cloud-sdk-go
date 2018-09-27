package live

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

// Room is a nested struct in live response
type Room struct {
	AnchorId    string `json:"AnchorId" xml:"AnchorId"`
	OpEndTime   string `json:"OpEndTime" xml:"OpEndTime"`
	OpStartTime string `json:"OpStartTime" xml:"OpStartTime"`
	CreateTime  string `json:"CreateTime" xml:"CreateTime"`
	RoomStatus  int    `json:"RoomStatus" xml:"RoomStatus"`
	RoomId      string `json:"RoomId" xml:"RoomId"`
}