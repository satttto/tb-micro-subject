package service

import (
	""
)


type SubjectService struct {}


func (s SubjectService) ListSubject(ctx context.Context, in *pb.SubjectListRequest) (*pb.SubjectListResponse, error) {
	log.Printf("Cursor: %v", in.GetCursor())

	// Retrieve
	var subjectList []SubjectModel
	db.Find(&subjectList)
	log.Println(&subjectList)

	subjects := []*pb.Subject{
		{Id: "1", Title: "title1", Members: []*pb.Member{{Id: "m1", Name: "Name"}}},
		{Id: "2", Title: "title2", Members: []*pb.Member{{Id: "m2", Name: "Name"}}},
	}
	return &pb.SubjectListResponse{Subjects: subjects, Size: 12, HasNext: true}, nil
}