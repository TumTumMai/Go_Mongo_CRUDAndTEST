package service

type UserResponse struct {
	ID   string `json:"Id,omitempty"`
	Name string `json:"name,omitempty"`
	City string `json:"city,omitempty"`
	Age  int    `json:"age,omitempty"`
}

type UserService interface {
	GetUsers() ([]UserResponse, error)
	GetUser(string) (*UserResponse, error)
	Insert(user UserResponse) (*UserResponse, error)
}
