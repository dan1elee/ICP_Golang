package main

import (
	"ICP_Golang/api"
	"ICP_Golang/conf"
	"flag"

	"github.com/gin-gonic/gin"
)

func init() {
	confFilePath := flag.String("configuration_file_path", "conf/conf.yaml", "The yaml configuration file")
	conf.InitConfiguration(*confFilePath)
	// db.AutoMigrate(&todoListModel{})
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
			register_api.POST("student", api.StudentRegister)
		}
	}

	router.Run()
}
