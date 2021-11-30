package main

import (
	"context"
	"log"
	"net"

	pb "github.com/satttto/tb-proto/subject"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedSubjectServiceServer
}

func (s *server) ListSubject(ctx context.Context, in *pb.SubjectListRequest) (*pb.SubjectListResponse, error) {
	log.Printf("Cursor: %v", in.GetCursor())
	subjects := []*pb.Subject{
		{Id: "1", Title: "title1", Members: []*pb.Member{{Id: "m1", Name: "Name"}}},
		{Id: "2", Title: "title2", Members: []*pb.Member{{Id: "m1", Name: "Name"}}},
	}
	return &pb.SubjectListResponse{Subjects: subjects, Size: 12, HasNext: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSubjectServiceServer(s, &server{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
