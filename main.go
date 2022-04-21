// @Title : main
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2022/4/14 20:58

package main

import (
	"CourseSeletionSystem/cmd"
	"CourseSeletionSystem/dao"
	"CourseSeletionSystem/util"
)

func main() {
	dao.MysqlInit()
	dao.RedisInit()
	util.NewEmailPool()
	cmd.Router()
}
