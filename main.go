package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Name   string
	Age    int
	Course int
}

func NewStudent(Name string, Age, Course int) Student {
	Freshman := Student{
		Name:   Name,
		Age:    Age,
		Course: Course,
	}
	return Freshman
}

func main() {
	Freshman := NewStudent("Imomali", 20, 1)
	file, err := os.Create("./data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	JsonData, err1 := json.Marshal(Freshman)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	_, err = file.Write(JsonData)
	if err != nil {
		fmt.Println(err)
		return
	}

	var ThirdCourse Student
	err = json.Unmarshal(JsonData, &ThirdCourse)
	fmt.Println(ThirdCourse)
}
