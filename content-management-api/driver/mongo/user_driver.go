package mongo

import (
	"context"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver/model"
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

func (u UserDriver) Register(user model.User) (*model.User, error) {
	client, err := u.Client.Get()
	if err != nil {
		return nil, err
	}

	collections := client.Database("user").Collection("user")

	insert := User{
		Name: user.Name,
	}

	result, err := collections.InsertOne(context.Background(), insert)
	if err != nil {
		return nil, err
	}

	return &model.User{
		Id:   result.InsertedID.(primitive.ObjectID).Hex(),
		Name: insert.Name,
	}, nil
}
