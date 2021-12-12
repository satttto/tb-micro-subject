package main

import (
	"log"
	"net"

	database "github.com/satttto/tb-micro-subject/db"
	service "github.com/satttto/tb-micro-subject/service"
	pb "github.com/satttto/tb-proto/subject"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func main() {

	// Connect to DB
	db, err := database.SetUpDB()
	if err != nil {
		log.Fatalln("Failed to connect DB")
	}

	// Prep and Serve
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	subjectService := service.SetUpService(db)
	s := grpc.NewServer()
	pb.RegisterSubjectServiceServer(s, subjectService)
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
