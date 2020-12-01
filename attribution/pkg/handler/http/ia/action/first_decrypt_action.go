package action

import (
	"github.com/TencentAd/attribution/attribution/pkg/crypto"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/ia/data"
)

// 广告主第一次进行解密
type FirstDecryptAction struct {

}

func NewFirstDecryptAction() *FirstDecryptAction {
	return &FirstDecryptAction{}
}

func (action *FirstDecryptAction) name() string {
	return "FirstDecryptAction"
}

func (action *FirstDecryptAction) Run(i interface{}) {
	ExecuteNamedAction(action, i)
}

func (action *FirstDecryptAction) run(c *data.ImpAttributionContext) error {
	var p crypto.Parallel
	for _, s := range c.EncryptTwiceData {
		var d data.IdSet
		s.Decrypt(&p, c.GroupId(), &d)
		c.FirstDecryptData = append(c.FirstDecryptData, &d)
	}
	return p.WaitAndCheck()
}