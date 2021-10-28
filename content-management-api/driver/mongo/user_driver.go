package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver/model"
)

type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}

type UserDriver struct {
	Client *Client
}

func NewUserDriver(client *Client) driver.UserDriver {
	return &UserDriver{
		Client: client,
	}
}

func (u *UserDriver) Register(user model.User) (model.User, error) {
	client, err := u.Client.Get()
	if err != nil {
		return model.User{}, err
	}

	collections := client.Database("user").Collection("user")

	insert := User{
		Name: user.Name,
	}

	result, err := collections.InsertOne(context.Background(), insert)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		Id:   result.InsertedID.(primitive.ObjectID).Hex(),
		Name: insert.Name,
	}, nil
}

func (u *UserDriver) DeleteById(id string) (int64, error) {
	client, err := u.Client.Get()
	if err != nil {
		return 0, err
	}

	collections := client.Database("user").Collection("user")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	deletedResult, err := collections.DeleteOne(context.Background(), bson.M{"_id": objectId})

	if err != nil {
		return 0, err
	}

	if deletedResult.DeletedCount == 0 {
		return deletedResult.DeletedCount, driver.NewUserNotFound(fmt.Sprintf("user is not found: %s", id))
	}

	return deletedResult.DeletedCount, nil
}

func (u *UserDriver) FindById(userId string) (*model.User, error) {
	client, err := u.Client.Get()

	if err != nil {
		return nil, err
	}

	collections := client.Database("user").Collection("user")

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	findOne := collections.FindOne(context.Background(), bson.M{"_id": objectId})

	user := User{}
	err = findOne.Decode(&user)

	return &model.User{
		Id:   user.ID.Hex(),
		Name: user.Name,
	}, nil
}

func (u *UserDriver) Update(updatedUser model.User) (*model.User, error) {
	client, err := u.Client.Get()
	if err != nil {
		return nil, err
	}

	collections := client.Database("user").Collection("user")

	objectId, err := primitive.ObjectIDFromHex(updatedUser.Id)
	if err != nil {
		return nil, err
	}

	update := User{
		Name: updatedUser.Name,
	}

	updateResult, errUpdate := collections.UpdateOne(
		context.Background(),
		bson.M{
			"$and": []bson.M{
				{"_id": objectId},
			},
		},
		bson.D{{"$set", update}},
	)

	if errUpdate != nil {
		return nil, err
	}

	if updateResult.MatchedCount == 0 {
		return nil, driver.UserCannotUpdateError{}
	}

	return &model.User{
		Id:   updatedUser.Id,
		Name: update.Name,
	}, err
}

