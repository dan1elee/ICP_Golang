package model

import (
	"ICP_Golang/conf"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// type BaseModel struct {
// 	ID        uint `gorm:"primary_key"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

type Student struct {
	StudentId  string `gorm:"type:varchar(40);primary_key;unique_index"`
	Name       string `gorm:"type:varchar(128)"`
	Password   string `gorm:"type:varchar(128)"`
	EvalCnt    uint   `gorm:"default:0"`
	DiscussCnt uint   `gorm:"default:0"`
}

type Teacher struct {
	TeacherId   string `gorm:"type:varchar(40);primary_key;unique_index"`
	Name        string `gorm:"type:varchar(128)"`
	Password    string `gorm:"type:varchar(128)"`
	DisscussCnt uint   `gorm:"default:0"`
	CourseCnt   uint   `gorm:"default:0"`
}

type Admin struct {
	AdminId  string `gorm:"type:varchar(40);primary_key;unique_index"`
	Name     string `gorm:"type:varchar(128)"`
	Password string `gorm:"type:varchar(128)"`
}

type Course struct {
	CourseId     string `gorm:"type:varchar(40);primary_key;unique_index"`
	Name         string `gorm:"type:varchar(128)"`
	Score        float32
	Introduction string `gorm:"type:varchar(1000)"`
	TeacherId    string `gorm:"type:varchar(40)"`
}

type StudentCourse struct {
	gorm.Model
	StudentId string `gorm:"type:varchar(40)"`
	CourseId  string `gorm:"type:varchar(40)"`
}

type CourseEval struct {
	EvalId    string `gorm:"type:varchar(40);primary_key;unique_index"`
	Time      time.Time
	Content   string `gorm:"type:varchar(1000)"`
	Score     int    `gorm:"default:5"`
	StudentId int    `gorm:"type:varchar(40)"`
	CourseId  int    `gorm:"type:varchar(40)"`
}

type Homework struct {
	HomeworkId string `gorm:"type:varchar(40);primary_key;unique_index"`
	EndTime    time.Time
	Content    string `gorm:"type:varchar(1000)"`
	IsTeacher  int
	TeacherId  string `gorm:"type:varchar(40)"`
	StudentId  string `gorm:"type:varchar(40)"`
	CourseId   string `gorm:"type:varchar(40)"`
}

type StudentHomework struct {
	gorm.Model
	StudentId  string `gorm:"type:varchar(40)"`
	HomeworkId string `gorm:"type:varchar(40)"`
	State      int
}

type MainComment struct {
	CommentId string `gorm:"type:varchar(40);primary_key;unique_index"`
	Title     string `gorm:"type:varchar(1000)"`
	Time      time.Time
	IsTeacher int
	Content   string `gorm:"type:varchar(1000)"`
	StudentId string
	TeacherId string
	AdminId   string
}

type SecondComment struct {
	SecondCommentId string `gorm:"type:varchar(40);primary_key;unique_index"`
	Time            time.Time
	IsTeacher       int
	Content         string `gorm:"type:varchar(1000)"`
	MainCommentId   string
	StudentId       string `gorm:"type:varchar(40)"`
	TeacherId       string `gorm:"type:varchar(40)"`
	AdminId         string `gorm:"type:varchar(40)"`
}

type StudentMain struct {
	gorm.Model
	StudentId     string `gorm:"type:varchar(40)"`
	MainCommentId string `gorm:"type:varchar(40)"`
}

type TeacherMain struct {
	gorm.Model
	TeacherId     string `gorm:"type:varchar(40)"`
	MainCommentId string `gorm:"type:varchar(40)"`
}

type AdminMain struct {
	gorm.Model
	AdminId       string `gorm:"type:varchar(40)"`
	MainCommentId string `gorm:"type:varchar(40)"`
}

type StudentSecond struct {
	gorm.Model
	StudentId       string `gorm:"type:varchar(40)"`
	SecondCommentId string `gorm:"type:varchar(40)"`
}

type TeacherSecond struct {
	gorm.Model
	TeacherId       string `gorm:"type:varchar(40)"`
	SecondCommentId string `gorm:"type:varchar(40)"`
}

type AdminSecond struct {
	gorm.Model
	AdminId         string `gorm:"type:varchar(40)"`
	SecondCommentId string `gorm:"type:varchar(40)"`
}

type StudentAddHomework struct {
	gorm.Model
	StudentId  string `gorm:"type:varchar(40)"`
	HomeWorkId string `gorm:"type:varchar(40)"`
}

type TeacherHomework struct {
	gorm.Model
	TeacherId  string `gorm:"type:varchar(40)"`
	HomeWorkId string `gorm:"type:varchar(40)"`
}

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open("mysql", conf.GetConfiguration().Database)
	if err != nil {
		return nil, err
	}
	//migrate todo
	return DB, err
}

