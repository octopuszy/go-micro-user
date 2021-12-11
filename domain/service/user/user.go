package user

import (
	userModel "micro_test/domain/model/user"
	userRepository "micro_test/domain/repository/user"
	"micro_test/domain/util"
)

type UserService struct {
	UserRepository userRepository.IUserRepository
}

func NewUserService(r userRepository.IUserRepository) *UserService {
	return &UserService{
		UserRepository: r,
	}
}

func (s UserService) AddUser(user *userModel.User)  (int64, error){
	hashPassword, err := util.GeneratePassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = string(hashPassword)
	return s.UserRepository.CreateUser(user)
}

func (s UserService) GetUserInfo(username string)  (user *userModel.User, err error){
	return s.UserRepository.FindUserByName(username)
}

func (s UserService) CheckPassword(inputPasswd ,hashPassword string) bool {
	return util.ComparePasswords([]byte(hashPassword), inputPasswd)
}