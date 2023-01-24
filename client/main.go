package main

import (
	pb "github.com/surajjyoti-personal/grpc-project/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error in connecting to the server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	//
	//names:= pb.NameList{
	//	Names: []string{"Suraj", "Ruhit","Kohli"}
	//}
	//
	callSayHello(client)
}
