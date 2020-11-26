package action

import (
	"time"

	"github.com/golang/glog"

	"github.com/TencentAd/attribution/attribution/pkg/handler/http/ia/data"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/ia/metrics"
)

type NamedAction interface {
	name() string
	run(c *data.ImpAttributionContext) error
}

func ExecuteNamedAction(action NamedAction, i interface{}) {
	startTime := time.Now()
	var err error
	defer func() {
		metrics.CollectActionMetrics(action.name(), startTime, err)
	}()

	c := i.(*data.ImpAttributionContext)
	if err = action.run(c); err != nil {
		glog.Errorf("failed to execute %s, err: %v", action.name(), err)
		c.StopWithError(err)
	}
}
