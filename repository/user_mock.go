package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepositoryMock struct {
	Users []User
}

func NewUserRepositoryMock() UserRepository {
	userss := []User{
		{UserID: "1", Name: "HadesGod", City: "Hell", Age: 22},
		{UserID: "2", Name: "TitonGod", City: "Heaven", Age: 30},
	}
	return &userRepositoryMock{Users: userss}
}

func (cm userRepositoryMock) GetAll() ([]User, error) {
	return cm.Users, nil
}

func (cm userRepositoryMock) GetById(id string) (*User, error) {
	for _, user := range cm.Users {
		if user.UserID == id {
			return &user, nil
		}
	}
	return nil, mongo.ErrNoDocuments
}

func (cm *userRepositoryMock) Insert(user User) (*User, error) {
	cm.Users = append(cm.Users, user)
	return &user, nil
}
