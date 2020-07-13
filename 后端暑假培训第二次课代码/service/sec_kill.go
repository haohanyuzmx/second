package service

import (
	"log"
	"summerCourse/model"
	"sync"
	"time"
)

type User struct {
	UserId  string
	GoodsId uint
}

var OrderChan = make(chan User, 1024)

var ItemMap = make(map[uint]*Item)

type Item struct {
	ID        uint   // 商品id
	Name      string // 名字
	Total     int    // 商品总量
	Left      int    // 商品剩余数量
	IsSoldOut bool   // 是否售罄
	leftCh    chan int
	sellCh    chan int
	done      chan struct{}
	Lock      sync.Mutex
}

// TODO 写一个定时任务，每天定时从数据库加载数据到Map
func InitMap() {
	t := time.NewTimer(time.Hour * 24)
	for {
		<-t.C
		for _, i2 := range ItemMap {
			some := model.Goods{
				Name: i2.Name,
				Num:  i2.Left,
			}
			if err:=some.UpdateGoods();err!=nil{
				log.Println(err)
				return
			}
		}
		some := SelectGoods()
		for _, i2 := range some {
			item := &Item{
				ID:        i2.ID,
				Name:      i2.Name,
				Total:     i2.Num,
				Left:      i2.Num,
				IsSoldOut: false,
				leftCh:    make(chan int),
				sellCh:    make(chan int),
			}
			ItemMap[item.ID] = item
		}
		t.Reset(time.Hour * 24)
	}
}

func getItem(itemId uint) *Item {
	return ItemMap[itemId]
}

func order() {
	user := <-OrderChan
	item := getItem(user.GoodsId)
	item.SecKilling(user.UserId)
}

func (item *Item) SecKilling(userId string) {

	item.Lock.Lock()
	defer item.Lock.Unlock()
	// 等价
	// var lock = make(chan struct{}, 1}
	// lock <- struct{}{}
	// defer func() {
	// 		<- lock
	// }
	if item.IsSoldOut {
		return
	}
	item.BuyGoods(1)

	MakeOrder(userId, item.ID, 1)

}

// 定时下架
func (item *Item) OffShelve() {
	beginTime := time.Now()
	// 获取第二天时间
	//nextTime := beginTime.Add(time.Hour * 24)
	// 计算次日零点，即商品下架的时间
	//offShelveTime := time.Date(nextTime.Year(), nextTime.Month(), nextTime.Day(), 0, 0, 0, 0, nextTime.Location())
	offShelveTime := beginTime.Add(time.Second * 5)
	timer := time.NewTimer(offShelveTime.Sub(beginTime))

	<-timer.C
	delete(ItemMap, item.ID)
	close(item.done)

}

// 出售商品
func (item *Item) SalesGoods() {
	for {
		select {
		case num := <-item.sellCh:
			if item.Left -= num; item.Left <= 0 {
				item.IsSoldOut = true
			}

		case item.leftCh <- item.Left:
		case <-item.Done():
			log.Println("我自闭了")
			return
		}
	}
}
//当关闭通道后可以读出信息满足case
func (item *Item) Done() <-chan struct{} {
	if item.done == nil {
		item.done = make(chan struct{})
	}
	d := item.done
	return d
}
//监控请求
func (item *Item) Monitor() {
	go item.SalesGoods()
}

// 获取剩余库存
func (item *Item) GetLeft() int {
	var left int
	left = <-item.leftCh
	return left
}

// 购买商品
func (item *Item) BuyGoods(num int) {
	item.sellCh <- num
}
//开启服务
func InitService() {
	go InitMap()
	for _, item := range ItemMap {
		item.Monitor()
		go item.OffShelve()
	}
	for i := 0; i < 10; i++ {
		go order()
	}
}
