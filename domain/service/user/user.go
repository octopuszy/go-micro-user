package user

import (
	userModel "github.com/octopuszy/go-micro-user/domain/model/user"
	userRepository "github.com/octopuszy/go-micro-user/domain/repository/user"
	util "github.com/octopuszy/micro-util"
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