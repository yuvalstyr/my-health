package handlers

import "gorm.io/gorm"

type Factory struct {
	Counter counterHandlerInterface
}

func NewHandlersFactory(db *gorm.DB) Factory {
	counter := newCounter(db)
	return Factory{
		Counter: counter,
	}
}
