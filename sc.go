package tbk

import (
	"strconv"
	"time"
)

type sc struct {
	Client *Tbk
}

func (t *Tbk) Sc() *sc {
	return &sc{Client: t}
}

type TbkScNewuserOrderGetResponse struct {
	Response struct {
		Results struct {
			Data struct {
				Results struct {
					MapData []ScNewUserOrderData `json:"map_data"`
				} `json:"results"`
				PageNo   int  `json:"page_no"`
				PageSize int  `json:"page_size"`
				HasNext  bool `json:"has_next"`
			} `json:"data"`
		} `json:"results"`
	} `json:"tbk_sc_newuser_order_get_response"`
}

type ScNewUserOrderData struct {
	RegisterTime    string `json:"register_time"`
	BindTime        string `json:"bind_time"`
	BuyTime         string `json:"buy_time"`
	Status          int    `json:"status"`
	Mobile          string `json:"mobile"`
	OrderTkType     int    `json:"order_tk_type"`
	UnionID         string `json:"union_id"`
	MemberID        int    `json:"member_id"`
	MemberNick      string `json:"member_nick"`
	SiteID          int    `json:"site_id"`
	SiteName        string `json:"site_name"`
	AdzoneID        int    `json:"adzone_id"`
	AdzoneName      string `json:"adzone_name"`
	TbTradeParentID int    `json:"tb_trade_parent_id"`
	AcceptTime      string `json:"accept_time"`
	ReceiveTime     string `json:"receive_time"`
	SuccessTime     string `json:"success_time"`
	ActivityType    string `json:"activity_type"`
	ActivityID      string `json:"activity_id"`
	BizDate         string `json:"biz_date"`
	Orders          struct {
		OrderData []struct {
			Commission         string `json:"commission"`
			ConfirmReceiveTime string `json:"confirm_receive_time"`
			PayTime            string `json:"pay_time"`
			OrderNo            string `json:"order_no"`
		} `json:"order_data"`
	} `json:"orders"`
	BindCardTime  string `json:"bind_card_time"`
	LoginTime     string `json:"login_time"`
	IsCardSave    int    `json:"is_card_save"`
	GetRightsTime string `json:"get_rights_time"`
	UseRightsTime string `json:"use_rights_time"`
	RelationID    string `json:"relation_id"`
}

/**
 * (淘宝客新用户订单API--社交)
 * taobao.tbk.sc.newuser.order.get( 淘宝客-服务商-新用户订单明细查询 )
 * @line https://open.taobao.com/api.htm?docId=33897&docType=2
 */
func (s *sc) NewUserOrderGet(activityId string, other ...map[string]string) (res *TbkScNewuserOrderGetResponse, err error) {
	param := make(map[string]string)
	if len(other) > 0 {
		param = other[0]
	}
	param["activity_id"] = activityId
	_, err = s.Client.httpPost("taobao.tbk.sc.newuser.order.get", param, &res)
	return
}

type TbkScNewuserOrderSumResponse struct {
	Response struct {
		Results struct {
			Data struct {
				PageNo   int  `json:"page_no"`
				PageSize int  `json:"page_size"`
				HasNext  bool `json:"has_next"`
				Results  struct {
					Data []struct {
						ActivityID           string `json:"activity_id"`
						BizDate              string `json:"biz_date"`
						RegUserCnt           int    `json:"reg_user_cnt"`
						LoginUserCnt         int    `json:"login_user_cnt"`
						AlipayUserCnt        int    `json:"alipay_user_cnt"`
						RcvValidUserCnt      int    `json:"rcv_valid_user_cnt"`
						RcvUserCnt           int    `json:"rcv_user_cnt"`
						AlipayUserCpaPreAmt  string `json:"alipay_user_cpa_pre_amt"`
						BindBuyUserCpaPreAmt string `json:"bind_buy_user_cpa_pre_amt"`
						BindBuyValidUserCnt  int    `json:"bind_buy_valid_user_cnt"`
						BindCardValidUserCnt int    `json:"bind_card_valid_user_cnt"`
						ReBuyValidUserCnt    int    `json:"re_buy_valid_user_cnt"`
						ValidNum             int    `json:"valid_num"`
						RelationID           string `json:"relation_id"`
					} `json:"data"`
				} `json:"results"`
			} `json:"data"`
		} `json:"results"`
	} `json:"tbk_sc_newuser_order_sum_response"`
}

