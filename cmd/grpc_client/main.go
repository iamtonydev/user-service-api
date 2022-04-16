package main

import (
	"context"
	"fmt"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
	"google.golang.org/grpc"
	"log"
)

const address = "localhost:50051"

func main() {
	ctx := context.Background()

	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewUserV1Client(con)

	// add user
	res, err := client.AddUser(ctx, &desc.AddUserRequest{
		Name:  "Admin",
		Age:   100500,
		Email: "admin@no.no",
	})
	if err != nil {
		log.Fatalf("failed to add user %s", err.Error())
	}

	fmt.Println("=== user id ===")
	fmt.Printf("user id: %d\n", res.GetResult().GetId())

	//get user
	user, err := client.GetUser(ctx, &desc.GetUserRequest{Id: 1})
	if err != nil {
		log.Fatalf("failed to get user %s", err.Error())
	}

	fmt.Println("=== user info ===")
	fmt.Printf("name: %s\n", user.GetResult().GetName())
	fmt.Printf("age: %d\n", user.GetResult().GetAge())
	fmt.Printf("email: %s\n", user.GetResult().GetEmail())

	// multi add user
	reqData := []*desc.MultiAddUserRequest_User{
		{
			Name:  "Admin",
			Age:   100500,
			Email: "admin@no.no",
		},
		{
			Name:  "Admin2",
			Age:   100501,
			Email: "admin2@no.no",
		},
	}
	users, err := client.MultiAddUser(ctx, &desc.MultiAddUserRequest{Users: reqData})
	if err != nil {
		log.Fatalf("failed to multi add users %s", err.Error())
	}

	fmt.Println("=== user ids ===")
	fmt.Printf("user ids: %v\n", users.GetResult().GetId())

	// list users
	listUsers, err := client.ListUser(ctx, &desc.Empty{})
	if err != nil {
		log.Fatalf("failed to get list users %s", err.Error())
	}

	fmt.Println("=== users list ===")
	fmt.Printf("users info: %s\n", listUsers.GetResult())

	// delete user
	_, err = client.RemoveUser(ctx, &desc.RemoveUserRequest{Id: 1})
	if err != nil {
		log.Fatalf("failed to remove user %s", err.Error())
	}

	// update user
	updateUser, err := client.UpdateUser(ctx, &desc.UpdateUserRequest{
		Id:    1,
		Name:  "NewAdmin",
		Age:   100500,
		Email: "newadmin@no.no",
	})
	if err != nil {
		log.Fatalf("failed to update user %s", err.Error())
	}

	fmt.Println("=== user info ===")
	fmt.Printf("name: %s\n", updateUser.GetResult().GetName())
	fmt.Printf("age: %d\n", updateUser.GetResult().GetAge())
	fmt.Printf("email: %s\n", updateUser.GetResult().GetEmail())
}