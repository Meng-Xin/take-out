package common

import (
	"gorm.io/gorm"
	"take-out/common/enum"
)

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type PageResult struct {
	Total   int64       `json:"total"`   //总记录数
	Records interface{} `json:"records"` //当前页数据集合
}

// PageVerify 分页查询 过滤器
func PageVerify(page *int, pageSize *int) {
	// 过滤 当前页、单页数量
	if *page < 1 {
		*page = 1
	}
	switch {
	case *pageSize > 100:
		*pageSize = enum.MaxPageSize
	case *pageSize <= 0:
		*pageSize = enum.MinPageSize
	}
}

func (p *PageResult) Paginate(page *int, pageSize *int) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		// 分页校验
		PageVerify(page, pageSize)

		//// 计算总页数 Total / PageSize = Pages
		//p.Pages = p.Total / int64(p.PageSize)
		//// 如果还有余那么也可以查询
		//if p.Total%int64(p.PageSize) != 0 {
		//	p.Pages++
		//}

		// 拼接分页
		d.Offset((*page - 1) * *pageSize).Limit(*pageSize)
		return d
	}
}
