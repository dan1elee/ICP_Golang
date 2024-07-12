package model

import (
	"ICP_Golang/conf"
	"encoding/json"
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
	CourseId     string  `gorm:"type:varchar(40);primary_key;unique_index" json:"course_id"`
	Name         string  `gorm:"type:varchar(128)" json:"course_name"`
	Score        float32 `json:"score"`
	Introduction string  `gorm:"type:varchar(1000)" json:"course_intro"`
	TeacherId    string  `gorm:"type:varchar(40)" json:"teacher_id"`
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
	DB.AutoMigrate(&Student{}, &Teacher{}, &Admin{},
		&Course{}, &StudentCourse{}, &CourseEval{},
		&Homework{}, &StudentHomework{},
		&MainComment{}, &SecondComment{},
		&StudentMain{}, &TeacherMain{}, &AdminMain{},
		&StudentSecond{}, &TeacherSecond{}, &AdminSecond{},
		&StudentAddHomework{}, &TeacherHomework{})
	//migrate todo
	return DB, err
}

func (teacher *Teacher) Insert() error {
	return DB.Model(&Teacher{}).Create(teacher).Error
}

func HasTeacher(id string) bool {
	var existTeacher Teacher
	result := DB.Model(&Teacher{}).First(&existTeacher, id)
	return !result.RecordNotFound()
}

func GetExistTeacher(id string) (bool, *Teacher) {
	var existTeacher = new(Teacher)
	result := DB.Model(&Teacher{}).First(&existTeacher, id)
	exist := !result.RecordNotFound()
	if !exist {
		return false, nil
	} else {
		return true, existTeacher
	}
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
	return DB.Model(&Teacher{}).Delete(teacher).Error
}

func (student *Student) Insert() error {
	return DB.Model(&Student{}).Create(student).Error
}

func HasStudent(id string) bool {
	var existStudent Student
	result := DB.Model(&Student{}).First(&existStudent, id)
	return !result.RecordNotFound()
}

