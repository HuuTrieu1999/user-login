package request

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	UserName    string `json:"userName"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Birthday    string `json:"birthday"`
	Password    string `json:"password"`
}
