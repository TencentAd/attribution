/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  yuanweishi@tencent.com
 * Last Modify: 9/16/20, 4:26 PM
 */

package storage

import (
	"attribution/proto/click"
	"attribution/proto/conv"
)

type AttributionStore interface {
	Store(conv *conv.ConversionLog, click *click.ClickLog) error
}
