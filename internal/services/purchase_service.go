package services

import (
	"context"
	"time"

	"github.com/saurabhraut1212/kitchen_kirana/internal/models"
	"github.com/saurabhraut1212/kitchen_kirana/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PurchaseService struct {
	repo     *repository.PurchaseRepo
	itemRepo *repository.ItemRepo
}

func NewPurchaseService(db *mongo.Database) *PurchaseService {
	return &PurchaseService{
		repo:     repository.NewPurchaseRepo(db),
		itemRepo: repository.NewItemRepo(db),
	}
}

func (s *PurchaseService) Record(ctx context.Context, p *models.Purchase) (*models.Purchase, error) {
	if p.Date.IsZero() {
		p.Date = time.Now()
	}
	created, err := s.repo.Create(ctx, p)
	if err != nil {
		return nil, err
	}

	// atomically update item quantity
	it, err := s.itemRepo.GetByID(ctx, p.ItemID)
	if err != nil {
		return nil, err
	}
	newQty := it.Quantity + p.Quantity
	_, err = s.itemRepo.Update(ctx, p.ItemID, map[string]interface{}{"quantity": newQty})
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (s *PurchaseService) QuickBuy(ctx context.Context, itemID primitive.ObjectID, qty float64) (*models.Purchase, error) {
	p := &models.Purchase{ItemID: itemID, Quantity: qty, Date: time.Now()}
	return s.Record(ctx, p)
}
