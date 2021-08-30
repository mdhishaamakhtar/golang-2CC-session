package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name     string
	Email    string
	Phone    int
	Password string
	Friends  []string
}

type Animal struct {
	Name   string `requiredMax:"100"`
	Origin string
}

type Bird struct {
	Animal
	speed  int
	canFly bool
}

func main() {
	user1 := User{
		Name:     "Hishaam",
		Email:    "hishaamakhtar2001.mha@gmail.com",
		Phone:    9051816615,
		Password: "password",
		Friends: []string{
			"Hasan",
			"Sharanya",
			"Mayank",
		},
	}
	bird1 := Bird{}
	bird1.Name = "Emu"
	bird1.Origin = "Australia"
	bird1.speed = 45
	bird1.canFly = false
	t := reflect.TypeOf(bird1)
	field, _ := t.FieldByName("Name")
	fmt.Println(user1)
	fmt.Println(bird1)
	fmt.Println(field.Tag)
}
