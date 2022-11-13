package db

import (
	"context"
	"fmt"
	"fyne.io/fyne/v2/widget"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sort"
)

var _ = godotenv.Load()
var urlExample = os.Getenv("DBHOST")
var dbpool, err = pgxpool.New(context.Background(), urlExample)
var user string
var score float32
var TopText = widget.NewLabel("")
var TText string

func ConnectDb() {
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	GetAllFromDb()

}

type User struct {
	name  string
	score float32
	email string
}

type UserScore []User

func GetAllFromDb() *widget.Label {
	u := UserScore{}
	for i := 1; ; i++ {
		err = dbpool.QueryRow(context.Background(), "select user_name, user_time from tops where id=$1", i).Scan(&user, &score)
		if err != nil {
			log.Println("Error getting data or data in the database has run out.")
			break
		}
		u = append(u, User{
			name:  user,
			score: score,
		})
	}

	sort.Slice(u, func(i, j int) bool {
		return u[i].score < u[j].score
	})

	for i, us := range u {
		fmt.Printf("%d. %s, время: %.3f\n", i+1, us.name, us.score)
		TText += fmt.Sprintf("%d. %s, время: %.3f\n", i+1, us.name, us.score)
	}
	TopText.SetText(TText)
	return TopText
}

func GetUserFromDb(name string) *widget.Label {
	u := UserScore{}
	for i := 1; ; i++ {
		err = dbpool.QueryRow(context.Background(), "select user_name, user_time from tops where id=$1", i).Scan(&user, &score)
		if err != nil {
			log.Println("Error getting data or data in the database has run out.")
			break
		}
		u = append(u, User{
			name:  user,
			score: score,
		})
	}

	sort.Slice(u, func(i, j int) bool {
		return u[i].score < u[j].score
	})

	for i, us := range u {
		fmt.Printf("%d. %s, время: %.3f\n", i+1, us.name, us.score)
		TText += fmt.Sprintf("%d. %s, время: %.3f\n", i+1, us.name, us.score)
	}
	TopText.SetText(TText)
	return TopText
}

func EmailCheck(email string) bool {
	u := UserScore{}
	for i := 1; ; i++ {
		err = dbpool.QueryRow(context.Background(), "select email from tops where id=$1", email).Scan(&user, &score, &email)
		if err != nil {
			log.Println("Error getting data or data in the database has run out.")
			break
		}
		u = append(u, User{
			name:  user,
			score: score,
			email: email,
		})
	}

	sort.Slice(u, func(i, j int) bool {
		return u[i].score < u[j].score
	})
	for _, us := range u {
		if us.email == email {
			return true
		}
	}
	return false
}
