package main

import "fmt"

func main() {
  var names string
  var ages int
  fmt.Println("ingrese nombre y luego edad")
  fmt.Scanf("%s\n",&names)
  fmt.Scanf("%d\n",&ages)
  My := MyStruct{name:names,age:ages}
  fmt.Println("tu nombre es: ",My.name)
  fmt.Println("tienes : ",My.age)
}
type MyStruct struct{
  name string
  age int
}
