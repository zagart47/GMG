package grpc

import (
	"GMG/pkg/api"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
)

var creds, err2 = credentials.NewClientTLSFromFile("grpc/server.crt", "")
var err1 = godotenv.Load()
var Host = os.Getenv("HOST")
var Conn, err = grpc.Dial(Host, grpc.WithTransportCredentials(creds))
var Client = api.NewScoreClient(Conn)

func ConnectGRPC() {
	if err != nil {
		log.Fatal(err)
	}
	if err2 != nil {
		log.Fatal(err2)
	}
	defer Conn.Close()

}

// AddUserScore sends a new player and their data to the database
func AddUserScore(name string, score float32, email string) {
	res, err := Client.AddScore(context.Background(), &api.AddRequest{
		User:  name,
		Time:  score,
		Email: email,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.GetAdded())
}

type User struct {
	Id    int64
	Name  string
	Email string
	Score float64
}

type UserScore []User

// GetUserScore gets all players and their data from the database
func GetUserScore() UserScore {
	res, err := Client.GetScore(context.Background(), &api.GetRequest{})
	if err != nil {
		log.Fatal(err)
	}
	u := UserScore{}

	for _, v := range res.Score {
		u = append(u, User{
			Id:    v.GetId(),
			Name:  v.GetName(),
			Email: v.GetEmail(),
			Score: v.GetScore(),
		})
	}

	for _, v := range u {
		fmt.Printf("%d Имя: %s; Email: %s; Время: %.2f\n", v.Id, v.Name, v.Email, v.Score)
	}
	return u
}
