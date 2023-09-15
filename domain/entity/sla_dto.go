package entity

import (
	"errors"
	"time"
)

type SLARequest struct {
	Ref  string `json:"ref"`
	Date string `json:"date"`
	Hour int    `json:"hour"`
}

type SLADto struct {
	Ref  string
	Date time.Time
	Hour int
}

type SLA struct {
	Ref      string
	Dateline []SLADateline
}

type SLADateline struct {
	Percentage int
	Date       time.Time
}

type SLADatelineResponse struct {
	Percentage int    `json:"percentage"`
	Date       string `json:"date"`
}

type SLAResponse struct {
	Ref      string `json:"ref"`
	Dateline []SLADatelineResponse
}

func NewSlaDto(req *SLARequest) (*SLADto, error) {

	if req.Ref == "" {
		return nil, errors.New("invalid ref")
	}

	date, err := time.Parse("2006-01-02 15", req.Date)
	if err != nil {
		return nil, errors.New("date format is YYYY-MM-DD HH")
	}

	if req.Hour <= 0 {
		return nil, errors.New("invalid hour")
	}

	return &SLADto{
		Ref:  req.Ref,
		Date: date,
		Hour: req.Hour,
	}, nil
}
