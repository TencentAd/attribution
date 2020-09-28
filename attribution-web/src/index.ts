import "./index.d";

const post = (url: string, data: any, callback: IResponseCallback) => {
  const xhr = new XMLHttpRequest();
  xhr.open("POST", url, true);
  xhr.setRequestHeader("Content-Type", "text/plain");
  xhr.send(JSON.stringify(data));
  xhr.onload = () => {
    if (xhr.status === 200) {
      callback(JSON.parse(xhr.responseText));
    } else {
      callback({
        code: -1,
        message: `status code: ${xhr.status}, ${xhr.responseText}`
      });
    }
  };
  xhr.onerror = error => {
    callback({
      code: -1,
      message: `请求异常`
    });
  };
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
