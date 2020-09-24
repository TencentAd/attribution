/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 8/13/20, 4:26 PM
 */

package storage

import (
	"attribution/proto/click"
	"attribution/proto/user"
)

type ClickIndex interface {
	Set(idType user.IdType, key string, click *click.ClickLog) error
	Get(idtype user.IdType, key string) (*click.ClickLog, error)
	Remove(idType user.IdType, key string) error
}
