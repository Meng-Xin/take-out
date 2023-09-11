package response

type EmployeeLogin struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Token    string `json:"token"`
	UserName string `json:"userName"`
}
