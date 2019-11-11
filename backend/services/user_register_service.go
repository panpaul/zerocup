package services

import (
	"comment/e"
	"comment/models"
	"comment/serializers"
)

type RegisterService struct {
	UserName string `json:"username" form:"username" binding:"required,min=3,max=20"`
	PassWord string `json:"password" form:"password" binding:"required,min=6,max=20"`
	PassWordConfirm string `json:"passwordConfirm" form:"passwordConfirm" binding:"required,min=6,max=20"`
}

func (service *RegisterService) Valid() *serializers.Response {
	if service.PassWord != service.PassWordConfirm {
		return &serializers.Response{
			Status: e.USER_PASSWORD_NOT_CONSISTENT,
			Message: e.GetMsg(e.USER_PASSWORD_NOT_CONSISTENT),
		}
	}

	count := 0
	models.DB.Model(&models.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializers.Response{
			Status: e.USERNAME_ALREADY_EXISTED,
			Message: e.GetMsg(e.USERNAME_ALREADY_EXISTED),
		}
	}

	return nil
}

func (service *RegisterService) Register() *serializers.Response {
	user := models.User{
		UserName: service.UserName,
	}
	// valid check
	if response := service.Valid(); response != nil {
		return response
	}

	// set user password
	if err := user.SetPassWord(service.PassWord); err != nil {
		return &serializers.Response{
			Status: e.USER_SET_PASSWORD_ERROR,
			Message: e.GetMsg(e.USER_SET_PASSWORD_ERROR),
		}
	}

	// create user
	if err := models.DB.Create(&user).Error; err != nil {
		return &serializers.Response{
			Status: e.REGISTER_ERROR,
			Message: e.GetMsg(e.REGISTER_ERROR),
		}
	}

	// register ok
	return &serializers.Response{
		Status: e.SUCCESS,
		Message: e.GetMsg(e.SUCCESS),
		Data: serializers.BuildUser(user),
	}
}