/**
 *  (拉新活动汇总API--社交)
 * taobao.tbk.sc.newuser.order.sum( 淘宝客-服务商-拉新活动对应数据查询 )
 * @line https://open.taobao.com/api.htm?docId=36837&docType=2
 */
func (s *sc) NewUserOrderSum(activityID string, pageNo, pageSize int, other ...map[string]string) (res *TbkScNewuserOrderSumResponse, err error) {
	params := make(map[string]string)
	if len(other) > 0 {
		params = other[0]
	}
	params["activity_id"] = activityID
	params["page_size"] = strconv.Itoa(pageSize)
	params["page_no"] = strconv.Itoa(pageNo)
	_, err = s.Client.httpPost("taobao.tbk.item.info.get", params, &res)
	return
}

type TbkScMaterialOptionalResponse struct {
	Response struct {
		TotalResults int `json:"total_results"`
		ResultList   struct {
			MapData []MaterialOptionalData `json:"map_data"`
		} `json:"result_list"`
	} `json:"tbk_sc_material_optional_response"`
}

type MaterialOptionalData struct {
	CouponStartTime string      `json:"coupon_start_time"`
	CouponEndTime   string      `json:"coupon_end_time"`
	InfoDxjh        string      `json:"info_dxjh"`
	TkTotalSales    string      `json:"tk_total_sales"`
	TkTotalCommi    string      `json:"tk_total_commi"`
	CouponID        string      `json:"coupon_id"`
	NumIid          MixedItemID `json:"num_iid"`
	Title           string      `json:"title"`
	PictURL         string      `json:"pict_url"`
	SmallImages     struct {
		String []string `json:"string"`
	} `json:"small_images"`
	ReservePrice           string      `json:"reserve_price"`
	ZkFinalPrice           string      `json:"zk_final_price"`
	UserType               int         `json:"user_type"`
	Provcity               string      `json:"provcity"`
	ItemURL                string      `json:"item_url"`
	IncludeMkt             string      `json:"include_mkt"`
	IncludeDxjh            string      `json:"include_dxjh"`
	CommissionRate         string      `json:"commission_rate"`
	Volume                 int         `json:"volume"`
	SellerID               MixedItemID `json:"seller_id"`
	CouponTotalCount       int         `json:"coupon_total_count"`
	CouponRemainCount      int         `json:"coupon_remain_count"`
	CouponInfo             string      `json:"coupon_info"`
	CommissionType         string      `json:"commission_type"`
	ShopTitle              string      `json:"shop_title"`
	URL                    string      `json:"url"`
	CouponShareURL         string      `json:"coupon_share_url"`
	ShopDsr                int         `json:"shop_dsr"`
	WhiteImage             string      `json:"white_image"`
	ShortTitle             string      `json:"short_title"`
	CategoryID             int         `json:"category_id"`
	CategoryName           string      `json:"category_name"`
	LevelOneCategoryID     int         `json:"level_one_category_id"`
	LevelOneCategoryName   string      `json:"level_one_category_name"`
	Oetime                 string      `json:"oetime"`
	Ostime                 string      `json:"ostime"`
	JddNum                 int         `json:"jdd_num"`
	JddPrice               string      `json:"jdd_price"`
	UvSumPreSale           int         `json:"uv_sum_pre_sale"`
	CouponAmount           string      `json:"coupon_amount"`
	CouponStartFee         string      `json:"coupon_start_fee"`
	ItemDescription        string      `json:"item_description"`
	Nick                   string      `json:"nick"`
	XID                    string      `json:"x_id"`
	OrigPrice              string      `json:"orig_price"`
	TotalStock             int         `json:"total_stock"`
	SellNum                int         `json:"sell_num"`
	Stock                  int         `json:"stock"`
	TmallPlayActivityInfo  string      `json:"tmall_play_activity_info"`
	ItemID                 MixedItemID `json:"item_id"`
	RealPostFee            string      `json:"real_post_fee"`
	LockRateStartTime      int64       `json:"lock_rate_start_time"`
	LockRateEndTime        int64       `json:"lock_rate_end_time"`
	LockRate               string      `json:"lock_rate"`
	PresaleDiscountFeeText string      `json:"presale_discount_fee_text"`
	PresaleTailEndTime     int64       `json:"presale_tail_end_time"`
	PresaleTailStartTime   int64       `json:"presale_tail_start_time"`
	PresaleEndTime         int64       `json:"presale_end_time"`
	PresaleStartTime       int64       `json:"presale_start_time"`
	PresaleDeposit         string      `json:"presale_deposit"`
	YsylTljSendTime        string      `json:"ysyl_tlj_send_time"`
	YsylClickURL           string      `json:"ysyl_click_url"`
	YsylCommissionRate     string      `json:"ysyl_commission_rate"`
	YsylTljFace            string      `json:"ysyl_tlj_face"`
	YsylTljUseEndTime      string      `json:"ysyl_tlj_use_end_time"`
	YsylTljUseStartTime    string      `json:"ysyl_tlj_use_start_time"`
	SaleBeginTime          string      `json:"sale_begin_time"`
	SaleEndTime            string      `json:"sale_end_time"`
	Distance               string      `json:"distance"`
	UsableShopId           string      `json:"usable_shop_id"`
	UsableShopName         string      `json:"usable_shop_name"`
	SalePrice              string      `json:"sale_price"`
	KuadianPromotionInfo   string      `json:"kuadian_promotion_info"`
	SuperiorBrand          string      `json:"superior_brand"`
	RewardInfo             string      `json:"reward_info"`
	CpaRewardType          string      `json:"cpa_reward_type"`
}

