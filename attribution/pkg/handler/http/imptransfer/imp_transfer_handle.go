/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/6/20, 10:42 AM
 */

package imptransfer

import "net/http"

// HttpHandle 处理曝光转发的数据，加密后再转发，暂时不需要
type HttpHandle struct {

}

func NewHttpHandle() *HttpHandle {
	return &HttpHandle{}
}

func (h *HttpHandle) ServeHTTP(w http.ResponseWriter, r *http.Request){

}
