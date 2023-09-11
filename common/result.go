package common

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type PageResult struct {
	Total int //总记录数

	Records interface{} //当前页数据集合
}
