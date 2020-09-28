import Attribution from "../src";

Attribution.send({
  callback: response => {
    console.log(response.code);
  },
  click_id: "xxxxxx",
  leads_name: "zhihuilai",
  leads_telephone: "18814122222",
  url: "http://9.134.67.201:9083/leads"
});
