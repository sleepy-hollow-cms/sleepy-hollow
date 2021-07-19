package mongo

import (
	"content-management-api/driver"
	"content-management-api/driver/model"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entry struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}

type EntryDriver struct {
	Client *Client
}

func NewEntryDriver(client *Client) driver.Entry {
	return &EntryDriver{
		Client: client,
	}
}

func (e EntryDriver) Create() (*model.Entry, error) {
	client, err := e.Client.Get()
	if err != nil {
		return nil, err
	}

	collection := client.Database("models").Collection("entry")

	insert := Entry{
		Name: "Not implement yet",
	}

	result, err := collection.InsertOne(context.Background(), insert)

	if err != nil {
		return nil, err
	}

	return &model.Entry{ID: result.InsertedID.(primitive.ObjectID).Hex()}, nil
}
