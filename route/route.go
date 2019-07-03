package route

import (
	"github.com/qaqzzl/happy"
	"happy-demo/app/middleware"
	"happy-demo/app/server/v1"
)

func init() {
	route := happy.NewRoute()

	route.UseMiddleware(middleware.MiddlewareTest)

	//v1 := route.RouterGroup("v1")
	//v1.UseMiddleware(MiddlewareTestGroup)
	//{
	//	v1.RouteAny("route", ControllerGroup)
	//}

	route.RouteAny("route", v1.ControllerTest)
}