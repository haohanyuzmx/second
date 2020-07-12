package model

import "github.com/jinzhu/gorm"

type Goods struct {
	gorm.Model
	Name  string
	Price int
	Num   int
}


// 添加商品
func (goods *Goods)AddGoods() error{
	return DB.Create(goods).Error
}
//更新商品数据
func (goods *Goods)UpdateGoods() error{
	return DB.Table("goods").Where("name=?",goods.Name).Update("num",goods.Num).Error
}

// 查看商品
func SelectGoodsById(id uint) (goods Goods, err error){
	err = DB.Table("goods").Where("id = ?",id).First(&goods).Error
	if err != nil {
		return Goods{}, err
	}
	return goods, nil
}

// 查看所有的商品
func SelectGoods() (goods []Goods, err error){
	err = DB.Table("goods").Find(&goods).Error
	if err != nil {
		return nil, err
	}
	return goods, nil
}

