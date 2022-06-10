package service

import (
	"context"
	"eitems/app/user/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	pb "eitems/api/user"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserService) CreatUser(ctx context.Context, req *pb.CreatUserRequest) (*pb.CreatUserRes, error) {
	return &pb.CreatUserRes{}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserRes, error) {
	return &pb.GetUserRes{
		Id:       req.Id,
		Username: "pong",
	}, nil
}
