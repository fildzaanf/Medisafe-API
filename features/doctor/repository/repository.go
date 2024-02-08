package repository

import (
	"encoding/json"
	"errors"
	"log"
	"talkspace/features/doctor/entity"
	"talkspace/features/doctor/model"
	"talkspace/utils/constant"
	"talkspace/utils/helper/bcrypt"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	cacheExpiration = 10 * time.Minute
)

type doctorRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewDoctorRepository(db *gorm.DB, rdb *redis.Client) entity.DoctorRepositoryInterface {
	return &doctorRepository{
		db:  db,
		rdb: rdb,
	}
}

func (dr *doctorRepository) Register(doctorCore entity.Doctor) (entity.Doctor, error) {
	request := entity.DoctorCoreToDoctorModel(doctorCore)

	result := dr.db.Create(&request)
	if result.Error != nil {
		return entity.Doctor{}, result.Error
	}

	response := entity.DoctorModelToDoctorCore(request)
	return response, nil
}

func (dr *doctorRepository) Login(email, password string) (entity.Doctor, error) {
	doctorModel := model.Doctor{}

	result := dr.db.Where("email = ?", email).First(&doctorModel)
	if result.Error != nil {
		return entity.Doctor{}, result.Error
	}

	if result.RowsAffected == 0 {
		return entity.Doctor{}, errors.New(constant.ERROR_EMAIL_NOTFOUND)
	}

	if errComparePass := bcrypt.ComparePassword(doctorModel.Password, password); errComparePass != nil {
		return entity.Doctor{}, errors.New(constant.ERROR_PASSWORD_INVALID)
	}

	response := entity.DoctorModelToDoctorCore(doctorModel)
	return response, nil
}

func (dr *doctorRepository) GetByID(id string) (entity.Doctor, error) {
	doctorModel := model.Doctor{}

	cacheKey := "doctor:" + id
	cachedDoctor, err := dr.rdb.Get(cacheKey).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(cachedDoctor), &doctorModel); err != nil {
			log.Printf("error unmarshalling cached doctor data: %v", err)
		} else {
			response := entity.DoctorModelToDoctorCore(doctorModel)
			return response, nil
		}
	}

	result := dr.db.Where("id = ?", id).First(&doctorModel)
	if result.Error != nil {
		return entity.Doctor{}, result.Error
	}

	if result.RowsAffected == 0 {
		return entity.Doctor{}, errors.New(constant.ERROR_ID_NOTFOUND)
	}

	doctorModelJSON, err := json.Marshal(doctorModel)
	if err != nil {
		log.Printf("error marshalling doctor data to JSON: %v", err)
	} else {
		err := dr.rdb.Set(cacheKey, doctorModelJSON, cacheExpiration).Err()
		if err != nil {
			log.Printf("error caching doctor data: %v", err)
		}
	}

	response := entity.DoctorModelToDoctorCore(doctorModel)
	return response, nil
}

func (dr *doctorRepository) UpdateByID(id string, doctorCore entity.Doctor) error {
	request := entity.DoctorCoreToDoctorModel(doctorCore)

	result := dr.db.Where("id = ?", id).Updates(&request)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New(constant.ERROR_ID_NOTFOUND)
	}

	return nil
}

func (dr *doctorRepository) FindByEmail(email string) (entity.Doctor, error) {
	doctorModel := model.Doctor{}

	cacheKey := "doctor:email:" + email
	cachedDoctor, err := dr.rdb.Get(cacheKey).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(cachedDoctor), &doctorModel); err != nil {
			log.Printf("error unmarshalling cached doctor data: %v", err)
		} else {
			response := entity.DoctorModelToDoctorCore(doctorModel)
			return response, nil
		}
	}

	result := dr.db.Where("email = ?", email).First(&doctorModel)

	if result.RowsAffected == 0 {
		return entity.Doctor{}, errors.New(constant.ERROR_EMAIL_NOTFOUND)
	}

	if result.Error != nil {
		return entity.Doctor{}, result.Error
	}

	doctorModelJSON, err := json.Marshal(doctorModel)
	if err != nil {
		log.Printf("error marshalling doctor data to JSON: %v", err)
	} else {
		err := dr.rdb.Set(cacheKey, doctorModelJSON, cacheExpiration).Err()
		if err != nil {
			log.Printf("error caching doctor data: %v", err)
		}
	}

	response := entity.DoctorModelToDoctorCore(doctorModel)
	return response, nil
}

