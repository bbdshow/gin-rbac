package http

import (
	"github.com/bbdshow/bkit/ginutil"
	"github.com/bbdshow/gin-rabc/pkg/middleware"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"github.com/gin-gonic/gin"
)

// @Summary [RBAC 登录]
// @Description
// @Tags RBAC 登录/登出
// @Accept json
// @Produce json
// @Param Request body model.LoginAccountReq true "request param"
// @Success 200 {object} model.LoginAccountResp "success"
// @Router /rbac/login [post]
func loginAccount(c *gin.Context) {
	in := &model.LoginAccountReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	out := &model.LoginAccountResp{}
	err := svc.LoginAccount(c.Request.Context(), in, out)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespData(c, out)
}

// @Summary [RBAC 登出]
// @Description
// @Tags RBAC 登录/登出
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.LoginOutAccountReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/loginOut [post]
func loginOutAccount(c *gin.Context) {
	token, err := middleware.GetContextAccessToken(c)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err = svc.LoginOutAccount(c.Request.Context(), &model.LoginOutAccountReq{Token: token})
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespSuccess(c)
}
