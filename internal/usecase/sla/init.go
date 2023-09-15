package sla

import (
	"sla/domain/usecase"
)

type SlaUC struct {
}

func NewSlaUC() usecase.SLAUsecase {
	return &SlaUC{}
}
