// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by model-api-gen. DO NOT EDIT.

package yunionconf

import (
	"yunion.io/x/onecloud/pkg/apis"
)

// SParameter is an autogenerated struct via yunion.io/x/onecloud/pkg/yunionconf/models.SParameter.
type SParameter struct {
	apis.SResourceBase
	Id int64 `json:"id"`
	// = Column(BigInteger, primary_key=True)
	CreatedBy string `json:"created_by"`
	// Column(VARCHAR(length=128, charset='ascii'), nullable=False)
	UpdatedBy string `json:"updated_by"`
	// Column(VARCHAR(length=128, charset='ascii'), nullable=False)  "user"/ serviceName/ "admin"
	Namespace string `json:"namespace"`
	// Column(VARCHAR(length=128, charset='ascii'), nullable=False)  user_id / serviceid
	NamespaceId string `json:"namespace_id"`
	// Column(VARCHAR(length=128, charset='ascii'), nullable=False)
	Name string `json:"name"`
	// Column(VARCHAR(length=128, charset='ascii'), nullable=false)
	Value interface{} `json:"value"`
}

// SScopedPolicy is an autogenerated struct via yunion.io/x/onecloud/pkg/yunionconf/models.SScopedPolicy.
type SScopedPolicy struct {
	apis.SInfrasResourceBase
	// 策略类别
	Category string `json:"category"`
	// 策略内容
	Policies interface{} `json:"policies"`
}

// SScopedPolicyBinding is an autogenerated struct via yunion.io/x/onecloud/pkg/yunionconf/models.SScopedPolicyBinding.
type SScopedPolicyBinding struct {
	apis.SResourceBase
	Category  string `json:"category"`
	DomainId  string `json:"domain_id"`
	ProjectId string `json:"project_id"`
	PolicyId  string `json:"policy_id"`
	Priority  int    `json:"priority"`
}