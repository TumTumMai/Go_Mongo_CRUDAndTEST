package service

import (
	"fmt"
	"testgo/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewuserService(usRepo repository.UserRepository) userService {
	return userService{
		userRepository: usRepo,
	}
}

func (us userService) GetUsers() ([]UserResponse, error) {

	users, err := us.userRepository.GetAll()
	if err != nil {
		// this is technical error should not sent to user
		// should be keep in log file
		fmt.Printf("err: %v\n", err)
		return []UserResponse{}, err
	}

	// transfer data to response model
	cusResponses := []UserResponse{}
	for _, user := range users {
		cusRes := UserResponse{
			ID:   user.UserID,
			Name: user.Name,
			City: user.City,
			Age:  user.Age,
		}
		cusResponses = append(cusResponses, cusRes)
	}
	return cusResponses, nil
}

func (us userService) GetUser(id string) (*UserResponse, error) {
	user, err := us.userRepository.GetById(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("not found")
		}
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	cusResponse := UserResponse{
		Name: user.Name,
		City: user.City,
		Age:  user.Age,
	}
	return &cusResponse, nil
}

func (us userService) Insert(user UserResponse) (*UserResponse, error) {
	userInsert := repository.User{
		UserID: user.ID,
		Name:   user.Name,
		City:   user.City,
		Age:    user.Age,
	}
	_, err := us.userRepository.Insert(userInsert)
	if err != nil {
		panic(err)
	}
	return &user, nil
}
