package redis

import (
	"context"
	"time"

	data "github.com/ahsansaif47/home-kitchens/common/database"
)

type ICacheRepository interface {
	StoreOTP(email, otp string, expTime time.Duration) error
	RetrieveOTP(email string) (string, error)
}

type CacheRepository struct {
	cache data.Cache
}

func NewUserCache(cache data.Cache) ICacheRepository {
	return &CacheRepository{cache: cache}
}

func (r *CacheRepository) StoreOTP(email, otp string, expTime time.Duration) error {
	return r.cache.Store(context.Background(), email, otp, expTime)
}

func (r *CacheRepository) RetrieveOTP(email string) (string, error) {
	value, err := r.cache.Retrieve(context.Background(), email)
	if err != nil {
		return "", err
	}

	otp, ok := value.(string)
	if !ok {
		return "", nil
	}

	return otp, nil
}
