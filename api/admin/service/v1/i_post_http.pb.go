// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.2
// - protoc             v3.19.1
// source: i_post.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v1 "kratos-blog/api/content/service/v1"
	pagination "kratos-blog/third_party/pagination"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationPostServiceCreatePost = "/admin.service.v1.PostService/CreatePost"
const OperationPostServiceDeletePost = "/admin.service.v1.PostService/DeletePost"
const OperationPostServiceGetPost = "/admin.service.v1.PostService/GetPost"
const OperationPostServiceListPost = "/admin.service.v1.PostService/ListPost"
const OperationPostServiceUpdatePost = "/admin.service.v1.PostService/UpdatePost"

type PostServiceHTTPServer interface {
	CreatePost(context.Context, *v1.CreatePostRequest) (*v1.Post, error)
	DeletePost(context.Context, *v1.DeletePostRequest) (*emptypb.Empty, error)
	GetPost(context.Context, *v1.GetPostRequest) (*v1.Post, error)
	ListPost(context.Context, *pagination.PagingRequest) (*v1.ListPostResponse, error)
	UpdatePost(context.Context, *v1.UpdatePostRequest) (*v1.Post, error)
}

func RegisterPostServiceHTTPServer(s *http.Server, srv PostServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/blog/v1/posts", _PostService_ListPost0_HTTP_Handler(srv))
	r.GET("/blog/v1/posts/{id}", _PostService_GetPost0_HTTP_Handler(srv))
	r.POST("/blog/v1/posts", _PostService_CreatePost0_HTTP_Handler(srv))
	r.PUT("/blog/v1/posts/{id}", _PostService_UpdatePost0_HTTP_Handler(srv))
	r.DELETE("/blog/v1/posts/{id}", _PostService_DeletePost0_HTTP_Handler(srv))
}

func _PostService_ListPost0_HTTP_Handler(srv PostServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in pagination.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostServiceListPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListPost(ctx, req.(*pagination.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.ListPostResponse)
		return ctx.Result(200, reply)
	}
}

func _PostService_GetPost0_HTTP_Handler(srv PostServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.GetPostRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostServiceGetPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPost(ctx, req.(*v1.GetPostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Post)
		return ctx.Result(200, reply)
	}
}

func _PostService_CreatePost0_HTTP_Handler(srv PostServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.CreatePostRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostServiceCreatePost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePost(ctx, req.(*v1.CreatePostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Post)
		return ctx.Result(200, reply)
	}
}

func _PostService_UpdatePost0_HTTP_Handler(srv PostServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.UpdatePostRequest
		if err := ctx.Bind(&in.Post); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostServiceUpdatePost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePost(ctx, req.(*v1.UpdatePostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Post)
		return ctx.Result(200, reply)
	}
}

func _PostService_DeletePost0_HTTP_Handler(srv PostServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.DeletePostRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostServiceDeletePost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeletePost(ctx, req.(*v1.DeletePostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type PostServiceHTTPClient interface {
	CreatePost(ctx context.Context, req *v1.CreatePostRequest, opts ...http.CallOption) (rsp *v1.Post, err error)
	DeletePost(ctx context.Context, req *v1.DeletePostRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetPost(ctx context.Context, req *v1.GetPostRequest, opts ...http.CallOption) (rsp *v1.Post, err error)
	ListPost(ctx context.Context, req *pagination.PagingRequest, opts ...http.CallOption) (rsp *v1.ListPostResponse, err error)
	UpdatePost(ctx context.Context, req *v1.UpdatePostRequest, opts ...http.CallOption) (rsp *v1.Post, err error)
}

type PostServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewPostServiceHTTPClient(client *http.Client) PostServiceHTTPClient {
	return &PostServiceHTTPClientImpl{client}
}

func (c *PostServiceHTTPClientImpl) CreatePost(ctx context.Context, in *v1.CreatePostRequest, opts ...http.CallOption) (*v1.Post, error) {
	var out v1.Post
	pattern := "/blog/v1/posts"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPostServiceCreatePost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PostServiceHTTPClientImpl) DeletePost(ctx context.Context, in *v1.DeletePostRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/blog/v1/posts/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPostServiceDeletePost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PostServiceHTTPClientImpl) GetPost(ctx context.Context, in *v1.GetPostRequest, opts ...http.CallOption) (*v1.Post, error) {
	var out v1.Post
	pattern := "/blog/v1/posts/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPostServiceGetPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PostServiceHTTPClientImpl) ListPost(ctx context.Context, in *pagination.PagingRequest, opts ...http.CallOption) (*v1.ListPostResponse, error) {
	var out v1.ListPostResponse
	pattern := "/blog/v1/posts"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPostServiceListPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PostServiceHTTPClientImpl) UpdatePost(ctx context.Context, in *v1.UpdatePostRequest, opts ...http.CallOption) (*v1.Post, error) {
	var out v1.Post
	pattern := "/blog/v1/posts/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPostServiceUpdatePost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Post, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
