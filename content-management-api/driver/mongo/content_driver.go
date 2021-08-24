package mongo

import (
	"content-management-api/driver"
	"content-management-api/driver/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Space struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}

type ContentModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Fields    []Field            `bson:"fields"`
	Name      string             `bson:"name"`
	CreatedAt primitive.DateTime `bson:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}

type Field struct {
	Name     string `bson:"name"`
	Type     string `bson:"field_type"`
	Required bool   `bson:"required"`
}

type Entry struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	ContentModelID string             `bson:"content_model_id"`
	Items          []interface{}      `bson:"items"`
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

func (c ContentDriver) FindSpaceByID(id string) (*model.Space, error) {
	client, err := c.Client.Get()
	if err != nil {
		return nil, err
	}

	collections := client.Database("space").Collection("space")

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	findOne := collections.FindOne(context.Background(), bson.M{"_id": objectId})

	space := Space{}
	err = findOne.Decode(&space)

	if err != nil {
		return nil, err
	}

	return &model.Space{
		ID:   space.ID.Hex(),
		Name: space.Name,
	}, nil

}

func (c ContentDriver) CreateSpace(space model.Space) (*model.Space, error) {
	client, err := c.Client.Get()
	if err != nil {
		return nil, err
	}

	collections := client.Database("space").Collection("space")

	insert := Space{
		Name: space.Name,
	}

	result, err := collections.InsertOne(context.Background(), insert)

	if err != nil {
		return nil, err
	}

	return &model.Space{
		ID:   result.InsertedID.(primitive.ObjectID).Hex(),
		Name: insert.Name,
	}, nil
}

func (c ContentDriver) CreateModel(name string, createdAt time.Time, fields []model.Field) (*model.ContentModel, error) {
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
		Name:      name,
		Fields:    fieldsModel,
		CreatedAt: primitive.NewDateTimeFromTime(createdAt),
		UpdatedAt: primitive.NewDateTimeFromTime(createdAt),
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
		ID:        result.InsertedID.(primitive.ObjectID).Hex(),
		Name:      insert.Name,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
		Fields:    resultFields,
	}, err
}

func (c ContentDriver) UpdateModel(updatedModel model.ContentModel) (*model.ContentModel, error) {
	client, err := c.Client.Get()
	if err != nil {
		return nil, err
	}

	collections := client.Database("models").Collection("content_model")

	fieldsModel := make([]Field, len(updatedModel.Fields))
	for i, field := range updatedModel.Fields {
		fieldsModel[i] = Field{
			Name:     field.Name,
			Type:     field.Type,
			Required: field.Required,
		}
	}

	objectId, err := primitive.ObjectIDFromHex(updatedModel.ID)
	if err != nil {
		return nil, err
	}

	update := ContentModel{
		Name:      updatedModel.Name,
		Fields:    fieldsModel,
		CreatedAt: primitive.NewDateTimeFromTime(updatedModel.CreatedAt),
		UpdatedAt: primitive.NewDateTimeFromTime(updatedModel.UpdatedAt),
	}

	_, errUpdate := collections.UpdateOne(
		context.Background(),
		bson.D{{"_id", objectId}},
		bson.D{{"$set", update}},
	)

	if errUpdate != nil {
		return nil, err
	}

	resultFields := make([]model.Field, len(update.Fields))
	for i, field := range update.Fields {
		resultFields[i] = model.Field{
			Name:     field.Name,
			Type:     field.Type,
			Required: field.Required,
		}
	}

	return &model.ContentModel{
		ID:        updatedModel.ID,
		Name:      update.Name,
		CreatedAt: updatedModel.CreatedAt,
		UpdatedAt: updatedModel.UpdatedAt,
		Fields:    resultFields,
	}, err
}

func (c ContentDriver) CreateEntry(entry model.Entry) (*model.Entry, error) {
	client, err := c.Client.Get()
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

func (c ContentDriver) CreateEntryItems(entryId model.EntryID, items []model.EntryItem) ([]model.EntryItem, error) {
	client, err := c.Client.Get()

	if err != nil {
		return nil, err
	}

	collections := client.Database("models").Collection("entry")

	entryItems := make([]interface{}, len(items))
	for i, v := range items {
		entryItems[i] = v
	}

	ctx := context.Background()

	objectId, err := primitive.ObjectIDFromHex(entryId.String())

	if err != nil {
		return nil, err
	}

	_, err = collections.UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{
		"$set": bson.M{
			"items": entryItems,
		},
	})

	if err != nil {
		return nil, err
	}

	return items, nil
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
		ID:        contentModel.ID.Hex(),
		Name:      contentModel.Name,
		CreatedAt: contentModel.CreatedAt.Time(),
		UpdatedAt: contentModel.UpdatedAt.Time(),
		Fields:    resultFields,
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
			ID:        contentModel.ID.Hex(),
			Name:      contentModel.Name,
			CreatedAt: contentModel.CreatedAt.Time(),
			UpdatedAt: contentModel.UpdatedAt.Time(),
			Fields:    resultFields,
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

func (c ContentDriver) FindEntryByID(id string) (*model.Entry, error) {

	client, err := c.Client.Get()

	if err != nil {
		return nil, err
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := client.Database("models").Collection("entry")

	findOne := collection.FindOne(context.Background(), bson.M{"_id": objectId})

	entry := Entry{}
	err = findOne.Decode(&entry)

	if err != nil {
		return nil, err
	}

	items := make([]model.EntryItem, len(entry.Items))
	for i, item := range entry.Items {
		d := item.(primitive.D)
		mamp := map[string]interface{}{}
		for _, e := range d {
			mamp[e.Key] = e.Value
		}
		items[i] = model.EntryItem{
			Value: mamp["value"].(interface{}),
		}
	}

	return &model.Entry{
		ID:      entry.ID.Hex(),
		ModelID: entry.ContentModelID,
		Items:   items,
	}, nil
}
