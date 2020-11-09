/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/5/20, 8:17 PM
 */

package protocal

type CryptoResponse struct {
	CampaignId int64
	Data       []*ResponseData `json:"data"`
}

type ResponseData struct {
	Imei      string `json:"imei"`
	Idea      string `json:"idfa"`
	AndroidId string `json:"androidId"`
	Oaid      string `json:"oaid"`
}
