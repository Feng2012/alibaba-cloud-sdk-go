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

func (client *Client) ListMediaWorkflowExecutions(request *ListMediaWorkflowExecutionsRequest) (response *ListMediaWorkflowExecutionsResponse, err error) {
	response = CreateListMediaWorkflowExecutionsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListMediaWorkflowExecutionsWithChan(request *ListMediaWorkflowExecutionsRequest) (<-chan *ListMediaWorkflowExecutionsResponse, <-chan error) {
	responseChan := make(chan *ListMediaWorkflowExecutionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMediaWorkflowExecutions(request)
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

func (client *Client) ListMediaWorkflowExecutionsWithCallback(request *ListMediaWorkflowExecutionsRequest, callback func(response *ListMediaWorkflowExecutionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMediaWorkflowExecutionsResponse
		var err error
		defer close(result)
		response, err = client.ListMediaWorkflowExecutions(request)
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

type ListMediaWorkflowExecutionsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	InputFileURL         string `position:"Query" name:"InputFileURL"`
	NextPageToken        string `position:"Query" name:"NextPageToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Action               string `position:"Query" name:"Action"`
	MaximumPageSize      string `position:"Query" name:"MaximumPageSize"`
	MediaWorkflowId      string `position:"Query" name:"MediaWorkflowId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaWorkflowName    string `position:"Query" name:"MediaWorkflowName"`
	AccessKeyId          string `position:"Query" name:"AccessKeyId"`
}

type ListMediaWorkflowExecutionsResponse struct {
	*responses.BaseResponse
	RequestId                  string `json:"RequestId"`
	NextPageToken              string `json:"NextPageToken"`
	MediaWorkflowExecutionList []struct {
		RunId           string `json:"RunId"`
		MediaWorkflowId string `json:"MediaWorkflowId"`
		Name            string `json:"Name"`
		State           string `json:"State"`
		MediaId         string `json:"MediaId"`
		CreationTime    string `json:"CreationTime"`
		Input           struct {
			UserData  string `json:"UserData"`
			InputFile struct {
				Bucket   string `json:"Bucket"`
				Location string `json:"Location"`
				Object   string `json:"Object"`
			} `json:"InputFile"`
		} `json:"Input"`
		ActivityList []struct {
			Name             string `json:"Name"`
			Type             string `json:"Type"`
			JobId            string `json:"JobId"`
			State            string `json:"State"`
			Code             string `json:"Code"`
			Message          string `json:"Message"`
			StartTime        string `json:"StartTime"`
			EndTime          string `json:"EndTime"`
			MNSMessageResult struct {
				MessageId    string `json:"MessageId"`
				ErrorMessage string `json:"ErrorMessage"`
				ErrorCode    string `json:"ErrorCode"`
			} `json:"MNSMessageResult"`
		} `json:"ActivityList"`
	} `json:"MediaWorkflowExecutionList"`
}

func CreateListMediaWorkflowExecutionsRequest() (request *ListMediaWorkflowExecutionsRequest) {
	request = &ListMediaWorkflowExecutionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ListMediaWorkflowExecutions", "", "")
	return
}

func CreateListMediaWorkflowExecutionsResponse() (response *ListMediaWorkflowExecutionsResponse) {
	response = &ListMediaWorkflowExecutionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}