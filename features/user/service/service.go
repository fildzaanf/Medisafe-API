package service

import (
	"errors"
	"talkspace/features/user/entity"
	"talkspace/middlewares"
	"talkspace/utils/constant"
	"talkspace/utils/helper/bcrypt"
	"talkspace/utils/helper/email/mailer/onetimepassword"
	"talkspace/utils/helper/email/mailer/verification"
	"talkspace/utils/helper/generator"
	"talkspace/utils/validator"
	"time"
)

type userService struct {
	userRepository entity.UserRepositoryInterface
}

func NewUserService(ur entity.UserRepositoryInterface) entity.UserServiceInterface {
	return &userService{
		userRepository: ur,
	}
}

func (us *userService) Register(userCore entity.User) (entity.User, error) {

	errEmpty := validator.IsDataEmpty(
		userCore.Fullname, 
		userCore.Email, 
		userCore.Password, 
		userCore.ConfirmPassword)
	if errEmpty != nil {
		return entity.User{}, errEmpty
	}

	errEmailValid := validator.IsEmailValid(userCore.Email)
	if errEmailValid != nil {
		return entity.User{}, errEmailValid
	}

	errLength := validator.IsMinLengthValid(userCore.Password, 10)
	if errLength != nil {
		return entity.User{}, errLength
	}

	_, errFindEmail := us.userRepository.FindByEmail(userCore.Email)
	if errFindEmail == nil {
		return entity.User{}, errors.New(constant.ERROR_EMAIL_EXIST)
	}

	if userCore.Password != userCore.ConfirmPassword {
		return entity.User{}, errors.New(constant.ERROR_PASSWORD_CONFIRM)
	}

	hashedPassword, errHash := bcrypt.HashPassword(userCore.Password)
	if errHash != nil {
		return entity.User{}, errors.New(constant.ERROR_PASSWORD_HASH)
	}
	userCore.Password = hashedPassword

	token, errGenerateVerifyToken := generator.GenerateRandomBytes()
	if errGenerateVerifyToken != nil {
		errors.New(constant.ERROR_TOKEN_VERIFICATION)
	}
	userCore.VerifyAccount = token

	userData, errRegister := us.userRepository.Register(userCore)
	if errRegister != nil {
		return entity.User{}, errRegister
	}

	verification.SendEmailVerificationAccount(userCore.Email, token)

	return userData, nil
}

func (us *userService) Login(email, password string) (entity.User, string, error) {

	errEmpty := validator.IsDataEmpty(email, password)
	if errEmpty != nil {
		return entity.User{}, "", errEmpty
	}

	errEmailValid := validator.IsEmailValid(email)
	if errEmailValid != nil {
		return entity.User{}, "", errEmailValid
	}

	userData, errFindEmail := us.userRepository.FindByEmail(email)
	if errFindEmail != nil {
		return entity.User{}, "", errors.New(constant.ERROR_EMAIL_UNREGISTERED)
	}

	if !userData.IsVerified {
		return entity.User{}, "", errors.New(constant.ERROR_ACCOUNT_UNVERIFIED)
	}

	comparePassword := bcrypt.ComparePassword(userData.Password, password)
	if comparePassword != nil {
		return entity.User{}, "", errors.New(constant.ERROR_LOGIN)
	}

	token, errCreateToken := middlewares.CreateToken(userData.ID, "")
	if errCreateToken != nil {
		return entity.User{}, "", errors.New(constant.ERROR_TOKEN_GENERATE)
	}
	return userData, token, nil
}

