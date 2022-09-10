package models

import "github.com/beego/beego/v2/adapter/orm"

type CourseData struct {
	Id          int32  `json:"id"`
	CourseName  string `json:"course_name"`
	CourseDesc  string `json:"course_desc"`
	CourseRating float64 `json:"course_rating"`
	TotalReviews int32 `json:"total_reviews"`
}


func AddCourse(c *CourseData) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(c)
}

func GetCourses() ([]*CourseData, error){
	var courses []*CourseData
	_, err :=orm.NewOrm().QueryTable("course_data").All(&courses)
	return courses, err

}

func init() {
	orm.RegisterModel(new(CourseData))
}