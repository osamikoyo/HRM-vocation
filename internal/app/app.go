package app

import (
	"fmt"
	"net"

	"github.com/osamikoyo/hrm-vocation/internal/chacker"
	"github.com/osamikoyo/hrm-vocation/internal/server"
	"github.com/osamikoyo/hrm-vocation/pkg/config"
	"github.com/osamikoyo/hrm-vocation/pkg/loger"
	"github.com/osamikoyo/hrm-vocation/pkg/proto/pb"
	"google.golang.org/grpc"
)

type App struct{
	Config *config.Config
	Logger loger.Logger
	Checker *chacker.Checker
	grpc *grpc.Server
}

func Init(cfg *config.Config) (*App, error) {
	ch, err := chacker.New(cfg)
	if err != nil{
		return nil, err
	}

	return &App{
		Logger: loger.New(),
		Checker: ch,
		grpc: grpc.NewServer(),
		Config: cfg,
	}, nil
}

func (a *App) Run() error {
	var ch chan error

	go a.Checker.StartCheck(ch, 24)

	pb.RegisterVocationServiceServer(a.grpc, &server.Server{})

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.Config.Host, a.Config.Port))
	if err != nil{
		return err
	}

	return 
}