func GetExistStudent(id string) (bool, *Student) {
	var existStudent = new(Student)
	result := DB.Model(&Student{}).First(&existStudent, id)
	exist := !result.RecordNotFound()
	if !exist {
		return false, nil
	} else {
		return true, existStudent
	}
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

func GetExistAdmin(id string) (bool, *Admin) {
	var existAdmin = new(Admin)
	result := DB.Model(&Admin{}).First(&existAdmin, id)
	exist := !result.RecordNotFound()
	if !exist {
		return false, nil
	} else {
		return true, existAdmin
	}
}

func (course *Course) UpdateAvg(avg int) error {
	return DB.Model(course).Updates(map[string]interface{}{
		"score": avg,
	}).Error
}

func GetAllCourses() []map[string]interface{} {
	var courses []Course
	DB.Model(&Course{}).Find(&courses)
	var courseInfos []map[string]interface{}
	for _, course := range courses {
		courseBytes, _ := json.Marshal(&course)
		courseInfo := new(map[string]interface{})
		json.Unmarshal(courseBytes, courseInfo)
		courseInfos = append(courseInfos, *courseInfo)
	}
	return courseInfos
}

func (studentHomework *StudentHomework) Insert() error {
	return DB.Model(&StudentHomework{}).Create(studentHomework).Error
}

func (studentHomework *StudentHomework) Delete() error {
	return DB.Model(&StudentHomework{}).Delete(studentHomework).Error
}

func GetStudentSelectedCourse(id string) []string {
	var thisStudentCourses []StudentCourse
	DB.Model(&StudentCourse{}).Where(map[string]interface{}{
		"student_id": id,
	}).Find(&thisStudentCourses)
	var courseIds []string
	for _, thisStudentCourse := range thisStudentCourses {
		courseIds = append(courseIds, thisStudentCourse.CourseId)
	}
	return courseIds
}

func GetExtraCourses(ids []string) []map[string]interface{} {
	var courses []Course
	DB.Model(&Course{}).Not(map[string]interface{}{"course_id": ids}).Find(&courses)
	var courseInfos []map[string]interface{}
	for _, course := range courses {
		courseBytes, _ := json.Marshal(&course)
		courseInfo := new(map[string]interface{})
		json.Unmarshal(courseBytes, courseInfo)
		courseInfos = append(courseInfos, *courseInfo)
	}
	return courseInfos
}

func GetCourseHomeworks(id string) []string {
	var homework_ids []string
	DB.Model(&Homework{}).Where(map[string]interface{}{
		"course_id": id,
	}).Pluck("homework_id", &homework_ids)
	return homework_ids
}

func AddNewSelectHomework(studentId string, homework_ids []string) {
	var stuHomeworkInfos []map[string]interface{}
	for _, homework_id := range homework_ids {
		stuHomeworkInfos = append(stuHomeworkInfos, map[string]interface{}{
			"student_id":  studentId,
			"homework_id": homework_id,
		})
	}
	DB.Model(&StudentHomework{}).Create(stuHomeworkInfos)
}

//todo

func (course *Course) AfterSave(db *gorm.DB) error {
	var thisTeacher Teacher
	db.Model(&Teacher{}).First(&thisTeacher, course.TeacherId)
	return thisTeacher.IncreaseCourseCnt()
}

func (course *Course) BeforeDelete(db *gorm.DB) error {
	var thisTeacher Teacher
	db.Model(&Teacher{}).First(&thisTeacher, course.TeacherId)
	return thisTeacher.DecreaseCourseCnt()
}

func (courseEval *CourseEval) AfterSave(db *gorm.DB) error {
	var thisStudent Student
	var err error
	db.Model(&Student{}).First(&thisStudent, courseEval.StudentId)
	err = thisStudent.IncreaseEvalCnt()
	if err != nil {
		return err
	}
	type Result struct {
		CourseId string
		Average  int
	}
	var result Result
	db.Model(&CourseEval{}).Select("course_id", "avg(score)").Where(map[string]interface{}{"course_id": courseEval.CourseId}).First(&result)
	var thisCourse Course
	db.Model(&Course{}).First(&thisCourse, courseEval.CourseId)
	return thisCourse.UpdateAvg(result.Average)
}

func (courseEval *CourseEval) AfterDelete(db *gorm.DB) error {
	var thisStudent Student
	var err error
	db.Model(&Student{}).First(&thisStudent, courseEval.StudentId)
	err = thisStudent.DecreaseEvalCnt()
	if err != nil {
		return err
	}
	type Result struct {
		CourseId string
		Average  int
	}
	var result Result
	db.Model(&CourseEval{}).Select("course_id", "avg(score)").Where(map[string]interface{}{"course_id": courseEval.CourseId}).First(&result)
	var thisCourse Course
	db.Model(&Course{}).First(&thisCourse, courseEval.CourseId)
	return thisCourse.UpdateAvg(result.Average)
}

func (homework *Homework) AfterSave(db *gorm.DB) error {
	if homework.IsTeacher == 1 {
		var thisCourse Course
		db.Model(&Course{}).First(&thisCourse, homework.CourseId)
		var students []StudentCourse
		db.Model(&Student{}).Where(map[string]interface{}{
			"course_id": homework.CourseId,
		}).Find(&students)
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
		db.Model(&StudentHomework{}).Where(map[string]interface{}{
			"homework_id": homework.HomeworkId,
		}).Find(&studentHomeworks)
		for _, studentHomework := range studentHomeworks {
			err := studentHomework.Delete()
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		db.Model(&StudentHomework{}).Where(map[string]interface{}{
			"homework_id": homework.HomeworkId,
			"student_id":  homework.StudentId,
		}).Find(&studentHomeworks)
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
		db.Model(&Teacher{}).First(&thisTeacher, mainComment.TeacherId)
		return thisTeacher.IncreaseDiscussCnt()
	} else {
		var thisStudent Student
		db.Model(&Student{}).First(&thisStudent, mainComment.StudentId)
		return thisStudent.IncreaseDiscussCnt()
	}
}

func (mainComment *MainComment) BeforeDelete(db *gorm.DB) error {
	if mainComment.IsTeacher == 1 {
		var thisTeacher Teacher
		db.Model(&Teacher{}).First(&thisTeacher, mainComment.TeacherId)
		return thisTeacher.DecreaseDiscussCnt()
	} else {
		var thisStudent Student
		db.Model(&Student{}).First(&thisStudent, mainComment.StudentId)
		return thisStudent.DecreaseDiscussCnt()
	}
}

func (SecondComment *SecondComment) AfterSave(db *gorm.DB) error {
	if SecondComment.IsTeacher == 1 {
		var thisTeacher Teacher
		db.Model(&Teacher{}).First(&thisTeacher, SecondComment.TeacherId)
		return thisTeacher.IncreaseDiscussCnt()
	} else {
		var thisStudent Student
		db.Model(&Student{}).First(&thisStudent, SecondComment.StudentId)
		return thisStudent.IncreaseDiscussCnt()
	}
}

func (SecondComment *SecondComment) BeforeDelete(db *gorm.DB) error {
	if SecondComment.IsTeacher == 1 {
		var thisTeacher Teacher
		db.Model(&Teacher{}).First(&thisTeacher, SecondComment.TeacherId)
		return thisTeacher.DecreaseDiscussCnt()
	} else {
		var thisStudent Student
		db.Model(&Student{}).First(&thisStudent, SecondComment.StudentId)
		return thisStudent.DecreaseDiscussCnt()
	}
}
