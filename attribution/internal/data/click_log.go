/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 8/12/20, 4:54 PM
 */

package data

import "encoding/json"

// 点击日志信息
type ClickLog struct {
	*UserData
	AppId     string `json:"app_id,omitempty"`     // 广告对应的推广app_id
	ClickTime int64  `json:"click_time,omitempty"` // 点击时间
}

func ParseClickLogFromJson(buf []byte) (*ClickLog, error) {
	var log ClickLog
	if err := json.Unmarshal(buf, &log); err != nil {
		return nil, err
	}
	return &log, nil
}
