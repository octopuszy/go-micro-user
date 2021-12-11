package handler

import (
	"context"
	userModel "micro_test/domain/model/user"
	service "micro_test/domain/service/user"
	proto_user "micro_test/proto/user"
)

type User struct {
	UserService *service.UserService
}

func (u User) Register(ctx context.Context, req *proto_user.RegisterReq, rsp *proto_user.CommonRsp) error {
	user := userModel.User{
		UserName: req.UserName,
		Password: req.Password,
		Email: req.Email,
	}
	_, err := u.UserService.AddUser(&user)
	if err != nil {
		return err
	}
	rsp.ErrorNo = 0
	rsp.Message = "ok"
	return nil
}

func (u User)Login(ctx context.Context, req *proto_user.LoginReq, rsp *proto_user.CommonRsp) error {
	info, err := u.UserService.GetUserInfo(req.UserName)
	if err != nil {
		return err
	}
	if info == nil {
		rsp.ErrorNo = 1
		rsp.Message = "user not found"
		return  nil
	}
	ok := u.UserService.CheckPassword(req.Password, info.Password)
	if !ok {
		rsp.ErrorNo = 1
		rsp.Message = "password error"
		return nil
	}
	rsp.ErrorNo = 0
	rsp.Message = "ok"
	return nil
}

func (u User) GetUserInfo(ctx context.Context, req *proto_user.GetUserInfoReq, rsp *proto_user.GetUserInfoRsp) error {
	info, err := u.UserService.GetUserInfo(req.UserName)
	if err != nil {
		return err
	}
	rsp.UserId = info.Id
	rsp.Email = info.Email
	rsp.UserName = info.UserName
	return nil
}