package handler

import (
	"context"
	"github.com/shuheng-mo/Mogo/user/domain/model"
	"github.com/shuheng-mo/Mogo/user/domain/service"
	user "github.com/shuheng-mo/Mogo/user/proto/user"
)

type User struct {
	UserDataService service.IUserDataService
}

func (u *User) Register(ctx context.Context, userRegisterRequest *user.UserRegisterRequest, userRegisterResponse *user.UserRegisterResponse) error {
	userRegister := &model.User{
		UserName:     userRegisterRequest.UserName,
		FirstName:    userRegisterRequest.FirstName,
		HashPassword: userRegisterRequest.Pwd,
	}

	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}

	userRegisterResponse.Message = "Registered successfully!"
	return nil
}

func (u *User) Login(ctx context.Context, userLogin *user.UserLoginRequest, loginResponse *user.UserLoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(userLogin.UserName, userLogin.Pwd)
	if err != nil {
		return err
	}
	loginResponse.IsSuccess = isOk
	return nil
}

func (u *User) GetUserInfo(ctx context.Context, userInfoRequest *user.UserInfoRequest, userInfoResponse *user.UserInfoResponse) error {
	userInfo, err := u.UserDataService.FindUserByName(userInfoRequest.UserName)
	if err != nil {
		return nil
	}
	userInfoResponse = UserForResponse(userInfo)
	return nil
}

// data type conversion
func UserForResponse(userModel *model.User) *user.UserInfoResponse {
	response := &user.UserInfoResponse{}
	response.UserName = userModel.UserName
	response.FirstName = userModel.FirstName
	response.UserId = userModel.ID
	return response
}
