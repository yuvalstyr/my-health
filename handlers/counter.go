package handlers

import (
	"net/http"
	"personal/health-app/views/counter"
)

func HandleCountersIndex(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, counter.CounterIndex())
}
