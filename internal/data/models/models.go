package models

import "time"

type Vocation struct{
	UserID uint64
	DateStart time.Time
	DateEnd time.Time
}

func ToModels()