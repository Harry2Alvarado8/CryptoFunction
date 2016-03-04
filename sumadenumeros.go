package main

import(
  "fmt"
)

func main() {
  var a,b,c int
  fmt.Println("Ingrese el primer numero")
  fmt.Scanf("%d",&a)
  fmt.Println("Ingrese el segundo numero")
  fmt.Scanf("%d",&b)
  c = a + b
  fmt.Println("la suma es ", c)
}
