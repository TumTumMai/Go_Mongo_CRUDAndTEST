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

func (c userRepositoryDB) Update(id string, user User) (*User, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	_, err = c.collection.UpdateOne(c.ctx, bson.M{"_id": _id}, bson.M{"$set": &user})
	if err != nil {
		return &user, err
	}
	return &user, nil
}

//////(id string) พารามิเตอร์ ตัวรับค่า (string, error)เป็นที่จะส่งออกไป

func (c userRepositoryDB) DeleteById(id string) (string, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	_, err = c.collection.DeleteOne(c.ctx, bson.M{"_id": _id})
	if err != nil {
		panic(err)
	}

	return id, nil
}

// id := c.Param("id")

// 	coletion := Dbconnet.GetDatabase().Database("TESTMVC").Collection("user")
// 	// objid, err := primitive.ObjectIDFromHex("6112341143062e8d442b584c")
// 	objid, err := primitive.ObjectIDFromHex(id)

// 	if err != nil {
// 		panic(err)
// 	}
// 	// var result = *mongo.DeleteResult
// 	_, err = coletion.DeleteOne(Dbconnet.GetDatabasectx(), bson.M{"_id": objid})
// 	if err != nil {
// 		panic(err)
// 	}
// 	// fmt.Println(result)

// 	return c.String(http.StatusOK, "Delsusecfull")
