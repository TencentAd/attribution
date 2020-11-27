package action

import (
	"flag"

	"github.com/TencentAd/attribution/attribution/pkg/impression/handler"
	"github.com/TencentAd/attribution/attribution/proto/impression"
	"github.com/golang/glog"

	"github.com/TencentAd/attribution/attribution/pkg/handler/http/ia/data"
	"github.com/TencentAd/attribution/attribution/pkg/impression/kv"
)

var (
	impAttributionWindowSecond = flag.Int64("imp_attribution_window_second", 86400, "")
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
		} else if ok {
			c.IntersectData = append(c.IntersectData, fd)
		}
	}

	return nil
}

func (action *AttributionAction) checkIntersect(c *data.ImpAttributionContext, fd *data.IdSet) (bool, error) {
	var resp impression.Response
	req := &impression.Request{
		CampaignId: c.CampaignIdStr,
	}
	if fd.Imei != "" {
		req.Records = append(req.Records, &impression.Record{
			IdType:  impression.IdType_encrypted_hash_imei,
			IdValue: fd.Imei,
		})
	}

	if fd.Idfa != "" {
		req.Records = append(req.Records, &impression.Record{
			IdType:  impression.IdType_encrypted_hash_idfa,
			IdValue: fd.Idfa,
		})
	}
	if fd.Oaid != "" {
		req.Records = append(req.Records, &impression.Record{
			IdType:  impression.IdType_encrypted_hash_oaid,
			IdValue: fd.Oaid,
		})
	}
	if fd.AndroidId != "" {
		req.Records = append(req.Records, &impression.Record{
			IdType:  impression.IdType_encrypted_hash_android_id,
			IdValue: fd.AndroidId,
		})
	}

	if err := handler.Query(action.impKV, req, &resp); err != nil {
		return false, err
	}

	for _, record := range resp.Records {
		if checkImpTimeValid(c.MinActionTime, c.MaxActionTime,
			int64(record.ImpressionTime), *impAttributionWindowSecond) {

			return true, nil
		}
	}

	return false, nil
}

func checkImpTimeValid(minActionTime int64, maxActionTime int64, impTime int64, window int64) bool {
	if impTime > maxActionTime {
		return false
	}

	return minActionTime <= impTime+window
}
