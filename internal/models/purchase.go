package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Purchase struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ItemID    primitive.ObjectID `bson:"item_id" json:"item_id" validate:"required"`
	Quantity  float64            `bson:"quantity" json:"quantity" validate:"required,gt=0"`
	Date      time.Time          `bson:"date" json:"date"`
	Note      string             `bson:"note,omitempty" json:"note,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
