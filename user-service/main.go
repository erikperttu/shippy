package main

import (
	"fmt"
	"log"
	pb "github.com/erikperttu/shippy/user-service/proto/user"
	"github.com/micro/go-micro"
)

func main() {
	// Create the database connection
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	// Auto migrate the user struct into the database columns/types.
	// Check for changes and migrate them each time the service is restarted
	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	// Create a new micro service
	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	// Initialize
	srv.Init()

	// Register the handler
	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService})


	// Run the server
	if err := srv.Run(); err != nil{
		fmt.Println(err)
	}

}