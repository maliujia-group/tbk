package tbk

import "strconv"

type coupon struct {
	Client *Tbk
}

func (t *Tbk) Coupon() *coupon {
	return &coupon{Client: t}
}

type TbkCouponGetResponse struct {
	Response struct {
		Data struct {
			CouponStartFee    string `json:"coupon_start_fee"`
			CouponRemainCount int    `json:"coupon_remain_count"`
			CouponTotalCount  int    `json:"coupon_total_count"`
			CouponEndTime     string `json:"coupon_end_time"`
			CouponStartTime   string `json:"coupon_start_time"`
			CouponAmount      string `json:"coupon_amount"`
			CouponSrcScene    int    `json:"coupon_src_scene"`
			CouponType        int    `json:"coupon_type"`
			CouponActivityID  string `json:"coupon_activity_id"`
		} `json:"data"`
	} `json:"tbk_coupon_get_response"`
}

/**
 * ( 淘宝客-公用-阿里妈妈推广券详情查询 )
 * taobao.tbk.coupon.get
 * @line https://open.taobao.com/api.htm?docId=31106&docType=2
 */
func (c *coupon) Get(other ...map[string]string) (res *TbkContentGetResponse, err error) {
	param := make(map[string]string)
	if len(other) > 0 {
		param = other[0]
	}
	_, err = c.Client.httpPost("taobao.tbk.coupon.get", param, &res)
	return
}

type TbkCouponConvertResponse struct {
	Response struct {
		Result struct {
			TotalResults int `json:"total_results"`
			Results      struct {
				CategoryID          int         `json:"category_id"`
				CouponClickURL      string      `json:"coupon_click_url"`
				CouponEndTime       string      `json:"coupon_end_time"`
				CouponInfo          string      `json:"coupon_info"`
				CouponStartTime     string      `json:"coupon_start_time"`
				ItemID              MixedItemID `json:"item_id"`
				MaxCommissionRate   string      `json:"max_commission_rate"`
				CouponTotalCount    int         `json:"coupon_total_count"`
				CouponRemainCount   int         `json:"coupon_remain_count"`
				CampaignType        int         `json:"campaign_type"`
				ItemURL             string      `json:"item_url"`
				YsylClickURL        string      `json:"ysyl_click_url"`
				YsylTljFace         string      `json:"ysyl_tlj_face"`
				YsylTljSendTime     string      `json:"ysyl_tlj_send_time"`
				YsylCommissionRate  string      `json:"ysyl_commission_rate"`
				YsylTljUseStartTime string      `json:"ysyl_tlj_use_start_time"`
				YsylTljUseEndTime   string      `json:"ysyl_tlj_use_end_time"`
			} `json:"results"`
		} `json:"result"`
	} `json:"tbk_coupon_convert_response"`
}

/**
 * ( 淘宝客-推广者-单品券高效转链 )
 * taobao.tbk.coupon.convert
 * @line https://open.taobao.com/api.htm?spm=a219a.7386653.0.0.18bb669aEqdgG4&source=search&docId=29289&docType=2
 */
func (c *coupon) Convert(adzoneId int64, other ...map[string]string) (res *TbkCouponConvertResponse, err error) {
	param := make(map[string]string)
	if len(other) > 0 {
		param = other[0]
	}
	param["adzone_id"] = strconv.FormatInt(adzoneId, 10)
	_, err = c.Client.httpPost("taobao.tbk.coupon.convert", param, &res)
	return
}
