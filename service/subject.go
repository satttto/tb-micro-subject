package service

import (
	"context"
	"log"

	pb "github.com/satttto/tb-proto/subject"
)

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
