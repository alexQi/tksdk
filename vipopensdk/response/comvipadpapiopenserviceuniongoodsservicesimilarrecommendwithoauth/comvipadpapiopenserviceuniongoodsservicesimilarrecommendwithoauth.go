package comvipadpapiopenserviceuniongoodsservicesimilarrecommendwithoauth

import (
	"encoding/json"
	response2 "github.com/alexQi/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionGoodsService 相似推荐商品列表-需要oauth授权
type Response struct {
	response2.TopResponse
	Result Result `json:"result"`
}

// 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ReturnCode = "-1"
		t.ReturnMessage = unmarshal.Error()
	}
}

type Result struct {
	Total         int64           `json:"total"`
	GoodsInfoList []GoodsInfoList `json:"goodsInfoList"`
	PageSize      int64           `json:"pageSize"`
	Page          int64           `json:"page"`
}

type GoodsInfoList struct {
	MarketPrice      string     `json:"marketPrice"`
	CommissionRate   string     `json:"commissionRate"`
	GoodsID          string     `json:"goodsId"`
	Discount         string     `json:"discount"`
	CategoryName     string     `json:"categoryName"`
	HaiTAO           int64      `json:"haiTao"`
	Cat2NdName       string     `json:"cat2ndName"`
	Cat1StName       string     `json:"cat1stName"`
	VipPrice         string     `json:"vipPrice"`
	Commission       string     `json:"commission"`
	Cat1StID         int64      `json:"cat1stId"`
	GoodsName        string     `json:"goodsName"`
	BrandName        string     `json:"brandName"`
	BrandLogoFull    string     `json:"brandLogoFull"`
	BrandStoreSn     string     `json:"brandStoreSn"`
	SellTimeFrom     int64      `json:"sellTimeFrom"`
	SchemeStartTime  int64      `json:"schemeStartTime"`
	SchemeEndTime    int64      `json:"schemeEndTime"`
	SourceType       int64      `json:"sourceType"`
	SellTimeTo       int64      `json:"sellTimeTo"`
	BrandID          int64      `json:"brandId"`
	GoodsThumbURL    string     `json:"goodsThumbUrl"`
	Cat2NdID         int64      `json:"cat2ndId"`
	StoreInfo        StoreInfo  `json:"storeInfo"`
	GoodsMainPicture string     `json:"goodsMainPicture"`
	DestURL          string     `json:"destUrl"`
	CategoryID       int64      `json:"categoryId"`
	Status           int64      `json:"status"`
	CouponInfo       CouponInfo `json:"couponInfo"`
}

type StoreInfo struct {
	StoreName string `json:"storeName"`
	StoreID   string `json:"storeId"`
}

type CouponInfo struct {
	UseEndTime        int64  `json:"useEndTime"`
	TotalAmount       int64  `json:"totalAmount"`
	CouponName        string `json:"couponName"`
	ActivateEndTime   int64  `json:"activateEndTime"`
	Buy               string `json:"buy"`
	UseBeginTime      int64  `json:"useBeginTime"`
	CouponNo          string `json:"couponNo"`
	Fav               string `json:"fav"`
	ActivateBeginTime int64  `json:"activateBeginTime"`
	ActivedAmount     int64  `json:"activedAmount"`
}
