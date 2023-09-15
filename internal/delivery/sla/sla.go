package sla

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sla/domain/entity"
	"sla/internal/delivery/mapper"
)

// CalculateDateline implements delivery.SLADelivery.
func (s *SLADelivery) CalculateDateline(ctx context.Context) (path string, router func(w http.ResponseWriter, r *http.Request)) {

	route := "/sla/dateline"
	method := http.MethodPost

	fmt.Printf("%s %s\n", method, route)

	return route, func(w http.ResponseWriter, r *http.Request) {

		// validate method
		if r.Method != method {
			http.Error(w, "invalid method", http.StatusBadRequest)
			return
		}

		// decode payload
		req := entity.SLARequest{}

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		dto, err := entity.NewSlaDto(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		et, err := s.slaUC.Calculate(ctx, dto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := mapper.ToSLAResponse(et)

		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