func (us *userService) GetByID(id string) (entity.User, error) {
	if id == "" {
		return entity.User{}, errors.New(constant.ERROR_ID_INVALID)
	}

	userData, errGetID := us.userRepository.GetByID(id)
	if errGetID != nil {
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

	errEmailValid := validator.IsEmailValid(userCore.Email)
	if errEmailValid != nil {
		return errEmailValid
	}

	errBirthdate := validator.IsDateValid(userCore.Birthdate)
	if errBirthdate != nil {
		return errBirthdate
	}

	validGender := []interface{}{"male", "female"}
	errGender := validator.IsDataValid(userCore.Gender, validGender, true)
	if errGender != nil {
		return errGender
	}

	validBloodType := []interface{}{"A", "B", "O", "AB"}
	errBloodType := validator.IsDataValid(userCore.BloodType, validBloodType, true)
	if errBloodType != nil {
		return errBloodType
	}

	errUpdate := us.userRepository.UpdateByID(id, userCore)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

func (us *userService) UpdatePassword(id string, userCore entity.User) error {
	if id == "" {
		return errors.New(constant.ERROR_ID_INVALID)
	}

	result, errGetID := us.GetByID(id)
	if errGetID != nil {
		return errGetID
	}

	errEmpty := validator.IsDataEmpty(userCore.Password, userCore.NewPassword, userCore.ConfirmPassword)
	if errEmpty != nil {
		return errEmpty
	}

	errLength := validator.IsMinLengthValid(userCore.NewPassword, 10)
	if errLength != nil {
		return errLength
	}

	comparePassword := bcrypt.ComparePassword(result.Password, userCore.Password)
	if comparePassword != nil {
		return errors.New(constant.ERROR_OLDPASSWORD_INVALID)
	}

	if userCore.NewPassword != userCore.ConfirmPassword {
		return errors.New(constant.ERROR_PASSWORD_CONFIRM)
	}

	HashPassword, errHash := bcrypt.HashPassword(userCore.NewPassword)
	if errHash != nil {
		return errors.New(constant.ERROR_PASSWORD_HASH)
	}
	userCore.Password = HashPassword

	errUpdate := us.userRepository.UpdatePassword(id, userCore)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

func (us *userService) VerifyUser(token string) (bool, error) {
	if token == "" {
		return false, errors.New(constant.ERROR_TOKEN_INVALID)
	}

	user, errGetVerifyToken := us.userRepository.GetByVerificationToken(token)
	if errGetVerifyToken != nil {
		return false, errors.New(constant.ERROR_DATA_RETRIEVED)
	}

	if user.IsVerified {
		return true, nil
	}

	errUpdate := us.userRepository.UpdateIsVerified(user.ID, true)
	if errUpdate != nil {
		return false, errors.New(constant.ERROR_ACCOUNT_VERIFICATION)
	}

	return false, nil
}

func (us *userService) UpdateIsVerified(id string, isVerified bool) error {
	if id == "" {
		return errors.New(constant.ERROR_ID_INVALID)
	}

	return us.userRepository.UpdateIsVerified(id, isVerified)
}

func (us *userService) SendOTP(email string) error {

	errEmpty := validator.IsDataEmpty(email)
	if errEmpty != nil {
		return errEmpty
	}

	errEmailValid := validator.IsEmailValid(email)
	if errEmailValid != nil {
		return errEmailValid
	}

	otp, errGenerateOTP := generator.GenerateRandomCode()
	if errGenerateOTP != nil {
		return errors.New(constant.ERROR_OTP_GENERATE)
	}

	expired := time.Now().Add(5 * time.Minute).Unix()

	_, errSend := us.userRepository.SendOTP(email, otp, expired)
	if errSend != nil {
		return errSend
	}

	onetimepassword.SendEmailOTP(email, otp)
	return nil
}

func (us *userService) VerifyOTP(email, otp string) (string, error) {

	errEmpty := validator.IsDataEmpty(email, otp)
	if errEmpty != nil {
		return "", errEmpty
	}

	userData, err := us.userRepository.VerifyOTP(email, otp)
	if err != nil {
		return "", errors.New(constant.ERROR_EMAIL_OTP)
	}

	if userData.OTPExpiration <= time.Now().Unix() {
		return "", errors.New(constant.ERROR_OTP_EXPIRED)
	}

	if userData.OTP != otp {
		return "", errors.New(constant.ERROR_OTP_INVALID)
	}

	token, err := middlewares.CreateVerifyToken(email)
	if err != nil {
		return "", errors.New(constant.ERROR_TOKEN_GENERATE)
	}

	_, errReset := us.userRepository.ResetOTP(otp)
	if errReset != nil {
		return "", errors.New(constant.ERROR_OTP_RESET)
	}

	return token, nil
}

func (us *userService) NewPassword(email string, userCore entity.User) error {

	errEmpty := validator.IsDataEmpty(email, userCore.Password, userCore.ConfirmPassword)
	if errEmpty != nil {
		return errEmpty
	}

	errEmailValid := validator.IsEmailValid(email)
	if errEmailValid != nil {
		return errEmailValid
	}

	errLength := validator.IsMinLengthValid(userCore.Password, 10)
	if errLength != nil {
		return errLength
	}

	if userCore.Password != userCore.ConfirmPassword {
		return errors.New(constant.ERROR_OLDPASSWORD_INVALID)
	}

	HashPassword, errHash := bcrypt.HashPassword(userCore.Password)
	if errHash != nil {
		return errors.New(constant.ERROR_PASSWORD_HASH)
	}
	userCore.Password = HashPassword

	_, errNewPass := us.userRepository.NewPassword(email, userCore)
	if errNewPass != nil {
		return errNewPass
	}

	return nil
}
