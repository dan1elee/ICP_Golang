package api

import (
	"ICP_Golang/encrypt"
	"ICP_Golang/enums"
	"ICP_Golang/idhandler"
	"ICP_Golang/model"
	"ICP_Golang/token"
	"fmt"
	"net/http"
	"strconv"
	"time"

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
	err := tokenValidation(c, enums.LEVELNORMAL)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "err": err})
		return
	}
	userName, exist := c.GetQuery("userName")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	selectedCourses := model.GetStudentSelectedCourse(userName)
	availableCourses := model.GetExtraCourses(selectedCourses)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "courses": availableCourses})
	return
}

func GetAllCourses(c *gin.Context) {
	err := tokenValidation(c, enums.LEVELNORMAL)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "err": err})
		return
	}
	// todo 权限验证具体文件权限
	allCourses := model.GetAllCourses()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "courses": allCourses})
	return
}

func AddSelectCourse(c *gin.Context) {
	err := tokenValidation(c, enums.LEVELNORMAL)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "err": err})
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
	err := tokenValidation(c, enums.LEVELNORMAL)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "err": err})
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
	err := tokenValidation(c, enums.LEVELNORMAL)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "err": err})
		return
	}
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
	err = thisStudent.UpdatePassword(newPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	return
}

func TeacherPasswordChange(c *gin.Context) {
	err := tokenValidation(c, enums.LEVELSECRET)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "err": err})
		return
	}
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
	err = thisTeacher.UpdatePassword(newPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	return
}

func AdminPasswordChange(c *gin.Context) {
	err := tokenValidation(c, enums.LEVELTOP)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "err": err})
		return
	}
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
	err = thisAdmin.UpdatePassword(newPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	return
}

func GetStudentSelectedCourse(c *gin.Context) {
	err := tokenValidation(c, enums.LEVELNORMAL)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "err": err})
		return
	}
	studentId, exist := c.GetQuery("sid")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
	}
	courses := model.GetStudentSelectedCourse(studentId)
	var courseInfos []map[string]interface{}
	for _, course := range courses {
		courseInfos = append(courseInfos, course.ToMap())
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "courses": courseInfos})
	return
}

func DropSelectedCourse(c *gin.Context) {
	err := tokenValidation(c, enums.LEVELNORMAL)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "err": err})
		return
	}
	studentId, exist := c.GetQuery("sid")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	courseId, exist := c.GetQuery("cid")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	err = model.DropStudentCourse(studentId, courseId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	return
}

func GetTeacherCourseList(c *gin.Context) {
	err := tokenValidation(c, enums.LEVELSECRET)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "err": err})
		return
	}
	teacherId, exist := c.GetQuery("id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	coursesInfo := model.GetCoursesInfoByTeacherId(teacherId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "courses": coursesInfo})
	return
}

func BuildCourse(c *gin.Context) {
	err := tokenValidation(c, enums.LEVELSECRET)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "err": err})
		return
	}
	teacherId, exist := c.GetPostForm("id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	courseName, exist := c.GetPostForm("name")
	if !exist || courseName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	courseIntro, exist := c.GetPostForm("intro")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	newCourse := model.Course{CourseId: idhandler.GenCourseId(), Name: courseName, Introduction: courseIntro, TeacherId: teacherId}
	err = newCourse.Insert()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	return
}

func UpdateCourse(c *gin.Context) {
	err := tokenValidation(c, enums.LEVELSECRET)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "err": err})
		return
	}
	courseId, exist := c.GetQuery("cid")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	courseName, exist := c.GetQuery("name")
	if !exist || courseName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	courseIntro, exist := c.GetQuery("intro")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	err = model.CourseUpdate(courseId, courseName, courseIntro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func DeleteCourse(c *gin.Context) {
	err := tokenValidation(c, enums.LEVELSECRET)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "err": err})
		return
	}
	courseId, exist := c.GetQuery("cid")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	err = model.DeleteCourse(courseId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	// 可能需要删除一些连续表
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	return
}

func GetCommentList(c *gin.Context) {
	err := tokenValidation(c, enums.LEVELNORMAL)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "err": err})
		return
	}
	courseId, exist := c.GetQuery("cid")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	courseCommentList := model.GetCourseCommentList(courseId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "comments": courseCommentList})
}

func CommentCourse(c *gin.Context) {
	courseId, exist := c.GetQuery("cid")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	userId, exist := c.GetQuery("uid")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	content, exist := c.GetQuery("content")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	timeStr, exist := c.GetQuery("time")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	timeVal, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	score, exist := c.GetQuery("score")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	scoreValue, err := strconv.Atoi(score)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	var courseEval model.CourseEval
	courseEval.EvalId = idhandler.GenEvalId()
	courseEval.Time = timeVal
	courseEval.Content = content
	courseEval.Score = scoreValue
	courseEval.StudentId = userId
	courseEval.CourseId = courseId
	err = courseEval.Insert()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	return
}
