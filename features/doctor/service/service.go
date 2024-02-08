package service

import (
	"errors"
	"talkspace/features/doctor/entity"
	"talkspace/middlewares"
	"talkspace/utils/constant"
	"talkspace/utils/helper/bcrypt"
	"talkspace/utils/helper/email/mailer/onetimepassword"
	"talkspace/utils/helper/email/mailer/verification"
	"talkspace/utils/helper/generator"
	"talkspace/utils/validator"
	"time"
)

type doctorService struct {
	doctorRepository entity.DoctorRepositoryInterface
}

func NewDoctorService(dr entity.DoctorRepositoryInterface) entity.DoctorServiceInterface {
	return &doctorService{
		doctorRepository: dr,
	}
}

func (ds *doctorService) Register(doctorCore entity.Doctor) (entity.Doctor, error) {

	errEmpty := validator.IsDataEmpty(
		doctorCore.Fullname,
		doctorCore.Email,
		doctorCore.Password,
		doctorCore.ConfirmPassword,
		doctorCore.ProfilePicture,
		doctorCore.Gender,
		doctorCore.Specialist,
		doctorCore.NoSTR,
		doctorCore.Alumnus,
		doctorCore.AboutDoctor,
		doctorCore.LocationPractice,
		doctorCore.Experience)
	if errEmpty != nil {
		return entity.Doctor{}, errEmpty
	}

	errEmailValid := validator.IsEmailValid(doctorCore.Email)
	if errEmailValid != nil {
		return entity.Doctor{}, errEmailValid
	}

	errLength := validator.IsMinLengthValid(doctorCore.Password, 10)
	if errLength != nil {
		return entity.Doctor{}, errLength
	}

	validGender := []interface{}{"male", "female"}
	errGender := validator.IsDataValid(doctorCore.Gender, validGender, true)
	if errGender != nil {
		return entity.Doctor{}, errGender
	}

	_, errFindEmail := ds.doctorRepository.FindByEmail(doctorCore.Email)
	if errFindEmail == nil {
		return entity.Doctor{}, errors.New(constant.ERROR_EMAIL_EXIST)
	}

	if doctorCore.Password != doctorCore.ConfirmPassword {
		return entity.Doctor{}, errors.New(constant.ERROR_PASSWORD_CONFIRM)
	}

	hashedPassword, errHash := bcrypt.HashPassword(doctorCore.Password)
	if errHash != nil {
		return entity.Doctor{}, errors.New(constant.ERROR_PASSWORD_HASH)
	}
	doctorCore.Password = hashedPassword

	token, errGenerateVerifyToken := generator.GenerateRandomBytes()
	if errGenerateVerifyToken != nil {
		errors.New(constant.ERROR_TOKEN_VERIFICATION)
	}
	doctorCore.VerifyAccount = token

	doctorData, errRegister := ds.doctorRepository.Register(doctorCore)
	if errRegister != nil {
		return entity.Doctor{}, errRegister
	}

	verification.SendEmailVerificationAccount(doctorCore.Email, token)

	return doctorData, nil
}

func (ds *doctorService) Login(email, password string) (entity.Doctor, string, error) {

	errEmpty := validator.IsDataEmpty(email, password)
	if errEmpty != nil {
		return entity.Doctor{}, "", errEmpty
	}

	errEmailValid := validator.IsEmailValid(email)
	if errEmailValid != nil {
		return entity.Doctor{}, "", errEmailValid
	}

	doctorData, errFindEmail := ds.doctorRepository.FindByEmail(email)
	if errFindEmail != nil {
		return entity.Doctor{}, "", errors.New(constant.ERROR_EMAIL_UNREGISTERED)
	}

	if !doctorData.IsVerified {
		return entity.Doctor{}, "", errors.New(constant.ERROR_ACCOUNT_UNVERIFIED)
	}

	comparePassword := bcrypt.ComparePassword(doctorData.Password, password)
	if comparePassword != nil {
		return entity.Doctor{}, "", errors.New(constant.ERROR_LOGIN)
	}

	token, errCreateToken := middlewares.CreateToken(doctorData.ID, "")
	if errCreateToken != nil {
		return entity.Doctor{}, "", errors.New(constant.ERROR_TOKEN_GENERATE)
	}
	return doctorData, token, nil
}

func (ds *doctorService) GetByID(id string) (entity.Doctor, error) {
	if id == "" {
		return entity.Doctor{}, errors.New(constant.ERROR_ID_INVALID)
	}

	doctorData, errGetID := ds.doctorRepository.GetByID(id)
	if errGetID != nil {
		return entity.Doctor{}, errors.New(constant.ERROR_DATA_EMPTY)
	}
	return doctorData, nil
}

