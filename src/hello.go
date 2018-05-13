package main

import (
  "fmt"
  "strconv"  
)


func main() {
  // name := "Abrar"
  // fmt.Printf("%s",name)
  const pi float64 = 3.1415
  // pi = 900.12
  var name = "Abrar"
  text := " Is the legend"
  num := 1
  
  a, b, c, d := 1, 2, 3, "abrar"
  
  fmt.Printf("%T\n",pi)
  fmt.Println(name + " is a Go programmer!" + text + strconv.Itoa(num))
  
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)

}