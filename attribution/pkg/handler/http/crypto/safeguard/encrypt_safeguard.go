/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/3/20, 3:15 PM
 */

package safeguard

import (
	"errors"
	"flag"
	"strconv"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/handler/http/crypto/freq"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/crypto/metrics"
	"github.com/golang/glog"
)

var (
	convEncryptMaxMinuteFreq = flag.Int("conv_encrypt_max_minute_freq", 10, "")

	ErrConvEncryptSafeguardInternal = errors.New("conv encrypt safeguard internal error")
	ErrConvEncryptExceedFreq        = errors.New("conv encrypt exceed frequency")
)

type ConvEncryptSafeguard struct {
	minuteFreqStorage freq.MinuteFreqStorage
}

func NewConvEncryptSafeguard() *ConvEncryptSafeguard {
	return &ConvEncryptSafeguard{}
}

func (g *ConvEncryptSafeguard) WithCounter(storage freq.MinuteFreqStorage) *ConvEncryptSafeguard {
	g.minuteFreqStorage = storage
	return g
}

func (g *ConvEncryptSafeguard) Against(opt *Parameter) error {
	cid := opt.CryptoRequest.CampaignId
	cidStr := strconv.FormatInt(cid, 10)

	freqInfo, err := g.minuteFreqStorage.Get(g.formatResourceKey(cidStr))
	if err != nil {
		glog.Errorf("failed to get freq info, err: %v", err)
		metrics.ConvEncryptSafeguardErrCount.WithLabelValues("type", "get_freq").Add(1)
		return ErrConvEncryptSafeguardInternal
	}

	curMinute := time.Now().Unix() / 60
	if curMinute <= freqInfo.LastMinute {
		freqInfo.LastMinute = curMinute
		freqInfo.Count++
	} else {
		freqInfo.LastMinute = curMinute
		freqInfo.Count = 1
	}

	err = g.minuteFreqStorage.Set(g.formatResourceKey(cidStr), freqInfo)
	if err != nil {
		glog.Errorf("failed to set freq info, err: %v", err)
		metrics.ConvEncryptSafeguardErrCount.WithLabelValues("type", "set_freq").Add(1)
		return ErrConvEncryptSafeguardInternal
	}

	if freqInfo.Count > *convEncryptMaxMinuteFreq {
		metrics.ConvEncryptSafeguardErrCount.WithLabelValues("type", "exceed_freq").Add(1)
		return ErrConvEncryptExceedFreq
	}
	return nil
}

const ConvEncryptResourcePrefix = "conv_encrypt::"

func (g *ConvEncryptSafeguard) formatResourceKey(cidStr string) string {
	return ConvEncryptResourcePrefix + cidStr
}
