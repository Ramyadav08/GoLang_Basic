package main

import "fmt"
func main(){
	var n int
	flag:=0 
	fmt.Println("enter the number:")
	fmt.Scan(&n)
	for i:=2; i<=n/2;i++{
		if n%i==0{
			flag=1
			fmt.Println("this is not prime number")
			break
		}
		

	}

	if(flag==0){
		fmt.Println("this is prime number")
	}



}