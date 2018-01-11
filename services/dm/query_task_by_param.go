package dm

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) QueryTaskByParam(request *QueryTaskByParamRequest) (response *QueryTaskByParamResponse, err error) {
	response = CreateQueryTaskByParamResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryTaskByParamWithChan(request *QueryTaskByParamRequest) (<-chan *QueryTaskByParamResponse, <-chan error) {
	responseChan := make(chan *QueryTaskByParamResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTaskByParam(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) QueryTaskByParamWithCallback(request *QueryTaskByParamRequest, callback func(response *QueryTaskByParamResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTaskByParamResponse
		var err error
		defer close(result)
		response, err = client.QueryTaskByParam(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type QueryTaskByParamRequest struct {
	*requests.RpcRequest
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Status               requests.Integer `position:"Query" name:"Status"`
	KeyWord              string           `position:"Query" name:"KeyWord"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNo               requests.Integer `position:"Query" name:"PageNo"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type QueryTaskByParamResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	Data       struct {
		Task []struct {
			TaskId        string `json:"TaskId" xml:"TaskId"`
			ReceiversName string `json:"ReceiversName" xml:"ReceiversName"`
			TemplateName  string `json:"TemplateName" xml:"TemplateName"`
			TaskStatus    string `json:"TaskStatus" xml:"TaskStatus"`
			CreateTime    string `json:"CreateTime" xml:"CreateTime"`
			UtcCreateTime int    `json:"UtcCreateTime" xml:"UtcCreateTime"`
			AddressType   string `json:"AddressType" xml:"AddressType"`
			TagName       string `json:"TagName" xml:"TagName"`
			RequestCount  string `json:"RequestCount" xml:"RequestCount"`
		} `json:"task" xml:"task"`
	} `json:"data" xml:"data"`
}

func CreateQueryTaskByParamRequest() (request *QueryTaskByParamRequest) {
	request = &QueryTaskByParamRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "QueryTaskByParam", "", "")
	return
}

func CreateQueryTaskByParamResponse() (response *QueryTaskByParamResponse) {
	response = &QueryTaskByParamResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}