// @Title : router
// @Description ://TODO: Add Description
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
		student.PUT("/", controller.UpdateStudent) // 完善信息
		student.DELETE("/", controller.DeleteStudent)
		//student.POST("/course", controller.StudentCourseSelect)
		//student.GET("/course/:id", controller.GetStudentCourse)
		//student.DELETE("/course", controller.DeleteStudentCourse)
	}

	//course := r.Group("/course")
	//{
	//	course.POST("/", controller.CreateCourse)
	//	course.GET("/:id", controller.GetCourse)
	//	course.PUT("/", controller.UpdateCourse)
	//	course.DELETE("/", controller.DeleteCourse)
	//}

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
