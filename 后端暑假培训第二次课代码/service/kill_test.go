package service

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGoods(t *testing.T) {
	ti:=time.NewTimer(time.Second*1)
	for  {
		<-ti.C
		fmt.Println("hhh")
		ti.Reset(time.Second*1)
	}
}
