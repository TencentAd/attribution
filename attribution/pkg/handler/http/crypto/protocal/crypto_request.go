/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/5/20, 11:15 AM
 */

package protocal

const (
	Imp  CryptoDataType = "imp"
	Conv CryptoDataType = "conv"
)

type CryptoDataType string

type CryptoRequest struct {
	CampaignId int64          `json:"campaign_id"`
	Data       []*RequestData `json:"data"`
	DataType   CryptoDataType `json:"dataType"`
}

type RequestData struct {
	Imei      string `json:"imei"`
	Idea      string `json:"idfa"`
	AndroidId string `json:"androidId"`
	Oaid      string `json:"oaid"`
	EventTime int64  `json:"eventTime"`
}
