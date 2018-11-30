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

package models


type InstanceTemplateSpec struct {

    /* 实例规格，可查询<a href="http://docs.jdcloud.com/virtual-machines/api/describeinstancetypes">DescribeInstanceTypes</a>接口获得指定地域或可用区的规格信息。  */
    InstanceType string `json:"instanceType"`

    /* 镜像ID，可查询<a href="http://docs.jdcloud.com/virtual-machines/api/describeimages">DescribeImages</a>接口获得指定地域的镜像信息。  */
    ImageId string `json:"imageId"`

    /* 密码，<a href="http://docs.jdcloud.com/virtual-machines/api/general_parameters">参考公共参数规范</a>。 (Optional) */
    Password string `json:"password"`

    /* 密钥对名称；当前只支持一个 (Optional) */
    KeyNames []string `json:"keyNames"`

    /* 主网卡主IP关联的弹性IP规格 (Optional) */
    ElasticIp InstanceTemplateElasticIpSpec `json:"elasticIp"`

    /* 主网卡配置信息  */
    PrimaryNetworkInterface InstanceTemplateNetworkInterfaceAttachmentSpec `json:"primaryNetworkInterface"`

    /* 系统盘配置信息  */
    SystemDisk InstanceTemplateDiskAttachmentSpec `json:"systemDisk"`

    /* 数据盘配置信息 (Optional) */
    DataDisks []InstanceTemplateDiskAttachmentSpec `json:"dataDisks"`
}
