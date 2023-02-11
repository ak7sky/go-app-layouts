package api

import "go-app-layouts/v1/internal/core/model"

type AnythingService interface {
	CreateAnything(model *model.Anything) *model.Anything
	GetAnything(id string) *model.Anything
}

type AnythingStorage interface {
	Create(model *model.Anything) *model.Anything
	Get(id string) *model.Anything
}
