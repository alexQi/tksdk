package alibabaalscunionkbbbtitempromotionfilterlist

import (
	"encoding/json"
	"fmt"
	"github.com/alexQi/tksdk/alscopensdk/response"
)

// Response alibaba.alsc.union.kb.bbt.item.promotion.filter.list( 本地生活爆爆团选品筛选集合 )
type Response struct {
	response.TopResponse
	ResultResponse ResultResponse `json:"alibaba_alsc_union_kb_bbt_item_promotion_filter_list_response"`
}

// WrapResult 解析输出结果
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

// ResultResponse 响应结果
type ResultResponse struct {
	BizErrorDesc string `json:"biz_error_desc"`
	BizErrorCode int    `json:"biz_error_code"`
	RequestID    string `json:"request_id"`
	Result       Result `json:"result"`
}

func (rr ResultResponse) IsError() bool {
	return rr.BizErrorCode != 0
}

func (rr ResultResponse) Error() string {
	return fmt.Sprintf("biz_error_code:%d biz_error_desc:%s", rr.BizErrorCode, rr.BizErrorDesc)
}

type Result struct {
	FilterTableNameDTO []FilterTableNameDTO `json:"filter_table_name_d_t_o"`
}

type FilterTableNameDTO struct {
	FilterKey  string `json:"filter_key"`
	FilterName string `json:"filter_name"`
}
