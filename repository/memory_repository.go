package repository

import (
	modelling "axon-go/modeling"
	"context"
	"sync"
)

type MemoryRepository struct {
	db   map[string]modelling.Aggregate
	dbMu sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {
	r := &MemoryRepository{
		db: map[string]modelling.Aggregate{},
	}
	return r
}

func (m MemoryRepository) Load(ctx context.Context, uid modelling.UinId) (modelling.Aggregate, error) {
	m.dbMu.RLock()
	defer m.dbMu.RUnlock()
	model, ok := m.db[uid.Identifier]
	if !ok {
		return nil, ErrEntityNotFound
	}
	return model, nil
}

func (m MemoryRepository) Save(ctx context.Context, bizAggregate interface{}) (modelling.Aggregate, error) {
	m.dbMu.Lock()
	defer m.dbMu.Unlock()

	wrappedAggregate := NewAggregateWrapper(bizAggregate)

	id := wrappedAggregate.GetAggregateId()
	if id == "" {
		return nil, ErrMissingEntityID
	}
	m.db[id] = wrappedAggregate
	return wrappedAggregate, nil
}
