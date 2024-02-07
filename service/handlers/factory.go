package handlers

import dao "personal/health-app/service/daos"

type Factory struct {
	Counter counterHandlerInterface
}

func NewHandlersFactory(daoFactory dao.Factory) Factory {
	counter := newCounter(daoFactory)
	return Factory{
		Counter: counter,
	}
}
