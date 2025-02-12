package main

import (
	"fmt"

	"github.com/osamikoyo/hrm-vocation/internal/app"
	"github.com/osamikoyo/hrm-vocation/pkg/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil{
		fmt.Println(err)
	}

	app, err := app.Init(&cfg)
	if err != nil{
		fmt.Println(err)
	}

	if err = app.Run(); err != nil{
		app.Logger.Error().Err(err)
	}
}