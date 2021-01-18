package api

import (
	"fmt"
	"net/http"

	"time"

	"github.com/gorilla/mux"
)

// Until godoc
// @Summary Delay
// @Description waits for the specified period
// @Tags HTTP API
// @Accept json
// @Produce json
// @Router /delay/{seconds} [get]
// @Success 200 {object} api.MapResponse
func (s *Server) untilHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	until, err := time.Parse("150405",vars["until"])
	if err != nil {
		s.ErrorResponse(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	now := time.Now();

	s.logger.Info(fmt.Sprintf("now=%s, until %s",now,until));
	if now.After(until) {
		s.JSONResponse(w, r, map[string]string{"status": "ok"})
		return
	}
	time.Sleep(10 * time.Second)
	s.JSONResponseCode(w, r, map[string]string{"status": "ng"}, 502)
}
