package kafka

import (
	"flag"
	"github.com/TencentAd/attribution/attribution/proto/conv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAmsKafkaAttributionStore(t *testing.T) {
	flag.Parse()

	store, err := NewAmsKafkaAttributionStore()
	assert.NoError(t, err)

	c := &conv.ConversionLog{
		UserData:        nil,
		EventTime:       0,
		AppId:           "test appid",
		ConvId:          "test convid",
		CampaignId:      0,
		Index:           0,
		MatchClick:      nil,
		OriginalContent: "hello world",
	}

	assert.NoError(t, store.Store(c))
}
