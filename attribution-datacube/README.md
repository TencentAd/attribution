# 归因Batch任务

通过spark批量处理转化数据，获取转化对应的点击。



## quick start

通过local模式，快速的了解功能以及对应的输入和输出

env 

```
java: 1.8
scala: 2.11.8
maven: 3.6.3
```


```bash
cd attribution-datacube
mvn install

java -cp "conv-attribution/target/conv-attribution-1.0.0-SNAPSHOT-jar-with-dependencies.jar" com.tencent.attribution.ConversionAttribution input_path=../data/conversion.dat output_path=/tmp/attribution_output mode=local
```

输出的文件在 /tmp/attribution_output/

### input

json格式，每行一条转化数据记录，具体协议为 [conv.proto](https://git.code.oa.com/tssp/attribution/blob/master/proto/conv/conv.proto)

行数据示例
```json
{
  "user_data": {
    "idfa": "idfa3",
    "imei": "imei3"
  },
  "event_time": 1597318020,
  "app_id": "tt"
}
```

### output

同样使用[conv.proto](https://git.code.oa.com/tssp/attribution/blob/master/proto/conv/conv.proto)  
同时包含已经归因上的点击
- 假如没有归因上点击，则输出原始转化数据
- 假如归因上点击，则输出包含点击的转化数据

行数据示例，其中matchClick中的数据表示归因上的点击
```json
{
	"userData": {
		"imei": "imei1",
		"idfa": "idfa1"
	},
	"eventTime": "1597318010",
	"appId": "app_id1",
	"matchClick": {
		"userData": {
			"imei": "imei1"
		},
		"clickTime": "1597318001"
	}
}
```
