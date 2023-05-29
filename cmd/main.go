package main

import (
	"fmt"
	"log"
	"sync"

	"net"

	"github.com/i-akbarshoh/auth-service/internal/pkg/config"
	"github.com/i-akbarshoh/auth-service/internal/pkg/postgres"
	"github.com/i-akbarshoh/auth-service/internal/pkg/proto"
	"github.com/i-akbarshoh/auth-service/internal/repository/user"
	user1 "github.com/i-akbarshoh/auth-service/internal/service/user"
	"google.golang.org/grpc"
)

func main() {
	db := postgres.NewDB()
	repo := user.NewRepository(db)
	service := user1.NewService(repo)
	server := grpc.NewServer()
	proto.RegisterAuthServer(server, service)
	listen, err := net.Listen("tcp", config.C.Server.Host + ":" + config.C.Server.Port)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		if err := server.Serve(listen); err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Println("Program started...")
	wg.Wait()
}