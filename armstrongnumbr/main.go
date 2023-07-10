package main


import "fmt"

func main() {
   fmt.Println("Number = 153")
   
   var number, temp, remainder int
   var result int = 0
   
   number = 153
   temp = number
   
   for {
      remainder = temp % 10
      result += remainder * remainder * remainder
      temp /= 10
      if temp == 0 {
         break 
      }
   }
   
   if result == number {
      fmt.Printf("%d is an Armstrong number.", number)
   } else {
      fmt.Printf("%d is not an Armstrong number.", number)
   }
   
}