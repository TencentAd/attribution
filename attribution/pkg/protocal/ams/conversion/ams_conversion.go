/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 10/12/20, 2:18 PM
 */

package amsconversion

// 参考 转化归因 - api自归因接口
// https://developers.e.qq.com/docs/guide/conversion/api?version=1.1
type Request struct {
	Actions []*Action `json:"actions"`
}

type Action struct {
	OuterActionId string                 `json:"outer_action_id"`
	ActionTime    int64                  `json:"action_time"`
	UserId        *UserId                `json:"user_id"`
	ActionType    string                 `json:"action_type"`
	ActionParam   map[string]interface{} `json:"action_param"`
}

type UserId struct {
	HashImei      string `json:"hash_imei"`
	HashIdfa      string `json:"hash_idfa"`
	HashAndroidId string `json:"hash_android_id"`
	Oaid          string `json:"oaid"`
	HashOaid      string `json:"hash_oaid"`
}
