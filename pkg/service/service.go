package service

import "github.com/Ubludor/restAPI_on_GO/pkg/repository"

type Authrization interface {
}

type DutyList interface {
}

type Service struct {
	Authrization
	DutyList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}