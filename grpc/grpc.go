package grpc

import (
	"GMG/pkg/api"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"os"
)

func ConnectGRPC(name string, score float32) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("HOST")
	Conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer Conn.Close()

	c := api.NewScoreAdderClient(Conn)
	res, err := c.AddScore(context.Background(), &api.AddRequest{
		User: name,
		Time: score,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.GetAdded())

}
