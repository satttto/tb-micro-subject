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
	db database.DB
}

func SetUpService(db database.DB) Service {
	return &SubjectService{
		db: db,
	}
}

func (s *SubjectService) ListSubject(ctx context.Context, in *pb.SubjectListRequest) (*pb.SubjectListResponse, error) {
	log.Printf("Cursor: %v", in.GetCursor())

	list, err := s.db.RetrieveSubjectList()
	if err != nil {
		log.Println(err)
	}
	log.Println(list)

	subjectsPb := convertSubjectsToPbs(list)

	return &pb.SubjectListResponse{Subjects: subjectsPb, Total: 12, Cursor: 3, HasNext: true}, nil
}
