package main

import (
	"net"
	"time"
	"fmt"
)

func CheckErr(errors error){
	if  errors!=nil {
		fmt.Println("erro:"+errors.Error())
		panic(errors)
	}
}

func main() {
	tcpAdress, err := net.ResolveTCPAddr("tcp4", ":7777")
	CheckErr(err)
	listen, err := net.ListenTCP("tcp", tcpAdress)
	CheckErr(err)
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		dayTime := time.Now().String()+"\n"
		fmt.Println("get conn, the return:"+dayTime)
		conn.Write([]byte(dayTime))
		conn.Close()
	}
}
