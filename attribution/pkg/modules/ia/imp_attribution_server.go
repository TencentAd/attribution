package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/TencentAd/attribution/attribution/pkg/common/flagx"
	"github.com/TencentAd/attribution/attribution/pkg/common/metricutil"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/ia"
	"github.com/TencentAd/attribution/attribution/pkg/impression/handler"
	"github.com/TencentAd/attribution/attribution/pkg/impression/kv"
	"github.com/TencentAd/attribution/attribution/pkg/impression/kv/opt"
	"github.com/TencentAd/attribution/attribution/pkg/parser/ams"
)

var (
	serverAddress        = flag.String("server_address", ":80", "")
	metricsServerAddress = flag.String("metric_server_address", ":8080", "")
	impKvType            = flag.String("imp_kv_type", "LEVELDB", "")
	impKvAddress         = flag.String("imp_kv_address", "./db", "")
	impKvPassword        = flag.String("imp_kv_password", "", "")
)

func serveHttp() error {
	storage, err := kv.CreateKV(kv.StorageType(*impKvType), &opt.Option{
		Address:  *impKvAddress,
		Password: *impKvPassword,
	})
	if err != nil {
		return err
	}
	http.Handle("/impression", handler.NewSetHandler(storage))

	convParser := ams.NewUserActionAddRequestParser()
	impAttributionHandle := ia.NewImpAttributionHandle().
		WithConvParser(convParser).
		WithImpStorage(storage)
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
