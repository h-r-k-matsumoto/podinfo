package api

import (
	"math"
	"net/http"
)

// OutOfMemory godoc
// @Summary OutOfMemory
// @Description causes OOM
// @Tags HTTP API
// @Router /oom [get]
func (s *Server) oomHandler(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("OOM command received")

	val := "0123456789"
	for {
		val += val
		if len(val) > math.MaxInt64 {
			s.logger.Warn("This is dummy code for compiling that is not erase this block.")
		}
	}
}
