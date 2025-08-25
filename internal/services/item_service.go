package services

import (
	"context"
	"errors"
	"time"

	"github.com/saurabhraut1212/kitchen_kirana/internal/models"
	"github.com/saurabhraut1212/kitchen_kirana/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemService struct {
	repo  *repository.ItemRepo
	pRepo *repository.PurchaseRepo
}

func NewItemService(db *mongo.Database) *ItemService {
	return &ItemService{
		repo:  repository.NewItemRepo(db),
		pRepo: repository.NewPurchaseRepo(db),
	}
}

func (s *ItemService) Create(ctx context.Context, in *models.Item) (*models.Item, error) {
	return s.repo.Create(ctx, in)
}
func (s *ItemService) List(ctx context.Context) ([]models.Item, error) { return s.repo.List(ctx) }
func (s *ItemService) Get(ctx context.Context, id primitive.ObjectID) (*models.Item, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ItemService) Update(ctx context.Context, id primitive.ObjectID, upd map[string]interface{}) (*models.Item, error) {
	// ensure no negative quantity
	if q, ok := upd["quantity"].(float64); ok && q < 0 {
		return nil, errors.New("quantity cannot be negative")
	}
	return s.repo.Update(ctx, id, upd)
}
func (s *ItemService) Delete(ctx context.Context, id primitive.ObjectID) error {
	return s.repo.Delete(ctx, id)
}
func (s *ItemService) LowStock(ctx context.Context) ([]models.Item, error) {
	return s.repo.LowStock(ctx)
}

// Predict days to finish using purchases in window days
func (s *ItemService) PredictDaysToFinish(ctx context.Context, id primitive.ObjectID, days int) (float64, error) {
	if days <= 0 {
		days = 30
	}
	since := time.Now().AddDate(0, 0, -days)
	purs, err := s.pRepo.ListByItemSince(ctx, id, since)
	if err != nil {
		return 0, err
	}
	totalBought := 0.0
	for _, p := range purs {
		totalBought += p.Quantity
	}
	if totalBought == 0 {
		return -1, nil
	} // unknown
	avgPerDay := totalBought / float64(days)
	it, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return 0, err
	}
	if avgPerDay == 0 {
		return -1, nil
	}
	return it.Quantity / avgPerDay, nil
}
