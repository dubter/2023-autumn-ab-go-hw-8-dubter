package app

import (
	"homework/internal/devices"
)

//go:generate mockgen -package internal -destination ../mocks/repository.go . Repository
type Repository interface {
	Get(string) (*devices.Device, error)
	Create(*devices.Device) error
	Delete(string) error
	Update(*devices.Device) error
}

//go:generate mockgen -package internal -destination ../mocks/service.go . Service
type Service interface {
	GetDevice(string) (*devices.Device, error)
	CreateDevice(*devices.Device) error
	DeleteDevice(string) error
	UpdateDevice(*devices.Device) error
}

type deviceService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &deviceService{
		repo: repo,
	}
}

func (ds *deviceService) GetDevice(serialNum string) (*devices.Device, error) {
	return ds.repo.Get(serialNum)
}

func (ds *deviceService) CreateDevice(device *devices.Device) error {
	return ds.repo.Create(device)
}

func (ds *deviceService) DeleteDevice(serialNum string) error {
	return ds.repo.Delete(serialNum)
}

func (ds *deviceService) UpdateDevice(device *devices.Device) error {
	return ds.repo.Update(device)
}
