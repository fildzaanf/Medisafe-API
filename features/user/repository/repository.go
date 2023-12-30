package repository

import (
	"errors"
	"talkspace/features/user/entity"
	"talkspace/features/user/model"
	"talkspace/utils/constant"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) entity.UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Register(userCore entity.User) (entity.User, error) {
	request := entity.UserCoreToUserModel(userCore)

	result := ur.db.Create(&request)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	response := entity.UserModelToUserCore(request)
	return response, nil
}

func (ur *userRepository) GetByID(id string) (entity.User, error) {
	userModel := model.User{}

	result := ur.db.Where("id = ?", id).First(&userModel)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	if result.RowsAffected == 0 {
		return entity.User{}, errors.New(constant.ERROR_ID_NOTFOUND)
	}

	response := entity.UserModelToUserCore(userModel)
	return response, nil
}

func (ur *userRepository) UpdateByID(id string, userCore entity.User) error {
	request := entity.UserCoreToUserModel(userCore)

	result := ur.db.Where("id = ?", id).Updates(&request)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New(constant.ERROR_ID_NOTFOUND)
	}

	return nil
}

func (ur *userRepository) FindByEmail(email string) (entity.User, error) {
	userModel := model.User{}

	result := ur.db.Where("email = ?", email).First(&userModel)
	
	if result.RowsAffected == 0 {
		return entity.User{}, errors.New(constant.ERROR_EMAIL_NOTFOUND)
	}

	if result.Error != nil {
		return entity.User{},result.Error
	}

	response := entity.UserModelToUserCore(userModel)
	return response, nil
}

func (ur *userRepository) GetByVerificationToken(token string) (entity.User, error) {
	userModel := model.User{}

	result := ur.db.Where("verification_token = ?", token).First(&userModel)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	if result.RowsAffected == 0 {
		return entity.User{}, errors.New(constant.ERROR_TOKEN_NOTFOUND)
	}

	userToken := entity.UserModelToUserCore(userModel)
	return userToken, nil
}

func (ur *userRepository) UpdateIsVerified(id string, isVerified bool) error {
	userModel := model.User{}

	result := ur.db.Where("id = ?", id).First(&userModel)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New(constant.ERROR_ID_NOTFOUND)
	}

	userModel.IsVerified = isVerified

	errSave := ur.db.Save(&userModel)
	if errSave.Error != nil {
		return errSave.Error
	}

	return nil
}

func (ur *userRepository) SendOTP(email string, otp string, expired int64) (userCore entity.User, err error) {
	userModel := model.User{}

	result := ur.db.Where("email = ?", email).First(&userModel)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return entity.User{}, errors.New(constant.ERROR_EMAIL_NOTFOUND)
		}
		return entity.User{}, result.Error
	}

	userModel.OTP = otp
	userModel.OTPExpiration = expired

	errUpdate := ur.db.Save(&userModel).Error
	if errUpdate != nil {
		return entity.User{}, errUpdate
	}

	response := entity.UserModelToUserCore(userModel)

	return response, nil
}

func (ur *userRepository) VerifyOTP(email, otp string) (entity.User, error) {
	userModel := model.User{}

	result := ur.db.Where("otp = ? AND email = ?", otp, email).First(&userModel)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	if result.RowsAffected == 0 {
		return entity.User{}, errors.New(constant.ERROR_EMAIL_OTP)
	}

	response := entity.UserModelToUserCore(userModel)
	return response, nil
}

func (ur *userRepository) ResetOTP(otp string) (userCore entity.User, err error) {
	userModel := model.User{}

	result := ur.db.Where("otp = ?", otp).First(&userModel)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	if result.RowsAffected == 0 {
		return entity.User{}, errors.New(constant.ERROR_OTP_NOTFOUND)
	}

	userModel.OTP = ""
	userModel.OTPExpiration = 0

	errUpdate := ur.db.Save(&userModel).Error
	if errUpdate != nil {
		return entity.User{}, errUpdate
	}

	response := entity.UserModelToUserCore(userModel)
	return response, nil
}
