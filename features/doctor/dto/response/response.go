package response

type DoctorRegisterResponse struct {
	ID               string `json:"id"`
	Fullname         string `json:"fullname"`
	Email            string `json:"email"`
	Price            int    `json:"price"`
	ProfilePicture   string `json:"profile_picture"`
	Gender           string `json:"gender"`
	Status           bool   `json:"status"`
	Specialist       string `json:"specialist"`
	NoSTR            int    `json:"no_str"`
	Alumnus          string `json:"alumnus"`
	AboutDoctor      string `json:"about_doctor"`
	LocationPractice string `json:"location_practice"`
	Experience       string `json:"experience"`
}

type DoctorLoginResponse struct {
	ID       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type DoctorUpdateResponse struct {
	ID               string `json:"id"`
	Fullname         string `json:"fullname"`
	Email            string `json:"email"`
	Price            int    `json:"price"`
	ProfilePicture   string `json:"profile_picture"`
	Gender           string `json:"gender"`
	Status           bool   `json:"status"`
	Specialist       string `json:"specialist"`
	NoSTR            int    `json:"no_str"`
	Alumnus          string `json:"alumnus"`
	AboutDoctor      string `json:"about_doctor"`
	LocationPractice string `json:"location_practice"`
	Experience       string `json:"experience"`
}
