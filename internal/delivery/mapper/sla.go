package mapper

import "sla/domain/entity"

func ToSLAResponse(et *entity.SLA) *entity.SLAResponse {

	dateline := []entity.SLADatelineResponse{}

	for _, v := range et.Dateline {
		dateline = append(dateline, entity.SLADatelineResponse{
			Percentage: v.Percentage,
			Date:       v.Date.Format("2006-01-02 15:04:05"),
		})
	}

	return &entity.SLAResponse{
		Ref:      et.Ref,
		Dateline: dateline,
	}
}
