package main 

import "fmt"

func main() {

  var x [5]int = [5]int{1, 2, 3, 4, 5}
  var s [] = x[:]

  fmt.Println(s)  
  fmt.Println(len(s))
    fmt.Println(cap(s))
}

