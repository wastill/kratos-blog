package server

import (
	"context"
	userV1 "kratos-blog/api/user/service/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"

	jwtV4 "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"

	v1 "kratos-blog/api/blog/v1"
	commentV1 "kratos-blog/api/comment/service/v1"
	fileV1 "kratos-blog/api/file/service/v1"

	"kratos-blog/app/admin/gateway/internal/conf"
	"kratos-blog/app/admin/gateway/internal/service"
)

// NewWhiteListMatcher 创建jwt白名单
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]bool)
	whiteList["/blog.v1.UserService/Register"] = true
	whiteList["/blog.v1.UserService/Login"] = true
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewMiddleware 创建中间件
func NewMiddleware(ac *conf.Auth, logger log.Logger) http.ServerOption {
	return http.Middleware(
		recovery.Recovery(),
		tracing.Server(),
		logging.Server(logger),
		selector.Server(
			jwt.Server(func(token *jwtV4.Token) (interface{}, error) {
				return []byte(ac.ApiKey), nil
			}, jwt.WithSigningMethod(jwtV4.SigningMethodHS256)),
		).
			Match(NewWhiteListMatcher()).
			Build(),
	)
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	c *conf.Server, ac *conf.Auth, logger log.Logger,
	userSvc *service.UserService,
	postSvc *service.PostService,
	linkSvc *service.LinkService,
	cateSvc *service.CategoryService,
	commentSvc *service.CommentService,
	tagSvc *service.TagService,
	attachmentSvc *service.AttachmentService,
) *http.Server {
	var opts = []http.ServerOption{
		NewMiddleware(ac, logger),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	h := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", h)

	v1.RegisterPostServiceHTTPServer(srv, postSvc)
	v1.RegisterCategoryServiceHTTPServer(srv, cateSvc)
	v1.RegisterTagServiceHTTPServer(srv, tagSvc)
	v1.RegisterLinkServiceHTTPServer(srv, linkSvc)
	userV1.RegisterUserServiceHTTPServer(srv, userSvc)
	fileV1.RegisterAttachmentServiceHTTPServer(srv, attachmentSvc)
	commentV1.RegisterCommentServiceHTTPServer(srv, commentSvc)

	return srv
}