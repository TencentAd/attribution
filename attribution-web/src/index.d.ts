interface IResponse {
  code: number;
  message: string;
}

type IResponseCallback = (response: IResponse) => void;

interface IAttributionInfo {
  click_id: string;
  leads_name: string;
  leads_telephone: string;
  url: string;
  callback: IResponseCallback;
}
