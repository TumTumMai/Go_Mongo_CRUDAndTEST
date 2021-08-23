package main

import (
	"context"
	"fmt"
	"testgo/config"
	"testgo/controller"
	"testgo/repository"
	"testgo/service"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	config := config.GetConfig()
	initTimeZone()
	mongoPath := fmt.Sprintf("mongodb+srv://%v:%v@cluster0.mqkff.mongodb.net/TESTMVC?retryWrites=true&w=majority", config.Mongo.Username, config.Mongo.Password)
	connection := options.Client().ApplyURI(mongoPath)
	ctx := context.TODO()

	client, err := mongo.Connect(ctx, connection)
	if err != nil {
		panic(err)
	}

	db := client.Database("TESTMVC")
	collection := db.Collection("user")

	userRepository := repository.NewUserRepositoryDB(collection, ctx)
	_ = userRepository
	userRepositoryMock := repository.NewUserRepositoryMock()
	_ = userRepositoryMock

	userService := service.NewuserService(userRepository)
	userHandler := controller.NewuserHandler(userService)

	e := echo.New()
	e.GET("/users", userHandler.GetAllUser)
	e.GET("/user/:id", userHandler.GetUserById)
	e.POST("/user", userHandler.Insert)
	e.PUT("user/:id", userHandler.UpdateOne)

	e.Start(fmt.Sprintf(":%v", config.App.Port))

}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}
