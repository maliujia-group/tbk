package tbk

type tpwd struct {
	Client *Tbk
}

func (t *Tbk) Tpwd() *tpwd {
	return &tpwd{
		Client: t,
	}
}

type TpwdCreateResponse struct {
	Response struct {
		Data struct {
			Model          string `json:"model"`
			PasswordSimple string `json:"password_simple"`
		} `json:"data"`
	} `json:"tbk_tpwd_create_response"`
}

/**
 * (淘宝客淘口令)
 * taobao.tbk.tpwd.create
 * @line http://open.taobao.com/docs/api.htm?apiId=31127
 */
func (t *tpwd) Create(text, url string, other ...map[string]string) (res *TpwdCreateResponse, err error) {
	params := make(map[string]string)
	if len(other) > 0 {
		params = other[0]
	}
	params["text"] = text
	params["url"] = url
	_, err = t.Client.httpPost("taobao.tbk.tpwd.create", params, &res)
	return
}

type TpwQueryResponse struct {
	Response struct {
		Content     string `json:"content"`
		Title       string `json:"title"`
		Price       string `json:"price"`
		PicUrl      string `json:"pic_url"`
		Suc         bool   `json:"suc"`
		Url         string `json:"url"`
		NativeUrl   string `json:"native_url"`
		ThumbPicUrl string `json:"thumb_pic_url"`
	} `json:"wireless_share_tpwd_query_response"`
}

/**
 * ( 查询解析淘口令 )
 * taobao.wireless.share.tpwd.query
 * @line http://open.taobao.com/api.htm?docId=32461&docType=2
 */
func (t *tpwd) Query(passwordContent string) (res *TpwQueryResponse, err error) {
	_, err = t.Client.httpPost("taobao.wireless.share.tpwd.query", map[string]string{
		"password_content": passwordContent,
	}, &res)
	return
}

type TpwdConvertResponse struct {
	Response struct {
		Data struct {
			NumIid    MixedItemID `json:"num_iid"`
			ClickUrl  string      `json:"click_url"`
			SellerId  string      `json:"seller_id"`
			OriginUrl string      `json:"origin_url"`
			OriginPid string      `json:"origin_pid"`
		} `json:"data"`
	} `json:"tbk_tpwd_convert_response"`
}

/**
 * (淘宝客-推广者-淘口令解析&转链 )
 * taobao.tbk.tpwd.convert
 * @line https://aff-open.taobao.com/#!/document
 */
func (t *tpwd) Convert(passwordContent string, adzoneId string, dx string) (res *TpwdConvertResponse, err error) {
	_, err = t.Client.httpPost("taobao.tbk.tpwd.convert", map[string]string{
		"password_content": passwordContent,
		"adzone_id":        adzoneId,
		"dx":               dx,
	}, &res)
	return
}
