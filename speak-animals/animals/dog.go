package animals

type Dog struct {
    Age interface{}
}

func (d Dog) Speak() string {
    return "Woof!"
}

