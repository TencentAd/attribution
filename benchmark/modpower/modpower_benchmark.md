这里测试加解密函数的性能，方便评估机器



## 运行环境

**机器**

```
cpu: 8核， 2.5GHz， Intel(R) Xeon(R) Gold 61xx CPU， 
内存： 16G
```

**输入数据**

```
data: 256bit, 随机， 100万
```

## 参数选取

| prime长度 | 对应的值                                                     |
| --------- | ------------------------------------------------------------ |
| 2048      | CC5A0467495EC2506D699C9FFCE6ED5885AE09384671EE00B93A28712DA814240F2A471B2C77B120FE70DA02D33B0F85CD737B800942D5A2BF80DCCD290FB6553C9834197DC498F6AC69B5ECEF3FD8B05F231E3632E9AECA2B3F50977CF9033AF3005A9C0A339CFB4922971B3AF05A5955983C12B153BB78A2B1FB14C84A3C662ADDE5BCEBE8779FF9F97C6E73BD29D4044242581455EEB0543E2DB35F43997F46F8596A58080DC053BBB71F9A557185DF80738238713D3EDFD77D47B26977B373FB0D969920B3909CCC24792B5B4E94AD29F6AE6BD73ED5FE6528CDDBEA1560BBCD36E8B25008021A26E9A4E51BBCD8436F38D6A222E2138E7042A73A7877D7 |
|           |                                                              |
|           |                                                              |

| encode key 长度 | 对应的值                         |
| --------------- | -------------------------------- |
| 128             | 9cff5b8e1899bb3e7a356f3eb6b45a85 |
|                 |                                  |
|                 |                                  |

## golang

**程序使用说明:**

```
./mod_power_benchmark [options]

支持的参数如下：
  -encrypt_key 加密秘钥
  -file_path   输入文件路径（每行一个id）
  -prime			 加密的prime
  -worker_number 线程数
```

**测试结果**

| 素数p的长度(bit) | encrypt key长度（bit) | threads | 花费时间(s) | QPS   | 加密后的id长度(bit) | 机器负载情况  |
| ---------------- | --------------------- | ------- | ----------- | ----- | ------------------- | ------------- |
| 2048             | 128                   | 8       | 81          | 12345 | /                   | CPU利用率90%+ |
|                  |                       |         |             |       | /                   |               |
|                  |                       |         |             |       | /                   |               |



## Java
**程序使用说明:**
```
java -jar modPowerBenchmark [src_file_path] [encrypt_data_output_path] [decrypt_data_output_path] [thread_size] [prime] [encrypt_key]
```

**测试结果：**
<table>
    <tr>
        <th>素数p的长度(bit)</th>
        <th>密钥e的长度(bit)</th>
        <th>加密花费时间(s)</th>
        <th>加密QPS</th>
        <th>解密花费时间(s)</th>
        <th>解密QPS</th>
        <th>加密后的id长度</th>
        <th>备注</th>
    </tr>
    <tr>
        <td rowspan="4">512</td>
        <td>128</td>
        <td>789</td>
        <td>12674</td>
        <td>1873</td>
        <td>5339</td>
        <td rowspan="4">87</td>
        <td>CPU的利用率只有50%左右，解密的时候CPU的利用率达到80%左右</td>
    </tr>
    <tr>
        <td>224</td>
        <td>967</td>
        <td>10341</td>
        <td>1854</td>
        <td>5394</td>
        <td>CPU的利用率只有60%左右，解密的时候CPU的利用率达到80%左右</td>
    </tr>
    <tr>
        <td>256</td>
        <td>1109</td>
        <td>9017</td>
        <td>1840</td>
        <td>5435</td>
        <td>CPU的利用率达到67%左右，解密的时候CPU的利用率达到80%左右</td>
    </tr>
    <tr>
        <td>512</td>
        <td>1750</td>
        <td>5714</td>
        <td>1846</td>
        <td>5417</td>
        <td>CPU的利用率达到80%左右，解密的时候CPU的利用率达到80%左右</td>
    </tr>
    <tr>
        <td rowspan="4">1024</td>
        <td>128</td>
        <td>1741</td>
        <td>5744</td>
        <td>10092</td>
        <td>991</td>
        <td rowspan="4">172</td>
        <td>CPU的利用率达到90%左右</td>
    </tr>
    <tr>
        <td>224</td>
        <td>2711</td>
        <td>3689</td>
        <td>10352</td>
        <td>966</td>
        <td>CPU的利用率达到90%左右</td>
    </tr>
    <tr>
        <td>256</td>
        <td>3005</td>
        <td>3328</td>
        <td>10288</td>
        <td>972</td>
        <td>CPU的利用率达到95%左右</td>
    </tr>
    <tr>
        <td>512</td>
        <td>5574</td>
        <td>1794</td>
        <td>10758</td>
        <td>930</td>
        <td>CPU的利用率接近100%</td>
    </tr>
    <tr>
        <td rowspan="4">2048</td>
        <td>128</td>
        <td>5944</td>
        <td>1682</td>
        <td>71953</td>
        <td>139</td>
        <td rowspan="4">344</td>
        <td>CPU的利用率接近100%</td>
    </tr>
    <tr>
        <td>224</td>
        <td>9170</td>
        <td>1091</td>
        <td>71854</td>
        <td>139</td>
        <td>CPU的利用率接近100%</td>
    </tr>
    <tr>
        <td>256</td>
        <td>10725</td>
        <td>932</td>
        <td>71773</td>
        <td>139</td>
        <td>CPU的利用率接近100%</td>
    </tr>
    <tr>
        <td>512</td>
        <td>19373</td>
        <td>516</td>
        <td>71628</td>
        <td>140</td>
        <td>CPU的利用率接近100%</td>
    </tr>
</table>

| 长度 | 素数                                                         |
| ---- | ------------------------------------------------------------ |
| 512  | FC41606BC9ACA092B221E2BCED7DDBEE0925C5F28EA3C58FB0EBD752D1CD6AA4BC91073E3BE17B2190F0BDC0B47AF8C5005C87E937CB5B63D1BBF830B00F7813 |
| 1024 | C65A85A99D89B1ED4E955D91777127E65E4AB45A2F13B85329794CA37A361E4257CEF53D24BBD86AB8046FFAA0ECB0CD0B4BE47941E4D6EE915F1C01267B55C844EB8DA0BD30AF66A4C15836BFD07A2397876BF1ACEA4950E8A3F06BC902ADA6C03E7FAA27865D0CB9A65DE38E5841F715E2E17E3D123E55A907229102A8C85F |
| 2048 | DBB981C695EA850A12B0DBC9515E4D4903042DB4B9F6B80966754799DE2202B76F3313E1446997507D55DAA4147415C83E7E06AC85CFDC2CDC43D5EDA1E82FB2F796E0CE694F91752E23EABB359C7575561C456FDC59D82682C43127167098E8B15539C1A7F631BDB0078E308262E36AC0140CA430D180E76545BDD92708524F64049B4DF06592C9885449FE2EB49B2B0F3BA5A18E799A358DA8D2BEEE168D1D14A91CBE43EA4D7A685F1930EBBF3AC9705921848161E131556FFC11720C7B8E24C2470838AA9BCFDDAEA03EB54DDD50E8043DBFE1351CEFAD0C1CB9B8FC54E7652AF450C4E9E49A0FD7F7CFFD7002FE50A405C5CCA0EB9B99CEC40168CBB537 |

