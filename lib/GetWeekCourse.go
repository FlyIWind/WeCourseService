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

func GetWeekCourse(UserName string, PassWord string, WeekDay int) string {
	myWeekCourse = nil
	var allCourses []Course
	allCourses = nil
	var cstr string = GetCourse(UserName, PassWord)
	err := json.Unmarshal([]byte(cstr), &allCourses)
	if err != nil {
		fmt.Println(err)
	}
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
	js, _ := json.MarshalIndent(myWeekCourse, "", "\t")
	return B2S(js)
}