func (teacher *Teacher) Insert() error {
	return DB.Create(teacher).Error
}

func (teacher *Teacher) UpdateName() error {
	return DB.Model(teacher).Updates(map[string]interface{}{
		"password": teacher.Password,
	}).Error
}

func (teacher *Teacher) UpdatePassword() error {
	return DB.Model(teacher).Updates(map[string]interface{}{
		"password": teacher.Password,
	}).Error
}

func (teacher *Teacher) IncreaseCourseCnt() error {
	return DB.Model(teacher).Updates(map[string]interface{}{
		"course_cnt": teacher.CourseCnt + 1,
	}).Error
}

func (teacher *Teacher) DecreaseCourseCnt() error {
	return DB.Model(teacher).Updates(map[string]interface{}{
		"course_cnt": teacher.CourseCnt - 1,
	}).Error
}

func (teacher *Teacher) IncreaseDiscussCnt() error {
	return DB.Model(teacher).Updates(map[string]interface{}{
		"discuss_cnt": teacher.DisscussCnt + 1,
	}).Error
}

func (teacher *Teacher) DecreaseDiscussCnt() error {
	return DB.Model(teacher).Updates(map[string]interface{}{
		"discuss_cnt": teacher.DisscussCnt - 1,
	}).Error
}

func (teacher *Teacher) Delete() error {
	return DB.Delete(teacher).Error
}

func (student *Student) Insert() error {
	return DB.Create(student).Error
}

func HasStudent(id string) bool {
	var existStudent Student
	result := DB.First(&existStudent, id)
	return !result.RecordNotFound()
}

func (student *Student) IncreaseEvalCnt() error {
	return DB.Model(student).Updates(map[string]interface{}{
		"eval_cnt": student.EvalCnt + 1,
	}).Error
}

func (student *Student) DecreaseEvalCnt() error {
	return DB.Model(student).Updates(map[string]interface{}{
		"eval_cnt": student.EvalCnt - 1,
	}).Error
}

func (student *Student) IncreaseDiscussCnt() error {
	return DB.Model(student).Updates(map[string]interface{}{
		"discuss_cnt": student.DiscussCnt + 1,
	}).Error
}

func (student *Student) DecreaseDiscussCnt() error {
	return DB.Model(student).Updates(map[string]interface{}{
		"discuss_cnt": student.DiscussCnt - 1,
	}).Error
}

func (course *Course) UpdateAvg(avg int) error {
	return DB.Model(course).Updates(map[string]interface{}{
		"score": avg,
	}).Error
}

func (studentHomework *StudentHomework) Insert() error {
	return DB.Create(studentHomework).Error
}

func (studentHomework *StudentHomework) Delete() error {
	return DB.Delete(studentHomework).Error
}

//todo

func (course *Course) AfterSave(db *gorm.DB) error {
	var thisTeacher Teacher
	db.First(&thisTeacher, course.TeacherId)
	return thisTeacher.IncreaseCourseCnt()
}

func (course *Course) BeforeDelete(db *gorm.DB) error {
	var thisTeacher Teacher
	db.First(&thisTeacher, course.TeacherId)
	return thisTeacher.DecreaseCourseCnt()
}

