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
	Idfa      string `json:"idfa"`
	AndroidId string `json:"android_id"`
	Oaid      string `json:"oaid"`
}
