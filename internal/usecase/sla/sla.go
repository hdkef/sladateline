package sla

import (
	"context"
	"sla/domain/entity"
	"time"
)

// Calculate implements usecase.SLAUsecase.
func (s *SlaUC) Calculate(ctx context.Context, dto *entity.SLADto) (*entity.SLA, error) {
	hourOfWork := 0
	datePointer := dto.Date

	dateline := []entity.SLADateline{}
	dateLine50 := dto.Hour / 2
	dateLine75 := dto.Hour * 3 / 4
	dateLine100 := dto.Hour

	for hourOfWork < dto.Hour {
		// check if date is weekday
		if datePointer.Weekday() != time.Sunday && datePointer.Weekday() != time.Saturday {

			if datePointer.Hour() < 9 {
				// set hour to 9
				datePointer = time.Date(datePointer.Year(), datePointer.Month(), datePointer.Day(), 9, 0, 0, 0, datePointer.Location())
			}

			for i := 0; i < 8; i++ {

				if datePointer.Hour() >= 18 {
					// if hour > 18 break from loop
					break
				}

				if datePointer.Hour() == 12 {
					// if hour == 12, skip 1 hour
					datePointer = time.Date(datePointer.Year(), datePointer.Month(), datePointer.Day(), datePointer.Hour()+1, 0, 0, 0, datePointer.Location())
				}

				// if weekday increment by 1 hour
				datePointer = datePointer.Add(1 * time.Hour)
				hourOfWork += 1
				// check dateline
				switch hourOfWork {
				case dateLine50:
					dateline = append(dateline, entity.SLADateline{
						Percentage: 50,
						Date:       datePointer,
					})
				case dateLine75:
					dateline = append(dateline, entity.SLADateline{
						Percentage: 75,
						Date:       datePointer,
					})
				case dateLine100:
					dateline = append(dateline, entity.SLADateline{
						Percentage: 100,
						Date:       datePointer,
					})
				}
			}
			// set to next day
			datePointer = time.Date(datePointer.Year(), datePointer.Month(), datePointer.Day()+1, 9, 0, 0, 0, datePointer.Location())
		} else {
			// skip to next week day
			if datePointer.Weekday() == time.Sunday {
				// skip 1 day
				datePointer = time.Date(datePointer.Year(), datePointer.Month(), datePointer.Day()+1, 9, 0, 0, 0, datePointer.Location())
			} else if datePointer.Weekday() == time.Saturday {
				// skip 2 day
				datePointer = time.Date(datePointer.Year(), datePointer.Month(), datePointer.Day()+2, 9, 0, 0, 0, datePointer.Location())
			}
		}
	}

	return &entity.SLA{
		Ref:      dto.Ref,
		Dateline: dateline,
	}, nil

}
