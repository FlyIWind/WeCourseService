# WeCourseService

微课表服务端，基于[getMyCourses](https://github.com/whoisnian/getMyCourses)项目

本项目可以看作是一种第三方的树维教务系统开发SDK，可以根据自己的需求继承到项目当中 例如校园门禁、教室展示屏、电子宿舍门牌等

目前仅支持树维教务系统 后续会考虑增加其他教务系统

## 使用方法

* 将项目部署到服务器，防火墙开放25565端口
* 在uni-app中修改服务端地址
* 修改config.json中的内容以适配自己的学校

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

Type为必填项，表明请求类型 可选内容为：login（验证登录）、week（获取当前教学周）、teacher（获取教师列表）、account（获取学籍信息）、course（获取课程表）、photo（获取学籍照片） 日后还会增加更多接口例如获取考试安排、获取成绩

UserName为教务系统登陆账号，除了获取教学周外都需要提供

PassWord为教务系统登录密码，除了获取教学周外都需要提供

Week为需要获取的教学周 仅在获取课表时需要提供

### 1、验证登录

该功能用于验证账号密码是否能成功登录教务系统，传入JSON格式如下

```json
{
    "Type":"login",
    "UserName":"201808830303",
    "PassWord":"7355608"
}
```

> 登录成功后会返回四个大字：登录成功 只要没有返回登录成功都算是失败（懒）

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
		"CourseID": "A005217-2.06",
		"CourseName": "社会心理学",
		"CourseCredit": "2",
		"CourseTeacher": "刘悦"
	},
	{
		"CourseID": "A080311-4.11",
		"CourseName": "JavaScript程序设计",
		"CourseCredit": "4",
		"CourseTeacher": "薛现伟"
	},
	{
		"CourseID": "A080910-6.09",
		"CourseName": "HTML5混合App开发",
		"CourseCredit": "6",
		"CourseTeacher": "王永乾"
	},
	{
		"CourseID": "A080913-6.09",
		"CourseName": "PHP动态网站开发",
		"CourseCredit": "6",
		"CourseTeacher": "郑春光"
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
		"CourseID": "14290(A080311-4.11)",
		"CourseName": "JavaScript程序设计",
		"RoomID": "1556",
		"RoomName": "311,影视多媒体实训室",
		"Weeks": "01111111111111111110000000000000000000000000000000000",
		"CourseTimes": [
			{
				"DayOfTheWeek": 1,
				"TimeOfTheDay": 2
			},
			{
				"DayOfTheWeek": 1,
				"TimeOfTheDay": 3
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
	},
	{
		"CourseID": "19827(A080910-6.09)",
		"CourseName": "HTML5混合App开发",
		"RoomID": "-1",
		"RoomName": "停课",
		"Weeks": "00000000000000000110000000000000000000000000000000000",
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
	},
	{
		"CourseID": "19827(A080910-6.09)",
		"CourseName": "HTML5混合App开发",
		"RoomID": "1558",
		"RoomName": "317,信息决策实训室j",
		"Weeks": "01111111111111111110000000000000000000000000000000000",
		"CourseTimes": [
			{
				"DayOfTheWeek": 3,
				"TimeOfTheDay": 0
			},
			{
				"DayOfTheWeek": 3,
				"TimeOfTheDay": 1
			}
		]
	},
	{
		"CourseID": "19827(A080910-6.09)",
		"CourseName": "HTML5混合App开发",
		"RoomID": "1658",
		"RoomName": "303小,计算机基础实训室(二)",
		"Weeks": "01111111111101111000000000000000000000000000000000000",
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
	},
	{
		"CourseID": "19827(A080910-6.09)",
		"CourseName": "HTML5混合App开发",
		"RoomID": "1729",
		"RoomName": "304小,计算机基础实训室(三)",
		"Weeks": "01111111111111111110000000000000000000000000000000000",
		"CourseTimes": [
			{
				"DayOfTheWeek": 0,
				"TimeOfTheDay": 0
			},
			{
				"DayOfTheWeek": 0,
				"TimeOfTheDay": 1
			}
		]
	},
	{
		"CourseID": "19803(A080913-6.09)",
		"CourseName": "PHP动态网站开发",
		"RoomID": "1554",
		"RoomName": "308,软件开发实训室",
		"Weeks": "00000000000000010000000000000000000000000000000000000",
		"CourseTimes": [
			{
				"DayOfTheWeek": 1,
				"TimeOfTheDay": 4
			},
			{
				"DayOfTheWeek": 1,
				"TimeOfTheDay": 5
			}
		]
	},
	{
		"CourseID": "19803(A080913-6.09)",
		"CourseName": "PHP动态网站开发",
		"RoomID": "1553",
		"RoomName": "304大,计算机基础实训室(三)",
		"Weeks": "00000000000000010000000000000000000000000000000000000",
		"CourseTimes": [
			{
				"DayOfTheWeek": 4,
				"TimeOfTheDay": 0
			},
			{
				"DayOfTheWeek": 4,
				"TimeOfTheDay": 1
			}
		]
	},
	{
		"CourseID": "19803(A080913-6.09)",
		"CourseName": "PHP动态网站开发",
		"RoomID": "1457",
		"RoomName": "本部E202",
		"Weeks": "00000000000000001000000000000000000000000000000000000",
		"CourseTimes": [
			{
				"DayOfTheWeek": 2,
				"TimeOfTheDay": 8
			},
			{
				"DayOfTheWeek": 2,
				"TimeOfTheDay": 9
			}
		]
	},
	{
		"CourseID": "19803(A080913-6.09)",
		"CourseName": "PHP动态网站开发",
		"RoomID": "1729",
		"RoomName": "304小,计算机基础实训室(三)",
		"Weeks": "01111111111011111110000000000000000000000000000000000",
		"CourseTimes": [
			{
				"DayOfTheWeek": 4,
				"TimeOfTheDay": 4
			},
			{
				"DayOfTheWeek": 4,
				"TimeOfTheDay": 5
			}
		]
	},
	{
		"CourseID": "19803(A080913-6.09)",
		"CourseName": "PHP动态网站开发",
		"RoomID": "1558",
		"RoomName": "317,信息决策实训室j",
		"Weeks": "01111111111111011110000000000000000000000000000000000",
		"CourseTimes": [
			{
				"DayOfTheWeek": 3,
				"TimeOfTheDay": 2
			},
			{
				"DayOfTheWeek": 3,
				"TimeOfTheDay": 3
			}
		]
	},
	{
		"CourseID": "19803(A080913-6.09)",
		"CourseName": "PHP动态网站开发",
		"RoomID": "1613",
		"RoomName": "309,计算机基础实训室",
		"Weeks": "01111111111101111110000000000000000000000000000000000",
		"CourseTimes": [
			{
				"DayOfTheWeek": 2,
				"TimeOfTheDay": 4
			},
			{
				"DayOfTheWeek": 2,
				"TimeOfTheDay": 5
			}
		]
	},
	{
		"CourseID": "8892(A000032-.5.11)",
		"CourseName": "就业指导实务",
		"RoomID": "1487",
		"RoomName": "本部E206",
		"Weeks": "01111000000000000000000000000000000000000000000000000",
		"CourseTimes": [
			{
				"DayOfTheWeek": 2,
				"TimeOfTheDay": 2
			},
			{
				"DayOfTheWeek": 2,
				"TimeOfTheDay": 3
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
		"CourseName": "JavaScript程序设计",
		"TeacherName": "薛现伟",
		"RoomName": "311,影视多媒体实训室",
		"DayOfTheWeek": 1,
		"TimeOfTheDay": "3,4"
	},
	{
		"CourseName": "HTML5混合App开发",
		"TeacherName": "王永乾",
		"RoomName": "317,信息决策实训室j",
		"DayOfTheWeek": 3,
		"TimeOfTheDay": "1,2"
	},
	{
		"CourseName": "HTML5混合App开发",
		"TeacherName": "王永乾",
		"RoomName": "303小,计算机基础实训室(二)",
		"DayOfTheWeek": 2,
		"TimeOfTheDay": "1,2"
	},
	{
		"CourseName": "HTML5混合App开发",
		"TeacherName": "王永乾",
		"RoomName": "304小,计算机基础实训室(三)",
		"DayOfTheWeek": 0,
		"TimeOfTheDay": "1,2"
	},
	{
		"CourseName": "PHP动态网站开发",
		"TeacherName": "郑春光",
		"RoomName": "304小,计算机基础实训室(三)",
		"DayOfTheWeek": 4,
		"TimeOfTheDay": "5,6"
	},
	{
		"CourseName": "PHP动态网站开发",
		"TeacherName": "郑春光",
		"RoomName": "317,信息决策实训室j",
		"DayOfTheWeek": 3,
		"TimeOfTheDay": "3,4"
	},
	{
		"CourseName": "PHP动态网站开发",
		"TeacherName": "郑春光",
		"RoomName": "309,计算机基础实训室",
		"DayOfTheWeek": 2,
		"TimeOfTheDay": "5,6"
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

## 版权协议

本项目为MIT License协议，授权各类合法程序引用（需在法律信息中标注出处）

项目已获得软件著作权（登记号：2019SR0620279）

严谨出售倒卖 违者必究

## 联系方式

如果在使用过程中出现问题，您可以给我发issues（基本不看，不建议使用该方式）

联系QQ：77257474 电子邮箱：root@mchacker.cn

请注明来意，谢谢！