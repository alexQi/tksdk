package request

import (
	"alexQi/tksdk/utils"
	"net/url"
)

// taobao.tbk.dg.material.optional( 通用物料搜索API（导购） )
// https://open.taobao.com/api.htm?docId=62201&docType=2
type TbkDgMaterialRecommendRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkDgMaterialRecommendRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils.CheckNumber(tk.Parameters.Get("adzone_id"), "adzone_id")

}

// 添加请求参数
func (tk *TbkDgMaterialRecommendRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// 返回接口名称
func (tk *TbkDgMaterialRecommendRequest) GetApiName() string {
	return "taobao.tbk.dg.material.recommend"
}

// 返回请求参数
func (tk *TbkDgMaterialRecommendRequest) GetParameters() url.Values {
	return *tk.Parameters
}
