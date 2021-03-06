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
	Insert(UserResponse) (*UserResponse, error)
	Update(string, UserResponse) (*UserResponse, error)
	DeleteById(string) (*string, error)
}
