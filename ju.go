package tbk

import (
	"strconv"
	"time"
)

type ju struct {
	Client *Tbk
}

func (t *Tbk) Ju() *ju {
	return &ju{Client: t}
}

type TbkJuTqgGetResponse struct {
	Response struct {
		Results struct {
			Results []struct {
				Title        string      `json:"title"`
				TotalAmount  int         `json:"total_amount"`
				ClickURL     string      `json:"click_url"`
				CategoryName string      `json:"category_name"`
				ZkFinalPrice string      `json:"zk_final_price"`
				EndTime      string      `json:"end_time"`
				SoldNum      int         `json:"sold_num"`
				StartTime    string      `json:"start_time"`
				ReservePrice string      `json:"reserve_price"`
				PicURL       string      `json:"pic_url"`
				NumIid       MixedItemID `json:"num_iid"`
			} `json:"results"`
		} `json:"results"`
		TotalResults int `json:"total_results"`
	} `json:"tbk_ju_tqg_get_response"`
}

/**
 * ( 淘抢购api )
 * taobao.tbk.ju.tqg.get
 * @line https://open.taobao.com/api.htm?docId=27543&docType=2
 */
func (j *ju) TqgGet(adzoneId int64, fields string, startTime, endTime time.Time, other ...map[string]string) (res *TbkJuTqgGetResponse, err error) {
	param := make(map[string]string)
	if len(other) > 0 {
		param = other[0]
	}
	param["adzone_id"] = strconv.FormatInt(adzoneId, 10)
	param["fields"] = fields
	param["start_time"] = startTime.Format("2006-01-02 15:04:05")
	param["end_time"] = endTime.Format("2006-01-02 15:04:05")
	_, err = j.Client.httpPost("taobao.tbk.ju.tqg.get", param, &res)
	return
}
