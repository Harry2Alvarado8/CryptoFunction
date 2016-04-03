package main

import(
  "fmt"
  "strconv" // libreria para confertir datos de int a str or string a int
)
func main() {
    var a,b,exp,age int
    var name string
    menu := true
    // fmt.Println("\t1 : Suma\n\t2 : Tabla de multiplicacion\n\t3 : Hola")
    // fmt.Printf("\t\tQue opción: ")
    // fmt.Scanf("%d\n",&exp)

    for menu{
        fmt.Println("\n\n\t1 : Suma\n\t2 : Tabla de multiplicacion\n\t3 : Hola\n\t4 : Verificar notas de Estudiantes\n\t5 : Salir")
        fmt.Printf("\t\tQue opción: ")
        fmt.Scanf("%d\n",&exp)
    switch exp {
    case 1:
          fmt.Println("\t\tIngrese dos numeros")
          fmt.Printf("\n\tPrimer numero: ")
          fmt.Scanf("%d\n",&a)
          fmt.Printf("\n\tSegundo numero: ")
          fmt.Scanf("%d\n",&b)
          fmt.Println("\n\t\tla suma es ",suma(a,b))
    case 2: tabla()
    case 3:
            fmt.Printf("\n\tDime tu nombre : ")
            fmt.Scanf("%s\n",&name)
            fmt.Printf("\n\tQue edad tienes : ")
            fmt.Scanf("%d\n",&age)
            Tellamas(name,age)
    case 4:
            //fmt.Printf("\n\tCantidad de estudiantes: ")
            // fmt.Scanf("%d\n",&a)                    lograr que el usuario pueda
            // fmt.Printf("\n\tCantidad de NOTAS : ")  ser quien asigne la cantidad de estudiantes
            // fmt.Scanf("%d\n",&b)                    y notas de los estudiantes
            Alumnos()
    case 5: fmt.Println("\n\tGracias por a ver visitado mi menu")
                menu = false
    default: fmt.Println("\t\tOpción invalida")
    }
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

func Alumnos()  {
    var name [3]string
    var notas[3]string
    var CantNotas [4]int
    var acum [5]int
    for i := 0;i < 3;i++{
        fmt.Printf("\n\tNombre del alumno : ")
        fmt.Scanf("%s\n",&name[i])
        for n := 0; n < 4; n++ { //cantidad de notas
            fmt.Printf("\t:-> ")
            fmt.Scanf("%d\n",&CantNotas[n])
        }
        for verif := 0; verif < 4; verif++ {
            acum[i] += CantNotas[verif]
        }
    }
    for not := 0; not < 3; not++ {
        if (acum[not] / 3) < 61 {
            notas[not] = "F"
        }else if (acum[not] / 3) < 71 {
            notas[not] = "D"
        }else if (acum[not] / 3) < 81 {
            notas[not] = "C"
        }else if (acum[not] / 3) < 91{
            notas[not] = "B"
        }else{
            notas[not] = "A"
        }
    }
    for im := 0; im < 3; im++ {
        fmt.Println("\n\tEl Estudiante ",name[im]," tiene ",(acum[im] / 3)," puntos loque Equivale a : ",notas[im])
    }

    //lograr que el usuario pueda
    //ser quien asigne la cantidad de estudiantes
    //y notas de los estudiantes

}
