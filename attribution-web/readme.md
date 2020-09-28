# attribution-web

## 引入

npm 引入：

```
yarn install attribution-web
```

```
import AttributionWeb from 'attribution-web';
```

cdn 引入：

```html
<script src="//i.gtimg.cn/ams-web/attribution-web/attribution-web.min.js"></script>
<script>
  AttributionWeb.send({...});
</script>
```

## 具体使用方法

```js
import AttributionWeb from "attribution-web"; // 如果用 <script> 标签引入, 则不需要 import

AttributionWeb.send({
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
