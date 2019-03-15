package finmall

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

// GetCreditStatus invokes the finmall.GetCreditStatus API synchronously
// api document: https://help.aliyun.com/api/finmall/getcreditstatus.html
func (client *Client) GetCreditStatus(request *GetCreditStatusRequest) (response *GetCreditStatusResponse, err error) {
	response = CreateGetCreditStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetCreditStatusWithChan invokes the finmall.GetCreditStatus API asynchronously
// api document: https://help.aliyun.com/api/finmall/getcreditstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCreditStatusWithChan(request *GetCreditStatusRequest) (<-chan *GetCreditStatusResponse, <-chan error) {
	responseChan := make(chan *GetCreditStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCreditStatus(request)
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

// GetCreditStatusWithCallback invokes the finmall.GetCreditStatus API asynchronously
// api document: https://help.aliyun.com/api/finmall/getcreditstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCreditStatusWithCallback(request *GetCreditStatusRequest, callback func(response *GetCreditStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCreditStatusResponse
		var err error
		defer close(result)
		response, err = client.GetCreditStatus(request)
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

// GetCreditStatusRequest is the request struct for api GetCreditStatus
type GetCreditStatusRequest struct {
	*requests.RpcRequest
	CreditId string `position:"Query" name:"CreditId"`
	UserId   string `position:"Query" name:"UserId"`
}

// GetCreditStatusResponse is the response struct for api GetCreditStatus
type GetCreditStatusResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetCreditStatusRequest creates a request to invoke GetCreditStatus API
func CreateGetCreditStatusRequest() (request *GetCreditStatusRequest) {
	request = &GetCreditStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("finmall", "2018-07-23", "GetCreditStatus", "finmall", "openAPI")
	return
}

// CreateGetCreditStatusResponse creates a response to parse from GetCreditStatus response
func CreateGetCreditStatusResponse() (response *GetCreditStatusResponse) {
	response = &GetCreditStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}