package lib

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type WeekCourse struct {
	CourseName   string
	TeacherName  string
	RoomName     string
	DayOfTheWeek int
	TimeOfTheDay string
}

var myWeekCourse []WeekCourse
var tmpCourse WeekCourse
var myCourseResult WeekCourseResult

func GetWeekCourse(UserName string, PassWord string, WeekDay int) string {
	myWeekCourse = nil
	myCourseResult.Type = "course"
	var tmpResult CourseResult
	var allCourses []Course
	allCourses = nil
	var cstr string = GetCourse(UserName, PassWord)
	err := json.Unmarshal([]byte(cstr), &tmpResult)
	if err != nil {
		fmt.Println(err)
	}
	allCourses = tmpResult.Data
	for _, c1 := range allCourses {
		schoolWeek := strconv.Itoa(WeekDay)
		intWeek, _ := strconv.Atoi(schoolWeek)
		if c1.Weeks[intWeek] == '1' {
			arr := GetTeacherObj()
			for _, thisteacher := range arr {
				if strings.Contains(c1.CourseID, thisteacher.CourseID) {
					tmpCourse.TeacherName = thisteacher.CourseTeacher
					tmpCourse.CourseName = thisteacher.CourseName
					tmpCourse.RoomName = c1.RoomName
					tmpCourse.DayOfTheWeek = c1.CourseTimes[0].DayOfTheWeek
					tmpCourse.TimeOfTheDay = strconv.Itoa(c1.CourseTimes[0].TimeOfTheDay+1) + "," + strconv.Itoa(c1.CourseTimes[1].TimeOfTheDay+1)
					myWeekCourse = append(myWeekCourse, tmpCourse)
				}
			}
		}
	}
	myCourseResult.Data = myWeekCourse
	js, _ := json.MarshalIndent(myCourseResult, "", "\t")
	return B2S(js)
}
