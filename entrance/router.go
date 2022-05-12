// @Title : router
// @Description : 路由设置
// @Author : MX
// @Update : 2022/4/14 20:58

package entrance

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
		user.GET("/send-email/:id", controller.SendEmail)
		user.PUT("/find-password/:id", controller.AlterPassword)

		user.Use(middleware.LoginRequired)
		user.PUT("/bind-email/:id", controller.BindEmail)
	}

	student := r.Group("/student")
	{
		student.POST("/register", controller.StudentRegister) // 注册
		student.GET("/:id", controller.GetStudent)

		student.Use(middleware.LoginRequired, middleware.AdminAndStudentCheck)
		student.POST("/", controller.CreateStudent)
		student.PUT("/:id", controller.UpdateStudent)
		student.DELETE("/:id", controller.DeleteStudent)
		student.POST("/course", controller.StudentSelectCourse)   // 学生选课
		student.DELETE("/course", controller.StudentDeleteCourse) // 学生退课
		student.GET("/course", controller.StudentGetCourse)       // 查看自己的选课
	}

	course := r.Group("/course")
	{
		course.GET("/:id", controller.GetCourse)

		course.Use(middleware.LoginRequired, middleware.AdminCheck)
		course.POST("/", controller.CreateCourse)
		course.PUT("/:id", controller.UpdateCourse)
		course.DELETE("/:id", controller.DeleteCourse)
	}

	teacher := r.Group("/teacher")
	{
		teacher.POST("/register", controller.TeacherRegister)
		teacher.GET("/:id", controller.GetTeacher)

		teacher.Use(middleware.LoginRequired, middleware.AdminAndTeacherCheck)
		teacher.POST("/", controller.CreateTeacher)
		teacher.PUT("/:id", controller.UpdateTeacher)
		teacher.DELETE("/:id", controller.DeleteTeacher)
		teacher.GET("/selection", controller.GetStudentSelection)
	}

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
