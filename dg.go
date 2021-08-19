package tbk

import (
	"strconv"
)

type dg struct {
	Client *Tbk
}

func (t *Tbk) Dg() *dg {
	return &dg{Client: t}
}

// 搜索响应数据结构体
type MaterialOptionalResponse struct {
	Response struct {
		TotalResults int `json:"total_results"`
		ResultList   struct {
			MapData []MaterialOptionalData `json:"map_data"`
		} `json:"result_list"`
	} `json:"tbk_dg_material_optional_response"`
}

/**
 * (通用物料搜索API（导购）)
 * taobao.tbk.dg.material.optional
 * @line http://open.taobao.com/docs/api.htm?apiId=35896
 */
func (d *dg) MaterialOptional(adzoneId int64, others ...map[string]string) (res *MaterialOptionalResponse, err error) {
	params := make(map[string]string)
	if len(others) > 0 {
		params = others[0]
	}
	params["adzone_id"] = strconv.FormatInt(adzoneId, 10)
	_, err = d.Client.httpPost("taobao.tbk.dg.material.optional", params, &res)
	return
}

type OptimusMaterialResponse struct {
	Response struct {
		ResultList struct {
			MapData []struct {
				MaterialOptionalData
				JuPlayEndTime              string `json:"ju_play_end_time"`
				JuPlayStartTime            string `json:"ju_play_start_time"`
				PlayInfo                   string `json:"play_info"`
				TmallPlayActivityEndTime   int64  `json:"tmall_play_activity_end_time"`
				TmallPlayActivityStartTime int64  `json:"tmall_play_activity_start_time"`
				JuPreShowEndTime           int64  `json:"ju_pre_show_end_time"`
				JuPreShowStartTime         int64  `json:"ju_pre_show_start_time"`
				FavoritesInfo              struct {
					TotalCount    int               `json:"total_count"`
					FavoritesList []FavoritesDetail `json:"favorites_list"`
				} `json:"favorites_info"` // 选品库信息
			} `json:"map_data"`
		} `json:"result_list"`
		IsDefault  string `json:"is_default"`
		TotalCount int    `json:"total_count"`
	} `json:"tbk_dg_optimus_material_response"`
}

// 选品库详情
type FavoritesDetail struct {
	FavoritesId    int64  `json:"favorites_id"`
	FavoritesTitle string `json:"favorites_title"`
}

/**
 * ( 淘宝客-推广者-物料精选 )
 * taobao.tbk.dg.optimus.material
 * @line https://open.taobao.com/api.htm?docId=33947&docType=2
 */
func (d *dg) OptimusMaterial(adzoneId int64, others ...map[string]string) (res *OptimusMaterialResponse, err error) {
	params := make(map[string]string)
	if len(others) > 0 {
		params = others[0]
	}
	params["adzone_id"] = strconv.FormatInt(adzoneId, 10)
	_, err = d.Client.httpPost("taobao.tbk.dg.optimus.material", params, &res)
	return
}

type vegas struct {
	Client *Tbk
}

func (d *dg) Vegas() *vegas {
	return &vegas{Client: d.Client}
}

// 淘礼金创建配置
type TljConf struct {
	AdzoneId             int64  `json:"adzone_id"`                // 广告位id
	ItemId               int64  `json:"item_id"`                  // 商品id
	TotalNum             int64  `json:"total_num"`                // 总体个数
	Name                 string `json:"name"`                     // 淘礼金名称
	UserTotalWinNumLimit int    `json:"user_total_win_num_limit"` // 单用户累计中奖次数上限
	SecuritySwitch       bool   `json:"security_switch"`          // 安全开关
	PerFace              string `json:"per_face"`                 // 单个淘礼金面额
	SendStartTime        string `json:"send_start_time"`          // 发放开始时间
	SendEndTime          string `json:"send_end_time"`            // 发放截止时间
	UseEndTime           string `json:"use_end_time"`             // 使用结束日期
	UseEndTimeMode       int    `json:"use_end_time_mode"`        // 结束日期模式，1：相对时间 2：绝对时间
	UseStartTime         string `json:"use_start_time"`           // 开始时间
}

// 淘礼金创建完成响应
type TljResponse struct {
	Response struct {
		Result struct {
			Model struct {
				RightsId  string `json:"rights_id"`
				SendUrl   string `json:"send_url"`
				VegasCode string `json:"vegas_code"`
			} `json:"model"`
			MsgCode string `json:"msg_code"`
			MsgInfo string `json:"msg_info"`
			Success bool   `json:"success"`
		} `json:"result"`
	} `json:"tbk_dg_vegas_tlj_create_response"`
}

/**
 * (淘宝客-推广者-淘礼金创建)
 * taobao.tbk.dg.vegas.tlj.create
 * @line https://open.taobao.com/api.htm?docId=40173&docType=2
 */
