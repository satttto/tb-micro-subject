package service

import (
	"github.com/satttto/tb-micro-subject/model"
	pb "github.com/satttto/tb-proto/subject"
)

// ModelをPBに変換する

func convertMemberToPb(member model.MemberModel) *pb.Member {
	return &pb.Member{
		Id:   member.ID,
		Name: member.Name,
	}
}

func convertMembersToPb(members []model.MemberModel) []*pb.Member {
	membersPb := make([]*pb.Member, len(members))
	for i, member := range members {
		memberPb := convertMemberToPb(member)
		membersPb[i] = memberPb
	}
	return membersPb
}

func convertSubjectToPb(subject *model.SubjectModel) *pb.Subject {
	membersPb := convertMembersToPb(subject.Members)
	return &pb.Subject{
		Id:      subject.ID,
		Title:   subject.Title,
		Members: membersPb,
	}
}

func convertSubjectsToPbs(subjects []*model.SubjectModel) []*pb.Subject {
	subjectsPb := make([]*pb.Subject, len(subjects))
	for i, subject := range subjects {
		subjectPb := convertSubjectToPb(subject)
		subjectsPb[i] = subjectPb
	}
	return subjectsPb
}
