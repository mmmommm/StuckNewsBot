package main

import (
    "./mylib"
    "fmt"
)

func main() {
    s := []int{1,2,3,4,5}
    fmt.Println(mylib.Average(s))

    mylib.Say()
    person := mylib.Person{Name: "mike", Age: 20}
    fmt.Println(person)
}