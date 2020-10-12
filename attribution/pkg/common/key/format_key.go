/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 8/12/20, 8:25 PM
 */

package key

func FormatClickLogKey(appid string, uid string) string {
	return appid + "_" + uid
}
