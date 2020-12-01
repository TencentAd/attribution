package main

import (
	"flag"
	"net/http"

	"github.com/TencentAd/attribution/attribution/pkg/common/flagx"
	"github.com/TencentAd/attribution/attribution/pkg/common/metricutil"
	"github.com/TencentAd/attribution/attribution/pkg/crypto"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/decrypt"
	safeguard2 "github.com/TencentAd/attribution/attribution/pkg/handler/http/decrypt/safeguard"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/encrypt"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/encrypt/safeguard"
	"github.com/golang/glog"
)

var (
	serverAddress        = flag.String("crypto_server_address", ":9010", "")
	metricsServerAddress = flag.String("crypto_metrics_address", ":9015", "")
)

func serveHttp() error {
	if err := crypto.InitCrypto(); err != nil {
		glog.Errorf("failed to init crypto, err: %v", err)
		return err
	}

	encryptSafeguard, err := safeguard.NewConvEncryptSafeguard()
	if err != nil {
		glog.Errorf("failed to create encrypt safeguard, err: %v", err)
		return err
	}
	http.Handle("/encrypt", encrypt.NewHttpHandle().WithSafeguard(encryptSafeguard))

	decryptSafeguard, err := safeguard2.NewDecryptSafeguard()
	if err != nil {
		glog.Errorf("failed to create decrypt safeguard, err: %v", err)
		return err
	}
	http.Handle("/decrypt", decrypt.NewHttpHandle().WithSafeguard(decryptSafeguard))
	glog.Info("init done")

	return http.ListenAndServe(*serverAddress, nil)
}

func main() {
	if err := flagx.Parse(); err != nil {
		panic(err)
	}
	_ = metricutil.ServeMetrics(*metricsServerAddress)
	if err := serveHttp(); err != nil {
		glog.Fatal(err)
	}
}
