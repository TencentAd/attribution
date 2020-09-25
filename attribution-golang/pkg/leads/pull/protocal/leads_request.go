/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/25/20, 10:25 AM
 */

package protocal

// 参考 https://developers.e.qq.com/docs/api/insights/leads/lead_clues_get?version=1.3&_preview=1
type LeadsRequest struct {
	AccountId             int64                    `json:"account_id"`
	TimeRange             *TimeRange               `json:"time_range"`
	Filtering             []map[string]interface{} `json:"filtering"`
	Page                  int                      `json:"page"`
	PageSize              int                      `json:"page_size"`
	LastSearchAfterValues [2]string                `json:"last_search_after_values"`
}

type TimeRange struct {
	StartTime int64 `json:"start_time"`
	EndTime   int64 `json:"end_time"`
}
