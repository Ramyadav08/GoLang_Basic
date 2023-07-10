package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	// var age int
	// fmt.Println("enter the value here")
	// fmt.Scan(&age)
	// fmt.Println(age)
	var name string
	fmt.Println("enter the name here")
	reader:=bufio.NewReader(os.Stdin)
	name ,_= reader.ReadString('\n')
	fmt.Println(name)
	

	
}