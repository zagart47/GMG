package adder

import (
	"GMGgRPCServer/db"
	"GMGgRPCServer/pkg/api"
	"fmt"
	"golang.org/x/net/context"
	"log"
	"strings"
	"time"
)

type GRPCServer struct {
	api.UnimplementedScoreServer
}

func (S *GRPCServer) AddScore(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {

	dateTimeToUpdate := time.Now()
	s := fmt.Sprintf("%02d.%02d.%d %02d:%02d:%02d", dateTimeToUpdate.Day(), dateTimeToUpdate.Month(), dateTimeToUpdate.Year(), dateTimeToUpdate.Hour(), dateTimeToUpdate.Minute(), dateTimeToUpdate.Second())
	err := db.UserAdd(req.GetUser(), req.GetTime(), strings.ToLower(req.GetEmail()), s)
	if err != nil {
		log.Fatal(err)
	}
	return &api.AddResponse{Added: req.GetUser()}, nil
}

func (S *GRPCServer) GetScore(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	return &api.GetResponse{Scores: db.GetScore()}, nil
}