/**
 * (通用物料搜索API)
 * taobao.tbk.sc.material.optional( 淘宝客-服务商-物料搜索 )
 * @line https://open.taobao.com/api.htm?docId=35263&docType=2
 */
func (s *sc) MaterialOptional(adzoneID, siteID int64, other ...map[string]string) (res *TbkScMaterialOptionalResponse, err error) {
	params := make(map[string]string)
	if len(other) > 0 {
		params = other[0]
	}
	params["adzone_id"] = strconv.FormatInt(adzoneID, 10)
	params["site_id"] = strconv.FormatInt(siteID, 10)
	_, err = s.Client.httpPost("taobao.tbk.sc.material.optional", params, &res)
	return
}

type TbkScOptimusMaterialResponse struct {
	Response struct {
		ResultList struct {
			MapData []OptimusMaterialData `json:"map_data"`
		} `json:"result_list"`
		IsDefault string `json:"is_default"`
	} `json:"tbk_sc_optimus_material_response"`
}

type OptimusMaterialData struct {
	CouponAmount int `json:"coupon_amount"`
	SmallImages  struct {
		String []string `json:"string"`
	} `json:"small_images"`
	ShopTitle            string      `json:"shop_title"`
	CategoryID           int         `json:"category_id"`
	CouponStartFee       string      `json:"coupon_start_fee"`
	ItemID               MixedItemID `json:"item_id"`
	CouponTotalCount     int         `json:"coupon_total_count"`
	UserType             int         `json:"user_type"`
	ZkFinalPrice         string      `json:"zk_final_price"`
	CouponRemainCount    int         `json:"coupon_remain_count"`
	CommissionRate       string      `json:"commission_rate"`
	CouponStartTime      string      `json:"coupon_start_time"`
	Title                string      `json:"title"`
	ItemDescription      string      `json:"item_description"`
	SellerID             int         `json:"seller_id"`
	Volume               int         `json:"volume"`
	CouponEndTime        string      `json:"coupon_end_time"`
	CouponClickURL       string      `json:"coupon_click_url"`
	PictURL              string      `json:"pict_url"`
	ClickURL             string      `json:"click_url"`
	Stock                int         `json:"stock"`
	SellNum              int         `json:"sell_num"`
	TotalStock           int         `json:"total_stock"`
	Oetime               string      `json:"oetime"`
	Ostime               string      `json:"ostime"`
	JddNum               int         `json:"jdd_num"`
	JddPrice             string      `json:"jdd_price"`
	OrigPrice            string      `json:"orig_price"`
	LevelOneCategoryName string      `json:"level_one_category_name"`
	LevelOneCategoryID   int         `json:"level_one_category_id"`
	CategoryName         string      `json:"category_name"`
	ShortTitle           string      `json:"short_title"`
	WhiteImage           string      `json:"white_image"`
	WordList             struct {
		WordMapData []struct {
			Word string `json:"word"`
			URL  string `json:"url"`
		} `json:"word_map_data"`
	} `json:"word_list"`
	TmallPlayActivityInfo      string `json:"tmall_play_activity_info"`
	UvSumPreSale               int    `json:"uv_sum_pre_sale"`
	NewUserPrice               string `json:"new_user_price"`
	CouponInfo                 string `json:"coupon_info"`
	CouponShareURL             string `json:"coupon_share_url"`
	Nick                       string `json:"nick"`
	ReservePrice               string `json:"reserve_price"`
	JuOnlineEndTime            string `json:"ju_online_end_time"`
	JuOnlineStartTime          string `json:"ju_online_start_time"`
	MaochaoPlayFreePostFee     string `json:"maochao_play_free_post_fee"`
	MaochaoPlayEndTime         string `json:"maochao_play_end_time"`
	MaochaoPlayStartTime       string `json:"maochao_play_start_time"`
	MaochaoPlayConditions      string `json:"maochao_play_conditions"`
	MaochaoPlayDiscounts       string `json:"maochao_play_discounts"`
	MaochaoPlayDiscountType    string `json:"maochao_play_discount_type"`
	MultiCouponZkRate          string `json:"multi_coupon_zk_rate"`
	PriceAfterMultiCoupon      string `json:"price_after_multi_coupon"`
	MultiCouponItemCount       string `json:"multi_coupon_item_count"`
	LockRate                   string `json:"lock_rate"`
	LockRateEndTime            int64  `json:"lock_rate_end_time"`
	LockRateStartTime          int64  `json:"lock_rate_start_time"`
	PresaleDiscountFeeText     string `json:"presale_discount_fee_text"`
	PresaleTailEndTime         int64  `json:"presale_tail_end_time"`
	PresaleTailStartTime       int64  `json:"presale_tail_start_time"`
	PresaleEndTime             int64  `json:"presale_end_time"`
	PresaleStartTime           int64  `json:"presale_start_time"`
	PresaleDeposit             string `json:"presale_deposit"`
	YsylTljSendTime            string `json:"ysyl_tlj_send_time"`
	YsylClickURL               string `json:"ysyl_click_url"`
	YsylTljFace                string `json:"ysyl_tlj_face"`
	YsylCommissionRate         string `json:"ysyl_commission_rate"`
	YsylTljUseEndTime          string `json:"ysyl_tlj_use_end_time"`
	YsylTljUseStartTime        string `json:"ysyl_tlj_use_start_time"`
	TmallPlayActivityEndTime   int64  `json:"tmall_play_activity_end_time"`
	TmallPlayActivityStartTime int64  `json:"tmall_play_activity_start_time"`
	PlayInfo                   string `json:"play_info"`
	JuPlayEndTime              int64  `json:"ju_play_end_time"`
	JuPlayStartTime            int64  `json:"ju_play_start_time"`
	PromotionType              string `json:"promotion_type"`
	PromotionInfo              string `json:"promotion_info"`
	PromotionDiscount          string `json:"promotion_discount"`
	PromotionCondition         string `json:"promotion_condition"`
	JuPreShowEndTime           string `json:"ju_pre_show_end_time"`
	JuPreShowStartTime         string `json:"ju_pre_show_start_time"`
	CpaRewardType              string `json:"cpa_reward_type"`
}

