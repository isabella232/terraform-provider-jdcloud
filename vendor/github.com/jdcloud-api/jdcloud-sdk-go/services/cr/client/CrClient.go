// Copyright 2018 JDCLOUD.COM
//
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
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package client

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    cr "github.com/jdcloud-api/jdcloud-sdk-go/services/cr/apis"
    "encoding/json"
    "errors"
)

type CrClient struct {
    core.JDCloudClient
}

func NewCrClient(credential *core.Credential) *CrClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("cr.jdcloud-api.com")

    return &CrClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "cr",
            Revision:    "0.1.1",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *CrClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *CrClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 查询指定注册表名称是否已经存在以及是否符合命名规范。
 */
func (c *CrClient) CheckRegistryName(request *cr.CheckRegistryNameRequest) (*cr.CheckRegistryNameResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cr.CheckRegistryNameResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* <p>申请12小时有效期的令牌。 使用<code>docker</code> CLI push和pull镜像。</p>
<p><code>authorizationToken</code>为每个registry返回一个base64编码的字符串，解码后<code>docker login</code>命令
可完成指定registry的鉴权。JCR CLI提供<code>jcr get-login</code>进行认证处理。</p>
 */
func (c *CrClient) GetAuthorizationToken(request *cr.GetAuthorizationTokenRequest) (*cr.GetAuthorizationTokenResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cr.GetAuthorizationTokenResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 返回指定repository中images的元数据，包括image size, image tags和creation date。
 */
func (c *CrClient) DescribeImages(request *cr.DescribeImagesRequest) (*cr.DescribeImagesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cr.DescribeImagesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 通过参数创建注册表。
 */
func (c *CrClient) CreateRegistry(request *cr.CreateRegistryRequest) (*cr.CreateRegistryResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cr.CreateRegistryResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询配额 */
func (c *CrClient) DescribeQuotas(request *cr.DescribeQuotasRequest) (*cr.DescribeQuotasResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cr.DescribeQuotasResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 通过参数创建镜像仓库。
仓库名称可以分解为多个路径名，每个名称必须至少包含一个小写字母数字，考虑URL规范。
支持包含段划线或者下划线进行分割，但不允许点'.'，多个路径名之间通过("/")连接，总长度不超过256个字符，当前只支持二级目录。
 */
func (c *CrClient) CreateRepository(request *cr.CreateRepositoryRequest) (*cr.CreateRepositoryResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cr.CreateRepositoryResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}
