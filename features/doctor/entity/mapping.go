package entity

import "talkspace/features/doctor/model"

func DoctorCoreToDoctorModel(doctorCore Doctor) model.Doctor {
	doctorModel := model.Doctor{
		ID:               doctorCore.ID,
		Fullname:         doctorCore.Fullname,
		Email:            doctorCore.Email,
		Password:         doctorCore.Password,
		ProfilePicture:   doctorCore.ProfilePicture,
		Gender:           doctorCore.Gender,
		Status:           doctorCore.Status,
		Price:            doctorCore.Price,
		Specialist:       doctorCore.Specialist,
		Experience:       doctorCore.Experience,
		NoSTR:            doctorCore.NoSTR,
		Role:             doctorCore.Role,
		Alumnus:          doctorCore.Alumnus,
		AboutDoctor:      doctorCore.AboutDoctor,
		LocationPractice: doctorCore.LocationPractice,
		OTP:              doctorCore.OTP,
		OTPExpiration:    doctorCore.OTPExpiration,
		VerifyAccount:    doctorCore.VerifyAccount,
		IsVerified:       doctorCore.IsVerified,
		UpdatedAt:        doctorCore.UpdatedAt,
		CreatedAt:        doctorCore.CreatedAt,
		DeletedAt:        doctorCore.DeletedAt,
	}
	return doctorModel
}

func ListDoctorCoreToDoctorModel(doctorCore []Doctor) []model.Doctor {
	listDoctorModel := []model.Doctor{}
	for _, doctor := range doctorCore {
		doctorModel := DoctorCoreToDoctorModel(doctor)
		listDoctorModel = append(listDoctorModel, doctorModel)
	}
	return listDoctorModel
}

func DoctorModelToDoctorCore(doctorModel model.Doctor) Doctor {
	doctorCore := Doctor{
		ID:               doctorModel.ID,
		Fullname:         doctorModel.Fullname,
		Email:            doctorModel.Email,
		Password:         doctorModel.Password,
		ProfilePicture:   doctorModel.ProfilePicture,
		Gender:           doctorModel.Gender,
		Status:           doctorModel.Status,
		Price:            doctorModel.Price,
		Specialist:       doctorModel.Specialist,
		Experience:       doctorModel.Experience,
		NoSTR:            doctorModel.NoSTR,
		Role:             doctorModel.Role,
		Alumnus:          doctorModel.Alumnus,
		AboutDoctor:      doctorModel.AboutDoctor,
		LocationPractice: doctorModel.LocationPractice,
		OTP:              doctorModel.OTP,
		OTPExpiration:    doctorModel.OTPExpiration,
		VerifyAccount:    doctorModel.VerifyAccount,
		IsVerified:       doctorModel.IsVerified,
		UpdatedAt:        doctorModel.UpdatedAt,
		CreatedAt:        doctorModel.CreatedAt,
		DeletedAt:        doctorModel.DeletedAt,
	}
	return doctorCore
}

func ListDoctorModelToDoctorCore(doctorModel []model.Doctor) []Doctor {
	listDoctorCore := []Doctor{}
	for _, doctor := range doctorModel {
		doctorCore := DoctorModelToDoctorCore(doctor)
		listDoctorCore = append(listDoctorCore, doctorCore)
	}
	return listDoctorCore
}
