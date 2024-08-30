package db

import (
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
)

func ProvideCacheDB() DB {
	// TODO Retrieve default value from config
	cache := cache.New(5*time.Minute, 10*time.Minute)
	return &db{db: cache}
}

type db struct {
	db *cache.Cache
}

// SavePoints saves the points associated with the given ID into the cache
func (d *db) SavePoints(id string, points int) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}
	d.db.Set(id, points, cache.DefaultExpiration)
	return nil
}

// GetPointByReceiptID retrieves the points associated with the given ID from the cache
func (d *db) GetPointByReceiptID(id string) (int, error) {
	if id == "" {
		return 0, errors.New("id cannot be empty")
	}
	value, found := d.db.Get(id)
	if !found {
		return 0, errors.New("points not found for the given id")
	}
	points, ok := value.(int)
	if !ok {
		return 0, errors.New("invalid data type for points")
	}
	return points, nil
}