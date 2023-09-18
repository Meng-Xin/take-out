package model

type DishFlavor struct {
	Id     uint64 `json:"id"`      //口味id
	DishId uint64 `json:"dish_id"` //菜品id
	Name   string `json:"name"`    //口味主题 温度|甜度|辣度
	Value  string `json:"value"`   //口味信息 可多个
}

func (d *DishFlavor) TableName() string {
	return "dish_flavor"
}
