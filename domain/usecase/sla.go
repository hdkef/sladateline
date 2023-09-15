package usecase

import (
	"context"
	"sla/domain/entity"
)

type SLAUsecase interface {
	Calculate(ctx context.Context, dto *entity.SLADto) (*entity.SLA, error)
}
