package service

import (
	"context"
	"log"

	database "github.com/satttto/tb-micro-subject/db"
	pb "github.com/satttto/tb-proto/subject"
)

type Service interface {
	ListSubject(ctx context.Context, in *pb.SubjectListRequest) (*pb.SubjectListResponse, error)
}

type SubjectService struct {
	pb.UnimplementedSubjectServiceServer
	database database.DB
}

func InitService(db database.DB) Service {
	return &SubjectService{
		database: db,
	}
}

func (s *SubjectService) ListSubject(ctx context.Context, in *pb.SubjectListRequest) (*pb.SubjectListResponse, error) {
	log.Printf("Cursor: %v", in.GetCursor())

	// Retrieve
	//var subjectList []SubjectModel
	//db.Find(&subjectList)
	//log.Println(&subjectList)

	subjects := []*pb.Subject{
		{Id: "1", Title: "title1", Members: []*pb.Member{{Id: "m1", Name: "Name"}}},
		{Id: "2", Title: "title2", Members: []*pb.Member{{Id: "m2", Name: "Name"}}},
	}
	return &pb.SubjectListResponse{Subjects: subjects, Size: 12, HasNext: true}, nil
}
