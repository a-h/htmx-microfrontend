package result

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func NewHandler() http.Handler {
	return &Handler{}
}

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.Get(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

var nameToCost = map[string]int64{
	"apple":  1011,
	"banana": 2012,
	"cherry": 3013,
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	cost := nameToCost[name]
	m := &Model{
		Name: name,
		Cost: cost,
	}
	templ.Handler(View(m)).ServeHTTP(w, r)
}

func displayCost(cost int64) string {
	return fmt.Sprintf("£%.2f", (float64(cost) / 100.0))
}
