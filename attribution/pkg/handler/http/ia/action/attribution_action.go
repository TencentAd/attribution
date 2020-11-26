package action

import (
	"github.com/golang/glog"

	"github.com/TencentAd/attribution/attribution/pkg/handler/http/ia/data"
	"github.com/TencentAd/attribution/attribution/pkg/impression/kv"
)

// 曝光归因逻辑
type AttributionAction struct {
	impKV kv.KV
}

func NewAttributionAction() *AttributionAction {
	return &AttributionAction{}
}

func (action *AttributionAction) WithImpStorage(impKV kv.KV) *AttributionAction {
	action.impKV = impKV
	return action
}

func (action *AttributionAction) name() string {
	return "AttributionAction"
}

func (action *AttributionAction) Run(i interface{}) {
	ExecuteNamedAction(action, i)
}

func (action *AttributionAction) run(c *data.ImpAttributionContext) error {
	for _, fd := range c.FirstDecryptData {
		if ok, err := action.checkIntersect(c, fd); err != nil {
			glog.Errorf("failed to check intersect, %v, err: %v", fd, err)
		} else if ok  {
			c.IntersectData = append(c.IntersectData, fd)
		}
	}

	return nil
}

// TODO
// 1. 设置时间窗口
// 2. 从kv中获取campaignId
func (action *AttributionAction) checkIntersect(c *data.ImpAttributionContext, fd *data.IdSet) (bool, error) {
	ok, err := action.impKV.Has(fd.Imei)
	if err != nil {
		glog.Errorf("failed to check imei, err: %v", err)
	}
	if ok {
		return true, nil
	}

	ok, err = action.impKV.Has(fd.Idfa)
	if err != nil {
		glog.Errorf("failed to check idfa, err: %v", err)
	}
	if ok {
		return true, nil
	}

	ok, err = action.impKV.Has(fd.AndroidId)
	if err != nil {
		glog.Errorf("failed to check androidId, err: %v", err)
	}
	if ok {
		return true, nil
	}

	ok, err = action.impKV.Has(fd.Oaid)
	if err != nil {
		glog.Errorf("failed to check oaid, err: %v", err)
	}
	if ok {
		return true, nil
	}

	return false, nil
}
