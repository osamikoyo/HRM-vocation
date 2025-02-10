package data

import "gorm.io/gorm"

type Data struct{
	db *gorm.DB
}

func New() (*Data, error) {
	
}

