package api

import (
	"ICP_Golang/encrypt"
	"ICP_Golang/model"
	"ICP_Golang/token"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StudentLogin(c *gin.Context) {
	userName, exist := c.GetPostForm("userName")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "description": "Lack userName or userPassword"})
		return
	}
	userPassword, exist := c.GetPostForm("userPassword")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "description": "Lack userName or userPassword"})
		return
	}
	fmt.Println("StudentLogin得到的学号和密码是 ", userName, " ", userPassword)
	var thisStudent model.Student
	result := model.DB.First(&thisStudent, userName)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "description": "Wrong userName or userPassword"})
		fmt.Println("找不到该学生")
		return
	}
	fmt.Println("请求登录的学生姓名为", thisStudent.Name)
	if userPassword == thisStudent.Password {
		tokenString, err := token.GenToken(userName)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": http.StatusServiceUnavailable, "description": "JWT process error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "user_nickname": thisStudent.Name, "token": tokenString})
		return
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "description": "Wrong userName or userPassword"})
		return
	}
}

func TeacherLogin(c *gin.Context) {
	userName, exist := c.GetPostForm("userName")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "description": "Lack userName or userPassword"})
		return
	}
	userPassword, exist := c.GetPostForm("userPassword")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "description": "Lack userName or userPassword"})
		return
	}
	fmt.Println("TeacherLogin得到的学号和密码是 ", userName, " ", userPassword)
	var thisTeacher model.Teacher
	result := model.DB.First(&thisTeacher, userName)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "description": "Wrong userName or userPassword"})
		fmt.Println("找不到该教师")
		return
	}
	fmt.Println("请求登录的教师姓名为", thisTeacher.Name)
	if userPassword == thisTeacher.Password {
		tokenString, err := token.GenToken(userName)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": http.StatusServiceUnavailable, "description": "JWT process error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "user_nickname": thisTeacher.Name, "token": tokenString})
		return
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "description": "Wrong userName or userPassword"})
		return
	}
}

func AdminLogin(c *gin.Context) {
	userName, exist := c.GetPostForm("userName")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "description": "Lack userName or userPassword"})
		return
	}
	userPassword, exist := c.GetPostForm("userPassword")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "description": "Lack userName or userPassword"})
		return
	}
	fmt.Println("AdminLogin得到的学号和密码是 ", userName, " ", userPassword)
	var thisAdmin model.Admin
	result := model.DB.First(&thisAdmin, userName)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "description": "Wrong userName or userPassword"})
		fmt.Println("找不到该管理员")
		return
	}
	fmt.Println("请求登录的管理员姓名为", thisAdmin.Name)
	if userPassword == thisAdmin.Password {
		tokenString, err := token.GenToken(userName)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": http.StatusServiceUnavailable, "description": "JWT process error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "user_nickname": thisAdmin.Name, "token": tokenString})
		return
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "description": "Wrong userName or userPassword"})
		return
	}
}

func StudentRegister(c *gin.Context) {
	var thisStudent model.Student
	if userNickName, exist := c.GetPostForm("userNickName"); exist {
		thisStudent.Name = userNickName
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "description": "Nickname field cannot be empty"})
		return
	}
	if userName, exist := c.GetPostForm("userName"); exist {
		thisStudent.StudentId = userName
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "description": "Name field cannot be empty"})
		return
	}
	if userPassword, exist := c.GetPostForm("userPassword"); exist {
		thisStudent.Password = encrypt.EncryptPassword(userPassword)
	}
	if model.HasStudent(thisStudent.StudentId) {
		c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict})
		return
	}
	thisStudent.Insert()
}