func (courseEval *CourseEval) AfterSave(db *gorm.DB) error {
	var thisStudent Student
	var err error
	db.First(&thisStudent, courseEval.StudentId)
	err = thisStudent.IncreaseEvalCnt()
	if err != nil {
		return err
	}
	type Result struct {
		CourseId string
		Average  int
	}
	var result Result
	db.Model(&CourseEval{}).Select("course_id, avg(score)").Where("course_id = ?",
		courseEval.CourseId).First(&result)
	var thisCourse Course
	db.First(&thisCourse, courseEval.CourseId)
	return thisCourse.UpdateAvg(result.Average)
}

func (courseEval *CourseEval) AfterDelete(db *gorm.DB) error {
	var thisStudent Student
	var err error
	db.First(&thisStudent, courseEval.StudentId)
	err = thisStudent.DecreaseEvalCnt()
	if err != nil {
		return err
	}
	type Result struct {
		CourseId string
		Average  int
	}
	var result Result
	db.Model(&CourseEval{}).Select("course_id, avg(score)").Where("course_id = ?",
		courseEval.CourseId).First(&result)
	var thisCourse Course
	db.First(&thisCourse, courseEval.CourseId)
	return thisCourse.UpdateAvg(result.Average)
}

func (homework *Homework) AfterSave(db *gorm.DB) error {
	if homework.IsTeacher == 1 {
		var thisCourse Course
		db.First(&thisCourse, homework.CourseId)
		var students []StudentCourse
		db.Find(&students, "course_id = ?", homework.CourseId)
		for _, student := range students {
			studentHomework := &StudentHomework{StudentId: student.StudentId, HomeworkId: homework.HomeworkId}
			err := studentHomework.Insert()
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		studentHomework := &StudentHomework{StudentId: homework.StudentId, HomeworkId: homework.HomeworkId}
		err := studentHomework.Insert()
		return err
	}
}

func (homework *Homework) BeforeDelete(db *gorm.DB) error {
	var studentHomeworks []StudentHomework
	if homework.IsTeacher == 1 {
		db.Find(&studentHomeworks, "homework_id = ?", homework.HomeworkId)
		for _, studentHomework := range studentHomeworks {
			err := studentHomework.Delete()
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		db.Find(&studentHomeworks, "homework_id = ? and student_id = ?", homework.HomeworkId, homework.StudentId)
		for _, studentHomework := range studentHomeworks {
			err := studentHomework.Delete()
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func (mainComment *MainComment) AfterSave(db *gorm.DB) error {
	if mainComment.IsTeacher == 1 {
		var thisTeacher Teacher
		db.First(&thisTeacher, mainComment.TeacherId)
		return thisTeacher.IncreaseDiscussCnt()
	} else {
		var thisStudent Student
		db.First(&thisStudent, mainComment.StudentId)
		return thisStudent.IncreaseDiscussCnt()
	}
}

func (mainComment *MainComment) BeforeDelete(db *gorm.DB) error {
	if mainComment.IsTeacher == 1 {
		var thisTeacher Teacher
		db.First(&thisTeacher, mainComment.TeacherId)
		return thisTeacher.DecreaseDiscussCnt()
	} else {
		var thisStudent Student
		db.First(&thisStudent, mainComment.StudentId)
		return thisStudent.DecreaseDiscussCnt()
	}
}

func (SecondComment *SecondComment) AfterSave(db *gorm.DB) error {
	if SecondComment.IsTeacher == 1 {
		var thisTeacher Teacher
		db.First(&thisTeacher, SecondComment.TeacherId)
		return thisTeacher.IncreaseDiscussCnt()
	} else {
		var thisStudent Student
		db.First(&thisStudent, SecondComment.StudentId)
		return thisStudent.IncreaseDiscussCnt()
	}
}

func (SecondComment *SecondComment) BeforeDelete(db *gorm.DB) error {
	if SecondComment.IsTeacher == 1 {
		var thisTeacher Teacher
		db.First(&thisTeacher, SecondComment.TeacherId)
		return thisTeacher.DecreaseDiscussCnt()
	} else {
		var thisStudent Student
		db.First(&thisStudent, SecondComment.StudentId)
		return thisStudent.DecreaseDiscussCnt()
	}
}
