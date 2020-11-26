package action

import "github.com/TencentAd/attribution/attribution/pkg/handler/http/ia/data"

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
	for _, s := range c.EncryptTwiceData {
		d, err := s.Decrypt(c.GroupId())
		if err != nil {
			return err
		}
		c.FirstDecryptData = append(c.FirstDecryptData, d)
	}

	return nil
}