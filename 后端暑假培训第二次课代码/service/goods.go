package service

import (
	"log"
	"summerCourse/model"
)

type Goods struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Num   int    `json:"num"`
}

// 添加商品
func AddGoods(some Goods)(error) {
	// TODO
	dbsome:=model.Goods{
		Name: some.Name,
		Num: some.Num,
		Price: some.Price,
	}
	return dbsome.AddGoods()
}
//获取model的goods，返回service的goods
func SelectGoods() (goods []Goods) {
	_goods, err := model.SelectGoods()
	if err != nil {
		log.Printf("Error get goods info. Error: %s", err)
	}
	for _, v := range _goods {
		good := Goods{
			ID:    v.ID,
			Name:  v.Name,
			Price: v.Price,
			Num:   v.Num,
		}
		goods = append(goods, good)
	}
	return goods
}
