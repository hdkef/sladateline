package sla

import (
	"sla/domain/delivery"
	"sla/domain/usecase"
)

type SLADelivery struct {
	slaUC usecase.SLAUsecase
}

func NewSLADelivery(slaUC usecase.SLAUsecase) delivery.SLADelivery {
	return &SLADelivery{
		slaUC: slaUC,
	}
}
