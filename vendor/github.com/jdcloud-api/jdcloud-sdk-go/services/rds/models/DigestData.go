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


type DigestData struct {

    /* 表示执行结果中95% 数据小于或等于此数值 (Optional) */
    Pct95 float32 `json:"pct95"`

    /* 执行结果的最大值 (Optional) */
    Max float32 `json:"max"`

    /* 执行结果的平均值 (Optional) */
    Avg float32 `json:"avg"`

    /* 执行结果的最小值 (Optional) */
    Min float32 `json:"min"`

    /* 执行结果的合计值 (Optional) */
    Total float64 `json:"total"`
}
