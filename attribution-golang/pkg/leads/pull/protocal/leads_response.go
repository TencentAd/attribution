/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/25/20, 4:46 PM
 */

package protocal

type LeadsResponse struct {
	Code      int          `json:"code"`
	Message   string       `json:"message"`
	MessageCn string       `json:"message_cn"`
	Data      []*LeadsInfo `json:"data"`
}

type LeadsInfo struct {
	AccountId       int64  `json:"account_id"`
	ClickId         string `json:"click_id"`
	LeadsId         int64  `json:"leads_id"`
	LeadsActionTime int64  `json:"leads_action_time"`
	LeadsName       string `json:"leads_name"`
	LeadsTelephone  string `json:"leads_telephone"`
}
