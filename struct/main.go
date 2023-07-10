package main
import "fmt"
type User struct {
	name string
	age int
	email string
  
}

func main(){
	Ram:=User{"ramrekha",23,"ryadav370@rku.ac.im"}
	fmt.Println("The details of Ram Is:",Ram)

}