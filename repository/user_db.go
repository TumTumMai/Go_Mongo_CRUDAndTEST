package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepositoryDB struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewUserRepositoryDB(collection *mongo.Collection, ctx context.Context) UserRepository {
	return userRepositoryDB{
		collection: collection,
		ctx:        ctx,
	}
}

func (c userRepositoryDB) GetAll() ([]User, error) {
	users := []User{}
	cur, err := c.collection.Find(c.ctx, bson.M{})
	if err != nil {
		return users, err
	}
	for cur.Next(c.ctx) {
		user := User{}
		cur.Decode(&user)
		users = append(users, user)
	}
	return users, nil
}

func (c userRepositoryDB) GetById(id string) (*User, error) {
	user := User{}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &user, err
	}
	err = c.collection.FindOne(c.ctx, bson.M{"_id": _id}).Decode(&user)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

///////ใส่*ถ้าไม่เป็น อเร
func (c userRepositoryDB) Insert(user User) (*User, error) {

	_, err := c.collection.InsertOne(c.ctx, user)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

// func Insert(c echo.Context) error {
// 	coletion := Dbconnet.GetDatabase().Database("TESTMVC").Collection("user")
// 	testuser := model.User{
// 		// Name: "",
// 		// City: "Samut",
// 		// Age:  22, //////,ต้องมีอยุ่ตัวสุดท้าย ถ้าเป็น Go
// 	}

// 	if err := c.Bind(&testuser); err != nil {
// 		return c.NoContent(http.StatusBadRequest)
// 	}
// 	res, err := coletion.InsertOne(Dbconnet.GetDatabasectx(), testuser)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Print(res)
// 	return c.String(http.StatusOK, "InsertOk")
// }
