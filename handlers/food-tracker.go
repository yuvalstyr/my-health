package handlers

import (
	"net/http"
	"personal/health-app/daos"
	"personal/health-app/views/tracker"
)

type FoodTrackerHandler struct {
	daos *daos.Factory
}

func NewFoodTrackerHandler(daos daos.Factory) *FoodTrackerHandler {
	return &FoodTrackerHandler{daos: &daos}
}

func (hc *CounterHandler) HandleFoodTrackerIndex(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, tracker.FoodTracker())
}
