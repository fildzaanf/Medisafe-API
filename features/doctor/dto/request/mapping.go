package request

import "talkspace/features/doctor/entity"

func DoctorRegisterRequestToDoctorCore(request DoctorRegisterRequest) entity.Doctor {
	return entity.Doctor{
		Fullname:         request.Fullname,
		Email:            request.Email,
		Password:         request.Password,
		ConfirmPassword:  request.ConfirmPassword,
		ProfilePicture:   request.ProfilePicture,
		Gender:           request.Gender,
		Specialist:       request.Specialist,
		NoSTR:            request.NoSTR,
		Alumnus:          request.Alumnus,
		AboutDoctor:      request.AboutDoctor,
		LocationPractice: request.LocationPractice,
		Experience:       request.Experience,
	}
}

func DoctorLoginRequestToDoctorCore(request DoctorLoginRequest) entity.Doctor {
	return entity.Doctor{
		Email:    request.Email,
		Password: request.Password,
	}
}

func DoctorUpdateRequestToDoctorCore(request DoctorUpdateRequest) entity.Doctor {
	return entity.Doctor{
		Fullname:         request.Fullname,
		Email:            request.Email,
		ProfilePicture:   request.ProfilePicture,
		Gender:           request.Gender,
		Status:           request.Status,
		Specialist:       request.Specialist,
		NoSTR:            request.NoSTR,
		Alumnus:          request.Alumnus,
		AboutDoctor:      request.AboutDoctor,
		LocationPractice: request.LocationPractice,
		Experience:       request.Experience,
	}
}

func DoctorNewPasswordRequestToDoctorCore(request DoctorNewPasswordRequest) entity.Doctor {
	return entity.Doctor{
		Password:        request.Password,
		ConfirmPassword: request.ConfirmPassword,
	}
}

func DoctorUpdatePasswordRequestToDoctorCore(request DoctorUpdatePasswordRequest) entity.Doctor {
	return entity.Doctor{
		Password:        request.Password,
		NewPassword:     request.NewPassword,
		ConfirmPassword: request.ConfirmPassword,
	}
}

func DoctorSendOTPRequestToDoctorCore(request DoctorSendOTPRequest) entity.Doctor {
	return entity.Doctor{
		Email: request.Email,
	}
}

func DoctorVerifyOTPRequestToDoctorCore(request DoctorVerifyOTPRequest) entity.Doctor {
	return entity.Doctor{
		Email: request.Email,
		OTP:   request.OTP,
	}
}
