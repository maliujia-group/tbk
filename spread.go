package tbk

type spread struct {
	Client *Tbk
}

func (t *Tbk) Spread() *spread {
	return &spread{Client: t}
}

type TbkSpreadGetResponse struct {
	Response struct {
		Results struct {
			TbkSpread []struct {
				Content string `json:"content"`
				ErrMsg  string `json:"err_msg"`
			} `json:"tbk_spread"`
		} `json:"results"`
		TotalResults int `json:"total_results"`
	} `json:"tbk_spread_get_response"`
}

/**
 * (物料传播方式获取)
 * taobao.tbk.spread.get( 淘宝客-公用-长链转短链 )
 * @line https://open.taobao.com/api.htm?docId=27832&docType=2
 */
func (s *spread) Get(urls ...string) (res *TbkSpreadGetResponse, err error) {
	data := make([]struct{ Url string }, len(urls))
	for i, url := range urls {
		data[i] = struct{ Url string }{Url: url}
	}
	_, err = s.Client.httpPostJson("taobao.tbk.spread.get", "", &data, &res)
	return
}
