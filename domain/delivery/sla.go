package delivery

import (
	"context"
	"net/http"
)

type SLADelivery interface {
	CalculateDateline(ctx context.Context) (path string, router func(w http.ResponseWriter, r *http.Request))
}
