/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 10/12/20, 3:55 PM
 */

package http

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/TencentAd/attribution/attribution/proto/conv"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	AmsConvReportCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   "attribution",
			Subsystem:   "",
			Name:        "ams_conv_report_count",
			Help:        "ams conv report count",
		},
		[]string{"conv_id", "status"},
		)

	amsConvServerUrl = flag.String("ams_conv_server_url",
		"http://tracking.e.qq.com/conv?cb=__CALLBACK__&conv_id=__CONV_ID__", "")
)

func init() {
	prometheus.MustRegister(AmsConvReportCount)
}

type AmsAttributionForward struct {
	client *http.Client
}

func NewAmsAttributionForward() interface{} {
	return &AmsAttributionForward{
		client: &http.Client{},
	}
}

type convResponse struct {
	Code    int
	Message string
}

func (f *AmsAttributionForward) Store(convLog *conv.ConversionLog) error {
	 // 没有归因上点击
	if convLog.MatchClick == nil {
		AmsConvReportCount.WithLabelValues(convLog.ConvId, "no_match").Add(1)
		return nil
	}

	if err := f.doStore(convLog); err != nil {
		AmsConvReportCount.WithLabelValues(convLog.ConvId, "fail").Add(1)
		glog.Errorf("failed to store attribution result, err: %s", err)
		return err
	} else {
		AmsConvReportCount.WithLabelValues(convLog.ConvId, "success").Add(1)
		return nil
	}
}

func (f *AmsAttributionForward) doStore(convLog *conv.ConversionLog) error {
	matchClick := convLog.MatchClick
	url := formatConvRequestUrl(matchClick.ClickLog.Callback, convLog.ConvId)

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(convLog.OriginalContent))
	if err != nil {
		return err
	}

	resp, err := f.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("http code[%d] not valid", resp.StatusCode)
	}

	var cr convResponse
	if err := json.Unmarshal(body, &cr); err != nil {
		return err
	}

	if cr.Code != 0 {
		return fmt.Errorf("ams conv response code[%d] not valid", cr.Code)
	}

	return nil
}

func formatConvRequestUrl(callback string, convId string) string {
	tmp := strings.Replace(*amsConvServerUrl, "__CALLBACK__", callback, 1)
	tmp = strings.Replace(tmp, "__CONV_ID__", convId, 1)
	return tmp
}
