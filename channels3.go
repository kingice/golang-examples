package main

import (
	"fmt"
	"time"
	"strconv"
	"strings"
)

func makeCakeAndSend(cs chan string, count int) {
	for i := 1; i <= count; i++ {
		cakeName := "Strawberry Cake " + strconv.Itoa(i)
		cs <- cakeName // 传递一个 cake
	}
}

func receiveCakeAndPack(cs chan string) {
	for s := range cs {
		fmt.Println("Packing received cake: ", s)
	}
}
func receiveMoney(money chan int) {
	for temp := range money {
		println("receive money:", temp)
	}
}
func emitMoney(money chan int, time int) {
	for i := 0; i < time; i++ {
		println("emmit money:", i)
		money<-i
	}
}

func main() {
	testChan := make(chan int)
	go emitMoney(testChan, 5)
	go receiveMoney(testChan)
	cs := make(chan string)
	go makeCakeAndSend(cs, 5)
	go receiveCakeAndPack(cs)

	du, _ := time.ParseDuration("3s")
	result := strings.Split("a  b c", " ")//严格用分隔符分隔
	result=strings.Fields("a   b c")//用空格分隔，
	for a:= range result{
		println("one:",result[a])

	}

		time.Sleep(du)
}