func (v *vegas) TljCreate(conf *TljConf) (res *TljResponse, err error) {
	_, err = v.Client.httpPost("taobao.tbk.dg.vegas.tlj.create", v.Client.Struct2MapString(conf), &res)
	return
}

type TljReportResponse struct {
	Response struct {
		Result struct {
			MsgCode string `json:"msg_code"`
			MsgInfo string `json:"msg_info"`
			Success bool   `json:"success"`
			Model   struct {
				UnfreezeAmount      float64 `json:"unfreeze_amount"`
				UnfreezeNum         int64   `json:"unfreeze_num"`
				RefundAmount        float64 `json:"refund_amount"`
				RefundNum           int64   `json:"refund_num"`
				AlipayAmount        float64 `json:"alipay_amount"`
				UseAmount           float64 `json:"use_amount"`
				UseNum              int64   `json:"use_num"`
				WinAmount           float64 `json:"win_amount"`
				WinNum              int64   `json:"win_num"`
				PerCommissionAmount float64 `json:"per_commission_amount"`
				FpRefundAmount      float64 `json:"fp_refund_amount"`
				FpRefundNum         int64   `json:"fp_refund_num"`
			} `json:"model"`
		} `json:"result"`
	} `json:"tbk_dg_vegas_tlj_instance_report_response"`
}

/**
 * ( 淘宝客-推广者-淘礼金发放及使用报表 )
 * taobao.tbk.dg.vegas.tlj.instance.report
 * @line https://open.taobao.com/api.htm?docId=43317&docType=2&scopeId=15029
 */
func (v *vegas) TljInstanceReport(RightsId string) (res *TljReportResponse, err error) {
	_, err = v.Client.httpPost("taobao.tbk.dg.vegas.tlj.instance.report", map[string]string{
		"rights_id": RightsId,
	}, &res)
	return
}

type LbTljConf struct {
	SecurityLevel        int    `json:"security_level"`           //安全等级
	UseStartTime         string `json:"use_start_time"`           // 使用开始日期
	UseEndTime           string `json:"use_end_time"`             // 使用结束日期
	UseEndTimeMode       int    `json:"use_end_time_mode"`        // 结束日期的模式 1:相对时间 2:绝对时间
	AcceptStartTime      string `json:"accept_start_time"`        // 裂变任务领取开始时间
	AcceptEndTime        string `json:"accept_end_time"`          // 裂变任务领取截止时间
	RightsPerFace        string `json:"rights_per_face"`          // 单个淘礼金面额，支持两位小数，单位元
	SecuritySwitch       bool   `json:"security_switch"`          // 安全开关 true启用 false不启用
	UserTotalWinNumLimit int    `json:"user_total_win_num_limit"` // 单用户累计中奖次数上限
	Name                 string `json:"name"`                     // 淘礼金名称
	RightsNum            int    `json:"rights_num"`               // 淘礼金总个数
	ItemId               string `json:"item_id"`                  // 商品ID
	CampaignType         string `json:"campaign_type"`
	TaskRightNum         int    `json:"task_right_num"`       // 裂变淘礼金总个数
	TaskRightsPerFace    string `json:"task_rights_per_face"` // 裂变单个淘礼金面额
	InviteNum            int    `json:"invite_num"`           // 裂变淘礼金邀请人数
	InviteTimeLimit      int    `json:"invite_time_limit"`    // 裂变淘礼金邀请时长，单位分钟，最大120分钟
	AdzoneId             int    `json:"adzone_id"`            // 推广位ID
}

type LbTljResponse struct {
	Response struct {
		Result struct {
			Model struct {
				RightsID     string `json:"rights_id"`      // 直接领取淘礼金ID 即小额淘礼金
				SendURL      string `json:"send_url"`       // 发放地址
				TaskRightsID string `json:"task_rights_id"` // 裂变淘礼金ID 即大额淘礼金
				TaskID       string `json:"task_id"`        // 裂变任务ID
			} `json:"model"`
			MsgCode string `json:"msg_code"`
			MsgInfo string `json:"msg_info"`
			Success bool   `json:"success"`
		} `json:"result"`
	} `json:"tbk_dg_vegas_lbtlj_create_response"`
}

/**
 * ( 淘宝客-推广者-裂变淘礼金创建 )
 * taobao.tbk.dg.vegas.lbtlj.create
 * @line https://open.taobao.com/api.htm?docId=57710&docType=2&scopeId=23995
 */
func (v *vegas) LbTljCreate(conf *LbTljConf) (res *LbTljResponse, err error) {
	_, err = v.Client.httpPost("taobao.tbk.dg.vegas.lbtlj.create", v.Client.Struct2MapString(conf), &res)
	return
}

type newUser struct {
	Client *Tbk
}

func (d *dg) NewUser() *newUser {
	return &newUser{Client: d.Client}
}

