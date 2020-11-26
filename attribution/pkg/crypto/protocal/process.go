package protocal

func ProcessData(groupId string, reqData *RequestData,
	cryptoFunc func(string, string) (string, error)) (*ResponseData, error) {

	var err error
	var resp ResponseData
	resp.Imei, err = cryptoFunc(groupId, reqData.Imei)
	if err != nil {
		return nil, err
	}
	resp.Idfa, err = cryptoFunc(groupId, reqData.Idfa)
	if err != nil {
		return nil, err
	}
	resp.AndroidId, err = cryptoFunc(groupId, reqData.AndroidId)
	if err != nil {
		return nil, err
	}
	resp.Oaid, err = cryptoFunc(groupId, reqData.Oaid)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
