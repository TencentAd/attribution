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

	"github.com/TencentAd/attribution/attribution/pkg/common/httpx"
	"github.com/TencentAd/attribution/attribution/pkg/protocal/ams/conversion"
	"github.com/TencentAd/attribution/attribution/pkg/protocal/parse"
	"github.com/TencentAd/attribution/attribution/proto/conv"
	"github.com/TencentAd/attribution/attribution/proto/user"

	"github.com/golang/glog"
)

// 自归因转化上报接口
type ConvParser struct {
}

func NewConvParser() *ConvParser {
	return &ConvParser{}
}

func (p *ConvParser) Parse(data interface{}) (*parse.ConvParseResult, error) {
	r := data.(*http.Request)

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	req := &amsconversion.Request{}
	if err := json.Unmarshal(requestBody, req); err != nil {
		return nil, err
	}

	query := r.URL.Query()
	appId, err := httpx.HttpMustQueryStringParam(query, "app_id")
	if err != nil {
		return nil, err
	}
	convId, err := httpx.HttpMustQueryStringParam(query, "conv_id")
	if err != nil {
		return nil, err
	}

	convs := make([]*conv.ConversionLog, 0, len(req.Actions))
	for _, action := range req.Actions {
		// 协议转换
		convLog := &conv.ConversionLog{
			UserData: &user.UserData{
				Imei:      action.UserId.HashImei,
				Idfa:      action.UserId.HashIdfa,
				AndroidId: action.UserId.HashAndroidId,
				HashOaid:  action.UserId.HashOaid,
				Oaid:      action.UserId.Oaid,
			},
			AppId:     appId,
			EventTime: action.ActionTime,
			ConvId:    convId,
		}

		// 保留原始的信息
		tmp := &amsconversion.Request{
			Actions: []*amsconversion.Action{action},
		}
		content, err := json.Marshal(tmp)
		if err != nil {
			glog.Errorf("failed to marshal, err: %v", err)
		}
		convLog.OriginalContent = string(content)
		convs = append(convs, convLog)
	}

	return &parse.ConvParseResult{
		AppId:    appId,
		ConvId:   convId,
		ConvLogs: convs,
		Actions:  req.Actions,
	}, nil
}
