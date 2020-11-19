/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/25/20, 2:20 PM
 */

package oauth

import (
	"github.com/google/uuid"
)

var (
	// TODO(提供方法持续更新token)
	Token string
)

// 生成随机字符串标识，global unique
func GenNonce() string {
	id, _ := uuid.NewRandom()
	return id.String()
}

func GetToken() (string, error) {
	return Token, nil
}
