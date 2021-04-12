package kafka

import (
	"context"
	"flag"
	click2 "github.com/TencentAd/attribution/attribution/proto/click"
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
			Help:      "ams conv kafka store count",
		},
		[]string{"conv_id", "status"},
	)

	AmsClickKafkaStoreCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "attribution",
			Subsystem: "",
			Name:      "ams_click_kafka_store_count",
			Help:      "ams click kafka store count",
		},
		[]string{"click_id", "status"})

	Address          = flag.String("kafka_address", "9.134.188.241:9093", "kafka address, split with comma")
	AttributionTopic = flag.String("attribution_kafka_topic", "conversion_test", "")
	ClickTopic       = flag.String("click_kafka_topic", "click_test", "")
	ConversionTopic  = flag.String("conversion_kafka_topic", "conversion_test", "")
)

func init() {
	prometheus.MustRegister(AmsConvKafkaStoreCount)
}

type AmsStoreInterface interface {
	Store(message interface{}) error
}

type AmsKafkaStore struct {
	writer *kafka.Writer
}

type AmsKafkaClickStore struct {
	store *AmsKafkaStore
}

type AmsKafkaAttributionStore struct {
	store *AmsKafkaStore
}

type AmsKafkaConversionStore struct {
	store *AmsKafkaStore
}

func NewAmsKafkaAttributionStore() interface{} {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(*Address),
		Topic:    *AttributionTopic,
		Balancer: &kafka.Hash{},
	}

	return &AmsKafkaAttributionStore{store: &AmsKafkaStore{writer: writer}}
}

func NewAmsKafkaClickStore() interface{} {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(*Address),
		Topic:    *ClickTopic,
		Balancer: &kafka.Hash{},
	}

	return &AmsKafkaClickStore{store: &AmsKafkaStore{writer: writer}}
}

func NewAmsKafkaConversionStore() interface{} {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(*Address),
		Topic:    *ConversionTopic,
		Balancer: &kafka.Hash{},
	}

	return &AmsKafkaConversionStore{store: &AmsKafkaStore{writer: writer}}
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
	err = a.store.writer.WriteMessages(context.Background(), kafka.Message{Value: value})
	if err != nil {
		return err
	}
	return nil
}

func (a *AmsKafkaClickStore) Store(message interface{}) error {
	clickLog := message.(*click2.ClickLog)
	if err := a.doStore(clickLog); err != nil {
		AmsClickKafkaStoreCount.WithLabelValues(clickLog.GetClickId(), "fail").Add(1)
		glog.Errorf("fail to store click result, err: %s", err)
		return err
	}
	AmsClickKafkaStoreCount.WithLabelValues(clickLog.ClickId, "success").Add(1)
	return nil
}

func (a *AmsKafkaClickStore) doStore(click *click2.ClickLog) error {
	value, err := proto.Marshal(click)
	if err != nil {
		return err
	}
	err = a.store.writer.WriteMessages(context.Background(), kafka.Message{Value: value})
	if err != nil {
		return err
	}
	return nil
}
