package service

import (
	"context"

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
