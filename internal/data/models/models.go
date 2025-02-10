package models

import (
	"time"

	"github.com/osamikoyo/hrm-vocation/pkg/proto/pb"
)

const TIME_TAMPLATE = "2006.01.02"

type Vocation struct{
	VocId uint64 `gorm:"primaryKey;autoIncrement"`
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
		DateStart: startTime,
		DateEnd: endTime,
		UserID: voc.UserID,
		VocId: voc.VocID,
	}, nil
}

func ToPB(voc *Vocation) *pb.Vocation {
	return &pb.Vocation{
		StartTime: voc.DateStart.Format(TIME_TAMPLATE),
		EndTime: voc.DateStart.Format(TIME_TAMPLATE),
		UserID: voc.UserID,
		VocID: voc.VocId,
	}
}