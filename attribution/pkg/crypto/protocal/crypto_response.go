/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/5/20, 8:17 PM
 */

package protocal

type CryptoResponse struct {
	Code       int             `json:"code"`
	Message    string          `json:"message"`
	Data       []*ResponseData `json:"data,omitempty"`
}

type ResponseData struct {
	Imei      string `json:"imei"`
	Idfa      string `json:"idfa"`
	AndroidId string `json:"android_id"`
	Oaid      string `json:"oaid"`
}

func CreateErrCryptoResponse(err error) *CryptoResponse {
	return &CryptoResponse{
		Code:    -1,
		Message: err.Error(),
	}
}
