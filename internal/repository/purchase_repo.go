package repository

import (
	"context"
	"time"

	"github.com/saurabhraut1212/kitchen_kirana/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PurchaseRepo struct{ col *mongo.Collection }

func NewPurchaseRepo(db *mongo.Database) *PurchaseRepo {
	return &PurchaseRepo{col: db.Collection("purchases")}
}

func (r *PurchaseRepo) Create(ctx context.Context, p *models.Purchase) (*models.Purchase, error) {
	if p.Date.IsZero() {
		p.Date = time.Now()
	}
	p.CreatedAt = time.Now()
	res, err := r.col.InsertOne(ctx, p)
	if err != nil {
		return nil, err
	}
	p.ID = res.InsertedID.(primitive.ObjectID)
	return p, nil
}

func (r *PurchaseRepo) ListByItemSince(ctx context.Context, itemID primitive.ObjectID, since time.Time) ([]models.Purchase, error) {
	cur, err := r.col.Find(ctx, bson.M{"item_id": itemID, "date": bson.M{"$gte": since}})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var out []models.Purchase
	for cur.Next(ctx) {
		var p models.Purchase
		if err := cur.Decode(&p); err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, nil
}
