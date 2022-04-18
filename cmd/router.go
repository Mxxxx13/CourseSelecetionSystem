// @Title : router
// @Description : 路由设置
// @Author : MX
// @Update : 2022/4/14 20:58

package cmd

import (
	"CourseSeletionSystem/controller"
	"CourseSeletionSystem/middleware"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.GET("/login", controller.Login)
		user.PUT("", controller.AlterUser)
	}

	student := r.Group("/student")
	{
		student.POST("/register", controller.StudentRegister) // 注册
		student.GET("/:id", controller.GetStudent)

		student.Use(middleware.LoginRequired)
		student.POST("/", controller.CreateStudent)
		student.PUT("/:id", controller.UpdateStudent)
		student.DELETE("/:id", controller.DeleteStudent)
		student.POST("/course", controller.StudentSelectCourse)       // 学生选课
		student.DELETE("/course/:id", controller.StudentDeleteCourse) // 学生退课
		student.GET("/course", controller.StudentGetCourse)
	}

	course := r.Group("/course")
	{
		course.POST("/", controller.CreateCourse)
		course.GET("/:id", controller.GetCourse)
		course.PUT("/:id", controller.UpdateCourse)
		course.DELETE("/:id", controller.DeleteCourse)
	}

	teacher := r.Group("/teacher")
	{
		teacher.POST("/register", controller.TeacherRegister)
		teacher.POST("/", controller.CreateTeacher)
		teacher.GET("/:id", controller.GetTeacher)
		teacher.PUT("/:id", controller.UpdateTeacher)
		teacher.DELETE("/:id", controller.DeleteTeacher)
	}

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
