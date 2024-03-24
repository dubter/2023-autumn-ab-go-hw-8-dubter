package hashmap

import (
	"sync"

	"homework/internal/app"
	"homework/internal/devices"
	"homework/internal/errors"
)

type hash struct {
	hashTable map[string]*devices.Device
	mu        sync.RWMutex
}

func NewHash() app.Repository {
	return &hash{
		hashTable: make(map[string]*devices.Device),
		mu:        sync.RWMutex{},
	}
}

func (h *hash) Get(serialNum string) (*devices.Device, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if _, ok := h.hashTable[serialNum]; !ok {
		return nil, errors.NewNotFoundError(serialNum)
	}

	return h.hashTable[serialNum], nil
}

func (h *hash) Create(device *devices.Device) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.hashTable[device.SerialNum]; ok {
		return errors.NewAlreadyExistDeviceError(device.SerialNum)
	}

	h.hashTable[device.SerialNum] = device

	return nil
}

func (h *hash) Delete(serialNum string) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.hashTable[serialNum]; !ok {
		return errors.NewNotFoundError(serialNum)
	}

	delete(h.hashTable, serialNum)

	return nil
}

func (h *hash) Update(device *devices.Device) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.hashTable[device.SerialNum]; !ok {
		return errors.NewNotFoundError(device.SerialNum)
	}

	h.hashTable[device.SerialNum] = device

	return nil
}
