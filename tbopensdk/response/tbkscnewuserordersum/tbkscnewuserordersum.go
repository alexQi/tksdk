package tbkscnewuserordersum

import (
	"encoding/json"
	"github.com/alexQi/tksdk/tbopensdk/response"
)

// taobao.tbk.sc.newuser.tborder.sum( 拉新活动汇总API--社交 )
type Response struct {
	response.TopResponse
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
