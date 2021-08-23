package repository

type User struct {
	UserID string `json:"customerId,omitempty" bson:"_id,omitempty"`
	Name   string `json:"name,omitempty" bson:"name,omitempty"`
	City   string `json:"city,omitempty" bson:"city,omitempty"`
	Age    int    `json:"age,omitempty" bson:"age,omitempty"`
}

type UserRepository interface { /////(รับเเค่ type)
	GetAll() ([]User, error)
	GetById(string) (*User, error)
	Insert(User) (*User, error)
	Update(string, User) (*User, error)
}
