package repository

import (
	"context"
	"time"

	"github.com/saurabhraut1212/kitchen_kirana/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemRepo struct {
	col *mongo.Collection
}

func NewItemRepo(db *mongo.Database) *ItemRepo {
	return &ItemRepo{
		col: db.Collection("items"),
	}
}

func (r *ItemRepo) create(ctx context.Context, it *models.Item) (*models.Item, error) {
	now := time.Now()
	it.CreatedAt = now
	it.LastUpdated = now
	res, err := r.col.InsertOne(ctx, it)
	if err != nil {
		return nil, err
	}
	it.ID = res.InsertedID.(primitive.ObjectID)
	return it, err

}

func (r *ItemRepo) List(ctx context.Context) ([]models.Item, error) {
	cur, err := r.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var out []models.Item
	for cur.Next(ctx) {
		var it models.Item
		if err := cur.Decode(&it); err != nil {
			return nil, err
		}
		out = append(out, it)
	}
	return out, cur.Err()

}

func (r *ItemRepo) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Item, error) {
	var it models.Item
	if err := r.col.FindOne(ctx, bson.M{"_id": id}).Decode(&it); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &it, nil
}

func (r *ItemRepo) Update(ctx context.Context, id primitive.ObjectID, upd bson.M) (*models.Item, error) {
	upd["last_updated"] = time.Now()
	_, err := r.col.UpdateByID(ctx, id, bson.M{"$set": upd})
	if err != nil {
		return nil, err
	}
	return r.GetByID(ctx, id)
}

func (r *ItemRepo) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.col.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *ItemRepo) LowStock(ctx context.Context) ([]models.Item, error) {
	cur, err := r.col.Find(ctx, bson.M{"$expr": bson.M{"$lt": []interface{}{"$quantity", "$threshold"}}})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var out []models.Item
	for cur.Next(ctx) {
		var it models.Item
		if err := cur.Decode(&it); err != nil {
			return nil, err
		}
		out = append(out, it)
	}
	return out, nil
}
