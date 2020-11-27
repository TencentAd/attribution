package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/common/flagx"
	"github.com/TencentAd/attribution/attribution/pkg/common/metricutil"
	"github.com/TencentAd/attribution/attribution/pkg/common/workflow"
	"github.com/TencentAd/attribution/attribution/pkg/crypto"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/ia"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/token"
	"github.com/TencentAd/attribution/attribution/pkg/impression/handler"
	"github.com/TencentAd/attribution/attribution/pkg/impression/kv"
	"github.com/TencentAd/attribution/attribution/pkg/impression/kv/opt"
	"github.com/TencentAd/attribution/attribution/pkg/oauth"
	"github.com/TencentAd/attribution/attribution/pkg/parser/ams"
	"github.com/TencentAd/attribution/attribution/pkg/storage/metadata"
	"github.com/golang/glog"
)

var (
	serverAddress        = flag.String("server_address", ":9020", "")
	metricsServerAddress = flag.String("metric_server_address", ":9025", "")
	impKvType            = flag.String("imp_kv_type", "LEVELDB", "")
	impKvAddress         = flag.String("imp_kv_address", "./db", "")
	impKvPassword        = flag.String("imp_kv_password", "", "")

	workerCount    = flag.Int("imp_attribution_worker_count", 50, "")
	queueSize      = flag.Int("imp_attribution_queue_size", 200, "")
	queueTimeoutMS = flag.Int("imp_attribution_queue_timeout_ms", 1000, "")
)

func serveHttp() error {
	if err := crypto.InitCrypto(); err != nil {
		glog.Errorf("failed to init crypto, err: %v", err)
		return err
	}

	storage, err := kv.CreateKV(kv.StorageType(*impKvType), &opt.Option{
		Address:  *impKvAddress,
		Password: *impKvPassword,
	})
	if err != nil {
		return err
	}

	store := metadata.GetStore(nil)
	t, err := oauth.NewToken(store)
	if err != nil {
		return err
	}

	t.FetchBackGround(context.Background())

	http.Handle("/impression", handler.NewSetHandler(storage))
	http.Handle("/token/set", token.NewHandler(t))

	convParser := ams.NewUserActionAddRequestParser()
	jq := workflow.NewDefaultJobQueue(
		&workflow.QueueOption{
			WorkerCount: *workerCount,
			QueueSize:   *queueSize,
			PushTimeout: time.Duration(*queueTimeoutMS) * time.Millisecond,
		})
	jq.Start()

	impAttributionHandle := ia.NewImpAttributionHandle().
		WithConvParser(convParser).
		WithImpStorage(storage).
		WithJobQueue(jq).
		WithToken(t)

	http.Handle("/conv", impAttributionHandle)

	return http.ListenAndServe(*serverAddress, nil)
}

func main() {
	if err := flagx.Parse(); err != nil {
		panic(err)
	}
	_ = metricutil.ServeMetrics(*metricsServerAddress)
	if err := serveHttp(); err != nil {
		log.Fatal(err)
	}
}
