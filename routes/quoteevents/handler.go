package quoteevents

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/a-h/templ"
)

func NewHandler() http.Handler {
	return &Handler{}
}

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.WriteHeader(http.StatusOK)

	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		m1 := &Model{
			Name: fmt.Sprintf("apple-%d", i),
			Cost: 1011 + int64(i),
		}
		m1s := renderToString(View(m1))
		// Write event to HTTP response and flush.
		fmt.Fprintf(w, "event: %s\ndata: %s\n\n", "quote", m1s)
		w.(http.Flusher).Flush()
	}
}

func renderToString(c templ.Component) string {
	var b strings.Builder
	c.Render(context.Background(), &b)
	return b.String()
}

func displayCost(cost int64) string {
	return fmt.Sprintf("£%.2f", (float64(cost) / 100.0))
}
