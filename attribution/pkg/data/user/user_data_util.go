/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/9/20, 10:08 AM
 */

package user

import "github.com/TencentAd/attribution/attribution/proto/user"

type Id struct {
	Id string      // id值
	T  user.IdType // id类型
}

func GenerateNormalIdsByPriority(ud *user.UserData) ([]*Id, error) {
	if ud == nil {
		return nil, nil
	}

	ids := make([]*Id, 0)
	ids = appendValid(ids, user.IdType_Device, ud.Idfa)
	ids = appendValid(ids, user.IdType_Device, ud.Imei)
	ids = appendValid(ids, user.IdType_QQOpenId, ud.QqOpenid)
	ids = appendValid(ids, user.IdType_Wuid, ud.Wuid)
	ids = appendValid(ids, user.IdType_Mac, ud.Mac)
	ids = appendValid(ids, user.IdType_AndroidId, ud.AndroidId)
	ids = appendValid(ids, user.IdType_Qadid, ud.Qadid)
	ids = appendValid(ids, user.IdType_MappedAndroidId, ud.MappedAndroidId...)
	ids = appendValid(ids, user.IdType_MappedMac, ud.MappedMac...)
	ids = appendValid(ids, user.IdType_Taid, ud.Taid)
	ids = appendValid(ids, user.IdType_Oaid, ud.Oaid)
	ids = appendValid(ids, user.IdType_HashOaid, ud.HashOaid)
	ids = appendValid(ids, user.IdType_Uuid, ud.Uuid)

	return ids, nil
}

func appendValid(ids []*Id, t user.IdType, id ...string) []*Id {
	for _, iter := range id {
		if isUserIdValid(iter) {
			ids = append(ids, &Id{Id: iter, T: t})
		}
	}
	return ids
}

func isUserIdValid(id string) bool {
	if id == "" ||  isBadId(id) {
		return false
	}

	return true
}

var (
	/**   设备号明文      |             设备号MD5
	 *    unknown       |  ad921d60486366258809553a3db49a4a
	 *    00000000      |  dd4b21e9ef71e1291183a46b913ae6f2
	 *    0             |  cfcd208495d565ef66e7dff9f98764da
	 */
	badDeviceIdMd5Set = map[string]struct{}{
		"d41d8cd98f00b204e9800998ecf8427e": {},
		"e3f5536a141811db40efd6400f1d0a4e": {},
		"5284047f4ffb4e04824a2fd1d1f0cd62": {},
		"37a6259cc0c1dae299a7866489dff0bd": {},
		"6c3e226b4d4795d518ab341b0824ec29": {},
		"dd4b21e9ef71e1291183a46b913ae6f2": {},
		"1e4a1b03d1b6cd8a174a826f76e009f4": {},
		"9f89c84a559f573636a47ff8daed0d33": {},
		"4b997118e2db4480b95d28cc741b3917": {},
		"3a6bff0799c7389f522f3847c33a468f": {},
		"ad921d60486366258809553a3db49a4a": {},
		"cfcd208495d565ef66e7dff9f98764da": {},
	}
)

func isBadId(deviceId string) bool {
	_, ok := badDeviceIdMd5Set[deviceId]
	return ok
}
