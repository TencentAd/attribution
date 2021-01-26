package kafka

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/TencentAd/attribution/attribution/proto/conv"
	"github.com/golang/protobuf/proto"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestNewAmsKafkaAttributionStore(t *testing.T) {
	flag.Parse()

	store := NewAmsKafkaAttributionStore().(*AmsKafkaAttributionStore)

	c := &conv.ConversionLog{
		UserData:   nil,
		EventTime:  0,
		AppId:      "test appid!!!",
		ConvId:     "test convid!!!",
		CampaignId: 0,
		Index:      0,
		MatchClick: &conv.MatchClick{
			ClickLog:    nil,
			MatchIdType: 1,
		},
		OriginalContent: "hello world",
	}

	assert.NoError(t, store.Store(c))
}

func testConsumer(t *testing.T) {
	topic := "flink_result_test"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9810", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	_ = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		read, err := batch.Read(b)
		if err != nil {
			assert.Error(t, err)
			break
		}
		conversionLog := &conv.ConversionLog{}
		if err := proto.Unmarshal(b[0:read], conversionLog); err != nil {
			fmt.Println(string(b[0:read]))
		} else {
			value, _ := json.Marshal(conversionLog)
			fmt.Println(string(value))
		}
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}
