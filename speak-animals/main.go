package main

import (
	"fmt"

	"github.com/mmullis/learning-go/speak-animals/animals"
)

func mainV1() {
    animals := []animals.Animal{animals.Dog{}, animals.Cat{}, animals.Llama{}, animals.JavaProgrammer{}}
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}

func main() {
    dog := animals.Dog{}
    dog.Age = "3"
    fmt.Printf("%#v %T\n", dog.Age, dog.Age)

    dog.Age = 3
    fmt.Printf("%#v %T\n", dog.Age, dog.Age)

    dog.Age = "not really an age"
    fmt.Printf("%#v %T", dog.Age, dog.Age)

    fmt.Println("Calling mainv1")
    mainV1()
}

