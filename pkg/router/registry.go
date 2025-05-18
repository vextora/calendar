package router

import (
	"github.com/gin-gonic/gin"
)

type RouteRegistrar func(*gin.RouterGroup)

var registrars []RouteRegistrar

func Register(f RouteRegistrar) {
	registrars = append(registrars, f)
}

func RegisterRoutes(group *gin.RouterGroup) {
	for _, r := range registrars {
		r(group)
	}
}
