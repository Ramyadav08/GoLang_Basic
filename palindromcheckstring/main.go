package main

import "fmt"

func main(){
	str:=[]string{"ram","madam","abcdcba","nurse"}
	fmt.Println(findnumber(str))
}

func findnumber(str []string) int{
	var count int=0
	for _,i:=range str{
		if palindrom(i){
			count ++
		}
		
	}
	return count

}

func palindrom(str string) bool {
	reversedStr := ""
    for i := len(str)-1; i >= 0; i-- {
        reversedStr += string(str[i])
    }
    for i := range(str) {
        if str[i] != reversedStr[i] {
            return false
        }
    }
    return true
}