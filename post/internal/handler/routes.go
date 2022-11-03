// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/yiGmMk/zero-paopao/post/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/v1/posts",
				Handler: GetPostsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/tags",
				Handler: GetTagsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/v1/auth/regiter",
				Handler: UserRegisterHandler(serverCtx),
			},
		},
	)
}
