package main

import (
	"fmt"
)

// Stringfyインターフェース
type Stringfy interface {
	ToString() string
}

// Person構造体
type Person struct {
	Name   string
	Age    int
	Gender string
}

// Person構造体のメソッドToString
func (person *Person) ToString() string {
	return fmt.Sprintf("%s(%d歳:%s)", person.Name, person.Age, person.Gender)
}

// Animal構造体
type Animal struct {
	Name string
	Kind string
}

// Animal構造体のメソッドToString
func (animal *Animal) ToString() string {
	return fmt.Sprintf("%s[%s]", animal.Name, animal.Kind)
}

func main() {
	vs := []Stringfy{
		&Person{Name: "山田太郎", Age: 20, Gender: "男"},
		&Animal{Name: "牛", Kind: "哺乳類"},
	}
	for _, v := range vs {
		fmt.Println(v.ToString())
	}
}
