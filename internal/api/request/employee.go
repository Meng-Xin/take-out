package request

type EmployeeLogin struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type EmployeeEditPassword struct {
	EmpId       uint64 `json:"empId"`
	NewPassword string `json:"newPassword" binding:"required"`
	OldPassword string `json:"oldPassword" binding:"required"`
}

type EmployeeDTO struct {
	Id       uint64 `json:"id"`                          //员工id
	IdNumber string `json:"idNumber" binding:"required"` //身份证
	Name     string `json:"name" binding:"required"`     //姓名
	Phone    string `json:"phone" binding:"required"`    //手机号
	Sex      string `json:"sex" binding:"required"`      //性别
	UserName string `json:"username" binding:"required"` //用户名
}

type EmployeePageQueryDTO struct {
	Name     string `form:"name"`     // 分页查询的name
	Page     int    `form:"page"`     // 分页查询的页数
	PageSize int    `form:"pageSize"` // 分页查询的页容量
}
