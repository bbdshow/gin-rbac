package http

import (
	"github.com/bbdshow/admin-rabc/pkg/model"
	"github.com/bbdshow/bkit/ginutil"
	"github.com/bbdshow/bkit/typ"
	"github.com/gin-gonic/gin"
)

// @Summary [角色配置列表]
// @Description
// @Tags RBAC 角色配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.ListRoleConfigReq true "request param"
// @Success 200 {object} model.ListRoleConfig "success"
// @Router /rbac/v1/role/list [get]
func listRoleConfig(c *gin.Context) {
	in := &model.ListRoleConfigReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	out := &typ.ListResp{}
	err := svc.ListRoleConfig(c.Request.Context(), in, out)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespData(c, out)
}

// @Summary [角色配置创建]
// @Description
// @Tags RBAC 角色配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.CreateRoleConfigReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/role/create [post]
func createRoleConfig(c *gin.Context) {
	in := &model.CreateRoleConfigReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.CreateRoleConfig(c.Request.Context(), in)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespSuccess(c)
}

// @Summary [角色配置更新]
// @Description
// @Tags RBAC 角色配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.UpdateRoleConfigReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/role/update [post]
func updateRoleConfig(c *gin.Context) {
	in := &model.UpdateRoleConfigReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.UpdateRoleConfig(c.Request.Context(), in)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespSuccess(c)
}