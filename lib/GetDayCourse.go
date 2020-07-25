package lib

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type DayCourse struct {
	CourseName   string
	TeacherName  string
	TimeOfTheDay string
	SchoolWeek   string
}

var ToDayCourse []DayCourse
var thisCourse DayCourse

func GetDayCourse(UserName, PassWord string) string {
	var allCourses []Course
	conf := ReadConfig()
	var cstr string = GetCourse(UserName, PassWord)
	err := json.Unmarshal([]byte(cstr), &allCourses)
	if err != nil {
		fmt.Println(err)
	}
	for _, c1 := range allCourses {
		schoolWeek := GetWeekTime(conf.School.CalendarFirst)
		serverWeek := GetWeekDay()
		intWeek, _ := strconv.Atoi(schoolWeek)
		if c1.Weeks[intWeek] == '1' {
			if c1.CourseTimes[0].DayOfTheWeek == serverWeek {
				arr := GetTeacherObj()
				for _, thisteacher := range arr {
					if strings.Contains(c1.CourseID, thisteacher.CourseID) {
						thisCourse.TeacherName = thisteacher.CourseTeacher
						thisCourse.CourseName = thisteacher.CourseName
						thisCourse.TimeOfTheDay = strconv.Itoa(c1.CourseTimes[0].TimeOfTheDay+1) + "," + strconv.Itoa(c1.CourseTimes[1].TimeOfTheDay+1)
						thisCourse.SchoolWeek = schoolWeek
						ToDayCourse = append(ToDayCourse, thisCourse)
					}
				}
			}
		}
	}
	js, _ := json.MarshalIndent(ToDayCourse, "", "\t")
	return B2S(js)
}

var WeekDayMap = map[string]int{
	"Monday":    0,
	"Tuesday":   1,
	"Wednesday": 2,
	"Thursday":  3,
	"Friday":    4,
	"Saturday":  5,
	"Sunday":    6,
}

func GetWeekDay() int {
	wd := time.Now().Weekday().String()
	return WeekDayMap[wd]
}
