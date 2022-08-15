package repository

import (
	modelling "axon-go/modeling"
	"axon-go/utils"
	"context"
	"reflect"
)

type AggregateWrapper struct {
	aggregateRoot interface{}
}

// NewAggregateWrapper creates an aggregate.
func NewAggregateWrapper(aggregate interface{}) modelling.Aggregate {
	return AggregateWrapper{
		aggregateRoot: aggregate,
	}
}

func (a AggregateWrapper) GetAggregateId() string {
	reflectValue := reflect.Indirect(reflect.ValueOf(a.aggregateRoot))
	aggregateIdIdentifier := reflectValue.FieldByName("AggregateId").FieldByName("Identifier").String()
	return aggregateIdIdentifier
}

func (a AggregateWrapper) GetAggregateType() string {
	return utils.GetObjectName(a.aggregateRoot)
}

func (a AggregateWrapper) ApplyEvent(ctx context.Context, event interface{}) error {
	return nil
}

func (a AggregateWrapper) GetAggregateRoot() interface{} {
	return a.aggregateRoot
}
