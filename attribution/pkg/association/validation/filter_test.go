/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 8/13/20, 4:06 PM
 */

package validation

import (
	"testing"

	"github.com/TencentAd/attribution/attribution/proto/click"
	"github.com/TencentAd/attribution/attribution/proto/conv"

	"github.com/stretchr/testify/assert"
)

func TestDefaultClickLogValidation_Check(t *testing.T) {
	v := &DefaultClickLogValidation{}
	{
		convLog := &conv.ConversionLog{
			EventTime: 100,
		}
		clickLog := &click.ClickLog{
			ClickTime: 101,
		}
		assert.False(t, v.Check(convLog, clickLog))
	}
	{
		convLog := &conv.ConversionLog{
			EventTime: 100,
			AppId:     "a",
		}
		clickLog := &click.ClickLog{
			ClickTime: 99,
			AppId:     "b",
		}
		assert.False(t, v.Check(convLog, clickLog))
	}
	{
		convLog := &conv.ConversionLog{
			EventTime: 10000000100,
			AppId:     "a",
		}
		clickLog := &click.ClickLog{
			ClickTime: 99,
			AppId:     "a",
		}
		assert.False(t, v.Check(convLog, clickLog))
	}
	{
		convLog := &conv.ConversionLog{
			EventTime: 100,
			AppId:     "a",
		}
		clickLog := &click.ClickLog{
			ClickTime: 99,
			AppId:     "a",
		}
		assert.True(t, v.Check(convLog, clickLog))
	}
}
