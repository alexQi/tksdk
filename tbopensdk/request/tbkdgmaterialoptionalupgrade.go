package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// taobao.tbk.dg.material.optional.upgrade( 通用物料搜索API（导购） )
// https://open.taobao.com/api.htm?docId=64759&docType=2
type TbkDgMaterialOptionalUpgradeRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkDgMaterialOptionalUpgradeRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils.CheckNumber(tk.Parameters.Get("adzone_id"), "adzone_id")

}

// 添加请求参数
func (tk *TbkDgMaterialOptionalUpgradeRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// 返回接口名称
func (tk *TbkDgMaterialOptionalUpgradeRequest) GetApiName() string {
	return "taobao.tbk.dg.material.optional.upgrade"
}

// 返回请求参数
func (tk *TbkDgMaterialOptionalUpgradeRequest) GetParameters() url.Values {
	return *tk.Parameters
}
