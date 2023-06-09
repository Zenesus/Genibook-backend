package main

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"
	api_v1 "webscrapper/apis/v1"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Static("/images", "./images")

	r.LoadHTMLFiles("static/howGPA.html", "static/pp.html", "static/download.html")
	//r.Use(api_v1.JsonLoggerMiddleware())

	r.POST("/hi/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})

	})

	r.GET("/app/:filename", func(ctx *gin.Context) {
		fileName := ctx.Param("filename")
		targetPath := filepath.Join("app/", fileName)
		//This ckeck is for example, I not sure is it can prevent all possible filename attacks - will be much better if real filename will not come from user side. I not even tryed this code
		if !strings.HasPrefix(filepath.Clean(targetPath), "app/") {
			ctx.String(403, "Look like you attacking me")
			return
		}
		//Seems this headers needed for some browsers (for example without this headers Chrome will download files as txt)
		ctx.Header("Content-Description", "File Transfer")
		ctx.Header("Content-Transfer-Encoding", "binary")
		ctx.Header("Content-Disposition", "attachment; filename="+fileName)
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.File(targetPath)
	})

	r.GET("/pp/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "pp.html", gin.H{})
	})

	r.GET("/howGPA/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "howGPA.html", gin.H{})
	})
	r.GET("/genibook/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "download.html", gin.H{})
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
