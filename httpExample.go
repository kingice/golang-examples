package  main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)
func HttpRoot(response http.ResponseWriter,request *http.Request){
	request.ParseForm();
	fmt.Println(request.Form)
	fmt.Println("path:",request.URL.Path)
	fmt.Println("scheme:",request.URL.Scheme)
	fmt.Println(request.Form["url_long"])
	for  k,v :=range request.Form{
		fmt.Println("key:",k)
		fmt.Println("value:",strings.Join(v,":"))
	}
	fmt.Fprintf(response,"hello first http server")
}
func main(){
	http.HandleFunc("/",HttpRoot)
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil{
		log.Fatal("http start error:",err)
	}

}
