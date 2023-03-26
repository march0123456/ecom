package service

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserService interface {
	SignUp(UserRequest) (*UserResponse, error)
	Login(UserRequest) (*UserResponse, error)
}
