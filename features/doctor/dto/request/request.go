package request

type DoctorRegisterRequest struct {
	Fullname         string `json:"fullname" form:"fullname"`
	Email            string `json:"email" form:"email"`
	Password         string `json:"password" form:"password"`
	ConfirmPassword  string `json:"confirm_password" form:"confirm_password"`
	ProfilePicture   string `json:"profile_picture" form:"profile_picture"`
	Gender           string `json:"gender" form:"gender"`
	Specialist       string `json:"specialist" form:"specialist"`
	NoSTR            int    `json:"no_str" form:"no_str"`
	Alumnus          string `json:"alumnus" form:"alumnus"`
	AboutDoctor      string `json:"about_doctor" form:"about_doctor"`
	LocationPractice string `json:"location_practice" form:"location_practice"`
	Experience       string `json:"experience" form:"experience"`
}

type DoctorLoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type DoctorProfileRequest struct {
	Fullname         string `json:"fullname" form:"fullname"`
	Email            string `json:"email" form:"email"`
	ProfilePicture   string `json:"profile_picture" form:"profile_picture"`
	Gender           string `json:"gender" form:"gender"`
	Status           bool   `json:"status" form:"status"`
	Specialist       string `json:"specialist" form:"specialist"`
	NoSTR            int    `json:"no_str" form:"no_str"`
	Alumnus          string `json:"alumnus" form:"alumnus"`
	AboutDoctor      string `json:"about_doctor" form:"about_doctor"`
	LocationPractice string `json:"location_practice" form:"location_practice"`
	Experience       string `json:"experience" form:"experience"`
}

type DoctorNewPasswordRequest struct {
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type DoctorUpdatePasswordRequest struct {
	Password        string `json:"password" form:"password"`
	NewPassword     string `json:"new_password" form:"new_password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type DoctorSendOTPRequest struct {
	Email string `json:"email" form:"email"`
}

type DoctorVerifyOTPRequest struct {
	Email string `json:"email" form:"email"`
	OTP   string `json:"otp" form:"otp"`
}