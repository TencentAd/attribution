import "./index.d";

const post = (url: string, data: any, callback: IResponseCallback) => {
  wx.request({
    data,
    fail (error: any) {
      callback({
        code: -1,
        error,
        message: "请求异常, 详见 error 字段",
      })
    },
    header: {
      'content-type': 'text/plain' // 默认值
    },
    method: "POST",
    success(res: any) {
      if (res.statusCode === 200) {
        callback(res.data);
      } else {
        callback({
          code: -1,
          message: `${res.statusCode}: ${res.errMsg}`,
        })
      }
    },
    url,
  })
};

class Attribution {
  public static send(attributionInfo: IAttributionInfo) {
    post(
      attributionInfo.url,
      {
        click_id: attributionInfo.click_id,
        leads_action_time: Date.now(),
        leads_name: attributionInfo.leads_name,
        leads_telephone: attributionInfo.leads_telephone
      },
      attributionInfo.callback
    );
  }
}

export default Attribution;
