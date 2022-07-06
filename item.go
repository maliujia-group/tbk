package tbk

import (
	"strconv"
	"strings"
)

type item struct {
	Client *Tbk
}

func (t *Tbk) Item() *item {
	return &item{Client: t}
}

type SmallImages struct {
	String []string `json:"string"`
}

type TbkItemInfoGetResponse struct {
	Response struct {
		Results struct {
			NTbkItem []NTbkItem `json:"n_tbk_item"`
		} `json:"results"`
	} `json:"tbk_item_info_get_response"`
}

// 淘宝客商品
type NTbkItem struct {
	CatName                    string      `json:"cat_name"`
	NumIid                     string      `json:"num_iid"`
	Title                      string      `json:"title"`
	PictURL                    string      `json:"pict_url"`
	SmallImages                SmallImages `json:"small_images"`
	ReservePrice               string      `json:"reserve_price"`
	ZkFinalPrice               string      `json:"zk_final_price"`
	UserType                   int         `json:"user_type"`
	Provcity                   string      `json:"provcity"`
	ItemURL                    string      `json:"item_url"`
	SellerID                   int         `json:"seller_id"`
	Volume                     int         `json:"volume"`
	Nick                       string      `json:"nick"`
	CatLeafName                string      `json:"cat_leaf_name"`
	IsPrepay                   bool        `json:"is_prepay"`
	ShopDsr                    int         `json:"shop_dsr"`
	Ratesum                    int         `json:"ratesum"`
	IRfdRate                   bool        `json:"i_rfd_rate"`
	HGoodRate                  bool        `json:"h_good_rate"`
	HPayRate30                 bool        `json:"h_pay_rate30"`
	FreeShipment               bool        `json:"free_shipment"`
	MaterialLibType            string      `json:"material_lib_type"`
	PresaleDiscountFeeText     string      `json:"presale_discount_fee_text"`
	PresaleTailEndTime         int64       `json:"presale_tail_end_time"`
	PresaleTailStartTime       int64       `json:"presale_tail_start_time"`
	PresaleEndTime             int64       `json:"presale_end_time"`
	PresaleStartTime           int64       `json:"presale_start_time"`
	PresaleDeposit             string      `json:"presale_deposit"`
	JuPlayEndTime              int64       `json:"ju_play_end_time"`
	JuPlayStartTime            int64       `json:"ju_play_start_time"`
	PlayInfo                   string      `json:"play_info"`
	TmallPlayActivityEndTime   int64       `json:"tmall_play_activity_end_time"`
	TmallPlayActivityStartTime int64       `json:"tmall_play_activity_start_time"`
	JuOnlineStartTime          string      `json:"ju_online_start_time"`
	JuOnlineEndTime            string      `json:"ju_online_end_time"`
	JuPreShowStartTime         string      `json:"ju_pre_show_start_time"`
	JuPreShowEndTime           string      `json:"ju_pre_show_end_time"`
}

/**
 * ( 淘宝客-公用-淘宝客商品详情查询(简版) )
 * taobao.tbk.item.info.get
 * @line https://open.taobao.com/api.htm?docId=24518&docType=2
 */
func (it *item) InfoGet(NumIids string, other ...map[string]string) (res *TbkItemInfoGetResponse, err error) {
	params := make(map[string]string)
	if len(other) > 0 {
		params = other[0]
	}
	params["num_iids"] = NumIids
	_, err = it.Client.httpPost("taobao.tbk.item.info.get", params, &res)
	return
}

type TbkItemConvertResponse struct {
	Response struct {
		Results struct {
			NTbkItem []struct {
				NumIid   string `json:"num_iid"`
				ClickUrl string `json:"click_url"`
			} `json:"n_tbk_item"`
		} `json:"results"`
	} `json:"tbk_item_convert_response"`
}

/**
 * ( 淘宝客-推广者-商品链接转换 )
 * taobao.tbk.item.convert
 * @line https://open.taobao.com/api.htm?docId=24516&docType=2
 */
func (it *item) Convert(fields, numIids string, adzoneId int64, other ...map[string]string) (res *TbkItemConvertResponse, err error) {
	params := make(map[string]string)
	if len(other) > 0 {
		params = other[0]
	}
	params["fields"] = fields
	params["num_iids"] = numIids
	params["adzone_id"] = strconv.FormatInt(adzoneId, 10)
	_, err = it.Client.httpPost("taobao.tbk.item.convert", params, &res)
	return
}

