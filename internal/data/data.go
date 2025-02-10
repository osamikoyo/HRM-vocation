package data

import (
	"github.com/osamikoyo/hrm-vocation/internal/data/models"
	"github.com/osamikoyo/hrm-vocation/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Data struct{
	db *gorm.DB
}

func New(cfg *config.Config) (*Data, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN))
	return &Data{db : db}, err
}

func (d *Data) Add(voc *models.Vocation) error {
	return d.db.Create(voc).Error
}

func (d *Data) Get(UserID uint64) (*models.Vocation, error) {
	var voc models.Vocation

	result := d.db.Where(&models.Vocation{
		UserID: UserID,
	}).Find(&voc)
	return &voc, result.Error
}

func (d *Data) Delete(UserId uint64) error {
	return d.db.Model(&models.Vocation{}).Where(
		&models.Vocation{
			UserID: UserId,
		},
	).Delete(&models.Vocation{}).Error
}