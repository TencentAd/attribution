/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/9/20, 10:41 AM
 */

package logic

import (
	"github.com/TencentAd/attribution/attribution/pkg/common/key"
	"github.com/TencentAd/attribution/attribution/pkg/data/user"
	"github.com/TencentAd/attribution/attribution/pkg/storage/clickindex"
	"github.com/TencentAd/attribution/attribution/proto/click"
)

func ProcessClickLog(clickLog *click.ClickLog, index clickindex.ClickIndex) error {
	userIds, err := user.GenerateNormalIdsByPriority(clickLog.UserData)
	if err != nil {
		return err
	}

	appId := clickLog.AppId

	for _, uid := range userIds {
		if err := index.Set(uid.T, key.FormatClickLogKey(appId, uid.Id), clickLog); err != nil {
			return err
		}
	}

	return nil
}


