package main

import (
	"web-student/server/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"web-student/server/services"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.ExposeHeaders = []string{"Authorization"}
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// TODO: static file middleware
	r.Use(middlewars.Frontend)
	api := r.Group("api")
	{
		// No login required
		user := api.Group("user")
		{
			// TODO: Login Service
			user.POST("login", services.Handler(&services.Login{}))
			// TODO: Register Service
			user.POST("register", services.Handler(&services.RegisterService{}))
		}

		auth := api.Group("")
		auth.Use(middlewars.JWTAuth())
		{
			course := auth.Group("course")
			{
				course.POST("", services.Handler(&services.CreateCourseService{}))
				course.GET("",services.Handler(&services.GetCoursesService{}))
				course.POST("select",services.Handler(&services.SelectCourseService{}))
				course.GET("select",services.Handler(&services.GetSelectedCoursesService{}))
			}
		}

	}

	return r
}

func main() {
	r := InitRouter() //返回
	r.Run(":9090")
}