/**
 * (淘宝客擎天柱通用物料API - 社交)
 * taobao.tbk.sc.optimus.material( 淘宝客-服务商-物料精选 )
 * @line https://open.taobao.com/api.htm?docId=37884&docType=2
 */
func (s *sc) OptionalMaterial(adzoneID, siteID int64, other ...map[string]string) (res *TbkScOptimusMaterialResponse, err error) {
	params := make(map[string]string)
	if len(other) > 0 {
		params = other[0]
	}
	params["adzone_id"] = strconv.FormatInt(adzoneID, 10)
	params["site_id"] = strconv.FormatInt(siteID, 10)
	_, err = s.Client.httpPost("taobao.tbk.sc.optional.material", params, &res)
	return
}

type TbkScPublisherInfoSaveResponse struct {
	Response struct {
		Data struct {
			RelationID  int    `json:"relation_id"`
			AccountName string `json:"account_name"`
			SpecialID   int    `json:"special_id"`
			Desc        string `json:"desc"`
		} `json:"data"`
	} `json:"tbk_sc_publisher_info_save_response"`
}

/**
 * (淘宝客渠道信息备案 - 社交)
 * taobao.tbk.sc.publisher.info.save( 淘宝客-公用-私域用户备案 )
 * @line https://open.taobao.com/api.htm?docId=37988&docType=2
 */
