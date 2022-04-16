// @Title : main
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2022/4/14 20:58

package main

import (
	"CourseSeletionSystem/cmd"
	"CourseSeletionSystem/dao"
)

func main() {
	dao.ConnDB()
	cmd.Router()
}
