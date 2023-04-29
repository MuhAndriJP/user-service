package main

import (
	"github.com/MuhAndriJP/user-service.git/grpc"
	"github.com/MuhAndriJP/user-service.git/repo/mysql"
)

func main() {
	mysql.InitDB()
	grpc.Run()
}
