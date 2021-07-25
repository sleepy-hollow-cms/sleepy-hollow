package mongo

import (
	"content-management-api/driver"
	"content-management-api/driver/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, driver.NewContentModelCannotFindById(fmt.Sprintf("%s is invalid format Id", id))
	}

	found := collections.FindOne(context.Background(), bson.M{"_id": objectId})

	var contentModel ContentModel
	if found.Err() == mongo.ErrNoDocuments {
		return nil, driver.NewContentModelCannotFindById(fmt.Sprintf("%s is not registered", id))
	}

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

func (c ContentDriver) FindContentModelBySpaceID(id string) ([]model.ContentModel, error) {

	client, err := c.Client.Get()
	if err != nil {
		return nil, err
	}

	collections := client.Database("models").Collection("content_model")

	found, err := collections.Find(context.Background(), bson.M{"space_id": id})
	if err != nil {
		return nil, err
	}
	defer found.Close(context.Background())
	contentModels := make([]ContentModel, found.RemainingBatchLength())
	for i := 0; found.Next(context.Background()); i++ {
		var contentModel ContentModel
		err := found.Decode(&contentModel)
		if err != nil {
			return nil, err
		}
		contentModels[i] = contentModel
	}

	resultModels := make([]model.ContentModel, len(contentModels))

	for i, contentModel := range contentModels {
		resultFields := make([]model.Field, len(contentModel.Fields))
		for i, field := range contentModel.Fields {
			resultFields[i] = model.Field{
				Name:     field.Name,
				Type:     field.Type,
				Required: field.Required,
			}
		}
		resultModels[i] = model.ContentModel{
			ID:     contentModel.ID.Hex(),
			Name:   contentModel.Name,
			Fields: resultFields,
		}
	}

	return resultModels, nil
}

func (c ContentDriver) DeleteContentModelByID(id string) error {

	client, err := c.Client.Get()
	if err != nil {
		return err
	}

	collections := client.Database("models").Collection("content_model")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	deletedResult, err := collections.DeleteOne(context.Background(), bson.M{"_id": objectId})

	if err != nil {
		return err
	}

	if deletedResult.DeletedCount == 0 {
		return driver.NewContentModelCannotFindById(fmt.Sprintf("Content model is not found: %s", id))
	}

	return nil
}
