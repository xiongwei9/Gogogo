package model

type User struct {
	Id         int    `json:"id" form:"id"`
	Username   string `json:"username" form:"username"`
	Password   string `json:"password" form:"password"`
	Status     int    `json:"status" form:"status"` // 0正常状态 1删除
	Createtime int64  `json:"createtime" form:"createtime"`
}
