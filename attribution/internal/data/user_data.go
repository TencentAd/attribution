/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 8/12/20, 4:55 PM
 */

package data

import "errors"

// 用户相关的信息，主要包含标识用户的id
type UserData struct {
	Imei            string   `json:"imei,omitempty"`              // android imei md5
	Idfa            string   `json:"idfa,omitempty"`              // ios idfa md5
	QQOpenId        string   `json:"qq_openid,omitempty"`         // QQ openid
	Wuid            string   `json:"wuid,omitempty"`              // wuid
	Mac             string   `json:"mac,omitempty"`               // mac
	AndroidId       string   `json:"android_id,omitempty"`        // android Id
	Qadid           string   `json:"qadid,omitempty"`             // qadid
	MappedAndroidId []string `json:"mapped_android_id,omitempty"` // 用多版本mappping到的androidid
	MappedMac       []string `json:"mapped_mac,omitempty"`        // 用多版本Mappping到的mac
	Taid            string   `json:"taid,omitempty"`              // taid
	Oaid            string   `json:"oaid,omitempty"`              // oaid
	HashOaid        string   `json:"hash_oaid,omitempty"`         // hash oaid
	Uuid            string   `json:"uuid,omitempty"`              // uuid
}

type IDType int

const (
	// -------------- begin -------------
	KRetrievalId IDType = iota
	KDevice
	KMQQOpenId
	KWuid
	KMac
	KAndroidId
	KQadid
	KMappedAndroidId
	KMappedMac
	KTaid
	KOaid
	KHashOaid
	KUuid
	// -----------------end --------------

	IDTypeMax
)

type UserId struct {
	Id string // id值
	T  IDType // id类型
}

// 按照匹配的优先级返回用户id信息
func (ud *UserData) GenerateIdsByPriority() ([]*UserId, error) {
	return ud.generateNormalIdsByPriority()
}

func (ud *UserData) generateCPSIdsByPriority() ([]*UserId, error) {
	// not support
	return nil, errors.New("not support cps")
}

func (ud *UserData) generateNormalIdsByPriority() ([]*UserId, error) {
	ids := make([]*UserId, 0)
	ids = appendValid(ids, KDevice, ud.Idfa)
	ids = appendValid(ids, KDevice, ud.Imei)
	ids = appendValid(ids, KMQQOpenId, ud.QQOpenId)
	ids = appendValid(ids, KWuid, ud.Wuid)
	ids = appendValid(ids, KMac, ud.Mac)
	ids = appendValid(ids, KAndroidId, ud.AndroidId)
	ids = appendValid(ids, KQadid, ud.Qadid)
	ids = appendValid(ids, KMappedAndroidId, ud.MappedAndroidId...)
	ids = appendValid(ids, KMappedMac, ud.MappedMac...)
	ids = appendValid(ids, KTaid, ud.Taid)
	ids = appendValid(ids, KOaid, ud.Oaid)
	ids = appendValid(ids, KHashOaid, ud.HashOaid)
	ids = appendValid(ids, KUuid, ud.Uuid)
	return ids, nil
}

func appendValid(ids []*UserId, t IDType, id ...string) []*UserId {
	for _, iter := range id {
		if isUserIdValid(iter) {
			ids = append(ids, &UserId{Id: iter, T: t})
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