type TbkItemRecommendGetResponse struct {
	Response struct {
		Results struct {
			NTbkItem []struct {
				NumIid       string      `json:"num_iid"`
				Title        string      `json:"title"`
				PictUrl      string      `json:"pict_url"`
				SmallImages  SmallImages `json:"small_images"`
				ReservePrice string      `json:"reserve_price"`
				ZkFinalPrice string      `json:"zk_final_price"`
				UserType     int         `json:"user_type"`
				Provcity     string      `json:"provcity"`
				ItemUrl      string      `json:"item_url"`
				Nick         string      `json:"nick"`
				SellerId     int         `json:"seller_id"`
				Volume       int         `json:"volume"`
			} `json:"n_tbk_item"`
		} `json:"results"`
	} `json:"tbk_item_recommend_get_response"`
}

/**
 * ( 淘宝客-公用-商品关联推荐 )
 * taobao.tbk.item.recommend.get
 * @line https://open.taobao.com/api.htm?docId=24517&docType=2
 */
func (it *item) RecommendGet(fields string, numIid string, other ...map[string]string) (res *TbkItemRecommendGetResponse, err error) {
	params := make(map[string]string)
	if len(other) > 0 {
		params = other[0]
	}
	params["fields"] = fields
	params["num_iid"] = numIid
	_, err = it.Client.httpPost("taobao.tbk.item.info.get", params, &res)
	return
}

type TbkItemGuessLikeResponse struct {
	Response struct {
		Results struct {
			TbkItemCoupon []TbkItemCoupon `json:"tbk_item_coupon"`
		} `json:"results"`
	} `json:"tbk_item_guess_like_response"`
}

type TbkItemCoupon struct {
	NumIid            string      `json:"num_iid"`
	Title             string      `json:"title"`
	PictURL           string      `json:"pict_url"`
	SmallImages       SmallImages `json:"small_images"`
	ReservePrice      string      `json:"reserve_price"`
	ZkFinalPrice      string      `json:"zk_final_price"`
	UserType          int         `json:"user_type"`
	Category          int         `json:"category"`
	Volume            int         `json:"volume"`
	CommissionRate    string      `json:"commission_rate"`
	ItemURL           string      `json:"item_url"`
	ShopTitle         string      `json:"shop_title"`
	ClickURL          string      `json:"click_url"`
	CouponInfo        string      `json:"coupon_info"`
	CouponAmount      string      `json:"coupon_amount"`
	CouponTotalCount  int         `json:"coupon_total_count"`
	CouponRemainCount int         `json:"coupon_remain_count"`
	CouponStartTime   string      `json:"coupon_start_time"`
	CouponEndTime     string      `json:"coupon_end_time"`
	CouponClickURL    string      `json:"coupon_click_url"`
}

/**
 * ( 淘宝客商品猜你喜欢 )
 * taobao.tbk.item.guess.like
 * @line https://open.taobao.com/api.htm?docId=29528&docType=2
 */
func (it *item) GuessLike(adzoneId int64, os, ip, ua, net string, other ...map[string]string) (res *TbkItemGuessLikeResponse, err error) {
	params := make(map[string]string)
	if len(other) > 0 {
		params = other[0]
	}
	params["ip"] = ip
	params["ua"] = ua
	params["os"] = strings.ToLower(os)
	params["net"] = strings.ToLower(net)
	params["adzone_id"] = strconv.FormatInt(adzoneId, 10)
	_, err = it.Client.httpPost("taobao.tbk.item.guess.like", params, &res)
	return
}

type TbkItemClickExtractResponse struct {
	Response struct {
		ItemID  string `json:"item_id"`
		OpenIid string `json:"open_iid"`
	} `json:"tbk_item_click_extract_response"`
}

/**
 * ( 淘宝客-公用-链接解析出商品id )
 * taobao.tbk.item.click.extract
 * @line https://open.taobao.com/api.htm?docId=28156&docType=2&scopeId=12316
 */
func (it *item) ClickExtract(url string) (res *TbkItemClickExtractResponse, err error) {
	_, err = it.Client.httpPost("taobao.tbk.item.click.extract", map[string]string{"click_url": url}, &res)
	return
}
