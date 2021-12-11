package user

import (
	"micro_test/domain/model/user"
)

type IUserRepository interface {
	// FindUserByName 根据用户名查找用户信息
	FindUserByName(string) (*user.User, error)

	// FindUserById 根据用户 id 查用户信息
	FindUserById(int64) (*user.User, error)

	// CreateUser 创建用户
	CreateUser(*user.User) (int64, error)

	// DelUserById 根据用户 id 删除用户信息
	DelUserById(int64) error

	// UpdateUser 更新用户信息
	UpdateUser(*user.User) error

	// FindAll 分页查找所有用户
	FindAll(pageIndex int, pageSize int) ([]user.User, error)
}
