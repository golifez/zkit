package server

import (
	"context"

	"github.com/golifez/zkit/internal/conf"
	"github.com/golifez/zkit/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/handlers"
)

type ServiceContainer struct {
	AuthService *service.AuthService
	// UserService  *service.UserService
	// 其他服务...
}

func NewServiceContainer(
	auth *service.AuthService,
	// register *service.RegisterService,
	// 其他服务...
) *ServiceContainer {
	return &ServiceContainer{
		AuthService: auth,
		// UserService:  user,
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, cg *conf.Config, service *ServiceContainer, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			selector.Server(
				jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
					return []byte(cg.Secretkey.Jwtkey), nil
				}),
			).Match(NewWhiteListMatcher()).Build(), // 设置白名单
			// metadata.Server(),
			logging.Server(logger),
		),
		http.Filter(handlers.CORS( // 浏览器跨域
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
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
	// v1.RegisterAuthHTTPServer(srv, service.AuthService)
	// v1.RegisterRegisterHTTPServer(srv, login)
	// v2.RegisterShopHTTPServer(srv, s)
	return srv
}

// NewWhiteListMatcher 设置白名单，不需要 token 验证的接口
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/auth.v1.Auth/Login"] = struct{}{}
	whiteList["/auth.v1.Auth/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
