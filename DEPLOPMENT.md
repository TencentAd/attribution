# Deployment

## Install

### install by docker

```shell
docker pull attribution:latest
docker run -d -p 80:80 attribution:latest -key_manager_name=redis -default_redis_config='{"address":["127.0.0.1:6379"],"is_cluster":0}' -ams_encrypt_url=http://localhost:9010/encrypt -ams_decrypt_url=http://localhost:9010/decrypt -v=100
```

### install by helm

[charts/attribution](charts/attribution/README.md) for more info

## API

### crypto

notice: this server is host by tencent ams team on production

#### **POST** `/encrypt`

body example:

```json
{
  "campaign_id": 12345,
    "data": [
      {
        "imei": "f9efca36a3c30e1cf28170d86ecbf5e9"
      }
    ]
}
```

expect response:

```json
{
  "code":0,
  "message":"success",
  "data":
  [
    {
      "imei":"6cf3a7f758c88f230cde2816d3923920814c64eceb9130bd0c3755013cc982923f416c48abdd6b22992fbba5bae76fd38cbbb3c9e7590cecc953789210c823270101a3d34e80fa8cfdc0414ef05fa777ece40ea74cae0f176b4caa309b49fde13021162325cd241f27497a50a8e1d7c46547ed957dfadad5de65655b10f7828077b19e7fe0ef0a0247b8906c5b662b03be88b056ab5a859b8d185f1ee6ed1f3d34e4b03b246791ebf4e353d6095ae2a652e47067d5688448fb2682312353c6e95a82f2cffc344485a8394e6c45bf5fb618192d7384c4412171524e15be2df960a2a20cc49c15560c580431ee0dbd96f7e49788c2f1f2578bcd13cf17fd5a90c3",
      "idfa":"",
      "android_id":"",
      "oaid":""
    }
  ]
}
```

#### **POST** `/decrypt`

body example:

```json
TODO
```

### impression and attribution

#### GET `/impression`

url example:

```http
http://localhost/impression?campaign_id=12345&impression_time=1606448648&encrypted_hash_imei=6cf3a7f758c88f230cde2816d3923920814c64eceb9130bd0c3755013cc982923f416c48abdd6b22992fbba5bae76fd38cbbb3c9e7590cecc953789210c823270101a3d34e80fa8cfdc0414ef05fa777ece40ea74cae0f176b4caa309b49fde13021162325cd241f27497a50a8e1d7c46547ed957dfadad5de65655b10f7828077b19e7fe0ef0a0247b8906c5b662b03be88b056ab5a859b8d185f1ee6ed1f3d34e4b03b246791ebf4e353d6095ae2a652e47067d5688448fb2682312353c6e95a82f2cffc344485a8394e6c45bf5fb618192d7384c4412171524e15be2df960a2a20cc49c15560c580431ee0dbd96f7e49788c2f1f2578bcd13cf17fd5a90c3
```

#### GET `/token/set`

url example:

```http
http://localhost/token/set?token=1234
```

#### POST `/conv`

url example:

```http
http://localhost/conv?campaign_id=12345
```

body example:

```json
{
  "account_id": 123,
  "user_action_set_id": 123456,
  "actions": [
    {
      "action_time": 1606448649,
      "action_type": "DOWNLOAD_APP",
      "user_id": {
        "hash_imei": "f9efca36a3c30e1cf28170d86ecbf5e91"
      }
    },
    {
      "action_time": 1606448659,
      "action_type": "DOWNLOAD_APP",
      "user_id": {
        "hash_imei": "f9efca36a3c30e1cf28170d86ecbf5e9"
      }
    }
  ]
}
```

expect response:

```json
{
  "code":0,
  "message":"success"
}
```

## Configration

[configration](CONFIGRATION.md) for more info
