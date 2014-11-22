package main

import (
	"net"
	"io/ioutil"
	"fmt"
)

const (
	URL = "localhost:12981"
)

type Args struct {
	A, B int
}
func CheckErr(errors error){
	if  errors!=nil {
		fmt.Println("erro:"+errors.Error())
		panic(errors)
	}
}

func main() {

	//	client, err := jsonrpc.Dial("tcp", URL)
	//	defer client.Close()
	//
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//	args :=Args{7, 2}
	//	var reply int
	//	err = client.Call("Arith.Multiply", &args, &reply)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(reply)
	//	addr:=net.ParseIP("127.0.0.1")
	//	fmt.Println("the address:"+addr.String())
//	args := os.Args
//	_=args[1]
	service:="www.qq.com:80"
//	service="127.0.0.1:80"
	tcpAddr,err:=net.ResolveTCPAddr("tcp4",service)
	CheckErr(err)
	conn,err:=net.DialTCP("tcp",nil,tcpAddr)
	CheckErr(err)
	_,err=conn.Write([]byte("HEAD / HTTP 1.2\r\n\r\n"))
	CheckErr(err)
	result,err:=ioutil.ReadAll(conn)
	CheckErr(err)
	fmt.Println(string(result))
}
