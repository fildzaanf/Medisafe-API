package request

import "talkspace/features/user/entity"

func UserRegisterRequestToUserCore(request UserRegisterRequest) entity.User {
	return entity.User{
		Fullname: request.Fullname,
		Email:    request.Email,
		Password: request.Password,
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
		Password:       request.Password,
		ProfilePicture: request.ProfilePicture,
		Gender:         request.Gender,
		Birthdate:      request.Birthdate,
		BloodType:      request.BloodType,
		Weight:         request.Weight,
		Height:         request.Height,
	}
}
