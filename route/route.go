package route

import (
	"github.com/qaqzzl/happy"
	"happy-demo/app/middleware"
	"happy-demo/app/server/apiv1"
)

func init() {
	route := happy.NewRoute()

	route.UseMiddleware(middleware.MiddlewareTest)

	api_v1 := route.RouterGroup("apiv1")
	{
		api_v1.RouteAny("article/list", apiv1.ArticleList)
	}

	route.RouteAny("/", apiv1.ControllerTest)
}