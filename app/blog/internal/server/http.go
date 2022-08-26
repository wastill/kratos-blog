package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"

	"github.com/gorilla/handlers"

	v1 "kratos-blog/api/blog/v1"
	"kratos-blog/app/blog/internal/conf"
	"kratos-blog/app/blog/internal/service"
)

// NewWhiteListMatcher 创建jwt白名单
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
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
		//selector.Server(
		//	jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
		//		return []byte(ac.ApiKey), nil
		//	}, jwt.WithSigningMethod(jwtv4.SigningMethodHS256)),
		//).
		//	Match(NewWhiteListMatcher()).
		//	Build(),
	)
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, logger log.Logger, userSvc *service.UserService, postSvc *service.PostService, linkSvc *service.LinkService, cateSvc *service.CategoryService, tagSvc *service.TagService) *http.Server {
	var opts = []http.ServerOption{
		NewMiddleware(ac, logger),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"" +
				"", "Content-Type", "Authorization"}),
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

	v1.RegisterUserServiceHTTPServer(srv, userSvc)
	v1.RegisterPostServiceHTTPServer(srv, postSvc)
	v1.RegisterCategoryServiceHTTPServer(srv, cateSvc)
	v1.RegisterTagServiceHTTPServer(srv, tagSvc)
	v1.RegisterLinkServiceHTTPServer(srv, linkSvc)

	return srv
}