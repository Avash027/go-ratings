package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
)

type UserRatings struct {
	UserId   int32 `json:"uid"`
	CourseId int32 `json:"cid"`
	Rating int32 `json:"rating"`
}


func UpdateRating(newUserRating *UserRatings) (int64, error) {
	o := orm.NewOrm()

	var userRating UserRatings
	userRating.UserId = newUserRating.UserId
	userRating.CourseId = newUserRating.CourseId

	var courseData CourseData
	courseData.Id = newUserRating.CourseId
	err := o.Read(&courseData)

	if err != nil {
		return 0, err
	}
	err = o.Read(&userRating, "UserId", "CourseId")

	//It means there are no rows, so we will add a new row
	if err != nil {
		courseData.TotalRatingNum++; // Since we are adding a new rating
		courseData.CourseRating += newUserRating.Rating
		courseData.AverageRating = float64(courseData.CourseRating) / float64(courseData.TotalRatingNum)
	
		err := o.Begin()

		if err != nil {
			logs.Critical("Error in begining transaction")
			return 0, err
		}

		_, err = o.Update(&courseData)
		
		if err != nil {
			o.Rollback()
			logs.Critical("Error in updating course data")
			return 0, err
		}

		// Adding new entry to user_ratings table
		_, err = o.Raw("INSERT INTO user_ratings (user_id, course_id, rating) VALUES (?, ?, ?)", newUserRating.UserId, newUserRating.CourseId, newUserRating.Rating).Exec()

		if err != nil {
			o.Rollback()
			logs.Critical("Error in inserting new rating")
			return 0, err
		}
		o.Commit()

		return 1, nil
	} else {
		// We remove the previous rating
		courseData.CourseRating -= userRating.Rating
		//We add the new rating
		courseData.CourseRating += newUserRating.Rating
		courseData.AverageRating = float64(courseData.CourseRating) / float64(courseData.TotalRatingNum)

		err := o.Begin()

		if err!= nil {
			logs.Critical("Failed to start transaction")
			return 0,err
		}

		_, err = o.Update(&courseData)

		if err != nil {
			o.Rollback()
			return 0, err
		}
		

		userRating.Rating = newUserRating.Rating
		
		_, err = o.Raw("UPDATE user_ratings SET rating = ? WHERE user_id = ? AND course_id = ?", userRating.Rating, userRating.UserId, userRating.CourseId).Exec()

		if err != nil {
			o.Rollback()
			return 0, err
		}

		o.Commit()

		return 1,nil
		
	}
	


}

func init() {
	orm.RegisterModel(new(UserRatings))
}
