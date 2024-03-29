package http

import (
	"github.com/bbdshow/gin-rabc/pkg/server/http/middleware"
	"github.com/gin-gonic/gin"
)

func midAccessTokenVerify() gin.HandlerFunc {
	return middleware.AccessTokenVerify(cfg.Admin.AuthEnable, svc.VerifyAccountToken)
}

func midRBACEnforce() gin.HandlerFunc {
	enable := cfg.Admin.AuthEnable && cfg.Casbin.Enable
	return middleware.RBACEnforce(enable, svc.GetSyncedEnforcer())
}

func midApiSign() gin.HandlerFunc {
	return middleware.ApiSign(cfg.Admin.SignEnable, svc.GetSecretKey)
}

func RegisterAdminRouter(e *gin.Engine) {
	e.POST("/rbac/login", loginAccount)

	e.POST("/rbac/loginOut", midAccessTokenVerify(), loginOutAccount)
	e.POST("/rbac/password/update", midAccessTokenVerify(), updateAccountPassword)

	rbac := e.Group("/rbac/v1")

	app := rbac.Group("/app").Use(midAccessTokenVerify()).Use(midRBACEnforce())
	app.GET("/list", listAppConfig)
	app.POST("/create", createAppConfig)
	app.POST("/update", updateAppConfig)
	app.POST("/delete", delAppConfig)

	rbac.GET("/app/select", selectAppConfig)

	action := rbac.Group("/action").Use(midAccessTokenVerify()).Use(midRBACEnforce())
	action.GET("/list", listActionConfig)
	action.POST("/create", createActionConfig)
	action.POST("/update", updateActionConfig)
	action.POST("/delete", delActionConfig)
	action.POST("/find", findActionConfig)
	action.POST("/import", importSwaggerToActions)

	menu := rbac.Group("/menu").Use(midAccessTokenVerify()).Use(midRBACEnforce())
	menu.GET("/tree", treeMenuConfig)
	menu.GET("/actions", menuActions)
	menu.POST("/create", createMenuConfig)
	menu.POST("/update", updateMenuConfig)
	menu.POST("/delete", delMenuConfig)
	menu.POST("/action/update", updateMenuConfigAction)

	account := rbac.Group("/account").Use(midAccessTokenVerify()).Use(midRBACEnforce())
	account.GET("/list", listAccount)
	account.POST("/create", createAccount)
	account.POST("/update", updateAccount)
	account.POST("/delete", delAccount)
	account.POST("/pwd/reset", resetAccountPassword)
	account.POST("/role/update", updateAccountRole)
	rbac.GET("/account/menu/auth", midAccessTokenVerify(), getAccountMenuAuth)

	role := rbac.Group("/role").Use(midAccessTokenVerify()).Use(midRBACEnforce())
	role.GET("/list", listRoleConfig)
	role.POST("/create", createRoleConfig)
	role.POST("/update", updateRoleConfig)
	role.POST("/delete", delRoleConfig)
	role.GET("/action", getRoleMenuAction)
	role.POST("/action/upsert", upsertRoleMenuAction)

	verify := rbac.Group("/verify").Use(midApiSign())
	verify.POST("/enforce", rbacEnforce)

	// Admin 单页应用
	e.StaticFile("/favicon.ico", "web/dist/favicon.ico")
	e.Static("/static", "web/dist/static")
	e.Static("/admin", "web/dist")
}
