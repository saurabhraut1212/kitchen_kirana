package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name" validate:"required",min=1`
	Unit        string             `bson:"unit" json:"unit" validate:"required"`
	Quantity    float64            `bson:"quantity" json:"quantity",validate:"gte=0"`
	Threshold   float64            `bson:"threshold" json:"threshold",validate:"gte=0"`
	LastUpdated time.Time          `bson:"last_updated" json:"last_updated"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}
