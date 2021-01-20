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

	Address   = flag.String("kafka_address", ":9810", "kafka address, split with comma")
	Topic     = flag.String("attribution_kafka_topic", "attribution_test", "")
	Partition = flag.Int("partition", 0, "")
)

func init() {
	prometheus.MustRegister(AmsConvKafkaStoreCount)
}

type AmsKafkaAttributionStore struct {
	connection *kafka.Conn
}

func NewAmsKafkaAttributionStore() (*AmsKafkaAttributionStore, error) {
	connection, err := kafka.DialLeader(context.Background(), "tcp", *Address, *Topic, *Partition)
	if err != nil {
		return nil, err
	}
	return &AmsKafkaAttributionStore{connection: connection}, nil
}

func (a *AmsKafkaAttributionStore) Store(conv *conv.ConversionLog) error {
	if conv.MatchClick == nil {
		AmsConvKafkaStoreCount.WithLabelValues(conv.ConvId, "no_match").Add(1)
		return nil
	}

	if err := a.doStore(conv); err != nil {
		AmsConvKafkaStoreCount.WithLabelValues(conv.ConvId, "fail").Add(1)
		glog.Errorf("fail to store attribution result, err: %s", err)
		return err
	}
	AmsConvKafkaStoreCount.WithLabelValues(conv.ConvId, "success").Add(1)
	return nil

}

func (a *AmsKafkaAttributionStore) doStore(conv *conv.ConversionLog) error {
	value, err := proto.Marshal(conv)
	if err != nil {
		return err
	}
	_, err = a.connection.WriteMessages(kafka.Message{Value: value})
	if err != nil {
		return err
	}
	return nil
}
