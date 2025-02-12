package server

import (
	"context"
	"net/http"

	"github.com/osamikoyo/hrm-vocation/internal/data"
	"github.com/osamikoyo/hrm-vocation/internal/data/models"
	"github.com/osamikoyo/hrm-vocation/pkg/proto/pb"
)

type Server struct{
	pb.UnimplementedVocationServiceServer
	Storage data.Data
}


func (s *Server) Add(_ context.Context, req *pb.AddVocationRequest) (*pb.Response, error) {
	voc, err := models.ToModels(req.Vocation)
	if err != nil{
		return nil, err
	}

	err = s.Storage.Add(voc)
	if err != nil{
		return &pb.Response{
			Error: err.Error(),
			Status: http.StatusInternalServerError,
		}, err
	}

	return &pb.Response{
		Error: "",
		Status: http.StatusOK,
	}, nil
}
func (s *Server) Delete(_ context.Context, req *pb.DeleteVocationRequest) (*pb.Response, error) {
	err := s.Storage.Delete(req.UserID)
	if err != nil{
		return &pb.Response{
			Error: err.Error(),
			Status: http.StatusInternalServerError,
		}, err
	}

	return &pb.Response{
		Error: "",
		Status: http.StatusOK,
	}, nil
}

func (s *Server) Client(_ context.Context, req *pb.GetVocationRequest) (*pb.GetVocationResponse, error) {
	voc, err := s.Storage.Get(req.UserID)
	if err != nil{
		return &pb.GetVocationResponse{
			Vocation: nil,
			Response: &pb.Response{
				Error: err.Error(),
				Status: http.StatusInternalServerError,
			},
		}, err
	}

	return &pb.GetVocationResponse{
		Response: &pb.Response{
			Error: "",
			Status: http.StatusOK,
		},
		Vocation: models.ToPB(voc),
	}, nil
}