package main

import (
	"fmt"
	"net/http"
)

func main(){
	add(1,2,3,4,5,6,7,8,9,10)
	mydefer()
	myserver()
}

func add( values ...int){
	total:=0
	for _,val1:=range values{
		total+=val1

	}
	fmt.Println(total)

}

func mydefer(){
	val:=0
	for val<10{
	defer fmt.Println(val)
	val++
	}
}

// # defer is execute at last of its scope
// # it help to print in revers way
// # it works on the lifo 

func myserver(){
	http.HandleFunc("/",func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w,"Welcome to my server......❤️❤️❤️❤️❤️❤️❤️❤️")
	})
	http.ListenAndServe(":4000", nil)
}