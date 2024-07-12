package api

import (
	"ICP_Golang/encrypt"
	"ICP_Golang/enums"
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
	exist, thisStudent := model.GetExistStudent(userName)
	if !exist {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "description": "Wrong userName or userPassword"})
		fmt.Println("找不到该学生")
		return
	}
	fmt.Println("请求登录的学生姓名为", thisStudent.Name)
	if userPassword == thisStudent.Password {
		tokenString, err := token.GenToken(userName, enums.STUDENT)
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
	exist, thisTeacher := model.GetExistTeacher(userName)
	if !exist {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "description": "Wrong userName or userPassword"})
		fmt.Println("找不到该教师")
		return
	}
	fmt.Println("请求登录的教师姓名为", thisTeacher.Name)
	if userPassword == thisTeacher.Password {
		tokenString, err := token.GenToken(userName, enums.TEACHER)
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
	exist, thisAdmin := model.GetExistAdmin(userName)
	if !exist {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "description": "Wrong userName or userPassword"})
		fmt.Println("找不到该管理员")
		return
	}
	fmt.Println("请求登录的管理员姓名为", thisAdmin.Name)
	if userPassword == thisAdmin.Password {
		tokenString, err := token.GenToken(userName, enums.ADMIN)
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
	if err := thisStudent.Insert(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	return
}

func TeacherRegister(c *gin.Context) {
	var thisTeacher model.Teacher
	if userNickName, exist := c.GetPostForm("userNickName"); exist {
		thisTeacher.Name = userNickName
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "description": "Nickname field cannot be empty"})
		return
	}
	if userName, exist := c.GetPostForm("userName"); exist {
		thisTeacher.TeacherId = userName
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "description": "Name field cannot be empty"})
		return
	}
	if userPassword, exist := c.GetPostForm("userPassword"); exist {
		thisTeacher.Password = encrypt.EncryptPassword(userPassword)
	}
	if model.HasTeacher(thisTeacher.TeacherId) {
		c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict})
		return
	}
	if err := thisTeacher.Insert(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	return
}

func GetAllAvailableCourses(c *gin.Context) {
	userName, exist := c.GetQuery("userName")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	userToken, exist := c.GetQuery("token")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	tokenStruct, err := token.ParseToken(userToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "err": err})
		return
	}
	if !token.HaveAccess(tokenStruct, enums.LEVELNORMAL) {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden})
		return
	}
	selectedCourseIds := model.GetStudentSelectedCourse(userName)
	availableCourses := model.GetExtraCourses(selectedCourseIds)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "courses": availableCourses})
	return
}

func GetAllCourses(c *gin.Context) {
	userToken, exist := c.GetQuery("token")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	tokenStruct, err := token.ParseToken(userToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "err": err})
		return
	}
	if !token.HaveAccess(tokenStruct, enums.LEVELNORMAL) {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden})
		return
	}
	// todo 权限验证具体文件权限
	allCourses := model.GetAllCourses()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "courses": allCourses})
	return
}

func AddSelectCourse(c *gin.Context) {
	userToken, exist := c.GetPostForm("token")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	tokenStruct, err := token.ParseToken(userToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "err": err})
		return
	}
	if !token.HaveAccess(tokenStruct, enums.LEVELNORMAL) {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden})
		return
	}
	// todo 权限验证具体文件权限
	studentId, exist := c.GetPostForm("sid")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	courseId, exist := c.GetPostForm("cid")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	homework_ids := model.GetCourseHomeworks(courseId)
	model.AddNewSelectHomework(studentId, homework_ids)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated})
}

func GetCourseInfo(c *gin.Context) {
	userToken, exist := c.GetPostForm("token")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	tokenStruct, err := token.ParseToken(userToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "err": err})
		return
	}
	if !token.HaveAccess(tokenStruct, enums.LEVELNORMAL) {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden})
		return
	}
	courseId, exist := c.GetQuery("courseId")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	res, err := model.GetCourseInfoById(courseId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "info": res})
	return
}

func StudentPasswordChange(c *gin.Context) {
	studentId, exist := c.GetQuery("id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	prevPassword, exist := c.GetQuery("PassWord0")
	newPassword, exist := c.GetQuery("Password1")
	_, thisStudent := model.GetExistStudent(studentId)
	if thisStudent.Password != prevPassword {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	err := thisStudent.UpdatePassword(newPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	return
}

func TeacherPasswordChange(c *gin.Context) {
	teacherId, exist := c.GetQuery("id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	prevPassword, exist := c.GetQuery("PassWord0")
	newPassword, exist := c.GetQuery("Password1")
	_, thisTeacher := model.GetExistTeacher(teacherId)
	if thisTeacher.Password != prevPassword {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	err := thisTeacher.UpdatePassword(newPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	return
}

func AdminPasswordChange(c *gin.Context) {
	adminId, exist := c.GetQuery("id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	prevPassword, exist := c.GetQuery("PassWord0")
	newPassword, exist := c.GetQuery("Password1")
	_, thisAdmin := model.GetExistAdmin(adminId)
	if thisAdmin.Password != prevPassword {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	err := thisAdmin.UpdatePassword(newPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	return
}
