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