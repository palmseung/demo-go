package main

import "fmt"
import "mypackage"
import "secondpackage"
//Test for renaming an imported function

func main() {

var n int = 5                               
	
fmt.Println("value of myfunction ",mypackage.MyFunction(n))     

if secondpackage.IsEven(n) {

  fmt.Println("n is even")
 
}
   

}
