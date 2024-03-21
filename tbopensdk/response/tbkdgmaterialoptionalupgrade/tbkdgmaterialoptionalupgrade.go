package tbkdgmaterialoptionalupgrade

import (
	"encoding/json"
	"github.com/alexQi/tksdk/tbopensdk/response"
)

// taobao.tbk.dg.material.optional.upgrade( 通用物料搜索API（导购） )
// https://open.taobao.com/api.htm?docId=64759&docType=2
type Response struct {
	response.TopResponse
	TbkDgMaterialOptionalUpgradeResult Result `json:"tbk_dg_material_optional_upgrade_response"`
}

// 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type Result struct {
	ResultList   ResultList `json:"result_list"`
	TotalResults int64      `json:"total_results"`
	RequestID    string     `json:"request_id"`
}

type ResultList struct {
	MapData []MapDatum `json:"map_data"`
}

type MapDatum struct {
	ItemId             string             `json:"item_id"`
	PublishInfo        PublishInfo        `json:"publish_info"`
	PricePromotionInfo PricePromotionInfo `json:"price_promotion_info"`
	ItemBasicInfo      ItemBasicInfo      `json:"item_basic_info"`
	TmallRankInfo      TmallRankInfo      `json:"tmall_rank_info"`
	PresaleInfo        PresaleInfo        `json:"presale_info"`
	ScopeInfo          ScopeInfo          `json:"scope_info"`
	MgcInfo            MgcInfo            `json:"mgc_info"`
	MatchScore         string             `json:"match_score"`
	CommiScore         string             `json:"commi_score"`
}

type TmallRankInfo struct {
	TmallRankText string `json:"tmall_rank_text"`
	TmallRankURL  string `json:"tmall_rank_url"`
}

type PresaleInfo struct {
	PresaleStartTime       int    `json:"presale_start_time"`
	PresaleEndTime         int    `json:"presale_end_time"`
	PresaleTailStartTime   int    `json:"presale_tail_start_time"`
	PresaleTailEndTime     int    `json:"presale_tail_end_time"`
	PresaleDeposit         string `json:"presale_deposit"`
	PresaleDiscountFeeText string `json:"presale_discount_fee_text"`
}

type ScopeInfo struct {
	SuperiorBrand string `json:"superior_brand"`
}

type MgcInfo struct {
	Price            string `json:"price"`
	PriceDesc        string `json:"price_desc"`
	PromotionSummary string `json:"promotion_summary"`
	PublishTime      string `json:"publish_time"`
	ValidTime        string `json:"valid_time"`
}

type SPCampaign struct {
	SPCid        string `json:"sp_cid"`
	SPName       string `json:"sp_name"`
	SPRate       string `json:"sp_rate"`
	SPLockStatus string `json:"sp_lock_status"`
	SPApplyLink  string `json:"sp_apply_link"`
	SPStatus     string `json:"sp_status"`
}

type SPCampaignList struct {
	SPCampaign []SPCampaign `json:"sp_campaign"`
}

type TopnInfo struct {
	TopnQuantity   int64  `json:"topn_quantity"`
	TopnTotalCount int64  `json:"topn_total_count"`
	TopnEndTime    string `json:"topn_end_time"`
	TopnStartTime  string `json:"topn_start_time"`
	TopnRate       string `json:"topn_rate"`
}

type IncomeInfo struct {
	CommissionRate    string `json:"commission_rate"`
	CommissionAmount  string `json:"commission_amount"`
	SubsidyRate       string `json:"subsidy_rate"`
	SubsidyAmount     string `json:"subsidy_amount"`
	SubsidyUpperLimit string `json:"subsidy_upper_limit"`
	SubsidyType       string `json:"subsidy_type"`
}

type PublishInfo struct {
	IncomeRate                   string         `json:"income_rate"`
	TopnInfo                     TopnInfo       `json:"topn_info"`
	ClickURL                     string         `json:"click_url"`
	CouponShareURL               string         `json:"coupon_share_url"`
	FutureActivityCommissionRate string         `json:"future_activity_commission_rate"`
	FutureActivityTime           string         `json:"future_activity_time"`
	SpCampaignList               SPCampaignList `json:"sp_campaign_list"`
	RankPageURL                  string         `json:"rank_page_url"`
	CommissionType               string         `json:"commission_type"`
	IncomeInfo                   IncomeInfo     `json:"income_info"`
	CpaRewardType                string         `json:"cpa_reward_type"`
	CpaRewardAmount              string         `json:"cpa_reward_amount"`
}

type PromotionMapData struct {
	PromotionTitle     string `json:"promotion_title"`
	PromotionDesc      string `json:"promotion_desc"`
	PromotionFee       string `json:"promotion_fee"`
	PromotionStartTime string `json:"promotion_start_time"`
	PromotionEndTime   string `json:"promotion_end_time"`
	PromotionId        string `json:"promotion_id"`
}

type PricePromotionInfo struct {
	PredictRoundingUpPrice     string `json:"predict_rounding_up_price"`
	PredictRoundingUpPriceDesc string `json:"predict_rounding_up_price_desc"`
	MorePromotionList          struct {
		MorePromotionMapData []PromotionMapData `json:"more_promotion_map_data"`
	} `json:"more_promotion_list"`
	ReservePrice           string `json:"reserve_price"`
	ZkFinalPrice           string `json:"zk_final_price"`
	FinalPromotionPrice    string `json:"final_promotion_price"`
	FinalPromotionPathList struct {
		FinalPromotionPathMapData []PromotionMapData `json:"final_promotion_path_map_data"`
	} `json:"final_promotion_path_list"`
	FutureActivityPromotionPrice    string `json:"future_activity_promotion_price"`
	FutureActivityPromotionPathList struct {
		FutureActivityPromotionPathMapData []PromotionMapData `json:"future_activity_promotion_path_map_data"`
	} `json:"future_activity_promotion_path_list"`
	PromotionTagList struct {
		PromotionTagMapData []struct {
			TagName string `json:"tag_name"`
		} `json:"promotion_tag_map_data"`
	} `json:"promotion_tag_list"`
}

type ItemBasicInfo struct {
	Title                string `json:"title"`
	ShortTitle           string `json:"short_title"`
	PictURL              string `json:"pict_url"`
	WhiteImage           string `json:"white_image"`
	LevelOneCategoryID   int    `json:"level_one_category_id"`
	CategoryID           int    `json:"category_id"`
	CategoryName         string `json:"category_name"`
	SellerID             int    `json:"seller_id"`
	UserType             int    `json:"user_type"`
	ShopTitle            string `json:"shop_title"`
	Volume               int    `json:"volume"`
	SubTitle             string `json:"sub_title"`
	BrandName            string `json:"brand_name"`
	LevelOneCategoryName string `json:"level_one_category_name"`
	TkTotalSales         string `json:"tk_total_sales"`
	Provcity             string `json:"provcity"`
}
