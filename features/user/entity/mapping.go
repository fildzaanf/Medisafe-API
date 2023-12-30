package entity

import "talkspace/features/user/model"

func UserCoreToUserModel(userCore User) model.User {
	userModel := model.User{
		Fullname:          userCore.Fullname,
		Email:             userCore.Email,
		Password:          userCore.Password,
		ProfilePicture:    userCore.ProfilePicture,
		Birthdate:         userCore.Birthdate,
		Gender:            userCore.Gender,
		BloodType:         userCore.BloodType,
		Height:            userCore.Height,
		Weight:            userCore.Weight,
		OTP:               userCore.OTP,
		OTPExpiration:     userCore.OTPExpiration,
		IsVerified:        userCore.IsVerified,
		VerificationToken: userCore.VerificationToken,
	}
	return userModel
}

func ListUserCoreToUserModel(userCore []User) []model.User {
	listUserModel := []model.User{}
	for _, user := range userCore {
		userModel := UserCoreToUserModel(user)
		listUserModel = append(listUserModel, userModel)
	}
	return listUserModel
}

func UserModelToUserCore(userModel model.User) User {
	userCore := User{
		ID:                userModel.ID,
		Fullname:          userModel.Fullname,
		Email:             userModel.Email,
		Password:          userModel.Password,
		ProfilePicture:    userModel.ProfilePicture,
		Birthdate:         userModel.Birthdate,
		Gender:            userModel.Gender,
		BloodType:         userModel.BloodType,
		Height:            userModel.Height,
		Weight:            userModel.Weight,
		OTP:               userModel.OTP,
		OTPExpiration:     userModel.OTPExpiration,
		IsVerified:        userModel.IsVerified,
		VerificationToken: userModel.VerificationToken,
		CreatedAt:         userModel.CreatedAt,
		UpdatedAt:         userModel.UpdatedAt,
		DeletedAt:         userModel.DeletedAt,
	}
	return userCore
}

func ListUserModelToUserCore(userModel []model.User) []User {
	listUserCore := []User{}
	for _, user := range userModel {
		userCore := UserModelToUserCore(user)
		listUserCore = append(listUserCore, userCore)
	}
	return listUserCore
}
