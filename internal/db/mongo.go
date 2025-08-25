package db

import (
	"context"
	"log"
	"time"

	"github.com/saurabhraut1212/kitchen_kirana/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(cfg config.Config) (*mongo.Client, *mongo.Database) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatalf("mongo connect: %v", err)
	}
	db := client.Database(cfg.MongoDB)
	// ensure indexes
	ensureIndexes(ctx, db)
	return client, db
}

func ensureIndexes(ctx context.Context, db *mongo.Database) {
	_, _ = db.Collection("items").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    map[string]interface{}{"name": 1},
		Options: options.Index().SetUnique(true),
	})
	_, _ = db.Collection("purchases").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: map[string]interface{}{"item_id": 1, "date": -1},
	})
	_, _ = db.Collection("attendances").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    map[string]interface{}{"user_id": 1, "meal_session_id": 1},
		Options: options.Index().SetUnique(true),
	})
}
