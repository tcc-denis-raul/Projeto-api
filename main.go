package main

import (
	"os"

	"Projeto-api/api"
	"github.com/gin-gonic/gin"
)

var defaultPort = "5000"

func main() {
	/*Port*/
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := gin.Default()

	/*Route*/
	router.GET("/", api.Hello)
	router.GET("/types/courses", api.GetTypeCourses)
	router.GET("/courses", api.GetCourses)
	router.GET("/courses/questions", api.GetQuestions)
	router.GET("/users", api.GetUser)
	router.GET("/users/profile", api.GetUserPreferences)
	router.GET("/course/detail", api.GetDetailCourse)

	router.POST("/indicate/course", api.IndicateCourse)
	router.POST("/course/feedback", api.Feedback)
	router.POST("/users", api.CreateUser)
	router.POST("/users/update", api.UpdateUser)
	router.POST("/users/profile", api.SaveUserPreferences)

	router.Run(":" + port)
}
