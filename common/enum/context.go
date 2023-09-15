package enum

const (
	CurrentId   = "currentId"
	CurrentName = "currentName"
)

type PageNum = int

const (
	MaxPageSize PageNum = 100 // 单页最大数量
	MinPageSize PageNum = 10  // 单页最小数量
)

type CommonInt = int

const (
	MaxUrl CommonInt = 1
)