func (s *sc) PublisherInfoSave(inviterCode string, infoType int, other ...map[string]string) (res *TbkScPublisherInfoSaveResponse, err error) {
	params := make(map[string]string)
	if len(other) > 0 {
		params = other[0]
	}
	params["inviter_code"] = inviterCode
	params["info_type"] = strconv.Itoa(infoType)
	_, err = s.Client.httpPost("taobao.tbk.sc.publisher.info.save", params, &res)
	return
}

type TbkScPublisherInfoGetResponse struct {
	Response struct {
		Data struct {
			RootPidChannelList struct {
				String []string `json:"string"`
			} `json:"root_pid_channel_list"`
			TotalCount  int `json:"total_count"`
			InviterList struct {
				MapData []PublisherInfoData `json:"map_data"`
			} `json:"inviter_list"`
		} `json:"data"`
	} `json:"tbk_sc_publisher_info_get_response"`
}

type PublisherInfoData struct {
	RelationApp  string `json:"relation_app"`
	CreateDate   string `json:"create_date"`
	AccountName  string `json:"account_name"`
	RealName     string `json:"real_name"`
	RelationID   int    `json:"relation_id"`
	OfflineScene string `json:"offline_scene"`
	OnlineScene  string `json:"online_scene"`
	Note         string `json:"note"`
	RootPid      string `json:"root_pid"`
	Rtag         string `json:"rtag"`
	OfflineInfo  struct {
		ShopName        string `json:"shop_name"`
		ShopType        string `json:"shop_type"`
		PhoneNumber     string `json:"phone_number"`
		DetailAddress   string `json:"detail_address"`
		Location        string `json:"location"`
		ShopCertifyType string `json:"shop_certify_type"`
		CertifyNumber   string `json:"certify_number"`
		Career          string `json:"career"`
	} `json:"offline_info"`
	SpecialID    int    `json:"special_id"`
	PunishStatus string `json:"punish_status"`
}

/**
 * (淘宝客信息查询 - 社交)
 * taobao.tbk.sc.publisher.info.get( 淘宝客-公用-私域用户备案信息查询 )
 * @line https://open.taobao.com/api.htm?docId=37989&docType=2
 */
func (s *sc) publisherInfoGet(infoType int, relationApp string, other ...map[string]string) (res *TbkScPublisherInfoGetResponse, err error) {
	params := make(map[string]string)
	if len(other) > 0 {
		params = other[0]
	}
	params["relation_app"] = relationApp
	params["info_type"] = strconv.Itoa(infoType)
	_, err = s.Client.httpPost("taobao.tbk.sc.publisher.info.save", params, &res)
	return
}

type TbkScInvitecodeGetResponse struct {
	Response struct {
		Data struct {
			InviterCode string `json:"inviter_code"`
		} `json:"data"`
	} `json:"tbk_sc_invitecode_get_response"`
}

/**
 * (淘宝客邀请码生成-社交)
 * taobao.tbk.sc.invitecode.get( 淘宝客-公用-私域用户邀请码生成 )
 * @line https://open.taobao.com/api.htm?docId=38046&docType=2
 */
