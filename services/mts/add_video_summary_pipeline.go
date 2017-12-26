package mts

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

func (client *Client) AddVideoSummaryPipeline(request *AddVideoSummaryPipelineRequest) (response *AddVideoSummaryPipelineResponse, err error) {
	response = CreateAddVideoSummaryPipelineResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AddVideoSummaryPipelineWithChan(request *AddVideoSummaryPipelineRequest) (<-chan *AddVideoSummaryPipelineResponse, <-chan error) {
	responseChan := make(chan *AddVideoSummaryPipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddVideoSummaryPipeline(request)
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

func (client *Client) AddVideoSummaryPipelineWithCallback(request *AddVideoSummaryPipelineRequest, callback func(response *AddVideoSummaryPipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddVideoSummaryPipelineResponse
		var err error
		defer close(result)
		response, err = client.AddVideoSummaryPipeline(request)
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

type AddVideoSummaryPipelineRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Action               string `position:"Query" name:"Action"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	Priority             string `position:"Query" name:"Priority"`
	AccessKeyId          string `position:"Query" name:"AccessKeyId"`
}

type AddVideoSummaryPipelineResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId"`
	Pipeline  struct {
		Id           string `json:"Id"`
		Name         string `json:"Name"`
		Priority     int    `json:"Priority"`
		State        string `json:"State"`
		NotifyConfig struct {
			Topic string `json:"Topic"`
			Queue string `json:"Queue"`
		} `json:"NotifyConfig"`
	} `json:"Pipeline"`
}

func CreateAddVideoSummaryPipelineRequest() (request *AddVideoSummaryPipelineRequest) {
	request = &AddVideoSummaryPipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "AddVideoSummaryPipeline", "", "")
	return
}

func CreateAddVideoSummaryPipelineResponse() (response *AddVideoSummaryPipelineResponse) {
	response = &AddVideoSummaryPipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}