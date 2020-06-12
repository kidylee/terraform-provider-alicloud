package ddoscoo

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

// ModifySchedulerRule invokes the ddoscoo.ModifySchedulerRule API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/modifyschedulerrule.html
func (client *Client) ModifySchedulerRule(request *ModifySchedulerRuleRequest) (response *ModifySchedulerRuleResponse, err error) {
	response = CreateModifySchedulerRuleResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySchedulerRuleWithChan invokes the ddoscoo.ModifySchedulerRule API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/modifyschedulerrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySchedulerRuleWithChan(request *ModifySchedulerRuleRequest) (<-chan *ModifySchedulerRuleResponse, <-chan error) {
	responseChan := make(chan *ModifySchedulerRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySchedulerRule(request)
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

// ModifySchedulerRuleWithCallback invokes the ddoscoo.ModifySchedulerRule API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/modifyschedulerrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySchedulerRuleWithCallback(request *ModifySchedulerRuleRequest, callback func(response *ModifySchedulerRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySchedulerRuleResponse
		var err error
		defer close(result)
		response, err = client.ModifySchedulerRule(request)
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

// ModifySchedulerRuleRequest is the request struct for api ModifySchedulerRule
type ModifySchedulerRuleRequest struct {
	*requests.RpcRequest
	Rules           string           `position:"Query" name:"Rules"`
	RuleName        string           `position:"Query" name:"RuleName"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	Param           string           `position:"Query" name:"Param"`
	RuleType        requests.Integer `position:"Query" name:"RuleType"`
}

// ModifySchedulerRuleResponse is the response struct for api ModifySchedulerRule
type ModifySchedulerRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Cname     string `json:"Cname" xml:"Cname"`
	RuleName  string `json:"RuleName" xml:"RuleName"`
}

// CreateModifySchedulerRuleRequest creates a request to invoke ModifySchedulerRule API
func CreateModifySchedulerRuleRequest() (request *ModifySchedulerRuleRequest) {
	request = &ModifySchedulerRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "ModifySchedulerRule", "ddoscoo", "openAPI")
	return
}

// CreateModifySchedulerRuleResponse creates a response to parse from ModifySchedulerRule response
func CreateModifySchedulerRuleResponse() (response *ModifySchedulerRuleResponse) {
	response = &ModifySchedulerRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}