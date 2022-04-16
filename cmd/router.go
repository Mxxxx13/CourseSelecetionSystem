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

		student.Use(middleware.LoginRequired)
		student.POST("/info", controller.CreateStudentInfo)
		student.PUT("/info", controller.UpdateStudentInfo) // 完善信息
		//student.GET("/:id", controller.GetStudent)
		//student.POST("/course", controller.StudentCourseSelect)
		//student.GET("/course/:id", controller.GetStudentCourse)
		//student.DELETE("/course", controller.DeleteStudentCourse)
	}
	r.Run(":8080")
}
