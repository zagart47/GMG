package db

import (
	"GMGgRPCServer/pkg/api"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"sort"
)

var dbhost = os.Getenv("DBHOST")
var Pool, err = pgxpool.New(context.Background(), dbhost)

func ConnectDb() *pgxpool.Pool {
	Pool, err = pgxpool.New(context.Background(), dbhost)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer Pool.Close()
	return nil
}

func UserAdd(name string, score float32, email string, dateTimeToUpdate string) error {
	CheckConnect()
	var emailDb interface{}
	err := Pool.QueryRow(context.Background(), "SELECT email FROM tops WHERE email = $1;", email).Scan(&emailDb)
	if err != nil {
		return err
	}
	if fmt.Sprint(emailDb) == email {
		_, err := Pool.Query(context.Background(), "UPDATE tops SET user_time=$1, user_name=$2, created_at=$3 WHERE email=$4", score, name, dateTimeToUpdate, email)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			return err
			os.Exit(1)
		}
	} else {
		_, err := Pool.Query(context.Background(), "INSERT INTO tops (user_name, user_time, email, created_at) VALUES ($1, $2, $3, $4)", name, score, email, dateTimeToUpdate)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			return err
			os.Exit(1)

		}
	}

	return nil
}

type User []*api.User

func GetScore() []*api.User {
	CheckConnect()
	u := User{}
	rows, err := Pool.Query(context.Background(), "SELECT * FROM tops")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		i := 1
		values, err := rows.Values()
		if err != nil {
			log.Fatal(err)
		}
		u = append(u, &api.User{
			Id:    values[0].(int64),
			Name:  values[1].(string),
			Score: values[2].(float64),
			Email: values[3].(string),
		})
		i++
	}
	sort.Slice(u, func(i, j int) bool {
		return u[i].Score < u[j].Score
	})
	for i, v := range u {
		v.Id = int64(i + 1)
	}
	return u
}

func CheckConnect() {
	if Pool.Ping(context.Background()) != nil {
		ConnectDb()
	}
	CheckTable()
}

func CheckTable() {
	var tableStatus bool
	err := Pool.QueryRow(context.Background(), "SELECT EXISTS(SELECT FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'tops');").Scan(&tableStatus)
	if err != nil {
		log.Fatal(err)
	}
	if tableStatus != true {
		queryAdd := fmt.Sprint("create table tops(id bigserial primary key, user_name text, user_time double precision, email text, created_at text); ")
		queryOwner := fmt.Sprint("alter table tops owner to postgres;")
		err := Pool.QueryRow(context.Background(), queryAdd)
		if err != nil {
			log.Fatal(err)
		}
		err = Pool.QueryRow(context.Background(), queryOwner)
		if err != nil {
			log.Fatal(err)
		}
	}
}
