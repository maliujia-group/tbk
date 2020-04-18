package tbk

import "strconv"

type uatm struct {
	Client *Tbk
}

func (t *Tbk) Uatm() *uatm {
	return &uatm{Client: t}
}

type UatmTbkItem struct {
	NumIid      int    `json:"num_iid"`
	Title       string `json:"title"`
	PictURL     string `json:"pict_url"`
	SmallImages struct {
		String []string `json:"string"`
	} `json:"small_images"`
	ReservePrice      string `json:"reserve_price"`
	ZkFinalPrice      string `json:"zk_final_price"`
	UserType          int    `json:"user_type"`
	Provcity          string `json:"provcity"`
	ItemURL           string `json:"item_url"`
	ClickURL          string `json:"click_url"`
	Nick              string `json:"nick"`
	SellerID          int    `json:"seller_id"`
	Volume            int    `json:"volume"`
	TkRate            string `json:"tk_rate"`
	ZkFinalPriceWap   string `json:"zk_final_price_wap"`
	ShopTitle         string `json:"shop_title"`
	EventStartTime    string `json:"event_start_time"`
	EventEndTime      string `json:"event_end_time"`
	Type              int    `json:"type"`
	Status            int    `json:"status"`
	Category          int    `json:"category"`
	CouponClickURL    string `json:"coupon_click_url"`
	CouponEndTime     string `json:"coupon_end_time"`
	CouponInfo        string `json:"coupon_info"`
	CouponStartTime   string `json:"coupon_start_time"`
	CouponTotalCount  int    `json:"coupon_total_count"`
	CouponRemainCount int    `json:"coupon_remain_count"`
}

type TbkUatmFavoritesItemGetResponse struct {
	Response struct {
		Results struct {
			UatmTbkItem []UatmTbkItem `json:"uatm_tbk_item"`
		} `json:"results"`
		TotalResults int `json:"total_results"`
	} `json:"tbk_uatm_favorites_item_get_response"`
}

/**
 * (获取淘宝联盟选品库的宝贝信息)
 * taobao.tbk.uatm.favorites.item.get( 淘宝客-推广者-选品库宝贝信息 )
 * @line https://open.taobao.com/api.htm?docId=26619&docType=2
 */
func (u *uatm) FavoritesItemGet(adzoneID int64, favoritesID int, other ...map[string]string) (res *TbkUatmFavoritesItemGetResponse, err error) {
	param := make(map[string]string)
	if len(other) > 0 {
		param = other[0]
	}
	param["adzone_id"] = strconv.FormatInt(adzoneID, 10)
	param["favorites_id"] = strconv.Itoa(favoritesID)
	_, err = u.Client.httpPost("taobao.tbk.uatm.favorites.item.get", param, &res)
	return
}

type TbkUatmFavoritesGetResponse struct {
	Response struct {
		Results struct {
			TbkFavorites []TbkFavorites `json:"tbk_favorites"`
		} `json:"results"`
		TotalResults int `json:"total_results"`
	} `json:"tbk_uatm_favorites_get_response"`
}

type TbkFavorites struct {
	Type           int    `json:"type"`
	FavoritesID    int    `json:"favorites_id"`
	FavoritesTitle string `json:"favorites_title"`
}

/**
 * (获取淘宝联盟选品库列表)
 * taobao.tbk.uatm.favorites.get( 淘宝客-推广者-选品库宝贝列表 )
 * @line https://open.taobao.com/api.htm?docId=26620&docType=2
 */
func (u *uatm) FavoritesGet(fields string, other ...map[string]string) (res *TbkUatmFavoritesGetResponse, err error) {
	param := make(map[string]string)
	if len(other) > 0 {
		param = other[0]
	}
	param["fields"] = fields
	_, err = u.Client.httpPost("taobao.tbk.uatm.favorites.get", param, &res)
	return
}
