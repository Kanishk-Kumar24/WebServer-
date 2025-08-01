package main

import (
	"fmt"
	"log"
	"net/http"
)
func formHandler(w http.ResponseWriter,r *http.Request){
	if r.Method!="POST"{
		http.Error(w,"wrong method called",http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err!=nil {
		fmt.Fprintf(w,"ParseForm() err: is %v", err)
		return
	}
	fmt.Fprintf(w,"POST req successful")
	name:=r.FormValue("name")
	address:=r.FormValue("address")
	fmt.Fprintf(w,"name is %s\n",name)
	fmt.Fprintf(w,"address is %s \n",address)
}
func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path!="/hello"{
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method!="GET"{
		http.Error(w,"method not supported",http.StatusNotFound)
		return
	}
	fmt.Fprintf(w,"hello world")
}
func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)
	http.Handle("/",fileServer)

	fmt.Printf("starting server at 8080\n")
	if err := http.ListenAndServe(":8080",nil); err!=nil{
		log.Fatal(err)
	}
}