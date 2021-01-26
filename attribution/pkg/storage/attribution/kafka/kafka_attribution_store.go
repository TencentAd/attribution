package kafka

import (
	"context"
	"flag"
	"github.com/TencentAd/attribution/attribution/proto/conv"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/segmentio/kafka-go"
)

var (
	AmsConvKafkaStoreCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "attribution",
			Subsystem: "",
			Name:      "ams_conv_kafka_store_count",
			Help:      "amd conv kafka store count",
		},
		[]string{"conv_id", "status"},
	)

	Address          = flag.String("kafka_address", ":9092", "kafka address, split with comma")
	AttributionTopic = flag.String("attribution_kafka_topic", "attribution_test", "")
	ClickTopic = flag.String("attribution_kafka_topic", "click_test", "")
	ConversionTopic = flag.String("attribution_kafka_topic", "conversion_test", "")
)

func init() {
	prometheus.MustRegister(AmsConvKafkaStoreCount)
}

type AmsKafkaAttributionStore struct {
	writer *kafka.Writer
}

func NewAmsKafkaAttributionStore() interface{} {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(*Address),
		Topic:    *AttributionTopic,
		Balancer: &kafka.Hash{},
	}

	return &AmsKafkaAttributionStore{writer: writer}
}

func NewAmsKafkaClickStore() interface{} {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(*Address),
		Topic:    *ClickTopic,
		Balancer: &kafka.Hash{},
	}

	return &AmsKafkaAttributionStore{writer: writer}
}

func NewAmsKafkaConversionStore() interface{} {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(*Address),
		Topic:    *ConversionTopic,
		Balancer: &kafka.Hash{},
	}

	return &AmsKafkaAttributionStore{writer: writer}
}

func (a *AmsKafkaAttributionStore) Store(message interface{}) error {
	conversionLog := message.(*conv.ConversionLog)
	if conversionLog.MatchClick == nil {
		AmsConvKafkaStoreCount.WithLabelValues(conversionLog.ConvId, "no_match").Add(1)
		return nil
	}

	if err := a.doStore(conversionLog); err != nil {
		AmsConvKafkaStoreCount.WithLabelValues(conversionLog.ConvId, "fail").Add(1)
		glog.Errorf("fail to store attribution result, err: %s", err)
		return err
	}
	AmsConvKafkaStoreCount.WithLabelValues(conversionLog.ConvId, "success").Add(1)
	return nil

}

func (a *AmsKafkaAttributionStore) doStore(conv *conv.ConversionLog) error {
	value, err := proto.Marshal(conv)
	if err != nil {
		return err
	}
	err = a.writer.WriteMessages(context.Background(), kafka.Message{Value: value})
	if err != nil {
		return err
	}
	return nil
}
