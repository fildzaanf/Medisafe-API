package response

import "talkspace/features/doctor/entity"

func DoctorCoreToDoctorRegisterResponse(response entity.Doctor) DoctorRegisterResponse {
	return DoctorRegisterResponse{
		ID:               response.ID,
		Fullname:         response.Fullname,
		Email:            response.Email,
		Price:            response.Price,
		ProfilePicture:   response.ProfilePicture,
		Gender:           response.Gender,
		Status:           response.Status,
		Specialist:       response.Specialist,
		NoSTR:            response.NoSTR,
		Alumnus:          response.Alumnus,
		AboutDoctor:      response.AboutDoctor,
		LocationPractice: response.LocationPractice,
		Experience:       response.Experience,
	}
}

func DoctorCoreToDoctorLoginResponse(response entity.Doctor, token string) DoctorLoginResponse {
	return DoctorLoginResponse{
		ID:       response.ID,
		Fullname: response.Fullname,
		Email:    response.Email,
		Token:    token,
	}
}

func DoctorCoreToDoctorUpdateResponse(response entity.Doctor) DoctorUpdateResponse {
	return DoctorUpdateResponse{
		ID:               response.ID,
		Fullname:         response.Fullname,
		Email:            response.Email,
		Price:            response.Price,
		ProfilePicture:   response.ProfilePicture,
		Gender:           response.Gender,
		Status:           response.Status,
		Specialist:       response.Specialist,
		NoSTR:            response.NoSTR,
		Alumnus:          response.Alumnus,
		AboutDoctor:      response.AboutDoctor,
		LocationPractice: response.LocationPractice,
		Experience:       response.Experience,
	}
}
