package main

import "fmt"

func main(){
	arr:=[]int{1,2,3,4,5}
	reverse(arr)
	fmt.Println(arr)
}

func reverse(arr []int){
	s:=0
	e:=len(arr)-1
	for s<=e {
		swap(arr,s,e)
		s++
		e--
	}
}

func swap(arr[]int, f,s int){
	temp:= arr[f]
	arr[f]=arr[s]
	arr[s]=temp
}