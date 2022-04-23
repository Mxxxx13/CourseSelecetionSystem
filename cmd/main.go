// @Title : main
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2022/4/14 20:58

package main

import (
	"CourseSeletionSystem/dao"
	"CourseSeletionSystem/entrance"
	"CourseSeletionSystem/util"
)

func main() {
	dao.MysqlInit()
	dao.RedisInit()
	util.NewEmailPool()
	entrance.Router()
}
