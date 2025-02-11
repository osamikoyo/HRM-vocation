package chacker

import (
	"fmt"
	"time"

	"github.com/osamikoyo/hrm-vocation/internal/data"
	"github.com/osamikoyo/hrm-vocation/internal/data/models"
	"github.com/osamikoyo/hrm-vocation/internal/sender"
	"github.com/osamikoyo/hrm-vocation/pkg/config"
	"github.com/osamikoyo/hrm-vocation/pkg/loger"
)

type Checker struct{
	Loger loger.Logger
	Sender *sender.Sender
	Data *data.Data
}

func New(cfg *config.Config) (*Checker, error) {
	sender, err := sender.Init(cfg)
	if err != nil{
		return nil, err
	}

	return &Checker{
		Sender: sender,
		Loger: loger.New(),
	}, nil
}

func (c *Checker) StartCheck(ch chan error, hoursduration uint8) {
	now := time.Now().Format(models.TIME_TAMPLATE)

	for {
		vocs, err := c.Data.GetAll()
		if err != nil{
			ch <- err
		}

		for _, v := range vocs {
			if v.DateEnd.String() == now{
				c.Data.Delete(v.UserID)
				message := models.Msg{
					From: models.MESSAGE_FROM,
					To: v.UserEmail,
					CC: []string{fmt.Sprintf("%d", v.UserID), v.UserEmail},
					Subject: "vocation",
					Body: models.EMAIL_ABOUT_DELETE_TEMPLATE,
				}
				err = c.Sender.Send(message)
				if err != nil{
					ch <- err
				}
			}
		}

		time.Sleep(time.Hour * time.Duration(hoursduration))
	}
}