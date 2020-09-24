/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/9/20, 11:24 AM
 */

package response

type ClickHttpResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
