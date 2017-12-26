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

func (client *Client) ListAllCategory(request *ListAllCategoryRequest) (response *ListAllCategoryResponse, err error) {
	response = CreateListAllCategoryResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListAllCategoryWithChan(request *ListAllCategoryRequest) (<-chan *ListAllCategoryResponse, <-chan error) {
	responseChan := make(chan *ListAllCategoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAllCategory(request)
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

func (client *Client) ListAllCategoryWithCallback(request *ListAllCategoryRequest, callback func(response *ListAllCategoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAllCategoryResponse
		var err error
		defer close(result)
		response, err = client.ListAllCategory(request)
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

type ListAllCategoryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Action               string `position:"Query" name:"Action"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	AccessKeyId          string `position:"Query" name:"AccessKeyId"`
}

type ListAllCategoryResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId"`
	CategoryList []struct {
		CateId   string `json:"CateId"`
		CateName string `json:"CateName"`
		ParentId string `json:"ParentId"`
		Level    string `json:"Level"`
	} `json:"CategoryList"`
}

func CreateListAllCategoryRequest() (request *ListAllCategoryRequest) {
	request = &ListAllCategoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ListAllCategory", "", "")
	return
}

func CreateListAllCategoryResponse() (response *ListAllCategoryResponse) {
	response = &ListAllCategoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}