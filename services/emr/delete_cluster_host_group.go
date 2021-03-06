package emr

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

// DeleteClusterHostGroup invokes the emr.DeleteClusterHostGroup API synchronously
// api document: https://help.aliyun.com/api/emr/deleteclusterhostgroup.html
func (client *Client) DeleteClusterHostGroup(request *DeleteClusterHostGroupRequest) (response *DeleteClusterHostGroupResponse, err error) {
	response = CreateDeleteClusterHostGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteClusterHostGroupWithChan invokes the emr.DeleteClusterHostGroup API asynchronously
// api document: https://help.aliyun.com/api/emr/deleteclusterhostgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteClusterHostGroupWithChan(request *DeleteClusterHostGroupRequest) (<-chan *DeleteClusterHostGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteClusterHostGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteClusterHostGroup(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DeleteClusterHostGroupWithCallback invokes the emr.DeleteClusterHostGroup API asynchronously
// api document: https://help.aliyun.com/api/emr/deleteclusterhostgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteClusterHostGroupWithCallback(request *DeleteClusterHostGroupRequest, callback func(response *DeleteClusterHostGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteClusterHostGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteClusterHostGroup(request)
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

// DeleteClusterHostGroupRequest is the request struct for api DeleteClusterHostGroup
type DeleteClusterHostGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	HostGroupId     string           `position:"Query" name:"HostGroupId"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
}

// DeleteClusterHostGroupResponse is the response struct for api DeleteClusterHostGroup
type DeleteClusterHostGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteClusterHostGroupRequest creates a request to invoke DeleteClusterHostGroup API
func CreateDeleteClusterHostGroupRequest() (request *DeleteClusterHostGroupRequest) {
	request = &DeleteClusterHostGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DeleteClusterHostGroup", "emr", "openAPI")
	return
}

// CreateDeleteClusterHostGroupResponse creates a response to parse from DeleteClusterHostGroup response
func CreateDeleteClusterHostGroupResponse() (response *DeleteClusterHostGroupResponse) {
	response = &DeleteClusterHostGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
