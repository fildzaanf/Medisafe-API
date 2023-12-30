package request

import "talkspace/features/user/entity"

func UserRegisterRequestToUserCore(request UserRegisterRequest) entity.User {
	return entity.User{
		Fullname:        request.Fullname,
		Email:           request.Email,
		Password:        request.Password,
		ConfirmPassword: request.ConfirmPassword,
	}
}

func UserLoginRequestToUserCore(request UserLoginRequest) entity.User {
	return entity.User{
		Email:    request.Email,
		Password: request.Password,
	}
}

func UserUpdateProfileRequestToUserCore(request UserUpdateProfileRequest) entity.User {
	return entity.User{
		Fullname:       request.Fullname,
		Email:          request.Email,
		ProfilePicture: request.ProfilePicture,
		Gender:         request.Gender,
		Birthdate:      request.Birthdate,
		BloodType:      request.BloodType,
		Weight:         request.Weight,
		Height:         request.Height,
	}
}

func UserNewPasswordRequestToUserCore(request UserNewPasswordRequest) entity.User {
	return entity.User{
		Password:        request.Password,
		ConfirmPassword: request.ConfirmPassword,
	}
}

func UserUpdatePasswordRequestToUserCore(request UserUpdatePasswordRequest) entity.User {
	return entity.User{
		Password:        request.Password,
		NewPassword:     request.NewPassword,
		ConfirmPassword: request.ConfirmPassword,
	}
}

func UserSendOTPRequestToUserCore(request UserSendOTPRequest) entity.User {
	return entity.User{
		Email: request.Email,
	}
}

func UserVerifyOTPRequestToUserCore(request UserVerifyOTPRequest) entity.User {
	return entity.User{
		Email: request.Email,
		OTP:   request.OTP,
	}
}
