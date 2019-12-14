package ecs

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

// ImportImage invokes the ecs.ImportImage API synchronously
// api document: https://help.aliyun.com/api/ecs/importimage.html
func (client *Client) ImportImage(request *ImportImageRequest) (response *ImportImageResponse, err error) {
	response = CreateImportImageResponse()
	err = client.DoAction(request, response)
	return
}

// ImportImageWithChan invokes the ecs.ImportImage API asynchronously
// api document: https://help.aliyun.com/api/ecs/importimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ImportImageWithChan(request *ImportImageRequest) (<-chan *ImportImageResponse, <-chan error) {
	responseChan := make(chan *ImportImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImportImage(request)
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

// ImportImageWithCallback invokes the ecs.ImportImage API asynchronously
// api document: https://help.aliyun.com/api/ecs/importimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ImportImageWithCallback(request *ImportImageRequest, callback func(response *ImportImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImportImageResponse
		var err error
		defer close(result)
		response, err = client.ImportImage(request)
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

// ImportImageRequest is the request struct for api ImportImage
type ImportImageRequest struct {
	*requests.RpcRequest
	DiskDeviceMapping    *[]ImportImageDiskDeviceMapping `position:"Query" name:"DiskDeviceMapping"  type:"Repeated"`
	ResourceOwnerId      requests.Integer                `position:"Query" name:"ResourceOwnerId"`
	Description          string                          `position:"Query" name:"Description"`
	Platform             string                          `position:"Query" name:"Platform"`
	ImageName            string                          `position:"Query" name:"ImageName"`
	Architecture         string                          `position:"Query" name:"Architecture"`
	LicenseType          string                          `position:"Query" name:"LicenseType"`
	ResourceOwnerAccount string                          `position:"Query" name:"ResourceOwnerAccount"`
	RoleName             string                          `position:"Query" name:"RoleName"`
	OSType               string                          `position:"Query" name:"OSType"`
	OwnerId              requests.Integer                `position:"Query" name:"OwnerId"`
}

// ImportImageDiskDeviceMapping is a repeated param struct in ImportImageRequest
type ImportImageDiskDeviceMapping struct {
	OSSBucket     string `name:"OSSBucket"`
	DiskImSize    string `name:"DiskImSize"`
	Format        string `name:"Format"`
	Device        string `name:"Device"`
	OSSObject     string `name:"OSSObject"`
	DiskImageSize string `name:"DiskImageSize"`
}

// ImportImageResponse is the response struct for api ImportImage
type ImportImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
	RegionId  string `json:"RegionId" xml:"RegionId"`
	ImageId   string `json:"ImageId" xml:"ImageId"`
}

// CreateImportImageRequest creates a request to invoke ImportImage API
func CreateImportImageRequest() (request *ImportImageRequest) {
	request = &ImportImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ImportImage", "ecs", "openAPI")
	return
}

// CreateImportImageResponse creates a response to parse from ImportImage response
func CreateImportImageResponse() (response *ImportImageResponse) {
	response = &ImportImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