type OrderGetResponse struct {
	Response struct {
		Result struct {
			Data struct {
				Result   []OrderGetData `json:"result"`
				PageNo   int            `json:"page_no"`
				PageSize int            `json:"page_size"`
				HasNext  bool           `json:"has_next"`
			} `json:"data"`
		} `json:"result"`
	} `json:"tbk_dg_newuser_order_get_response"`
}

type OrderGetData struct {
	RegisterTime    string      `json:"register_time"`
	BindTime        string      `json:"bind_time"`
	BuyTime         string      `json:"buy_time"`
	Status          int         `json:"status"`
	Mobile          string      `json:"mobile"`
	OrderTkType     int         `json:"order_tk_type"` // 订单淘客类型：1.淘客订单 2.非淘客订单，仅淘宝天猫拉新适用
	UnionId         string      `json:"union_id"`
	MemberId        int         `json:"member_id"`
	MemberNick      string      `json:"member_nick"`
	SiteId          int64       `json:"site_id"`
	SiteName        string      `json:"site_name"`
	AdzoneId        int64       `json:"adzone_id"`
	AdzoneName      string      `json:"adzone_name"`
	TbTradeParentId int64       `json:"tb_trade_parent_id"`
	AcceptTime      string      `json:"accept_time"`
	ReceiveTime     string      `json:"receive_time"`
	SuccessTime     string      `json:"success_time"`
	ActivityType    string      `json:"activity_type"`
	ActivityId      string      `json:"activity_id"`
	BizDate         string      `json:"biz_date"`
	BindCardTime    string      `json:"bind_card_time"`
	LoginTime       string      `json:"login_time"`
	IsCardSave      int         `json:"is_card_save"` // 银行卡绑定状态 1.绑定 0.未绑定
	UseRightsTime   string      `json:"use_rights_time"`
	GetRightsTime   string      `json:"get_rights_time"`
	RelationId      string      `json:"relation_id"`
	Orders          []OrderData `json:"orders"`
}

type OrderData struct {
	Commission         string `json:"commission"`
	ConfirmReceiveTime string `json:"confirm_receive_time"`
	PayTime            string `json:"pay_time"`
	OrderNo            string `json:"order_no"`
}

/**
 * (淘宝客-推广者-新用户订单明细查询)
 * taobao.tbk.dg.newuser.order.get
 * @line https://open.taobao.com/api.htm?docId=33892&docType=2
 */
func (n *newUser) OrderGet(activityId string, others ...map[string]string) (res *OrderGetResponse, err error) {
	params := make(map[string]string)
	if len(others) > 0 {
		params = others[0]
	}
	params["activity_id"] = activityId
	_, err = n.Client.httpPost("taobao.tbk.dg.newuser.order.get", params, &res)
	return
}

type OrderSumResponse struct {
	Response struct {
		Results struct {
			Data struct {
				PageNo   int  `json:"page_no"`
				PageSize int  `json:"page_size"`
				HasNext  bool `json:"has_next"`
				Results  struct {
					Data []SumData `json:"data"`
				} `json:"results"`
			} `json:"data"`
		} `json:"results"`
	} `json:"tbk_dg_newuser_order_sum_response"`
}

type SumData struct {
	ActivityId           string `json:"activity_id"`
	BizDate              string `json:"biz_date"`
	RegUserCnt           int64  `json:"reg_user_cnt"`
	LoginUserCnt         int64  `json:"login_user_cnt"`
	AlipayUserCnt        int64  `json:"alipay_user_cnt"`
	RcvValidUserCnt      int64  `json:"rcv_valid_user_cnt"`
	RcvUserCnt           int64  `json:"rcv_user_cnt"`
	AlipayUserCpaPreAmt  string `json:"alipay_user_cpa_pre_amt"`
	BindBuyUserCpaPreAmt string `json:"bind_buy_user_cpa_pre_amt"`
	BindBuyValidUserCnt  int64  `json:"bind_buy_valid_user_cnt"`
	BindCardValidUserCnt int64  `json:"bind_card_valid_user_cnt"`
	ReBuyValidUserCnt    int64  `json:"re_buy_valid_user_cnt"`
	ValidNum             int64  `json:"valid_num"`
	RelationId           string `json:"relation_id"`
}

/**
 * ( 淘宝客-推广者-拉新活动对应数据查询 )
 * taobao.tbk.dg.newuser.order.sum
 * @line https://open.taobao.com/api.htm?docId=36836&docType=2
 */
func (n *newUser) OrderSum(activityId string, pageNo, pageSize int, others ...map[string]string) (res *OrderSumResponse, err error) {
	params := make(map[string]string)
	if len(others) > 0 {
		params = others[0]
	}
	params["activity_id"] = activityId
	params["page_no"] = strconv.Itoa(pageNo)
	params["page_size"] = strconv.Itoa(pageSize)
	_, err = n.Client.httpPost("taobao.tbk.dg.newuser.order.sum", params, &res)
	return
}
