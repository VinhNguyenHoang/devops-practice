package collectors

import "cs/internal/tom/models"

type Collector interface {
	RunCollect() error
	GetCollection() []*models.Phone
}
