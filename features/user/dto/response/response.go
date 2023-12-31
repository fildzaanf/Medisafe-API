package response

type UserRegisterResponse struct {
	ID       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type UserLoginResponse struct {
	ID       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type UserProfileResponse struct {
	ID             string `json:"id"`
	Fullname       string `json:"fullname"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
	Gender         string `json:"gender"`
	Birthdate      string `json:"birthdate"`
	BloodType      string `json:"blood_type"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
}
