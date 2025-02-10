package models

import (
	"time"

	"github.com/osamikoyo/hrm-vocation/pkg/proto/pb"
)

const TIME_TAMPLATE = "2006.01.02"

type Vocation struct{
	UserID uint64
	DateStart time.Time
	DateEnd time.Time
}

func ToModels(voc *pb.Vocation) (*Vocation, error) {
	startTime, err := time.Parse(TIME_TAMPLATE, voc.StartTime)
	if err != nil{
		return nil, err
	}
	endTime, err := time.Parse(TIME_TAMPLATE, voc.EndTime)
	if err != nil{
		return nil, err
	}

	return &Vocation{

	}, nil
}