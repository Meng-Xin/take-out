package tx

// Transaction 定义事务管理器接口
// 因为不同数据库的事务管理方式不同，所以这里使用接口来定义事务管理器的功能
// 具体的实现可以由不同的数据库驱动来实现
// 例如，我们当前项目的Dao层为了引入事务破坏了 `依赖倒置` 原则，所以我们的目标是面向抽象
type Transaction interface {
	// Begin 开始一个事务
	Begin() error
	// Commit 提交事务
	Commit() error
	// Rollback 回滚事务
	Rollback()
}
