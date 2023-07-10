package main

import "fmt"

func main(){
	var n1 int
	var n2 int
	
	fmt.Println("enter the lower number")
	fmt.Scan(&n1)
	fmt.Println("enter the second nubmber")
	fmt.Scan(&n2)
	fmt.Println("this prime interval between %v and %v is :",n1,n2)

	for i:=n1;i<=n2;i++{
		flag:=0
		if n1>=2{
			for j:=2;j<=i/2;j++{
				if i%j==0{
					flag=1;
				
					break
				}
			}

			if(flag==0){
				fmt.Println(i)
			}
			
		}

	
	}


}