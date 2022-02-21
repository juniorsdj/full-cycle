package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/codeedu/fc2-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to the gRPC server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	// AddUser(client)
	// AddUserVerbose(client)
	// AddUsers(client)
	AddUserStreamBoth(client)

}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "001",
		Name:  "Manoel",
		Email: "mma@gmail.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatal("Fail on gRPC request %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "001",
		Name:  "Manoel",
		Email: "mma@gmail.com",
	}

	resStream, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		log.Fatal("Fail on gRPC request %v", err)
	}

	for {
		stream, err := resStream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the msg %v", err)

		}

		fmt.Println("status:", stream.Status)
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "001",
			Name:  "Manoel 1",
			Email: "mma@gmail.com",
		},
		&pb.User{
			Id:    "002",
			Name:  "Manoel 2",
			Email: "mma@gmail.com",
		},
		&pb.User{
			Id:    "003",
			Name:  "Manoel 3",
			Email: "mma@gmail.com",
		},
		&pb.User{
			Id:    "004",
			Name:  "Manoel 4",
			Email: "mma@gmail.com",
		},
		&pb.User{
			Id:    "005",
			Name:  "Manoel 5",
			Email: "mma@gmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Error creating request %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receiving response %v", err)
	}

	fmt.Println(res)

}

func AddUserStreamBoth(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "001",
			Name:  "Manoel 1",
			Email: "mma@gmail.com",
		},
		&pb.User{
			Id:    "002",
			Name:  "Manoel 2",
			Email: "mma@gmail.com",
		},
		&pb.User{
			Id:    "003",
			Name:  "Manoel 3",
			Email: "mma@gmail.com",
		},
		&pb.User{
			Id:    "004",
			Name:  "Manoel 4",
			Email: "mma@gmail.com",
		},
		&pb.User{
			Id:    "005",
			Name:  "Manoel 5",
			Email: "mma@gmail.com",
		},
	}

	stream, err := client.AddUserStreamBoth(context.Background())

	if err != nil {
		log.Fatalf("Error creating request %v", err)
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user ", req.GetId())
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			stream, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Could not receive the msg %v", err)
				break
			}

			fmt.Println("Receving user:", stream.User.GetId())
			fmt.Println("status:", stream.Status)
		}
		close(wait)
	}()
	<-wait
}
