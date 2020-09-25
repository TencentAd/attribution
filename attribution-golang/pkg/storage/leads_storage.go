/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/25/20, 5:05 PM
 */

package storage

import "attribution/pkg/leads/pull/protocal"

// 线索存储
type LeadsStorage interface {
	Store(leads *protocal.LeadsInfo) error
}