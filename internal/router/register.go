package router

import "github.com/cloudwego/hertz/pkg/app/server"

func GeneratedRegister(r *server.Hertz) {
	AuthRoutes(r)
	TagsRoutes(r)
}
