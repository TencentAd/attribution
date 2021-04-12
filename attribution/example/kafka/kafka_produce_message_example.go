package main

import (
	"flag"
	"fmt"
	"github.com/TencentAd/attribution/attribution/pkg/storage/attribution/kafka"
	click2 "github.com/TencentAd/attribution/attribution/proto/click"
	"strconv"
	"time"
)

func main() {
	flag.Parse()

	//conversionProducer := kafka.NewAmsKafkaAttributionStore().(*kafka.AmsKafkaAttributionStore)
	clickProducer := kafka.NewAmsKafkaClickStore().(*kafka.AmsKafkaClickStore)

	index := 1

	start := time.Now().Unix()

	for index <= 10 {
		now := time.Now()
		clickLog := &click2.ClickLog{
			ClickTime: now.Unix(),
			ClickId:   strconv.Itoa(index),
			AdId:      int64(index),
			Platform:  click2.Platform(index%3 + 1),
		}

		if err := clickProducer.Store(clickLog); err != nil {
			break
		}
		//// 转化的时间是点击的时间过10秒
		//conversionLog := &conv.ConversionLog{
		//	EventTime: now.Add(10 * time.Second).Unix(),
		//	AppId:     "appId " + strconv.Itoa(index),
		//	ConvId:    "convId" + strconv.Itoa(index),
		//	Index:     int32(index),
		//	MatchClick: &conv.MatchClick{
		//		ClickLog: clickLog,
		//	},
		//}
		//if err := conversionProducer.Store(conversionLog); err != nil {
		//	fmt.Println(err)
		//	fmt.Println("send message to kafka failed")
		//	break
		//}

		time.Sleep(2 * time.Second)

		index++
	}

	//for index <= 15 {
	//	conversionLog := &conv.ConversionLog{
	//		EventTime: time.Now().Unix(),
	//		AppId:     "appId " + strconv.Itoa(index),
	//		ConvId:    "convId" + strconv.Itoa(index),
	//		Index:     int32(index),
	//		MatchClick: &conv.MatchClick{
	//			ClickLog: nil,
	//		},
	//	}
	//	if err := conversionProducer.Store(conversionLog); err != nil {
	//		fmt.Println(err)
	//		fmt.Println("send message to kafka failed")
	//		break
	//	}
	//
	//	time.Sleep(2 * time.Second)
	//
	//	index++
	//}

	end := time.Now().Unix()

	fmt.Println("cost time: ", end-start)
}
