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

func (client *Client) AddTemplate(request *AddTemplateRequest) (response *AddTemplateResponse, err error) {
	response = CreateAddTemplateResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AddTemplateWithChan(request *AddTemplateRequest) (<-chan *AddTemplateResponse, <-chan error) {
	responseChan := make(chan *AddTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddTemplate(request)
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

func (client *Client) AddTemplateWithCallback(request *AddTemplateRequest, callback func(response *AddTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddTemplateResponse
		var err error
		defer close(result)
		response, err = client.AddTemplate(request)
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

type AddTemplateRequest struct {
	*requests.RpcRequest
	Container            string `position:"Query" name:"Container"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Action               string `position:"Query" name:"Action"`
	TransConfig          string `position:"Query" name:"TransConfig"`
	MuxConfig            string `position:"Query" name:"MuxConfig"`
	Video                string `position:"Query" name:"Video"`
	Audio                string `position:"Query" name:"Audio"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	AccessKeyId          string `position:"Query" name:"AccessKeyId"`
}

type AddTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId"`
	Template  struct {
		Id        string `json:"Id"`
		Name      string `json:"Name"`
		State     string `json:"State"`
		Container struct {
			Format string `json:"Format"`
		} `json:"Container"`
		Video struct {
			Codec      string `json:"Codec"`
			Profile    string `json:"Profile"`
			Bitrate    string `json:"Bitrate"`
			Crf        string `json:"Crf"`
			Width      string `json:"Width"`
			Height     string `json:"Height"`
			Fps        string `json:"Fps"`
			Gop        string `json:"Gop"`
			Preset     string `json:"Preset"`
			ScanMode   string `json:"ScanMode"`
			Bufsize    string `json:"Bufsize"`
			Maxrate    string `json:"Maxrate"`
			PixFmt     string `json:"PixFmt"`
			Degrain    string `json:"Degrain"`
			Qscale     string `json:"Qscale"`
			Remove     string `json:"Remove"`
			Crop       string `json:"Crop"`
			Pad        string `json:"Pad"`
			MaxFps     string `json:"MaxFps"`
			BitrateBnd struct {
				Max string `json:"Max"`
				Min string `json:"Min"`
			} `json:"BitrateBnd"`
		} `json:"Video"`
		Audio struct {
			Codec      string `json:"Codec"`
			Profile    string `json:"Profile"`
			Samplerate string `json:"Samplerate"`
			Bitrate    string `json:"Bitrate"`
			Channels   string `json:"Channels"`
			Qscale     string `json:"Qscale"`
			Remove     string `json:"Remove"`
			Volume     struct {
				Level  string `json:"Level"`
				Method string `json:"Method"`
			} `json:"Volume"`
		} `json:"Audio"`
		TransConfig struct {
			TransMode               string `json:"TransMode"`
			IsCheckReso             string `json:"IsCheckReso"`
			IsCheckResoFail         string `json:"IsCheckResoFail"`
			IsCheckVideoBitrate     string `json:"IsCheckVideoBitrate"`
			IsCheckAudioBitrate     string `json:"IsCheckAudioBitrate"`
			AdjDarMethod            string `json:"AdjDarMethod"`
			IsCheckVideoBitrateFail string `json:"IsCheckVideoBitrateFail"`
			IsCheckAudioBitrateFail string `json:"IsCheckAudioBitrateFail"`
		} `json:"TransConfig"`
		MuxConfig struct {
			Segment struct {
				Duration string `json:"Duration"`
			} `json:"Segment"`
			Gif struct {
				Loop            string `json:"Loop"`
				FinalDelay      string `json:"FinalDelay"`
				IsCustomPalette string `json:"IsCustomPalette"`
				DitherMode      string `json:"DitherMode"`
			} `json:"Gif"`
		} `json:"MuxConfig"`
	} `json:"Template"`
}

func CreateAddTemplateRequest() (request *AddTemplateRequest) {
	request = &AddTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "AddTemplate", "", "")
	return
}

func CreateAddTemplateResponse() (response *AddTemplateResponse) {
	response = &AddTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}