func (ds *doctorService) UpdateByID(id string, doctorCore entity.Doctor) error {
	if id == "" {
		return errors.New(constant.ERROR_ID_INVALID)
	}

	_, errGetID := ds.doctorRepository.GetByID(id)
	if errGetID != nil {
		return errGetID
	}

	errEmailValid := validator.IsEmailValid(doctorCore.Email)
	if errEmailValid != nil {
		return errEmailValid
	}

	validGender := []interface{}{"male", "female"}
	errGender := validator.IsDataValid(doctorCore.Gender, validGender, true)
	if errGender != nil {
		return errGender
	}

	errUpdate := ds.doctorRepository.UpdateByID(id, doctorCore)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

func (ds *doctorService) UpdatePassword(id string, doctorCore entity.Doctor) error {
	if id == "" {
		return errors.New(constant.ERROR_ID_INVALID)
	}

	result, errGetID := ds.GetByID(id)
	if errGetID != nil {
		return errGetID
	}

	errEmpty := validator.IsDataEmpty(doctorCore.Password, doctorCore.NewPassword, doctorCore.ConfirmPassword)
	if errEmpty != nil {
		return errEmpty
	}

	errLength := validator.IsMinLengthValid(doctorCore.NewPassword, 10)
	if errLength != nil {
		return errLength
	}

	comparePassword := bcrypt.ComparePassword(result.Password, doctorCore.Password)
	if comparePassword != nil {
		return errors.New(constant.ERROR_OLDPASSWORD_INVALID)
	}

	if doctorCore.NewPassword != doctorCore.ConfirmPassword {
		return errors.New(constant.ERROR_PASSWORD_CONFIRM)
	}

	HashPassword, errHash := bcrypt.HashPassword(doctorCore.NewPassword)
	if errHash != nil {
		return errors.New(constant.ERROR_PASSWORD_HASH)
	}
	doctorCore.Password = HashPassword

	errUpdate := ds.doctorRepository.UpdatePassword(id, doctorCore)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

func (ds *doctorService) VerifyDoctor(token string) (bool, error) {
	if token == "" {
		return false, errors.New(constant.ERROR_TOKEN_INVALID)
	}

	doctor, errGetVerifyToken := ds.doctorRepository.GetByVerificationToken(token)
	if errGetVerifyToken != nil {
		return false, errors.New(constant.ERROR_DATA_RETRIEVED)
	}

	if doctor.IsVerified {
		return true, nil
	}

	errUpdate := ds.doctorRepository.UpdateIsVerified(doctor.ID, true)
	if errUpdate != nil {
		return false, errors.New(constant.ERROR_ACCOUNT_VERIFICATION)
	}

	return false, nil
}

func (ds *doctorService) UpdateIsVerified(id string, isVerified bool) error {
	if id == "" {
		return errors.New(constant.ERROR_ID_INVALID)
	}

	return ds.doctorRepository.UpdateIsVerified(id, isVerified)
}

func (ds *doctorService) SendOTP(email string) error {

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

	_, errSend := ds.doctorRepository.SendOTP(email, otp, expired)
	if errSend != nil {
		return errSend
	}

	onetimepassword.SendEmailOTP(email, otp)
	return nil
}

func (ds *doctorService) VerifyOTP(email, otp string) (string, error) {

	errEmpty := validator.IsDataEmpty(email, otp)
	if errEmpty != nil {
		return "", errEmpty
	}

	doctorData, err := ds.doctorRepository.VerifyOTP(email, otp)
	if err != nil {
		return "", errors.New(constant.ERROR_EMAIL_OTP)
	}

	if doctorData.OTPExpiration <= time.Now().Unix() {
		return "", errors.New(constant.ERROR_OTP_EXPIRED)
	}

	if doctorData.OTP != otp {
		return "", errors.New(constant.ERROR_OTP_INVALID)
	}

	token, err := middlewares.CreateVerifyToken(email)
	if err != nil {
		return "", errors.New(constant.ERROR_TOKEN_GENERATE)
	}

	_, errReset := ds.doctorRepository.ResetOTP(otp)
	if errReset != nil {
		return "", errors.New(constant.ERROR_OTP_RESET)
	}

	return token, nil
}

func (ds *doctorService) NewPassword(email string, doctorCore entity.Doctor) error {

	errEmpty := validator.IsDataEmpty(email, doctorCore.Password, doctorCore.ConfirmPassword)
	if errEmpty != nil {
		return errEmpty
	}

	errEmailValid := validator.IsEmailValid(email)
	if errEmailValid != nil {
		return errEmailValid
	}

	errLength := validator.IsMinLengthValid(doctorCore.Password, 10)
	if errLength != nil {
		return errLength
	}

	if doctorCore.Password != doctorCore.ConfirmPassword {
		return errors.New(constant.ERROR_OLDPASSWORD_INVALID)
	}

	HashPassword, errHash := bcrypt.HashPassword(doctorCore.Password)
	if errHash != nil {
		return errors.New(constant.ERROR_PASSWORD_HASH)
	}
	doctorCore.Password = HashPassword

	_, errNewPass := ds.doctorRepository.NewPassword(email, doctorCore)
	if errNewPass != nil {
		return errNewPass
	}

	return nil
}
