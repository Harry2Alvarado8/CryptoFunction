package main

import(
  "fmt"
  "strconv" // libreria para confertir datos de int a str or string a int
)
func main() {
    var a,b,exp,age int
    var name string
    fmt.Println("\t1 : Suma\n\t2 : Tabla de multiplicacion\n\t3 : Hola")
    fmt.Printf("\t\tQue opción: ")
    fmt.Scanf("%d\n",&exp)
    switch exp {
    case 1: fmt.Println("\t\tIngrese dos numeros")
          fmt.Printf("\n\tPrimer numero: ")
          fmt.Scanf("%d\n",&a)
          fmt.Printf("\n\tSegundo numero: ")
          fmt.Scanf("%d\n",&b)
          fmt.Println("\n\t\tla suma es ",suma(a,b))
    case 2: tabla()
    case 3: fmt.Printf("\n\tDime tu nombre : ")
            fmt.Scanf("%s\n",&name)
            fmt.Printf("\n\tQue edad tienes : ")
            fmt.Scanf("%d\n",&age)
            Tellamas(name,age)
    default: fmt.Println("\t\tOpción invalida")
    }
}

func tabla()  {
  var num string
  var contenedor [11]int
  fmt.Println("\tMultiplicacion")
  fmt.Printf("\tIngrese un numero: ")
  fmt.Scanf("%s\n",&num)
  Nnum,_ := strconv.Atoi(num)
  // fmt.Println(Nnum + 2)
  for i:=0;i<=len(contenedor) - 1;i++ {
    fmt.Println("\t",Nnum,"*",i,"=",Nnum * i)
    contenedor[i] = Nnum * i
  }
  fmt.Println("\n\tEsta es tu tabla de multiplicar, y aqui abajo un vestor cargador")
  fmt.Println("\t",contenedor)
}

func suma(a int,b int)  int{
  c := a + b
  return c
}

func Tellamas(n string,a int)  {
  tellamas := MyStruct{name:n,age:a}
  fmt.Println("\n\t\tHola ",tellamas.name," tienes ",tellamas.age," años.")
}
type MyStruct struct{
  name string
  age int
}
