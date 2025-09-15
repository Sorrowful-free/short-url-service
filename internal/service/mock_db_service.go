package service

import "errors"

type MockDBService struct {
	FireError bool
}

func NewMockDBService(fireError bool) DBService {
	return &MockDBService{
		FireError: fireError,
	}
}

func (m *MockDBService) Ping() error {
	if m.FireError {
		return errors.New("test error")
	}
	return nil
}
