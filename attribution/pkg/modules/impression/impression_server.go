package main

import (
	"flag"
	"github.com/TencentAd/attribution/attribution/pkg/impression/handler"
	"github.com/TencentAd/attribution/attribution/pkg/impression/kv"
	"log"
	"net/http"

	"github.com/TencentAd/attribution/attribution/pkg/common/flagx"
	"github.com/TencentAd/attribution/attribution/pkg/common/metricutil"
)

var (
	serverAddress        = flag.String("server_address", ":80", "")
	metricsServerAddress = flag.String("metric_server_address", ":8080", "")
	kvType               = flag.String("kv_type", "LEVELDB", "")
	kvAddress            = flag.String("kv_address", "./db", "")
	kvPassword           = flag.String("kv_password", "", "")
)

func serveHttp() error {
	storage, err := kv.CreateKV(kv.StorageType(*kvType), &kv.Option{
		Address: *kvAddress,
		Password: *kvPassword,
	})
	if err != nil {
		return err
	}

	http.Handle("/impression", handler.NewSetHandler(storage))
	http.Handle("/impression/has", handler.NewHasHandler(storage))
	return http.ListenAndServe(*serverAddress, nil)
}

func main() {
	_ = flagx.Parse()
	_ = metricutil.ServeMetrics(*metricsServerAddress)
	if err := serveHttp(); err != nil {
		log.Fatal(err)
	}
}
