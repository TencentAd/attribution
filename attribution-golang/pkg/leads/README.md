## About

获取落地页的线索，目前支持两种方式
- 线索拉取
  - 对于平台生成落地页，如蹊径等，从AMS线索平台拉取
- 线索上报
  - 对于广告主自己生成的落地页，支持通过客户端的形式进行上报

## 线索拉取

官方线索拉取接口[参考](https://developers.e.qq.com/docs/api/insights/leads/lead_clues_get?version=1.3&_preview=1)

### 目的

 获取所有的线索，同时保证广告主的使用成本更低。所以协议和官方的协议保持一致，同时屏蔽分页等功能。

### 使用

#### 1. 修改线索拉取配置文件

```json
{
    "leads_config_id": 1,
    "account_id": 123456,
    "filtering": [
        {
            "field": "campaign_id",
            "operator": "EQUALS",
            "values": [
                "11111111"
            ]
        }
    ]
}
```

#### 2. 启动服务



#### 3. 确定线索条数

##### api方式

##### prometheus方式

## 线索上报

广告主在落地页调用我们的lib，上报如注册、预约等信息。