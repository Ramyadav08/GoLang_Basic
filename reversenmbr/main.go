package main


import "fmt"

func main(){
	var n,rem int
	var res int =0

	fmt.Println("enter the number you want to print")
	fmt.Scan(&n)

	for{
		rem=n%10
		res=(res*10)+rem
		n=n/10
		if n==0{
			break
		}

	}

	
	fmt.Println(`the reverse number of is:=>`,res)
}