/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/5/20, 4:48 PM
 */

package info

import "encoding/json"

type MinuteFreqInfo struct {
	LastMinute int64 `json:"last_minute"`
	Count      int   `json:"count"`
}

func ParseMinuteFreqInfo(value []byte) (*MinuteFreqInfo, error) {
	var fi MinuteFreqInfo
	if err := json.Unmarshal(value, &fi); err != nil {
		return nil, err
	}

	return &fi, nil
}

func (fi *MinuteFreqInfo) String() (string, error) {
	data, err := json.Marshal(fi)
	return string(data), err
}
