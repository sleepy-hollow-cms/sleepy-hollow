package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entry struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	ContentModelID string             `bson:"content_model_id"`
}
