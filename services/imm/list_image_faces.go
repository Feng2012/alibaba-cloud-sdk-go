package imm

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

// ListImageFaces invokes the imm.ListImageFaces API synchronously
// api document: https://help.aliyun.com/api/imm/listimagefaces.html
func (client *Client) ListImageFaces(request *ListImageFacesRequest) (response *ListImageFacesResponse, err error) {
	response = CreateListImageFacesResponse()
	err = client.DoAction(request, response)
	return
}

// ListImageFacesWithChan invokes the imm.ListImageFaces API asynchronously
// api document: https://help.aliyun.com/api/imm/listimagefaces.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListImageFacesWithChan(request *ListImageFacesRequest) (<-chan *ListImageFacesResponse, <-chan error) {
	responseChan := make(chan *ListImageFacesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListImageFaces(request)
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

// ListImageFacesWithCallback invokes the imm.ListImageFaces API asynchronously
// api document: https://help.aliyun.com/api/imm/listimagefaces.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListImageFacesWithCallback(request *ListImageFacesRequest, callback func(response *ListImageFacesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListImageFacesResponse
		var err error
		defer close(result)
		response, err = client.ListImageFaces(request)
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

// ListImageFacesRequest is the request struct for api ListImageFaces
type ListImageFacesRequest struct {
	*requests.RpcRequest
	MaxKeys requests.Integer `position:"Query" name:"MaxKeys"`
	Marker  string           `position:"Query" name:"Marker"`
	Project string           `position:"Query" name:"Project"`
	SetId   string           `position:"Query" name:"SetId"`
	SrcUri  string           `position:"Query" name:"SrcUri"`
}

// ListImageFacesResponse is the response struct for api ListImageFaces
type ListImageFacesResponse struct {
	*responses.BaseResponse
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	NextMarker string      `json:"NextMarker" xml:"NextMarker"`
	Faces      []FacesItem `json:"Faces" xml:"Faces"`
}

// CreateListImageFacesRequest creates a request to invoke ListImageFaces API
func CreateListImageFacesRequest() (request *ListImageFacesRequest) {
	request = &ListImageFacesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "ListImageFaces", "imm", "openAPI")
	return
}

// CreateListImageFacesResponse creates a response to parse from ListImageFaces response
func CreateListImageFacesResponse() (response *ListImageFacesResponse) {
	response = &ListImageFacesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
