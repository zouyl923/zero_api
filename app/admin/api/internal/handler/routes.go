// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	v1articlearticle "blog/app/admin/api/internal/handler/v1/article/article"
	v1articlecategory "blog/app/admin/api/internal/handler/v1/article/category"
	v1login "blog/app/admin/api/internal/handler/v1/login"
	v1setadmin "blog/app/admin/api/internal/handler/v1/set/admin"
	v1setmenu "blog/app/admin/api/internal/handler/v1/set/menu"
	v1setpermission "blog/app/admin/api/internal/handler/v1/set/permission"
	v1setrole "blog/app/admin/api/internal/handler/v1/set/role"
	"blog/app/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CorsMiddleware, serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/v1/article/article/delete",
					Handler: v1articlearticle.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/article/article/info",
					Handler: v1articlearticle.InfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/article/article/page_list",
					Handler: v1articlearticle.PageHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/article/article/update",
					Handler: v1articlearticle.UpdateHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/admin"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CorsMiddleware, serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/v1/article/category/delete",
					Handler: v1articlecategory.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/article/category/info",
					Handler: v1articlecategory.InfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/article/category/tree_list",
					Handler: v1articlecategory.TreeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/article/category/update",
					Handler: v1articlecategory.UpdateHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/admin"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CorsMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/v1/login/env",
					Handler: v1login.EnvHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/login/login",
					Handler: v1login.LoginHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/admin"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CorsMiddleware, serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/v1/admin/info",
					Handler: v1login.InfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/login/login_out",
					Handler: v1login.LoginOutHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/admin"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CorsMiddleware, serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/v1/set/admin/all_list",
					Handler: v1setadmin.AllListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/set/admin/delete",
					Handler: v1setadmin.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/set/admin/info",
					Handler: v1setadmin.InfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/set/admin/page_list",
					Handler: v1setadmin.PageListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/set/admin/update",
					Handler: v1setadmin.UpdateHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/admin"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CorsMiddleware, serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/v1/set/menu/delete",
					Handler: v1setmenu.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/set/menu/info",
					Handler: v1setmenu.InfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/set/menu/tree_list",
					Handler: v1setmenu.TreeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/set/menu/update",
					Handler: v1setmenu.UpdateHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/admin"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CorsMiddleware, serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/v1/set/permission/delete",
					Handler: v1setpermission.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/set/permission/info",
					Handler: v1setpermission.InfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/set/permission/tree_list",
					Handler: v1setpermission.TreeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/set/permission/update",
					Handler: v1setpermission.UpdateHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/admin"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CorsMiddleware, serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/v1/set/role/all_list",
					Handler: v1setrole.AllListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/set/role/delete",
					Handler: v1setrole.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/set/role/info",
					Handler: v1setrole.InfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/set/role/page_list",
					Handler: v1setrole.ListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/set/role/update",
					Handler: v1setrole.UpdateHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/admin"),
	)
}
