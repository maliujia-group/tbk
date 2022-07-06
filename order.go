package tbk

import "time"

type order struct {
	Client *Tbk
}

func (t *Tbk) Other() *order {
	return &order{Client: t}
}

type TbkOrderDetailsGetResponse struct {
	Response struct {
		Data struct {
			Results struct {
				PublisherOrderDto []PublisherOrderDto `json:"publisher_order_dto"`
			} `json:"results"`
			HasPre        bool   `json:"has_pre"`
			PositionIndex string `json:"position_index"`
			HasNext       bool   `json:"has_next"`
			PageNo        int    `json:"page_no"`
			PageSize      int    `json:"page_size"`
		} `json:"data"`
	} `json:"tbk_order_details_get_response"`
}

type PublisherOrderDto struct {
	TbPaidTime                         string `json:"tb_paid_time"`
	TkPaidTime                         string `json:"tk_paid_time"`
	PayPrice                           string `json:"pay_price"`
	PubShareFee                        string `json:"pub_share_fee"`
	TradeID                            string `json:"trade_id"`
	TkOrderRole                        int    `json:"tk_order_role"`
	TkEarningTime                      string `json:"tk_earning_time"`
	AdzoneID                           int    `json:"adzone_id"`
	PubShareRate                       string `json:"pub_share_rate"`
	Unid                               string `json:"unid"`
	RefundTag                          int    `json:"refund_tag"`
	SubsidyRate                        string `json:"subsidy_rate"`
	TkTotalRate                        string `json:"tk_total_rate"`
	ItemCategoryName                   string `json:"item_category_name"`
	SellerNick                         string `json:"seller_nick"`
	PubID                              int    `json:"pub_id"`
	AlimamaRate                        string `json:"alimama_rate"`
	SubsidyType                        string `json:"subsidy_type"`
	ItemImg                            string `json:"item_img"`
	PubSharePreFee                     string `json:"pub_share_pre_fee"`
	AlipayTotalPrice                   string `json:"alipay_total_price"`
	ItemTitle                          string `json:"item_title"`
	SiteName                           string `json:"site_name"`
	ItemNum                            int    `json:"item_num"`
	SubsidyFee                         string `json:"subsidy_fee"`
	AlimamaShareFee                    string `json:"alimama_share_fee"`
	TradeParentID                      string `json:"trade_parent_id"`
	OrderType                          string `json:"order_type"`
	TkCreateTime                       string `json:"tk_create_time"`
	FlowSource                         string `json:"flow_source"`
	TerminalType                       string `json:"terminal_type"`
	ClickTime                          string `json:"click_time"`
	TkStatus                           int    `json:"tk_status"`
	ItemPrice                          string `json:"item_price"`
	ItemID                             string `json:"item_id"`
	AdzoneName                         string `json:"adzone_name"`
	TotalCommissionRate                string `json:"total_commission_rate"`
	ItemLink                           string `json:"item_link"`
	SiteID                             int    `json:"site_id"`
	SellerShopTitle                    string `json:"seller_shop_title"`
	IncomeRate                         string `json:"income_rate"`
	TotalCommissionFee                 string `json:"total_commission_fee"`
	TkCommissionPreFeeForMediaPlatform string `json:"tk_commission_pre_fee_for_media_platform"`
	TkCommissionFeeForMediaPlatform    string `json:"tk_commission_fee_for_media_platform"`
	TkCommissionRateForMediaPlatform   string `json:"tk_commission_rate_for_media_platform"`
	SpecialID                          int    `json:"special_id"`
	RelationID                         int    `json:"relation_id"`
	DepositPrice                       string `json:"deposit_price"`
	TbDepositTime                      string `json:"tb_deposit_time"`
	TkDepositTime                      string `json:"tk_deposit_time"`
	AlscID                             string `json:"alsc_id"`
	AlscPid                            string `json:"alsc_pid"`
}

/**
 * ( 淘宝客-推广者-所有订单查询 )
 * taobao.tbk.order.details.get
 * @line https://open.taobao.com/api.htm?docId=43328&docType=2
 */
func (o *order) DetailsGet(startTime, endTime time.Time, other ...map[string]string) (res *TbkOrderDetailsGetResponse, err error) {
	param := make(map[string]string)
	if len(other) > 0 {
		param = other[0]
	}
	param["start_time"] = startTime.Format("2006-01-02 15:04:05")
	param["end_time"] = endTime.Format("2006-01-02 15:04:05")
	_, err = o.Client.httpPost("taobao.tbk.order.details.get", param, &res)
	return
}
