package main

import (
	"log"
	"net/http"
	api_v1 "webscrapper/apis/v1"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLFiles("h_pp.html")
	r.LoadHTMLFiles("h_howGPA.html")
	//r.Use(api_v1.JsonLoggerMiddleware())

	r.POST("/hi/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})

	})

	r.GET("/pp/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "h_pp.html", gin.H{})
	})

	r.GET("/howGPA/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "h_howGPA.html", gin.H{})
	})

	r.POST("/apiv1/mps/", api_v1.MakeHandler(api_v1.MpsHandlerV1))
	r.POST("/apiv1/student/", api_v1.MakeHandler(api_v1.StudentHandlerV1))
	r.POST("/apiv1/schedule/", api_v1.MakeHandler(api_v1.ScheduleAssignmentHandlerV1))
	r.POST("/apiv1/assignments/", api_v1.MakeHandler(api_v1.AssignmentHandlerV1))
	r.POST("/apiv1/grades/", api_v1.MakeHandler(api_v1.GradesHandlerV1))
	r.POST("/apiv1/profile/", api_v1.MakeHandler(api_v1.ProfileHandlerV1))
	r.POST("/apiv1/login/", api_v1.MakeHandler(api_v1.LoginHandlerV1))
	r.POST("/apiv1/gpas/", api_v1.MakeHandler(api_v1.GPAshandlerV1))
	r.POST("/apiv1/gpas_his/", api_v1.MakeHandler(api_v1.GPAHistoryHandlerV1))
	r.POST("/apiv1/grade_of_students/", api_v1.MakeHandler(api_v1.StudentGradesHandlerV1))
	r.POST("/apiv1/ids/", api_v1.MakeHandler(api_v1.StudentIDHandlerV1))

	log.Fatal(r.Run(":6969"))

}
