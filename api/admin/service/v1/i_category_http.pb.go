// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.2
// - protoc             v3.19.1
// source: i_category.proto

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

const OperationCategoryServiceCreateCategory = "/admin.service.v1.CategoryService/CreateCategory"
const OperationCategoryServiceDeleteCategory = "/admin.service.v1.CategoryService/DeleteCategory"
const OperationCategoryServiceGetCategory = "/admin.service.v1.CategoryService/GetCategory"
const OperationCategoryServiceListCategory = "/admin.service.v1.CategoryService/ListCategory"
const OperationCategoryServiceUpdateCategory = "/admin.service.v1.CategoryService/UpdateCategory"

type CategoryServiceHTTPServer interface {
	CreateCategory(context.Context, *v1.CreateCategoryRequest) (*v1.Category, error)
	DeleteCategory(context.Context, *v1.DeleteCategoryRequest) (*emptypb.Empty, error)
	GetCategory(context.Context, *v1.GetCategoryRequest) (*v1.Category, error)
	ListCategory(context.Context, *pagination.PagingRequest) (*v1.ListCategoryResponse, error)
	UpdateCategory(context.Context, *v1.UpdateCategoryRequest) (*v1.Category, error)
}

func RegisterCategoryServiceHTTPServer(s *http.Server, srv CategoryServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/blog/v1/categories", _CategoryService_ListCategory0_HTTP_Handler(srv))
	r.GET("/blog/v1/categories/{id}", _CategoryService_GetCategory0_HTTP_Handler(srv))
	r.POST("/blog/v1/categories", _CategoryService_CreateCategory0_HTTP_Handler(srv))
	r.PUT("/blog/v1/categories/{id}", _CategoryService_UpdateCategory0_HTTP_Handler(srv))
	r.DELETE("/blog/v1/categories/{id}", _CategoryService_DeleteCategory0_HTTP_Handler(srv))
}

func _CategoryService_ListCategory0_HTTP_Handler(srv CategoryServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in pagination.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCategoryServiceListCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCategory(ctx, req.(*pagination.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.ListCategoryResponse)
		return ctx.Result(200, reply)
	}
}

func _CategoryService_GetCategory0_HTTP_Handler(srv CategoryServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.GetCategoryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCategoryServiceGetCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCategory(ctx, req.(*v1.GetCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Category)
		return ctx.Result(200, reply)
	}
}

func _CategoryService_CreateCategory0_HTTP_Handler(srv CategoryServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.CreateCategoryRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCategoryServiceCreateCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCategory(ctx, req.(*v1.CreateCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Category)
		return ctx.Result(200, reply)
	}
}

func _CategoryService_UpdateCategory0_HTTP_Handler(srv CategoryServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.UpdateCategoryRequest
		if err := ctx.Bind(&in.Category); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCategoryServiceUpdateCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCategory(ctx, req.(*v1.UpdateCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Category)
		return ctx.Result(200, reply)
	}
}

func _CategoryService_DeleteCategory0_HTTP_Handler(srv CategoryServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.DeleteCategoryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCategoryServiceDeleteCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteCategory(ctx, req.(*v1.DeleteCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type CategoryServiceHTTPClient interface {
	CreateCategory(ctx context.Context, req *v1.CreateCategoryRequest, opts ...http.CallOption) (rsp *v1.Category, err error)
	DeleteCategory(ctx context.Context, req *v1.DeleteCategoryRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetCategory(ctx context.Context, req *v1.GetCategoryRequest, opts ...http.CallOption) (rsp *v1.Category, err error)
	ListCategory(ctx context.Context, req *pagination.PagingRequest, opts ...http.CallOption) (rsp *v1.ListCategoryResponse, err error)
	UpdateCategory(ctx context.Context, req *v1.UpdateCategoryRequest, opts ...http.CallOption) (rsp *v1.Category, err error)
}

type CategoryServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewCategoryServiceHTTPClient(client *http.Client) CategoryServiceHTTPClient {
	return &CategoryServiceHTTPClientImpl{client}
}

func (c *CategoryServiceHTTPClientImpl) CreateCategory(ctx context.Context, in *v1.CreateCategoryRequest, opts ...http.CallOption) (*v1.Category, error) {
	var out v1.Category
	pattern := "/blog/v1/categories"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCategoryServiceCreateCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CategoryServiceHTTPClientImpl) DeleteCategory(ctx context.Context, in *v1.DeleteCategoryRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/blog/v1/categories/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCategoryServiceDeleteCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CategoryServiceHTTPClientImpl) GetCategory(ctx context.Context, in *v1.GetCategoryRequest, opts ...http.CallOption) (*v1.Category, error) {
	var out v1.Category
	pattern := "/blog/v1/categories/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCategoryServiceGetCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CategoryServiceHTTPClientImpl) ListCategory(ctx context.Context, in *pagination.PagingRequest, opts ...http.CallOption) (*v1.ListCategoryResponse, error) {
	var out v1.ListCategoryResponse
	pattern := "/blog/v1/categories"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCategoryServiceListCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CategoryServiceHTTPClientImpl) UpdateCategory(ctx context.Context, in *v1.UpdateCategoryRequest, opts ...http.CallOption) (*v1.Category, error) {
	var out v1.Category
	pattern := "/blog/v1/categories/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCategoryServiceUpdateCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Category, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
