package handlers

import dao "personal/health-app/service/daos"

type Factory struct {
	Counter     counterHandlerInterface
	DishHandler dishHandlerInterface
}

func NewHandlersFactory(daoFactory dao.Factory) Factory {
	return Factory{
		Counter:     newCounter(daoFactory),
		DishHandler: newDish(daoFactory),
	}
}
