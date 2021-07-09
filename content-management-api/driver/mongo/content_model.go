package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContentModel struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Fields []Field            `bson:"fields"`
}

type Field struct {
	Type string `bson:"field_type"`
}

type ContentModelDriverInterface interface {
	Create([]string) (*ContentModel, error)
}

//ContentModelDriver ContentModel Collection on MongoDB
type ContentModelDriver struct {
	Client *Client
}

func NewContentModelDriver(client *Client) *ContentModelDriver {
	return &ContentModelDriver{
		Client: client,
	}
}

func (c ContentModelDriver) Create(fields []string) (*ContentModel, error) {
	collections := c.Client.Get().Database("models").Collection("content_model")

	fieldsModel := make([]Field, len(fields))
	for i, field := range fields {
		fieldsModel[i] = Field{
			Type: field,
		}
	}

	model := ContentModel{
		Fields: fieldsModel,
	}

	result, err := collections.InsertOne(context.Background(), model)

	return &ContentModel{
		ID:     result.InsertedID.(primitive.ObjectID),
		Fields: model.Fields,
	}, err
}

func (c ContentModelDriver) Read() (ContentModel, error) {
	panic("implement me")
}

func (c ContentModelDriver) Update() (ContentModel, error) {
	panic("implement me")
}

func (c ContentModelDriver) Delete() error {
	panic("implement me")
}
