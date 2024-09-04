/*
Copyright (c) 2023 OceanBase
ob-operator is licensed under Mulan PSL v2.
You can use this software according to the terms and conditions of the Mulan PSL v2.
You may obtain a copy of Mulan PSL v2 at:
         http://license.coscl.org.cn/MulanPSL2
THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
See the Mulan PSL v2 for more details.
*/

package sql

type MetricCategory string

const (
	Meta      MetricCategory = "meta"
	Latency   MetricCategory = "latency"
	Execution MetricCategory = "execution"
)

type SqlMetricMeta struct {
	Name             string         `json:"name" binding:"required"`
	Description      string         `json:"description" binding:"required"`
	DisplayByDefault bool           `json:"displayByDefault" binding:"required"`
	Category         MetricCategory `json:"category" binding:"required"`
}

type SqlMetricMetaCategory struct {
	Category MetricCategory  `json:"category" binding:"required"`
	Metrics  []SqlMetricMeta `json:"metrics" binding:"required"`
}
