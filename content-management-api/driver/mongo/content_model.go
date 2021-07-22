package mongo

import (
	"content-management-api/driver"
	"content-management-api/driver/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"

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

//ContentDriver ContentModel Collection on MongoDB
type ContentDriver struct {
	Client *Client
}

func NewContentDriver(client *Client) driver.ContentDriver {
	return &ContentDriver{
		Client: client,
	}
}

func (c ContentDriver) CreateModel(name string, fields []model.Field) (*model.ContentModel, error) {
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

	if err != nil {
		return nil, err
	}

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

func (e ContentDriver) CreateEntry(entry model.Entry) (*model.Entry, error) {
	client, err := e.Client.Get()
	if err != nil {
		return nil, err
	}

	collection := client.Database("models").Collection("entry")

	insert := Entry{
		ContentModelID: entry.ModelID,
	}

	result, err := collection.InsertOne(context.Background(), insert)

	if err != nil {
		return nil, err
	}

	return &model.Entry{ID: result.InsertedID.(primitive.ObjectID).Hex()}, nil
}

func (c ContentDriver) Select() (*model.ContentModel, error) {
	panic("implement me")
}

func (c ContentDriver) Update() (*model.ContentModel, error) {
	panic("implement me")
}

func (c ContentDriver) Delete() error {
	panic("implement me")
}

func (c ContentDriver) FindContentModelByID(id string) (*model.ContentModel, error) {

	client, err := c.Client.Get()
	if err != nil {
		return nil, err
	}

	collections := client.Database("models").Collection("content_model")

	found := collections.FindOne(context.Background(), bson.M{"_id": id})

	var contentModel ContentModel
	err = found.Decode(&contentModel)
	if err != nil {
		return nil, err
	}

	resultFields := make([]model.Field, len(contentModel.Fields))
	for i, field := range contentModel.Fields {
		resultFields[i] = model.Field{
			Name:     field.Name,
			Type:     field.Type,
			Required: field.Required,
		}
	}

	return &model.ContentModel{
		ID:     contentModel.ID.Hex(),
		Name:   contentModel.Name,
		Fields: resultFields,
	}, nil
}
