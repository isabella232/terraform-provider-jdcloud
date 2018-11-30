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


type TopicListInfo struct {

    /* 是否归档（0：未归档，1：已归档） (Optional) */
    Archived int `json:"archived"`

    /* 创建topic的时间戳 (Optional) */
    CreatedTime string `json:"createdTime"`

    /* topic是否已删除（0：未删除，1：删除） (Optional) */
    Deleted int `json:"deleted"`

    /* topic的id值 (Optional) */
    Id int `json:"id"`

    /* 数据写入后的保留时间 (Optional) */
    Lifecycle int `json:"lifecycle"`

    /* 流数据总线中topic的名字 (Optional) */
    Name string `json:"name"`

    /* 对应kafka中的分区数 (Optional) */
    PartitionNum string `json:"partitionNum"`

    /* 备注 (Optional) */
    Remark string `json:"remark"`

    /* 流数据总线shard的数量 (Optional) */
    ShardNum string `json:"shardNum"`

    /* 0: 已创建, 1: 创建中 (Optional) */
    Status int `json:"status"`

    /* 对应kafka中的topic名字 (Optional) */
    TopicName string `json:"topicName"`

    /* 对应kafka中的uid (Optional) */
    Uid string `json:"uid"`

    /* 更新topic的时间戳 (Optional) */
    UpdatedTime string `json:"updatedTime"`

    /* 用户的userPin (Optional) */
    UserPin string `json:"userPin"`

    /*  (Optional) */
    DataSize string `json:"dataSize"`
}
