package tbk

import "strconv"

type contents struct {
	Client *Tbk
}

func (t *Tbk) Contents() *contents {
	return &contents{Client: t}
}

type TbkContentGetResponse struct {
	Response struct {
		Result struct {
			Data struct {
				Count    int `json:"count"`
				Contents struct {
					MapData []ContentsData `json:"map_data"`
				} `json:"contents"`
				LastTimestamp string `json:"last_timestamp"`
			} `json:"data"`
		} `json:"result"`
	} `json:"tbk_content_get_response"`
}

type ContentsData struct {
	Summary      string `json:"summary"`
	Score        int    `json:"score"`
	Tags         string `json:"tags"`
	Title        string `json:"title"`
	AuthorID     string `json:"author_id"`
	AuthorNick   string `json:"author_nick"`
	AuthorAvatar string `json:"author_avatar"`
	Clink        string `json:"clink"`
	Type         string `json:"type"`
	UIStyle      string `json:"ui_style"`
	Images       struct {
		String []string `json:"string"`
	} `json:"images"`
	ContentCategories string `json:"content_categories"`
	PublishTime       string `json:"publish_time"`
	ContentID         int    `json:"content_id"`
	PromotionTag      string `json:"promotion_tag"`
	ItemIds           struct {
		Number []int `json:"number"`
	} `json:"item_ids"`
	UpdateTime string `json:"update_time"`
}

/**
 * ( 淘宝客-推广者-图文内容输出 )
 * taobao.tbk.content.get
 * @line https://open.taobao.com/api.htm?docId=31137&docType=2
 */
func (c *contents) Get(adzoneId int64, other ...map[string]string) (res *TbkContentGetResponse, err error) {
	params := make(map[string]string)
	if len(other) > 0 {
		params = other[0]
	}
	params["adzone_id"] = strconv.FormatInt(adzoneId, 10)
	_, err = c.Client.httpPost("taobao.tbk.item.info.get", params, &res)
	return
}
