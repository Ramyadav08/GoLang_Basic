package main

import "fmt"

func main(){
	var n1 int=0
	var n2 int=1
	var n3, i, num int

	fmt.Println("enter the number :")
	fmt.Scan(&num)

	fmt.Print(n1," ",n2," ")
	for i=2;i<num;i++{
		n3=n1+n2
		fmt.Print(n3," ")
		n1=n2
		n2=n3
	}

}