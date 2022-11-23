package main

import (
	"GMG/cmd"
	"GMG/grpc"
)

func main() {
	grpc.GetUserScore()
	cmd.StartApp()
}
