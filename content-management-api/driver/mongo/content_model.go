package mongo

import (
	"content-management-api/driver"
	"content-management-api/driver/model"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContentModel struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Fields []Field            `bson:"fields"`
	Name   string             `bson:"name"`
}

type Field struct {
	Name     string `bson:"name"`
	Type     string `bson:"field_type"`
	Required bool   `bson:"required"`
}

//ContentModelDriver ContentModel Collection on MongoDB
type ContentModelDriver struct {
	Client *Client
}

func NewContentModelDriver(client *Client) driver.ContentModel {
	return &ContentModelDriver{
		Client: client,
	}
}

func (c ContentModelDriver) Create(name string, fields []model.Field) (*model.ContentModel, error) {
	client, err := c.Client.Get()
	if err != nil {
		return nil, err
	}

	collections := client.Database("models").Collection("content_model")

	fieldsModel := make([]Field, len(fields))
	for i, field := range fields {
		fieldsModel[i] = Field{
			Name:     field.Name,
			Type:     field.Type,
			Required: field.Required,
		}
	}

	insert := ContentModel{
		Name:   name,
		Fields: fieldsModel,
	}

	result, err := collections.InsertOne(context.Background(), insert)

	resultFields := make([]model.Field, len(insert.Fields))
	for i, field := range insert.Fields {
		resultFields[i] = model.Field{
			Name:     field.Name,
			Type:     field.Type,
			Required: field.Required,
		}
	}

	return &model.ContentModel{
		ID:     result.InsertedID.(primitive.ObjectID).Hex(),
		Name:   insert.Name,
		Fields: resultFields,
	}, err
}

func (c ContentModelDriver) Select() (*model.ContentModel, error) {
	panic("implement me")
}

func (c ContentModelDriver) Update() (*model.ContentModel, error) {
	panic("implement me")
}

func (c ContentModelDriver) Delete() error {
	panic("implement me")
}
