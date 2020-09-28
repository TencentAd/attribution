# attribution-miniprogram

## QuickStart

```
yarn install attribution-miniprogram
```

```js
import AttributionMiniprogram from "attribution-miniprogram";

AttributionMiniprogram.send({
  callback: response => {
    // code 为 0 表示上报成功, 非 0 则失败, response.message 中有失败原因
    console.log(response.code);
  },
  click_id: "xxxxxx", // 点击 id
  leads_name: "姓名", // 用户提交的姓名
  leads_telephone: "手机号", // 用户填写的手机号
  url: "http://9.134.xxx.xxx:9083/leads" // 自行部署的归因服务
});
```
