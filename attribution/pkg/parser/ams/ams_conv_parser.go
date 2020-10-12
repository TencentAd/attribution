/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 10/10/20, 11:05 AM
 */

package ams

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"attribution/pkg/common/httpx"
	"attribution/pkg/protocal/ams/conversion"
	"attribution/proto/conv"
	"attribution/proto/user"
)

// 自归因转化上报接口
type ConvParser struct {
}

func NewConvParser() *ConvParser {
	return &ConvParser{}
}

func (p *ConvParser) Parse(data interface{}) (*conv.ConversionLog, error) {
	r := data.(*http.Request)

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	req := &amsconversion.Request{}
	if err := json.Unmarshal(requestBody, req); err != nil {
		return nil, err
	}

	appId, err := httpx.HttpMustQueryStringParam(r, "app_id")
	if err != nil {
		return nil, err
	}

	// 协议转换
	convLog := &conv.ConversionLog{
		UserData: &user.UserData{
			Imei:      req.UserId.HashImei,
			Idfa:      req.UserId.HashIdfa,
			AndroidId: req.UserId.HashAndroidId,
			HashOaid:  req.UserId.HashOaid,
			Oaid:      req.UserId.Oaid,
		},
		AppId: appId,
		EventTime: req.ActionTime,

		OriginalContent: string(requestBody),
	}

	return convLog, nil
}
