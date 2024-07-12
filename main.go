package main

import (
	"ICP_Golang/api"
	"ICP_Golang/conf"
	"flag"

	"github.com/gin-gonic/gin"
)

func init() {
	confFilePath := flag.String("conf", "conf/conf.yaml", "The yaml configuration file")
	flag.Parse()
	conf.InitConfiguration(*confFilePath)
}

func main() {
	router := gin.Default()

	icp_api := router.Group("/")
	{
		login_api := icp_api.Group("login/")
		{
			login_api.POST("student/", api.StudentLogin)
			login_api.POST("teacher/", api.TeacherLogin)
			login_api.POST("admin/", api.AdminLogin)
		}
		register_api := icp_api.Group("register/")
		{
			register_api.POST("student/", api.StudentRegister)
			register_api.POST("teacher/", api.TeacherRegister)
		}
		courses_api := icp_api.Group("courses/")
		{
			courses_api.GET("available/", api.GetAllAvailableCourses)
			courses_api.GET("all/", api.GetAllCourses)
			courses_api.PUT("select/", api.AddSelectCourse)
		}
	}

	router.Run()
}
