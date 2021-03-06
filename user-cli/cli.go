package main

import (
	"log"
	"os"

	pb "github.com/erikperttu/shippy/user-service/proto/user"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"golang.org/x/net/context"
)

func main() {

	cmd.Init()

	// Create a new greeter client
	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	// Define flags
	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "Your full name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "Your email",
			},
			cli.StringFlag{
				Name:  "Password",
				Usage: "Your password",
			},
			cli.StringFlag{
				Name:  "Company",
				Usage: "Your company",
			},
		),
	)

	// Start the service
	service.Init(
		micro.Action(func(c *cli.Context) {
			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			company := c.String("company")

			// Call the user service
			r, err := client.Create(context.TODO(), &pb.User{
				Name:     name,
				Email:    email,
				Password: password,
				Company:  company,
			})
			if err != nil {
				log.Fatalf("Failed to create: %v", err)
			}
			log.Printf("Created: %s", r.User.Id)

			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				log.Fatalf("Failed to list users: %v", err)
			}

			for _, v := range getAll.Users {
				log.Println(v)
			}
			os.Exit(0)
		}),
	)

	// Run the service
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
