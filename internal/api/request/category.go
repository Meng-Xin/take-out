package request

type CategoryDTO struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Sort int    `json:"sort"`
	Cate int    `json:"cate"`
}
