package main

import "fmt"

func main(){
	var i int =1
	var res,rem,num int
	fmt.Println("enter the number")
	fmt.Scan(&num)
	for i<num{
		n:=i
		res=0
		for {
			rem=n%10
			res=res+(rem*rem*rem)
			n=n/10
			if n==0{
				break
			}


		}
		if(res==i){
			fmt.Println(i)
		}
		i++
	}
	

}