package service

import (
	"errors"
	"talkspace/features/user/entity"
	"talkspace/middlewares"
	"talkspace/utils/constant"
	"talkspace/utils/helper/bcrypt"
	"talkspace/utils/validator"
)

type userService struct {
	userRepository entity.UserRepositoryInterface
}

func NewUserService(userRepository entity.UserRepositoryInterface) entity.UserServiceInterface {
	return &userService{
		userRepository: userRepository,
	}
}

func (us *userService) Register(userCore entity.User) (entity.User, error) {

	errEmpty := validator.IsDataEmpty(userCore.Fullname, userCore.Email, userCore.Password)
	if errEmpty != nil {
		return entity.User{}, errEmpty
	}

	errEmail := validator.IsEmailValid(userCore.Email)
	if errEmail != nil {
		return entity.User{}, errEmail
	}

	errLength := validator.IsMinLengthValid(userCore.Password, 8)
	if errLength != nil {
		return entity.User{}, errLength
	}

	_, err := us.userRepository.FindByEmail(userCore.Email)
	if err == nil {
		return entity.User{}, errors.New(constant.ERROR_EMAIL_EXIST)
	}

	hashedPassword, err := bcrypt.HashPassword(userCore.Password)
	if err != nil {
		return entity.User{}, errors.New(constant.ERROR_PASSWORD_HASH)
	}

	userCore.Password = hashedPassword

	dataUsers, err := us.userRepository.Register(userCore)
	if err != nil {
		return entity.User{}, err
	}

	return dataUsers, nil
}

func (us *userService) Login(email, password string) (entity.User, string, error) {

	errEmpty := validator.IsDataEmpty(email, password)
	if errEmpty != nil {
		return entity.User{}, "", errEmpty
	}

	errEmail := validator.IsEmailValid(email)
	if errEmail != nil {
		return entity.User{}, "", errEmail
	}

	userData, errEmail := us.userRepository.FindByEmail(email)
	if errEmail != nil {
		return entity.User{}, "", errors.New(constant.ERROR_EMAIL_UNREGISTERED)
	}

	comparePassword := bcrypt.ComparePassword(userData.Password, password)
	if comparePassword != nil {
		return entity.User{}, "", errors.New(constant.ERROR_LOGIN)
	}

	token, err := middlewares.CreateToken(userData.ID, "")
	if err != nil {
		return entity.User{}, "", errors.New(constant.ERROR_TOKEN_GENERATE)
	}
	return userData, token, nil
}

func (us *userService) GetByID(id string) (entity.User, error) {
	if id == "" {
		return entity.User{}, errors.New(constant.ERROR_ID_INVALID)
	}

	userData, err := us.userRepository.GetByID(id)
	if err != nil {
		return entity.User{}, errors.New(constant.ERROR_DATA_EMPTY)
	}
	return userData, nil
}

func (us *userService) UpdateByID(id string, userCore entity.User) error {
	if id == "" {
		return errors.New(constant.ERROR_ID_INVALID)
	}

	_, errGetID := us.userRepository.GetByID(id)
	if errGetID != nil {
		return errGetID
	}

	errEmail := validator.IsEmailValid(userCore.Email)
	if errEmail != nil {
		return errEmail
	}

	errBirthdate := validator.IsDateValid(userCore.Birthdate)
	if errBirthdate != nil {
		return errBirthdate
	}

	err := us.userRepository.UpdateByID(id, userCore)
	if err != nil {
		return err
	}

	return nil
}