func (s *sc) InviteCodeGet(relationApp string, codeType int, other ...map[string]string) (string, error) {
	res := TbkScInvitecodeGetResponse{}
	params := make(map[string]string)
	if len(other) > 0 {
		params = other[0]
	}
	params["relation_app"] = relationApp
	params["code_type"] = strconv.Itoa(codeType)
	_, err := s.Client.httpPost("taobao.tbk.sc.publisher.info.save", params, &res)
	if err != nil {
		return "", err
	}
	return res.Response.Data.InviterCode, nil
}

type scGroupChat struct {
	Client *Tbk
}

func (s *sc) GroupChat() *scGroupChat {
	return &scGroupChat{Client: s.Client}
}

type TbkScGroupchatMessageSendResponse struct {
	Response struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	} `json:"tbk_sc_groupchat_message_send_response"`
}

/**
 * (手淘群发单)
 * taobao.tbk.sc.groupchat.message.send( 淘宝客-服务商-手淘群发单 )
 * @line https://open.taobao.com/api.htm?spm=a219a.7386653.0.0.22e6669ams0A9b&source=search&docId=38243&docType=2
 */
func (sg *scGroupChat) MessageSend(pid, groupIds string, other ...map[string]string) (res *TbkScGroupchatMessageSendResponse, err error) {
	params := make(map[string]string)
	if len(other) > 0 {
		params = other[0]
	}
	params["pid"] = pid
	params["group_ids"] = groupIds
	_, err = sg.Client.httpPost("taobao.tbk.sc.groupchat.message.send", params, &res)
	return
}

type TbkScGroupchatCreateResponse struct {
	Response struct {
		Status   string `json:"status"`
		Message  string `json:"message"`
		GroupID  string `json:"group_id"`
		GroupURL string `json:"group_url"`
	} `json:"tbk_sc_groupchat_create_response"`
}

/**
 * (手淘群创建)
 * taobao.tbk.sc.groupchat.create( 淘宝客-服务商-手淘群创建 )
 * @line https://open.taobao.com/api.htm?spm=a219a.7386653.0.0.1ecf669aiQu2uE&source=search&docId=38262&docType=2
 */
func (sg *scGroupChat) MessageCreate(title string, subGroupNum ...int) (res *TbkScGroupchatCreateResponse, err error) {
	params := make(map[string]string)
	if len(subGroupNum) > 0 {
		params["sub_group_num"] = strconv.Itoa(subGroupNum[0])
	}
	params["title"] = title
	_, err = sg.Client.httpPost("taobao.tbk.sc.groupchat.create", params, &res)
	return
}

type TbkScGroupchatGetResponse struct {
	Response struct {
		Status    string `json:"status"`
		Message   string `json:"message"`
		GroupList struct {
			Grouplist []struct {
				TotalUserCount int    `json:"total_user_count"`
				GroupID        string `json:"group_id"`
				Title          string `json:"title"`
				SubGroupNum    int    `json:"sub_group_num"`
				GroupURL       string `json:"group_url"`
			} `json:"grouplist"`
		} `json:"group_list"`
	} `json:"tbk_sc_groupchat_get_response"`
}

/**
 * (手淘群查询)
 * taobao.tbk.sc.groupchat.get( 淘宝客-服务商-手淘群查询 )
 * @line https://open.taobao.com/api.htm?spm=a219a.7386653.0.0.4398669a4yVNpV&source=search&docId=38263&docType=2
 */
func (sg *scGroupChat) MessageGet() (res *TbkScGroupchatGetResponse, err error) {
	_, err = sg.Client.httpPost("taobao.tbk.sc.groupchat.get", map[string]string{}, &res)
	return
}

/**
 * ( 淘宝客订单查询 - 社交 新api )
 * taobao.tbk.sc.order.details.get( 淘宝客-服务商-所有订单查询 )
 * @line https://open.taobao.com/api.htm?docId=43755&docType=2
 */
func (s *sc) OrderDetailsGet(startTime, endTime time.Time, other ...map[string]string) (res *TbkOrderDetailsGetResponse, err error) {
	param := make(map[string]string)
	if len(other) > 0 {
		param = other[0]
	}
	param["start_time"] = startTime.Format("2006-01-02 15:04:05")
	param["end_time"] = endTime.Format("2006-01-02 15:04:05")
	_, err = s.Client.httpPost("taobao.tbk.sc.order.details.get", param, &res)
	return
}
