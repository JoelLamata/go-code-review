package memdb

import (
	"coupon_service/internal/service/entity"
	"fmt"
)

type Config struct{}

type Repository struct {
	entries map[string]entity.Coupon
}

func New() *Repository {
	return &Repository{
		entries: make(map[string]entity.Coupon),
	}
}

func (r *Repository) FindByCode(code string) (*entity.Coupon, error) {
	coupon, found := r.entries[code]
	if !found {
		return nil, fmt.Errorf("coupon not found")
	}
	return &coupon, nil
}

func (r *Repository) Save(coupon entity.Coupon) error {
	_, found := r.entries[coupon.Code]
	if found {
		return fmt.Errorf("coupon with code %s already exists", coupon.Code)
	}
	r.entries[coupon.Code] = coupon
	return nil
}
