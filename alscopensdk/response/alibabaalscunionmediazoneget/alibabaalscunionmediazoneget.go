package alibabaalscunionmediazoneget

import (
	"encoding/json"
	"fmt"
	"github.com/alexQi/tksdk/alscopensdk/response"
)

// Response alibaba.alsc.union.media.zone.get( 本地生活媒体推广位查询 )
type Response struct {
	response.TopResponse
	ResultResponse ResultResponse `json:"alibaba_alsc_union_media_zone_get_response"`
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
	TotalCount   int    `json:"total_count"`
	Result       Result `json:"result"`
}

func (rr ResultResponse) IsError() bool {
	return rr.BizErrorCode != 0
}

func (rr ResultResponse) Error() string {
	return fmt.Sprintf("biz_error_code:%d biz_error_desc:%s", rr.BizErrorCode, rr.BizErrorDesc)
}

type Result struct {
	ZoneInfoDTO []ZoneInfoDTOItem `json:"zone_info_d_t_o"`
}

type ZoneInfoDTOItem struct {
	Name string `json:"name"`
	Pid  string `json:"pid"`
}
