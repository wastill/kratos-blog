package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-blog/api/content/service/v1"
	"kratos-blog/app/content/service/internal/biz"
	"kratos-blog/third_party/pagination"
)

type LinkService struct {
	v1.UnimplementedLinkServiceServer

	uc  *biz.LinkUseCase
	log *log.Helper
}

func NewLinkService(logger log.Logger, uc *biz.LinkUseCase) *LinkService {
	l := log.NewHelper(log.With(logger, "module", "link/service/content-service"))
	return &LinkService{
		log: l,
		uc:  uc,
	}
}

// ListLink 获取链接列表
func (s *LinkService) ListLink(ctx context.Context, req *pagination.PagingRequest) (*v1.ListLinkResponse, error) {
	return s.uc.List(ctx, req)
}

// GetLink 获取链接数据
func (s *LinkService) GetLink(ctx context.Context, req *v1.GetLinkRequest) (*v1.Link, error) {
	return s.uc.Get(ctx, req)
}

// CreateLink 创建链接
func (s *LinkService) CreateLink(ctx context.Context, req *v1.CreateLinkRequest) (*v1.Link, error) {
	return s.uc.Create(ctx, req)
}

// UpdateLink 更新链接
func (s *LinkService) UpdateLink(ctx context.Context, req *v1.UpdateLinkRequest) (*v1.Link, error) {
	return s.uc.Update(ctx, req)
}

// DeleteLink 删除链接
func (s *LinkService) DeleteLink(ctx context.Context, req *v1.DeleteLinkRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
