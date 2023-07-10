package main 

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	var input_srt string
	
	fmt.Println("enter the string")
	reader:= bufio.NewReader(os.Stdin)
	input_srt , _=reader.ReadString('\n')
	palindrom(input_srt)
}

func palindrom(str string){
	reversedStr:=""
	for i := len(str)-1; i >= 0; i-- {
        reversedStr += string(str[i])
    }
   
	//fmt.Println(reversedStr)
	if reversedStr == str {
		fmt.Println("this is palindrom string")
	} else {
		fmt.Println("this is not palindrome string")
	}

}