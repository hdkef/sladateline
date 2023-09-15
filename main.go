package main

import (
	"context"
	"fmt"
	"net/http"
	slaDelivery "sla/internal/delivery/sla"
	slaUC "sla/internal/usecase/sla"
)

func main() {

	// usecase
	slaUC := slaUC.NewSlaUC()

	// delivery
	slaDelivery := slaDelivery.NewSLADelivery(slaUC)

	http.HandleFunc(slaDelivery.CalculateDateline(context.Background()))

	// Start the HTTP server on port 8080
	fmt.Println("about to serve on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
