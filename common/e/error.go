package e

import "errors"

var (
	Error_PASSWORD_ERROR                 = errors.New("密码错误")
	Error_ACCOUNT_NOT_FOUND              = errors.New("账号不存在")
	Error_ACCOUNT_LOCKED                 = errors.New("账号被锁定")
	Error_ALREADY_EXISTS                 = errors.New("已存在")
	Error_UNKNOWN_ERROR                  = errors.New("未知错误")
	Error_USER_NOT_LOGIN                 = errors.New("用户未登录")
	Error_CATEGORY_BE_RELATED_BY_SETMEAL = errors.New("当前分类关联了套餐,不能删除")
	Error_CATEGORY_BE_RELATED_BY_DISH    = errors.New("当前分类关联了菜品,不能删除")
	Error_SHOPPING_CART_IS_NULL          = errors.New("购物车数据为空，不能下单")
	Error_ADDRESS_BOOK_IS_NULL           = errors.New("用户地址为空，不能下单")
	Error_LOGIN_FAILED                   = errors.New("登录失败")
	Error_UPLOAD_FAILED                  = errors.New("文件上传失败")
	Error_SETMEAL_ENABLE_FAILED          = errors.New("套餐内包含未启售菜品，无法启售")
	Error_PASSWORD_EDIT_FAILED           = errors.New("密码修改失败")
	Error_DISH_ON_SALE                   = errors.New("起售中的菜品不能删除")
	Error_SETMEAL_ON_SALE                = errors.New("起售中的套餐不能删除")
	Error_DISH_BE_RELATED_BY_SETMEAL     = errors.New("当前菜品关联了套餐,不能删除")
	Error_ORDER_STATUS_ERROR             = errors.New("订单状态错误")
	Error_ORDER_NOT_FOUND                = errors.New("订单不存在")
)
