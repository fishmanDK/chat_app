package service

import "realtime_chat_app/internal/storage"

type Service struct {
	db *storage.Storage
}

func NewService(db *storage.Storage) *Service {
	return &Service{
		db: db,
	}
}