func (dr *doctorRepository) GetByVerificationToken(token string) (entity.Doctor, error) {
	doctorModel := model.Doctor{}

	cacheKey := "doctor:token:" + token
	cachedDoctor, err := dr.rdb.Get(cacheKey).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(cachedDoctor), &doctorModel); err != nil {
			log.Printf("error unmarshalling cached doctor data: %v", err)
		} else {
			response := entity.DoctorModelToDoctorCore(doctorModel)
			return response, nil
		}
	}

	result := dr.db.Where("verification_token = ?", token).First(&doctorModel)
	if result.Error != nil {
		return entity.Doctor{}, result.Error
	}

	if result.RowsAffected == 0 {
		return entity.Doctor{}, errors.New(constant.ERROR_TOKEN_NOTFOUND)
	}

	doctorModelJSON, err := json.Marshal(doctorModel)
	if err != nil {
		log.Printf("error marshalling doctor data to JSON: %v", err)
	} else {
		err := dr.rdb.Set(cacheKey, doctorModelJSON, cacheExpiration).Err()
		if err != nil {
			log.Printf("error caching doctor data: %v", err)
		}
	}

	doctorToken := entity.DoctorModelToDoctorCore(doctorModel)
	return doctorToken, nil
}

func (dr *doctorRepository) UpdateIsVerified(id string, isVerified bool) error {
	doctorModel := model.Doctor{}

	result := dr.db.Where("id = ?", id).First(&doctorModel)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New(constant.ERROR_ID_NOTFOUND)
	}

	doctorModel.IsVerified = isVerified

	errSave := dr.db.Save(&doctorModel)
	if errSave.Error != nil {
		return errSave.Error
	}

	return nil
}

func (dr *doctorRepository) SendOTP(email string, otp string, expired int64) (doctorCore entity.Doctor, err error) {
	doctorModel := model.Doctor{}

	result := dr.db.Where("email = ?", email).First(&doctorModel)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return entity.Doctor{}, errors.New(constant.ERROR_EMAIL_NOTFOUND)
		}
		return entity.Doctor{}, result.Error
	}

	doctorModel.OTP = otp
	doctorModel.OTPExpiration = expired

	errUpdate := dr.db.Save(&doctorModel).Error
	if errUpdate != nil {
		return entity.Doctor{}, errUpdate
	}

	response := entity.DoctorModelToDoctorCore(doctorModel)

	return response, nil
}

func (dr *doctorRepository) VerifyOTP(email, otp string) (entity.Doctor, error) {
	doctorModel := model.Doctor{}

	result := dr.db.Where("otp = ? AND email = ?", otp, email).First(&doctorModel)
	if result.Error != nil {
		return entity.Doctor{}, result.Error
	}

	if result.RowsAffected == 0 {
		return entity.Doctor{}, errors.New(constant.ERROR_EMAIL_OTP)
	}

	response := entity.DoctorModelToDoctorCore(doctorModel)
	return response, nil
}

func (dr *doctorRepository) ResetOTP(otp string) (doctorCore entity.Doctor, err error) {
	doctorModel := model.Doctor{}

	result := dr.db.Where("otp = ?", otp).First(&doctorModel)
	if result.Error != nil {
		return entity.Doctor{}, result.Error
	}

	if result.RowsAffected == 0 {
		return entity.Doctor{}, errors.New(constant.ERROR_OTP_NOTFOUND)
	}

	doctorModel.OTP = ""
	doctorModel.OTPExpiration = 0

	errUpdate := dr.db.Save(&doctorModel).Error
	if errUpdate != nil {
		return entity.Doctor{}, errUpdate
	}

	response := entity.DoctorModelToDoctorCore(doctorModel)
	return response, nil
}

func (dr *doctorRepository) UpdatePassword(id string, doctorCore entity.Doctor) error {

	request := entity.DoctorCoreToDoctorModel(doctorCore)

	result := dr.db.Where("id = ?", id).Updates(&request)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New(constant.ERROR_ID_NOTFOUND)
	}

	return nil
}

func (dr *doctorRepository) NewPassword(email string, doctorCore entity.Doctor) (entity.Doctor, error) {
	doctorModel := model.Doctor{}

	result := dr.db.Where("email = ?", email).First(&doctorModel)
	if result.Error != nil {
		return entity.Doctor{}, result.Error
	}

	if result.RowsAffected == 0 {
		return entity.Doctor{}, errors.New(constant.ERROR_EMAIL_NOTFOUND)
	}

	errUpdate := dr.db.Model(&doctorModel).Updates(entity.DoctorCoreToDoctorModel(doctorCore))
	if errUpdate != nil {
		return entity.Doctor{}, errUpdate.Error
	}

	response := entity.DoctorModelToDoctorCore(doctorModel)

	return response, nil
}
