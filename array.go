package main 

import "fmt"

func main() {
arr  := [3] int{4, 5, 6 }
  
   //fmt.Println(len(arr))

  // to loop through

sum := 0 
  
  for  i := 0; i < len(arr); i++1 {
sum += arr[i]
    
  }  
  fmt.Println(sum)
  
  // multi array that is array inside array
    arr2D := [2][3]int {{1, 2, 3},{3, 3, 4}}
fmt.Println(arr2D)
  
 

}
