/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/3/20, 3:15 PM
 */

package safeguard

import (
	"context"
	"errors"
	"flag"
	"strconv"

	"github.com/TencentAd/attribution/attribution/pkg/common/redisx"
	"github.com/go-redis/redis/v8"
	"github.com/golang/glog"
)

var (
	impTransferCountRedisConfig  = flag.String("imp_transfer_count_redis_config", "{}", "")
	impAttributionRatioThreshold = flag.Float64("imp_attribution_threshold", 0.02, "")
	decryptCntRedisConfig        = flag.String("decrypt_cnt_redis_config", "{}", "")

	ErrDecryptSafeguardInternal = errors.New("decrypt safeguard internal error")
	ErrDecryptExceedThreshold   = errors.New("decrypt exceed frequency")
)

type DecryptSafeguard struct {
	impTransferRedisClient redis.Cmdable
	decryptCntRedisClient  redis.Cmdable
}

func NewDecryptSafeguard() (*DecryptSafeguard, error) {
	impTransferRedisClient, err := redisx.CreateRedisClientV2(*impTransferCountRedisConfig)
	if err != nil {
		return nil, err
	}

	decryptCntRedisClient, err := redisx.CreateRedisClientV2(*decryptCntRedisConfig)
	if err != nil {
		return nil, err
	}

	return &DecryptSafeguard{
		impTransferRedisClient: impTransferRedisClient,
		decryptCntRedisClient:  decryptCntRedisClient,
	}, nil
}

func (g *DecryptSafeguard) Against(campaignId int64, count int64) error {
	key := g.formatDecryptKey(campaignId)

	current, err := g.impTransferRedisClient.IncrBy(context.Background(), key, count).Result()
	if err != nil {
		glog.Errorf("failed to incr decrypt count")
		return ErrDecryptSafeguardInternal
	}

	if current > g.getThreshold(campaignId) {
		return ErrDecryptExceedThreshold
	}

	return nil
}

const ConvEncryptResourcePrefix = "decryptCnt::"

func (g *DecryptSafeguard) formatDecryptKey(campaignId int64) string {
	return ConvEncryptResourcePrefix + strconv.FormatInt(campaignId, 10)
}

const permitExtra int64 = 10000

func (g *DecryptSafeguard) getThreshold(campaignId int64) int64 {
	impTransferCnt, err := g.loadImpTransferCntFromRedis(campaignId)
	if err != nil {
		return permitExtra
	}

	return g.calculateThreshold(impTransferCnt)
}

const impTransferPrefix = "impTransfer::"

func (g *DecryptSafeguard) formatImpTransferKey(campaignId int64) string {
	return impTransferPrefix + strconv.FormatInt(campaignId, 10)
}

func (g *DecryptSafeguard) loadImpTransferCntFromRedis(campaignId int64) (int64, error) {
	cnt, err := g.impTransferRedisClient.Get(context.Background(), g.formatImpTransferKey(campaignId)).Result()
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		glog.Errorf("failed to get imp transfer count for campaignId[%d]", campaignId)
		return 0, err
	}
	impTransferCnt, err := strconv.ParseInt(cnt, 10, 64)
	if err != nil {
		return 0, err
	}
	return impTransferCnt, nil
}

func (g *DecryptSafeguard) calculateThreshold(impTransferCnt int64) int64 {
	return permitExtra + int64(float64(impTransferCnt)*(*impAttributionRatioThreshold))
}
