package request

type UserRegisterRequest struct {
	Fullname        string `json:"fullname" form:"fullname"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type UserLoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserUpdateProfileRequest struct {
	Fullname       string `json:"fullname" form:"fullname"`
	Email          string `json:"email" form:"email"`
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
	Gender         string `json:"gender" form:"gender"`
	Birthdate      string `json:"birthdate" form:"birthdate"`
	BloodType      string `json:"blood_type" form:"blood_type"`
	Height         int    `json:"height" form:"height"`
	Weight         int    `json:"weight" form:"weight"`
}

type UserNewPasswordRequest struct {
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type UserUpdatePasswordRequest struct {
	Password        string `json:"password" form:"password"`
	NewPassword     string `json:"new_password" form:"new_password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type UserSendOTPRequest struct {
	Email string `json:"email" form:"email"`
}

type UserVerifyOTPRequest struct {
	Email string `json:"email" form:"email"`
	OTP   string `json:"otp" form:"otp"`
}
