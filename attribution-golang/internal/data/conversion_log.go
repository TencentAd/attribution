/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 8/13/20, 3:49 PM
 */

package data

// 转化信息
type ConversionLog struct {
	*UserData
	AppId     string `json:"appid,omitempty"`      // 广告对应的推广app_id
	EventTime int64  `json:"event_time,omitempty"` // 转化事件发生的时间
}
