/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 10/10/20, 10:41 AM
 */

package httpx

import (
	"fmt"
	"net/http"
	"strconv"
)

func HttpMustQueryInt64Param(r *http.Request, key string) (int64, error) {
	values, ok := r.URL.Query()[key]
	if !ok || len(values) == 0 {
		return 0, fmt.Errorf("'%s' param not exist", key)
	}
	v, err := strconv.ParseInt(values[0], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("'%s' not valid integer", key)
	}
	return v, nil
}

func HttpQueryInt64Param(r *http.Request, key string, def int64) (int64, error) {
	values, ok := r.URL.Query()[key]
	if !ok || len(values) == 0 {
		return def, nil
	}
	v, err := strconv.ParseInt(values[0], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("'%s' not valid integer", key)
	}
	return v, nil
}

func HttpQueryInt32Param(r *http.Request, key string, def int32) (int32, error) {
	values, ok := r.URL.Query()[key]
	if !ok || len(values) == 0 {
		return def, nil
	}
	v, err := strconv.Atoi(values[0])
	if err != nil {
		return 0, fmt.Errorf("'%s' not valid integer", key)
	}
	return int32(v), nil
}

func HttpQueryStringParam(r *http.Request, key string, def string) (string, error) {
	values, ok := r.URL.Query()[key]
	if !ok || len(values) == 0 {
		return def, nil
	}
	return values[0], nil
}

func HttpMustQueryStringParam(r *http.Request, key string) (string, error) {
	values, ok := r.URL.Query()[key]
	if !ok || len(values) == 0 {
		return "", fmt.Errorf("'%s' param not exist", key)
	}
	return values[0], nil
}

