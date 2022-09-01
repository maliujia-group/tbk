package tbk

import "strconv"

type shop struct {
	Client *Tbk
}

func (t *Tbk) Shop() *shop {
	return &shop{Client: t}
}

type NTbkShopData struct {
	UserID     MixedItemID `json:"user_id"`
	ShopTitle  string      `json:"shop_title"`
	ShopType   string      `json:"shop_type"`
	SellerNick string      `json:"seller_nick"`
	PictURL    string      `json:"pict_url"`
	ShopURL    string      `json:"shop_url"`
}

type TbkShopGetResponse struct {
	Response struct {
		Results struct {
			NTbkShop []NTbkShopData `json:"n_tbk_shop"`
		} `json:"results"`
		TotalResults int `json:"total_results"`
	} `json:"tbk_shop_get_response"`
}

/**
 * (淘宝客店铺查询)
 * taobao.tbk.shop.get( 淘宝客-推广者-店铺搜索 )
 * @line https://open.taobao.com/api.htm?docId=24521&docType=2
 */
func (s *shop) Get(fields, queryword string, other ...map[string]string) (res *TbkShopGetResponse, err error) {
	param := make(map[string]string)
	if len(other) > 0 {
		param = other[0]
	}
	param["fields"] = fields
	param["q"] = queryword
	_, err = s.Client.httpPost("taobao.tbk.shop.get", param, &res)
	return
}

type TbkShopRecommendGetResponse struct {
	Response struct {
		Results struct {
			NTbkShop []NTbkShopData `json:"n_tbk_shop"`
		} `json:"results"`
	} `json:"tbk_shop_recommend_get_response"`
}

/**
 * (淘宝客店铺关联推荐查询)
 * taobao.tbk.shop.recommend.get( 淘宝客-公用-店铺关联推荐 )
 * @line https://open.taobao.com/api.htm?docId=24522&docType=2
 */
func (s *shop) RecommendGet(fields string, UserID int64, other ...map[string]string) (res *TbkShopRecommendGetResponse, err error) {
	param := make(map[string]string)
	if len(other) > 0 {
		param = other[0]
	}
	param["fields"] = fields
	param["user_id"] = strconv.FormatInt(UserID, 10)
	_, err = s.Client.httpPost("taobao.tbk.shop.recommend.get", param, &res)
	return
}
