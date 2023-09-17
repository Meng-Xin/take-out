package request

type CategoryDTO struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Sort string `json:"sort"`
	Cate string `json:"type"`
}

type CategoryPageQueryDTO struct {
	Name     string `json:"name"`     // 分页查询的name
	Page     int    `json:"page"`     // 分页查询的页数
	PageSize int    `json:"pageSize"` // 分页查询的页容量
	Cate     int    `json:"type"`     // 分类类型：1为菜品分类，2为套餐分类
}
