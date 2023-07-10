package main

import "fmt"

func main(){
	var n, rem, tem int
	var res int=0
	fmt.Println("enter the number:")
	fmt.Scan(&n)

	tem=n
	for{
		rem=tem%10
		res=(res*10)+rem
		tem=tem/10
		if tem==0{
			break
		}
	}

	if(res==n){
		fmt.Println("this is palindrome number")
	}else{
		fmt.Println("this is not palindrome number")
	}


}