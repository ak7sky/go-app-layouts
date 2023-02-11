package service

import (
	"go-app-layouts/v1/internal/core/api"
	"go-app-layouts/v1/internal/core/model"
)

type AnythingService struct {
	storage api.AnythingStorage
}

func New(storage api.AnythingStorage) *AnythingService {
	return &AnythingService{storage: storage}
}

func (modelSrv *AnythingService) CreateAnything(model *model.Anything) *model.Anything {
	createdModel := modelSrv.storage.Create(model)
	// + extra business logic if required
	return createdModel
}

func (modelSrv *AnythingService) GetAnything(id string) *model.Anything {
	// + extra business logic if required
	return modelSrv.storage.Get(id)
}
