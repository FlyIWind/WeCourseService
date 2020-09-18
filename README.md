# WeCourseService

微课表服务端，基于[getMyCourses](https://github.com/whoisnian/getMyCourses)项目

本项目可以看作是一种第三方的树维教务系统开发SDK，可以根据自己的需求继承到项目当中 例如校园门禁、教室展示屏、电子宿舍门牌等

目前仅支持树维教务系统 后续会考虑增加其他教务系统

本项目使用了GoCache缓存库，避免了频繁访问教务系统的问题

## 使用方法

* 将项目部署到服务器，防火墙开放25565（默认）端口

* 在uni-app中修改服务端地址

* 修改config.json中的内容以适配自己的学校

  config.json内容如下：

  ```json
  {
  	"SchoolName":"山东商业职业技术学院",
  	"MangerType":"supwisdom",
  	"MangerURL":"http://szyjxgl.sict.edu.cn:9000/",
  	"CalendarFirst":"2020-02-17",
  	"SocketPort":25565
  }
  ```

  SchoolName为学校名称

  MangerType为教务系统类型，目前只支持树维教务系统（supwisdom）

  MangerURL为教务系统地址 注意不能包含eams 必须为根路径

  CalendarFirst为校历上的第一个星期一，即第一周开始的日期

  SocketPort为websocket开放的端口 请使用nginx转发WebSocketSSL以便支持微信小程序

## 请求格式

本程序通信协议为WebSocket 数据交换格式为JSON

```json
{
    "Type":"course",
    "UserName":"201808830303",
    "PassWord":"7355608",
    "Week":0
}
```

Type为必填项，表明请求类型 可选内容为：login（验证登录）、week（获取当前教学周）、teacher（获取教师列表）、account（获取学籍信息）、course（获取课程表）、photo（获取学籍照片）、grade（获取成绩） 日后还会增加更多接口例如获取考试安排、获取成绩

UserName为教务系统登陆账号，除了获取教学周外都需要提供

PassWord为教务系统登录密码，除了获取教学周外都需要提供

Week为需要获取的教学周 仅在获取课表时需要提供

## 返回格式

```json
{
    "Type":"login",
    "Data":object
}
```

由于uni-app只支持同时连接一个websocket，为方便管理 特修改了返回格式

Type和请求时传入的Type一致 Data为返回内容 为object 数据类型不唯一

 以下例子中的返回示例均代表返回结果的Data内容

### 1、验证登录

该功能用于验证账号密码是否能成功登录教务系统，传入JSON格式如下

```json
{
    "Type":"login",
    "UserName":"201808830303",
    "PassWord":"7355608"
}
```

> 返回结果为“登录成功”或“登录失败”

### 2、获取教学周

这个功能主要是为了方便同学们获取现在是第几周（学校的奇葩校历看的头大） 请求格式也相当简单

```json
{
    "Type":"week"
}
```

> 返回结果也是非常简单 直接会返回当前周的数字

### 3、获取教师列表

这个功能是为了防止小伙伴们忘记自己的老师是谁而设计的（虽然这种情况出现的概率不大）
请求格式是这样的

```json
{
    "Type":"teacher",
    "UserName":"201808830303",
    "PassWord":"7355608"
}
```

返回示例

```json
[
	{
		"CourseID": "A000032-.5.11",
		"CourseName": "就业指导实务",
		"CourseCredit": "0.5",
		"CourseTeacher": "王滢"
	},
	{
		"CourseID": "A080311-4.11",
		"CourseName": "JavaScript程序设计",
		"CourseCredit": "4",
		"CourseTeacher": "薛现伟"
	}
]
```

CourseID为教务系统内部分配的课程ID，CourseName为课程名称，CourseCredit为该课程学分，CourseTeacher为该课程任课教师

### 4、获取学籍信息

这个功能可以用来做用户身份识别，比如说展示资料卡一类的

> 免责声明：请先阅读有关公民个人信息储存及使用的法律法规以免产生法律责任，无特殊要求不要储存或缓存用户个人信息 严禁将用户个人信息用于违法犯罪活动

请求格式是这样的

```json
{
    "Type":"account",
    "UserName":"201808830303",
    "PassWord":"7355608"
}
```

返回结果示例

```json
{
	"FullName": "高峰",
	"EnglishName": "Gao Feng",
	"Sex": "男",
	"StartTime": "2018-09-01",
	"EndTime": "2021-06-30",
	"SchoolYear": "3",
	"Type": "专科(普通全日制)",
	"System": "信息与艺术学院(系)",
	"Specialty": "软件技术对口",
	"Class": "软件1803"
}
```

FullName为中文全名，EnglishName为英文名称（留学生等特殊情况），Sex为性别，StartTime为入学时间，EndTime为毕业时间，SchoolYear为学年，Type为学历类型（专科、本科、普招、单招等），System为院系，Speacialty为专业，Class为班级。

### 5、获取课程表

本程序的核心功能，用于获取本人的本学期全部课程表或本周课程表

请求格式是这样的

```json
{
    "Type":"course",
    "UserName":"201808830303",
    "PassWord":"7355608",
    "Week":0
}
```

前三个参数与之前的用途一致，这里不再废话，Week为要查询的教学周 如果为0则返回本学期课程表，为其他数值则返回对应周的课程表

返回结果也有两种 如果是返回本学期课程表（Week为0时） 那么返回格式如下：

```json
[
	{
		"CourseID": "14290(A080311-4.11)",
		"CourseName": "JavaScript程序设计",
		"RoomID": "1526",
		"RoomName": "301,计算机基础实训室(一)",
		"Weeks": "01111111111111111110000000000000000000000000000000000",
		"CourseTimes": [
			{
				"DayOfTheWeek": 3,
				"TimeOfTheDay": 4
			},
			{
				"DayOfTheWeek": 3,
				"TimeOfTheDay": 5
			}
		]
	},
	{
		"CourseID": "19827(A080910-6.09)",
		"CourseName": "HTML5混合App开发",
		"RoomID": "-1",
		"RoomName": "停课",
		"Weeks": "00000000000010000000000000000000000000000000000000000",
		"CourseTimes": [
			{
				"DayOfTheWeek": 2,
				"TimeOfTheDay": 0
			},
			{
				"DayOfTheWeek": 2,
				"TimeOfTheDay": 1
			}
		]
	}
]
```

其中CourseID为教务系统分配的课程ID，数字ID括号内的文本与教师列表一致 CourseName为课程名称，RoomID为教务系统内部的教室ID，RoomName为教室名称，Weeks为上课的周，从第0周开始计算 有课为1无课为0 CourseTimes为课程上课时间，DayOfTheWeek为周几上课（0表示周一），TimeOfTheDay表示该课程在当天的第几节（0表示第一节）

> 由于早期开发的微课表（MUI版）对代码优化没有考虑（大一的时候写的，缺少社会的毒打），所以都是直接请求本学期全部的课程，所以这个方法在新版本的程序已经弃用了，为了做兼容考虑才把这个反人类的方法保留了下来

改善后的返回结果示例（Week不为0 获取某周课程表）：

```json
[
	{
		"CourseName": "JavaScript程序设计",
		"TeacherName": "薛现伟",
		"RoomName": "301,计算机基础实训室(一)",
		"DayOfTheWeek": 3,
		"TimeOfTheDay": "5,6"
	},
	{
		"CourseName": "HTML5混合App开发",
		"TeacherName": "王永乾",
		"RoomName": "317,信息决策实训室j",
		"DayOfTheWeek": 3,
		"TimeOfTheDay": "1,2"
	},
	{
		"CourseName": "PHP动态网站开发",
		"TeacherName": "郑春光",
		"RoomName": "317,信息决策实训室j",
		"DayOfTheWeek": 3,
		"TimeOfTheDay": "3,4"
	},
	{
		"CourseName": "就业指导实务",
		"TeacherName": "王滢",
		"RoomName": "本部E206",
		"DayOfTheWeek": 2,
		"TimeOfTheDay": "3,4"
	}
]
```

可以看到新版的返回结果变得清晰明了，CourseName为课程名称，TeacherName为教师姓名，RoomName为教室名称，DayOfTheWeek依然是表示在周几上课（从0开始，0表示周一），TimeOfTheDay自动整合当天有课的节次，例如3,4则表示当天第三节、第四节有课

### 6、获取学籍照片

> 免责声明：请先阅读有关公民个人信息储存及使用、公民肖像权相关的法律法规以免产生法律责任，无特殊要求不要储存或缓存用户个人信息 严禁将用户个人信息用于违法犯罪活动

本功能可用于用户身份识别（人脸识别）、资料头像、电子学籍卡等用途

请求格式是这样的

```json
{
    "Type":"photo",
    "UserName":"201808830303",
    "PassWord":"7355608"
}
```

返回结果为base64图片

```
data:image/jpg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/...NFaXMND/9k=
```

### 7、获取成绩

本功能也算是一个比较常用的功能，毕竟关乎着自己是否挂科 为了方便同学们查询自己的成绩才有的这个功能
请求格式是这样的

```json
{
    "Type":"grade",
    "UserName":"201808830303",
    "PassWord":"7355608"
}
```
返回结果示例

```json
[
	{
		"CourseID": "A000003-4",
		"CourseName": "大学英语（一）A",
		"CourseTerm": "2018-2019 1",
		"CourseCredit": "4",
		"CourseGrade": "64",
		"GradePoint": "1.5"
	},
	{
		"CourseID": "A080011-6",
		"CourseName": "Java程序设计",
		"CourseTerm": "2018-2019 2",
		"CourseCredit": "6",
		"CourseGrade": "95",
		"GradePoint": "4.5"
	},
	{
		"CourseID": "A080310-6",
		"CourseName": "JavaWeb程序设计",
		"CourseTerm": "2019-2020 1",
		"CourseCredit": "6",
		"CourseGrade": "93",
		"GradePoint": "4.5"
	}
]
```

返回的子项全部为String类型（教务系统的返回结果十分奇葩，天然反爬（一时间写不出合适的正则表达式）），CourseID和先前几个接口一样，都是表示课程ID(不过这个和教师列表的不一样，因为这个是院系课程统一的编号，而教师列表的是每个老师的课程ID都不一样) CourseName是课程名称 CourseTerm代表学期 例如2019-2020 1就代表是2019-2020学年第一学期 CourseCredit代表学分 CourseGrade代表最终成绩 GradePoint代表绩点


## 版权协议

本项目为MIT License协议，授权各类合法程序引用（需在法律信息中标注出处）

项目已获得软件著作权（登记号：2019SR0620279）

严谨出售倒卖 违者必究

## 联系方式

如果在使用过程中出现问题，您可以给我发issues（基本不看，不建议使用该方式）

联系QQ：77257474 电子邮箱：root@mchacker.cn

请注明来意，谢谢！