package handlers

import (
	"log"
	"net/http"
	"personal/health-app/daos"
	"personal/health-app/model"
	"personal/health-app/views/counter"

	"github.com/go-chi/chi/v5"
)

type CounterHandler struct {
	daos *daos.Factory
}

func NewCounterHandler(daos daos.Factory) *CounterHandler {
	return &CounterHandler{daos: &daos}
}

func (hc *CounterHandler) HandleCountersIndex(w http.ResponseWriter, r *http.Request) error {
	counters, err := hc.daos.Counter.GetCountersPerWeek(2)
	if err != nil {
		return err
	}
	log.Println("counters", counters)
	return render(w, r, counter.CounterIndex(counters))
}

func (hc *CounterHandler) HandleCounterIncrementUpdate(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	modelCounter, err := hc.daos.Counter.Get(id)
	if err != nil {
		return err
	}
	modelCounter.Value++
	err = hc.daos.Counter.UpdateCounter(modelCounter)
	if err != nil {
		return err
	}
	return render(w, r, counter.VisualIndicator(id, 2, modelCounter.Value))
}

func (hc *CounterHandler) HandleCounterDecrementUpdate(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	modelCounter, err := hc.daos.Counter.Get(id)
	if err != nil {
		return err
	}
	modelCounter.Value--
	err = hc.handleCounterUpdate(modelCounter)
	if err != nil {
		return err
	}
	return render(w, r, counter.VisualIndicator(id, 2, modelCounter.Value))
}

func (hc *CounterHandler) handleCounterUpdate(counter *model.Counter) error {
	return hc.daos.Counter.UpdateCounter(counter)
}
