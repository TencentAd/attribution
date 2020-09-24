/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/9/20, 10:41 AM
 */

package logic

import (
	"attribution/pkg/common/key"
	"attribution/pkg/data/user"
	"attribution/pkg/storage"
	"attribution/proto/click"
)

func ProcessClickLog(clickLog *click.ClickLog, index storage.ClickIndex) error {
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


