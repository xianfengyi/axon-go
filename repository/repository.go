package repository

import (
	modelling "axon-go/modeling"
	"context"
	"errors"
)

// ErrEntityNotFound is when a entity could not be found.
var ErrEntityNotFound = errors.New("could not find entity")

// ErrCouldNotSaveEntity is when a entity could not be saved.
var ErrCouldNotSaveEntity = errors.New("could not save entity")

// ErrMissingEntityID is when a entity has no ID.
var ErrMissingEntityID = errors.New("missing entity id")

// Repository a combined read and write repo.
type Repository interface {
	// Load an aggregate by aggregate id.
	Load(ctx context.Context, uid modelling.UinId) (modelling.Aggregate, error)

	// Save  a aggregate into storage.
	Save(ctx context.Context, aggregate interface{}) (modelling.Aggregate, error)
}
