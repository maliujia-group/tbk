package tbk

import "strconv"

type PrivilegeResponse struct {
	Response struct {
		RequestId string `json:"request_id"`
		Result    struct {
			Data struct {
				CategoryId          int64  `json:"category_id"`
				CouponClickUrl      string `json:"coupon_click_url"`
				CouponEndTime       string `json:"coupon_end_time"`
				CouponInfo          string `json:"coupon_info"`
				CouponStartTime     string `json:"coupon_start_time"`
				ItemId              string `json:"item_id"`
				MaxCommissionRate   string `json:"max_commission_rate"`
				CouponTotalCount    int64  `json:"coupon_total_count"`
				CouponRemainCount   int64  `json:"coupon_remain_count"`
				MmCouponRemainCount int64  `json:"mm_coupon_remain_count"`
				MmCouponTotalCount  int64  `json:"mm_coupon_total_count"`
				MmCouponClickUrl    string `json:"mm_coupon_click_url"`
				MmCouponEndTime     string `json:"mm_coupon_end_time"`
				MmCouponStartTime   string `json:"mm_coupon_start_time"`
				MmCouponInfo        string `json:"mm_coupon_info"`
				CouponType          int    `json:"coupon_type"`
				ItemUrl             string `json:"item_url"`
				YsylClickUrl        string `json:"ysyl_click_url"`
				YsylTljFace         string `json:"ysyl_tlj_face"`
				YsylTljSendTime     string `json:"ysyl_tlj_send_time"`
				YsylCommissionRate  string `json:"ysyl_commission_rate"`
				YsylTljUseStartTime string `json:"ysyl_tlj_use_start_time"`
				YsylTljUseEndTime   string `json:"ysyl_tlj_use_end_time"`
			} `json:"data"`
		} `json:"result"`
	} `json:"tbk_privilege_get_response"`
}

type privilege struct {
	Client *Tbk
}

func (t *Tbk) Privilege() *privilege {
	return &privilege{
		Client: t,
	}
}

/**
 * (单品券高效转链API)
 * taobao.tbk.privilege.get
 * @line http://open.taobao.com/api.htm?docId=28625&docType=2
 */
func (p *privilege) Get(session string, adzoneId, siteId int64, other ...map[string]string) (res *PrivilegeResponse, err error) {
	params := make(map[string]string)
	if len(other) > 0 {
		params = other[0]
	}
	params["session"] = session
	params["adzone_id"] = strconv.FormatInt(adzoneId, 10)
	params["site_id"] = strconv.FormatInt(siteId, 10)
	_, err = p.Client.httpPost("taobao.tbk.privilege.get", params, &res)
	return
}
