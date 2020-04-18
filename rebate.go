package tbk

import (
	"strconv"
	"time"
)

type rebate struct {
	Client *Tbk
}

func (t *Tbk) Rebate() *rebate {
	return &rebate{Client: t}
}

type TbkRebateAuthGetResponse struct {
	Response struct {
		Results struct {
			NTbkRebateAuth []struct {
				Param  string `json:"param"`
				Rebate bool   `json:"rebate"`
			} `json:"n_tbk_rebate_auth"`
		} `json:"results"`
	} `json:"tbk_rebate_auth_get_response"`
}

/**
 * ( 淘宝客-推广者-返利商家授权查询 )
 * taobao.tbk.rebate.auth.get
 * @line https://open.taobao.com/api.htm?docId=24525&docType=2
 */
func (r *rebate) AuthGet(fields, params string, searchType int) (res *TbkRebateAuthGetResponse, err error) {
	_, err = r.Client.httpPost("taobao.tbk.rebate.auth.get", map[string]string{
		"fields": fields,
		"params": params,
		"type":   strconv.Itoa(searchType),
	}, &res)
	return
}

type TbkRebateOrderGetResponse struct {
	Response struct {
		Results struct {
			NTbkOrder []struct {
				TradeParentID   int    `json:"trade_parent_id"`
				TradeID         int    `json:"trade_id"`
				NumIid          int    `json:"num_iid"`
				ItemTitle       string `json:"item_title"`
				ItemNum         int    `json:"item_num"`
				Price           string `json:"price"`
				PayPrice        string `json:"pay_price"`
				SellerNick      string `json:"seller_nick"`
				SellerShopTitle string `json:"seller_shop_title"`
				Commission      string `json:"commission"`
				CommissionRate  string `json:"commission_rate"`
				Unid            string `json:"unid"`
				CreateTime      string `json:"create_time"`
				EarningTime     string `json:"earning_time"`
			} `json:"n_tbk_order"`
		} `json:"results"`
	} `json:"tbk_rebate_order_get_response"`
}

/**
 * ( 淘宝客-推广者-返利订单查询 )
 * taobao.tbk.rebate.order.get
 * @line https://open.taobao.com/api.htm?docId=24526&docType=2
 */
func (r *rebate) OrderGet(fields string, startTime time.Time, span int, other ...map[string]string) (res *TbkRebateOrderGetResponse, err error) {
	param := make(map[string]string)
	if len(other) > 0 {
		param = other[0]
	}
	param["fields"] = fields
	param["start_time"] = startTime.Format("2006-01-02 15:04:05")
	param["span"] = strconv.Itoa(span)
	_, err = r.Client.httpPost("taobao.tbk.rebate.order.get", param, &res)
	return
}
