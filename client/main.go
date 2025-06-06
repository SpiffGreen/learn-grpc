package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/spiffgreen/learn-grpc/shared/genproto/learn"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var userId = flag.Uint("userId", 1, "Defines the id of the user to retrieve")
	flag.Parse()
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewHelloAppClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	r, err := c.GetUser(ctx, &pb.UserRequest{Id: uint32(*userId)})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	fmt.Printf("Response\n ID: %d \n Name: %s \n Age: %d\n", r.GetId(), r.GetName(), r.GetAge())

	fmt.Printf("\n\n\n\n***** STREAM DATA BELOW ******\n\n\n")

	stream, err := c.GetUsers(ctx, &pb.EmptyMessage{})

	if err != nil {
		log.Fatalf("could not stream data: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err == nil {
			valStr := fmt.Sprintf("Response\n ID: %d \n Name: %s \n Age: %d\n", resp.GetId(), resp.GetName(), resp.GetAge())
			log.Println(valStr)
		}

		if err != nil {
			panic(err) // dont use panic in your real project
		}

	}

	fmt.Printf("\n\n\n\n***** REPEATED DATA BELOW ******\n\n\n")

	userArr, err := c.GetUsersNoStream(ctx, &pb.EmptyMessage{})

	if err != nil {
		log.Fatalf("could not get repeated data: %v", err)
	}

	for _, user := range userArr.Users {
		fmt.Printf("%d. %s, %d years old\n", user.GetId(), user.GetName(), user.GetAge())
	}

}
