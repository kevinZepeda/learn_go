package main

import(
    "fmt"
    "strconv"
)
// Global vars
var jp int = 21



func main() {
  // -------------------- Declaration---------------------------------
  // var [Name of variable] [Type of variable]
  // print value (%v) and type (%T) of this variable
  // This is another format to declaration name := [value]
  // the vars have to be used
  var number_name int
  number_name = 42
  fmt.Printf("%v",number_name)
  var j float32 = 75
  fmt.Printf("%v,%T",j,j)

  // Another Declaration
  i := 42  // [name of variable] := [Value]
  // With this format can't redeclare variable
  fmt.Printf("%T",i)

  //Var blocks
  // This blocks its only for more cleaning code
  var (
    name string = "Kevin"
    fake_age int = 25
  )

  var (
    another_block string = "This random message"
    conuter_average float64 = 9.9
  )


  fmt.Printf("%v",name)
  fmt.Printf("%v",fake_age)
  fmt.Printf("%v",another_block)
  fmt.Printf("%v",conuter_average)
  fmt.Println("")

  // ----------------Redeclaration --------------------------------
  // We can redeclare variables if variable are not in the same block
  // The format to redeclare is -> var name_var [type]
  // Not this -> name_var := [value]
  // In this example "jp" is global var and print and redeclare intro main block
  fmt.Println(jp)
  var jp int = 32
  fmt.Println(jp)
  fmt.Println(" ")

  var entero int = 45
  fmt.Printf("%v,%T \n",entero,entero)

  var flotante float32
  flotante = float32(entero)
  fmt.Printf("%v,%T\n",flotante,flotante)

  var cadena string
  // cadena = string(entero)// get ascii caracter by integer
  cadena = strconv.Itoa(entero)
  fmt.Printf("%v,%T",cadena,cadena)

}
