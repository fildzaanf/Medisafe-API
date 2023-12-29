package response

import "talkspace/features/user/entity"

func UserCoreToUserRegisterResponse(response entity.User) UserRegisterResponse {
	return UserRegisterResponse{
		ID:       response.ID,
		Fullname: response.Fullname,
		Email:    response.Email,
	}
}

func UserCoreToUserLoginResponse(response entity.User, token string) UserLoginResponse {
	return UserLoginResponse{
		ID:       response.ID,
		Fullname: response.Fullname,
		Email:    response.Email,
		Token:    token,
	}
}

func UserCoreToUserProfileResponse(response entity.User) UserProfileResponse {
	return UserProfileResponse{
		Fullname:       response.Fullname,
		Email:          response.Email,
		ProfilePicture: response.ProfilePicture,
		Gender:         response.Gender,
		Birthdate:      response.Birthdate,
		BloodType:      response.BloodType,
		Weight:         response.Weight,
		Height:         response.Height,
	}
}
