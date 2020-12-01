package action

import (
	"github.com/TencentAd/attribution/attribution/pkg/crypto"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/ia/data"
)

// 第一次进行加密
type EncryptAction struct {
}

func NewEncryptAction() *EncryptAction {
	return &EncryptAction{}
}

func (action *EncryptAction) name() string {
	return "EncryptAction"
}

func (action *EncryptAction) Run(i interface{}) {
	ExecuteNamedAction(action, i)
}

func (action *EncryptAction) run(c *data.ImpAttributionContext) error {
	var p crypto.Parallel
	for _, convLog := range c.ConvParseResult.ConvLogs {
		userData := convLog.UserData
		idSet := &data.IdSet{
			Imei:      userData.Imei,
			Idfa:      userData.Idfa,
			AndroidId: userData.AndroidId,
			Oaid:      userData.Oaid,
		}

		var encryptIdSet data.IdSet
		idSet.Encrypt(&p, c.GroupId(), &encryptIdSet)
		c.EncryptData = append(c.EncryptData, &encryptIdSet)
	}
	return p.WaitAndCheck()
}
