package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"summerCourse/service"
)

type goods struct {
	name string `json:"name"`
	num int `json:"num"`
	price int `json:"price"`
}
//获取所有物品返回json
func SelectGoods(ctx *gin.Context) {
	goods := service.SelectGoods()
	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info": "success",
		"data": struct {
			Goods []service.Goods `json:"goods"`
		}{goods},
	})
}
//添加物品到数据库
func AddGoods(ctx *gin.Context)  {
	var some goods
	if err:=ctx.ShouldBindJSON(&some);err!=nil{
		ctx.JSON(http.StatusOK,gin.H{
			"code":001,
			"info":"读取物品信息失败",
			"date":some,
		})
		log.Print(err,some)
		return
	}
	ssome:=service.Goods{
		Name:  some.name,
		Price: some.price,
		Num:   some.num,
	}
	if err:=service.AddGoods(ssome);err!=nil {
		ctx.JSON(http.StatusOK,gin.H{
			"code":002,
			"info":"添加物品失败",
			"date":some,
		})
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"info":"添加完成",
		"date":some,
	})
}

