package mongo

import (
	"context"
	"fmt"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (u UserDriver) Register(user model.User) (model.User, error) {
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

func (u UserDriver) DeleteById(id string) (int64, error) {
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
