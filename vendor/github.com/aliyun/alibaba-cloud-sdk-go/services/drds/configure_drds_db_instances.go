package drds

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

// ConfigureDrdsDbInstances invokes the drds.ConfigureDrdsDbInstances API synchronously
// api document: https://help.aliyun.com/api/drds/configuredrdsdbinstances.html
func (client *Client) ConfigureDrdsDbInstances(request *ConfigureDrdsDbInstancesRequest) (response *ConfigureDrdsDbInstancesResponse, err error) {
	response = CreateConfigureDrdsDbInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigureDrdsDbInstancesWithChan invokes the drds.ConfigureDrdsDbInstances API asynchronously
// api document: https://help.aliyun.com/api/drds/configuredrdsdbinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConfigureDrdsDbInstancesWithChan(request *ConfigureDrdsDbInstancesRequest) (<-chan *ConfigureDrdsDbInstancesResponse, <-chan error) {
	responseChan := make(chan *ConfigureDrdsDbInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigureDrdsDbInstances(request)
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

// ConfigureDrdsDbInstancesWithCallback invokes the drds.ConfigureDrdsDbInstances API asynchronously
// api document: https://help.aliyun.com/api/drds/configuredrdsdbinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConfigureDrdsDbInstancesWithCallback(request *ConfigureDrdsDbInstancesRequest, callback func(response *ConfigureDrdsDbInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigureDrdsDbInstancesResponse
		var err error
		defer close(result)
		response, err = client.ConfigureDrdsDbInstances(request)
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

// ConfigureDrdsDbInstancesRequest is the request struct for api ConfigureDrdsDbInstances
type ConfigureDrdsDbInstancesRequest struct {
	*requests.RpcRequest
	DbName         string                                `position:"Query" name:"DbName"`
	DbInstance     *[]ConfigureDrdsDbInstancesDbInstance `position:"Query" name:"DbInstance"  type:"Repeated"`
	DrdsInstanceId string                                `position:"Query" name:"DrdsInstanceId"`
}

// ConfigureDrdsDbInstancesDbInstance is a repeated param struct in ConfigureDrdsDbInstancesRequest
type ConfigureDrdsDbInstancesDbInstance struct {
	SlaveDbInstanceType string `name:"SlaveDbInstanceType"`
	SlaveDbInstanceId   string `name:"SlaveDbInstanceId"`
	MasterDbInstanceId  string `name:"MasterDbInstanceId"`
}

// ConfigureDrdsDbInstancesResponse is the response struct for api ConfigureDrdsDbInstances
type ConfigureDrdsDbInstancesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateConfigureDrdsDbInstancesRequest creates a request to invoke ConfigureDrdsDbInstances API
func CreateConfigureDrdsDbInstancesRequest() (request *ConfigureDrdsDbInstancesRequest) {
	request = &ConfigureDrdsDbInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "ConfigureDrdsDbInstances", "drds", "openAPI")
	return
}

// CreateConfigureDrdsDbInstancesResponse creates a response to parse from ConfigureDrdsDbInstances response
func CreateConfigureDrdsDbInstancesResponse() (response *ConfigureDrdsDbInstancesResponse) {
	response = &ConfigureDrdsDbInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}