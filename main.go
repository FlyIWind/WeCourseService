package main

import (
	"fmt"
	"github.com/FlyIWind/WeCourseService/lib"
)

func main() {
	conf := lib.ReadConfig()
	fmt.Println("学校名称：" + conf.School.SchoolName)
	switch conf.School.MangerType {
	case "supwisdom":
		fmt.Println("教务系统：树维教务系统")
		break
	}
	nowWeek := lib.GetWeekTime(conf.School.CalendarFirst)
	fmt.Println("当前教学周：" + nowWeek)
	lib.StartWebSocket()
}
