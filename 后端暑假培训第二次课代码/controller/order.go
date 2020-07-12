package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"summerCourse/service"
)

//异步处理订单
func MakeOrder(ctx *gin.Context) {
	userId := ctx.PostForm("userId")
	goodsId := ctx.PostForm("goodsId")
	itemId,_ := strconv.Atoi(goodsId)
    service.OrderChan <- service.User{
		UserId:  userId,
		GoodsId: uint(itemId),
	}
    ctx.JSON(200, gin.H{
    	"status": 200,
    	"info": "success",
	})
}

 

