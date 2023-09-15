package sla_test

import (
	"context"
	"reflect"
	"sla/domain/entity"
	"sla/internal/usecase/sla"
	"testing"
	"time"
)

func TestSlaUC_Calculate(t *testing.T) {
	type args struct {
		ctx context.Context
		dto *entity.SLADto
	}
	tests := []struct {
		name    string
		s       *sla.SlaUC
		args    args
		want    *entity.SLA
		wantErr bool
	}{
		{
			name: "should be ok",
			s:    &sla.SlaUC{},
			args: args{
				ctx: context.Background(),
				dto: &entity.SLADto{
					Ref:  "A",
					Date: time.Date(2023, 9, 12, 13, 0, 0, 0, time.Local),
					Hour: 8,
				},
			},
			want: &entity.SLA{
				Ref: "A",
				Dateline: []entity.SLADateline{
					{
						Percentage: 50,
						Date:       time.Date(2023, 9, 12, 17, 0, 0, 0, time.Local),
					},
					{
						Percentage: 75,
						Date:       time.Date(2023, 9, 13, 10, 0, 0, 0, time.Local),
					},
					{
						Percentage: 100,
						Date:       time.Date(2023, 9, 13, 12, 0, 0, 0, time.Local),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sla.SlaUC{}
			got, err := s.Calculate(tt.args.ctx, tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlaUC.Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlaUC.Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
