package handler

import (
	"net/http"

	"github.com/yiGmMk/zero-paopao/api/internal/logic"
	"github.com/yiGmMk/zero-paopao/api/internal/svc"
	"github.com/yiGmMk/zero-paopao/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPostsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPostsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetPostsLogic(r.Context(), svcCtx)
		resp, err := l.GetPosts(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
