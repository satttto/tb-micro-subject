package main

import (
	"log"
	"net"

	pb "github.com/satttto/tb-proto/subject"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedSubjectServiceServer
}

func main() {

	// Connect to DB
	dsn := "host=127.0.0.1 user=user password=password dbname=subject port=5433 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect DB")
	}
	log.Println(db)

	// DB Migation
	if err := db.AutoMigrate(&SubjectModel{}); err != nil {
		log.Fatalln("Failed to migrate DB")
	}

	// Start server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSubjectServiceServer(s, &server{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
